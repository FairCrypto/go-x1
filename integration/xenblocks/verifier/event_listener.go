package verifier

import (
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"time"
)

const (
	numOfWorkers = 5
	backlog      = 5000

	blockStorageAddr = "0xb3753e9F40DD0Dfd039e8c4B12895e2f636693a2"
)

type EventListener struct {
	enabled            bool
	numOfWorkers       int
	backlog            int
	stack              *node.Node
	validatorId        uint32
	bs                 *block_storage.BlockStorage
	eventChannel       chan *block_storage.BlockStorageNewHash
	conn               *ethclient.Client
	sub                event.Subscription
	verifier           *Verifier
	ks                 *keystore.KeyStore
	account            accounts.Account
	chainId            uint64
	voter              *Voter
	currentBlockNumber uint64
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
	time.Sleep(5 * time.Second)
	e.enabled = true

	err := e.initializeEventSystem()
	if err != nil {
		panic(err)
	}

	e.eventChannel = make(chan *block_storage.BlockStorageNewHash, backlog)
	for w := 1; w <= numOfWorkers; w++ {
		go e.worker(e.eventChannel)
	}

	// Start a goroutine which watches new events
	//go func() {
	//	e.sub, err = e.bs.WatchNewHash(nil, e.eventChannel, nil, nil, nil)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	for {
	//		select {
	//		case err := <-e.sub.Err():
	//			if err != nil && err != io.EOF {
	//				log.Error("Error in BlockStorage watcher", "err", err)
	//			}
	//			break
	//		}
	//		time.Sleep(time.Second)
	//	}
	//}()
}

func (e *EventListener) initializeEventSystem() error {
	rpc, err := e.stack.Attach()
	if err != nil {
		return err
	}

	e.conn = ethclient.NewClient(rpc)
	e.bs, err = block_storage.NewBlockStorage(common.HexToAddress(blockStorageAddr), e.conn)

	e.verifier = NewVerifier(e.validatorId, e.conn, e.bs)
	e.voter = NewVoter(e.conn, e.ks, e.account, e.chainId)

	return err
}

func (e *EventListener) OnNewLog(l *types.Log) {
	if l.Address == common.HexToAddress(blockStorageAddr) {

		evt, err := e.bs.ParseNewHash(*l)
		if err != nil {
			return
		}

		//log.Info("NewHash event received", "hashId", evt.HashId, "epoch", evt.Epoch, "account", evt.Account)
		e.eventChannel <- evt
	}
}

func (e *EventListener) OnNewBlock(blockNumber uint64) {
	//log.Info("New block received", "blockNumber", blockNumber)
	e.currentBlockNumber = blockNumber
}

func (e *EventListener) worker(events <-chan *block_storage.BlockStorageNewHash) {
	for evt := range events {
		// Wait for the block to be confirmed
		for {
			if evt.Raw.BlockNumber < e.currentBlockNumber {
				break
			} else {
				log.Info("waiting for block to be confirmed", "blockNumber", evt.Raw.BlockNumber, "currentBlockNumber", e.currentBlockNumber)
				time.Sleep(100 * time.Millisecond)
			}
		}
		tokens := e.verifier.validateHashEvent(evt)
		for _, token := range tokens {
			e.voter.AddToQueue(evt.HashId, token.currencyCode)
		}
	}
}

func (e *EventListener) Close() {
	if e.enabled {
		log.Info("Closing Block storage watcher")
		time.Sleep(time.Second)
		close(e.eventChannel)
		e.conn.Close()
	}
}
