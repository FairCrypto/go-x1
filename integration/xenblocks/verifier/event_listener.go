package verifier

import (
	"context"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/tokenregistry"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/votemanager"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	numOfWorkers             = 5
	backlog                  = 5000
	checkpointSize           = 1000 // save the check point every n blocks
	VoteManagerStart         = 865652
	BlockStorageStart        = 832437
	updateCheckpointGasLimit = 80000
	BlockStorageAddr         = "0xf7E0CF7453ac619fD64b3D46D7De3638510F15eA"
	VoteManagerAddr          = "0x1e29fae73cbe681f35208d470db8c7113820f0c2"
	TokenRegistryAddr        = "0x830c235F1CCa0c6760931d2A5F7e9bC608E1c750"
)

type EventListener struct {
	enabled              bool
	numOfWorkers         int
	backlog              int
	stack                *node.Node
	validatorId          uint32
	bs                   *block_storage.BlockStorage
	vm                   *votemanager.Votemanager
	tr                   *tokenregistry.Tokenregistry
	eventChannel         chan *block_storage.BlockStorageNewHash
	conn                 *ethclient.Client
	verifier             *Verifier
	ks                   *keystore.KeyStore
	account              accounts.Account
	chainId              uint64
	voter                *Voter
	currentBlockNumber   uint64
	processedBlockNumber uint64
	syncingProcess       *ethereum.SyncProgress
	newHashTopic         common.Hash
	blockWorkerRunning   bool
}

func NewEventListener(stack *node.Node, validatorId idx.ValidatorID, ks *keystore.KeyStore, account accounts.Account, chainId uint64) *EventListener {
	logTransferSig := []byte("NewHash(uint256,uint256,address,uint256,bytes)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	return &EventListener{
		enabled:              false,
		stack:                stack,
		numOfWorkers:         numOfWorkers,
		backlog:              backlog,
		validatorId:          uint32(validatorId),
		ks:                   ks,
		account:              account,
		chainId:              chainId,
		currentBlockNumber:   0,
		processedBlockNumber: 0,
		newHashTopic:         logTransferSigHash,
		blockWorkerRunning:   false,
	}
}

func (e *EventListener) Start() {
	log.Info("Starting Block storage watcher")
	e.enabled = true

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)

	err := e.initializeEventSystem()
	if err != nil {
		panic(err)
	}

	e.eventChannel = make(chan *block_storage.BlockStorageNewHash, backlog)
	for w := 0; w < numOfWorkers; w++ {
		go e.newHashWorker(e.eventChannel)
	}

	<-sigc
	e.Close()
}

func (e *EventListener) initializeEventSystem() error {
	rpc, err := e.stack.Attach()
	if err != nil {
		return err
	}

	e.conn = ethclient.NewClient(rpc)
	e.bs, err = block_storage.NewBlockStorage(common.HexToAddress(BlockStorageAddr), e.conn)
	e.vm, err = votemanager.NewVotemanager(common.HexToAddress(VoteManagerAddr), e.conn)
	e.tr, err = tokenregistry.NewTokenregistry(common.HexToAddress(TokenRegistryAddr), e.conn)

	e.verifier = NewVerifier(e.validatorId, e.conn, e.bs, e.vm, e.tr)
	e.voter = NewVoter(e.conn, e.ks, e.account, e.chainId, e.vm)

	return err
}

func (e *EventListener) OnNewBlock(blockNumber uint64) {
	if e.enabled == false {
		return
	}

	progress, err := e.conn.SyncProgress(context.TODO())
	if err != nil {
		if !e.enabled {
			return
		}
		panic(err)
	}

	e.syncingProcess = progress
	e.currentBlockNumber = blockNumber

	// do not start the block worker if we are synced
	if !e.blockWorkerRunning && e.syncingProcess == nil {
		go e.blockWorker()
		e.blockWorkerRunning = true
	}
}

func (e *EventListener) getLogs(blockNumber uint64) error {
	bn := big.NewInt(int64(blockNumber))

	query := ethereum.FilterQuery{
		FromBlock: bn,
		ToBlock:   bn,
		Addresses: []common.Address{
			common.HexToAddress(BlockStorageAddr),
		},
		Topics: [][]common.Hash{
			{e.newHashTopic},
		},
	}

	logs, err := e.conn.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	for _, vLog := range logs {
		event, err := e.bs.ParseNewHash(vLog)
		if err != nil {
			return err
		}

		select {
		case e.eventChannel <- event:
			return nil
		default:
		}
	}

	return nil
}

func (e *EventListener) blockWorker() {
	e.processedBlockNumber = e.getCheckpoint()

	for {
		if e.enabled == false {
			return
		}

		if e.currentBlockNumber > e.processedBlockNumber {
			err := e.getLogs(e.processedBlockNumber + 1)
			if err != nil {
				if !e.enabled {
					return
				}
				panic(err)
			}

			if e.processedBlockNumber%checkpointSize == 0 {
				// move on to the next block only after all events are processed
				e.waitForEventsProcessed()
				e.updateCheckpoint()
			}

			e.processedBlockNumber++
		} else {
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func (e *EventListener) waitForEventsProcessed() {
	log.Info("waiting for events to be processed", "blockNumber", e.processedBlockNumber)
	for {
		if e.enabled == false {
			return
		}
		if len(e.eventChannel) == 0 {
			return
		}

		time.Sleep(250 * time.Millisecond)
	}
}

func (e *EventListener) updateCheckpoint() {
	if e.enabled && e.syncingProcess == nil && e.currentBlockNumber >= VoteManagerStart {
		log.Info("updating checkpoint", "blockNumber", e.processedBlockNumber)

		auth, err := NewTransactionOpts(context.Background(), *e.conn, e.chainId, e.ks, e.account, updateCheckpointGasLimit)
		if err != nil {
			panic(err)
		}

		bn := big.NewInt(int64(e.processedBlockNumber))
		tx, err := e.vm.UpdateCheckpoint(auth, bn)

		if err != nil {
			if !e.enabled {
				return
			}
			panic(err)
		}

		WaitTxConfirmed(context.Background(), *e.conn, tx.Hash())
	}
}

func (e *EventListener) getCheckpoint() uint64 {
	if e.currentBlockNumber < VoteManagerStart {
		return BlockStorageStart
	}

	blockNumber, err := e.vm.BlockCheckpointByValidatorId(nil, big.NewInt(int64(e.validatorId)))
	if err != nil {
		if !e.enabled {
			return 0
		}
		panic(err)
	}

	log.Info("checkpoint", "blockNumber", blockNumber.Uint64(), "validatorId", e.validatorId)

	if blockNumber.Uint64() < BlockStorageStart {
		return BlockStorageStart
	}

	return blockNumber.Uint64()
}

func (e *EventListener) newHashWorker(events <-chan *block_storage.BlockStorageNewHash) {
	for evt := range events {
		if e.enabled == false {
			return
		}

		tokens := e.verifier.validateHashEvent(evt)
		e.voter.Syncing = e.syncingProcess != nil // increase the batch size if we are syncing
		for _, token := range tokens {
			e.voter.AddToQueue(evt.HashId, evt.Raw.BlockNumber, token.CurrencyType, token.Version)
		}
	}
}

func (e *EventListener) SetValidatorId(validatorId idx.ValidatorID) {
	log.Info("setting validator id", "validatorId", validatorId)
	e.validatorId = uint32(validatorId)
}

func (e *EventListener) Close() {
	if e.enabled {
		e.enabled = false
		log.Info("Closing Block storage watcher")
		time.Sleep(2 * time.Second)
		close(e.eventChannel)
		e.conn.Close()
	}
}
