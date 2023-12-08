package verifier

import (
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"io"
	"time"
)

const (
	numOfWorkers = 3
	backlog      = 5
)

type EventListener struct {
	enabled      bool
	numOfWorkers int
	backlog      int
	stack        *node.Node
	validatorId  uint32
	bs           *block_storage.BlockStorage
	eventChannel chan *block_storage.BlockStorageNewHash
	conn         *ethclient.Client
	sub          event.Subscription
	verifier     *Verifier
}

func NewEventListener(stack *node.Node, validatorId idx.ValidatorID) *EventListener {
	return &EventListener{
		enabled:      false,
		stack:        stack,
		numOfWorkers: numOfWorkers,
		backlog:      backlog,
		validatorId:  uint32(validatorId),
	}
}

func (e *EventListener) Start() {
	log.Info("Starting Block storage watcher")
	time.Sleep(5 * time.Second)
	e.enabled = true

	var err error

	rpc, err := e.stack.Attach()
	if err != nil {
		panic(err)
	}
	e.conn = ethclient.NewClient(rpc)

	if err != nil {
		panic(err)
	}

	e.bs, err = block_storage.NewBlockStorage(common.HexToAddress(blockStorageAddr), e.conn)
	if err != nil {
		panic(err)
	}

	e.eventChannel = make(chan *block_storage.BlockStorageNewHash, backlog)
	for w := 1; w <= numOfWorkers; w++ {
		go e.worker(e.eventChannel)
	}

	e.verifier = NewVerifier(e.validatorId, e.conn, e.bs)

	// Start a goroutine which watches new events
	go func() {
		e.sub, err = e.bs.WatchNewHash(nil, e.eventChannel, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-e.sub.Err():
				if err != nil && err != io.EOF {
					log.Error("Error in BlockStorage watcher", "err", err)
				}
				break
			}
			time.Sleep(time.Second)
		}
	}()
}

func (e *EventListener) worker(events <-chan *block_storage.BlockStorageNewHash) {
	for evt := range events {
		e.verifier.handleEvent(evt)
	}
}

func (e *EventListener) Close() {
	if e.enabled {
		log.Info("Closing Block storage watcher")
		e.sub.Unsubscribe()
		time.Sleep(time.Second)
		close(e.eventChannel)
		e.conn.Close()
	}
}
