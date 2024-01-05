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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"time"
)

const (
	numOfWorkers = 5
	backlog      = 5000

	confirmations     = 3
	allowBlockBacklog = 1000
	BlockStorageAddr  = "0xf7E0CF7453ac619fD64b3D46D7De3638510F15eA"
	VoteManagerAddr   = "0x1e29fae73cbe681f35208d470db8c7113820f0c2"
	TokenRegistryAddr = "0x830c235F1CCa0c6760931d2A5F7e9bC608E1c750"
)

type EventListener struct {
	enabled            bool
	numOfWorkers       int
	backlog            int
	stack              *node.Node
	validatorId        uint32
	bs                 *block_storage.BlockStorage
	vm                 *votemanager.Votemanager
	tr                 *tokenregistry.Tokenregistry
	eventChannel       chan *block_storage.BlockStorageNewHash
	conn               *ethclient.Client
	verifier           *Verifier
	ks                 *keystore.KeyStore
	account            accounts.Account
	chainId            uint64
	voter              *Voter
	currentBlockNumber uint64
	syncingProcess     *ethereum.SyncProgress
}

func NewEventListener(stack *node.Node, validatorId idx.ValidatorID, ks *keystore.KeyStore, account accounts.Account, chainId uint64) *EventListener {
	return &EventListener{
		enabled:            false,
		stack:              stack,
		numOfWorkers:       numOfWorkers,
		backlog:            backlog,
		validatorId:        uint32(validatorId),
		ks:                 ks,
		account:            account,
		chainId:            chainId,
		currentBlockNumber: 0,
	}
}

func (e *EventListener) Start() {
	log.Info("Starting Block storage watcher")
	e.enabled = true

	err := e.initializeEventSystem()
	if err != nil {
		panic(err)
	}

	e.eventChannel = make(chan *block_storage.BlockStorageNewHash, backlog)
	for w := 0; w < numOfWorkers; w++ {
		go e.worker(e.eventChannel)
	}
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

func (e *EventListener) OnNewLog(l *types.Log) {
	if e.enabled == false {
		return
	}

	if e.syncingProcess != nil && e.syncingProcess.HighestBlock-e.syncingProcess.CurrentBlock > allowBlockBacklog {
		return
	}

	if l.Address == common.HexToAddress(BlockStorageAddr) {
		evt, err := e.bs.ParseNewHash(*l)
		if err != nil {
			return
		}
		e.eventChannel <- evt
	}
}

func (e *EventListener) OnNewBlock(blockNumber uint64) {
	if e.enabled == false {
		return
	}

	progress, err := e.conn.SyncProgress(context.TODO())
	if err != nil {
		panic(err)
	}

	e.syncingProcess = progress
	e.currentBlockNumber = blockNumber
}

func (e *EventListener) worker(events <-chan *block_storage.BlockStorageNewHash) {
	for evt := range events {
		if e.enabled == false {
			return
		}

		// Wait for the block to be confirmed
		for {
			if evt.Raw.BlockNumber+confirmations < e.currentBlockNumber {
				break
			} else {
				log.Debug("waiting for block to be confirmed", "blockNumber", evt.Raw.BlockNumber, "currentBlockNumber", e.currentBlockNumber)
				time.Sleep(250 * time.Millisecond)
			}
		}

		tokens := e.verifier.validateHashEvent(evt)
		e.voter.Syncing = e.syncingProcess != nil // increase the batch size if we are syncing
		for _, token := range tokens {
			e.voter.AddToQueue(evt.HashId, evt.Raw.BlockNumber, token.CurrencyType, token.Version)
		}
	}
}

func (e *EventListener) Close() {
	if e.enabled {
		e.enabled = false
		log.Info("Closing Block storage watcher")
		time.Sleep(time.Second)
		close(e.eventChannel)
		e.conn.Close()
	}
}
