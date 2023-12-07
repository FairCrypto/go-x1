package verifier

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Fantom-foundation/go-opera/gossip/contract/sfclib100"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/hashicorp/golang-lru"
	"github.com/tvdburgt/go-argon2"
	"math"
	"math/big"
	"os"
	"time"
)

var (
	numOfWorkers     = 3
	backlog          = 5
	hashLen          = 64
	validatorsFactor = 0.2
	pattern1Salt     = "WEVOMTAwODIwMjJYRU4"
	blockStorageAddr = "0x23213196F7d13153a906c456f5a1008E23EA94bA"
)

type Verifier struct {
	enabled           bool
	numOfWorkers      int
	backlog           int
	ipcPath           string
	validatorId       uint32
	bs                *block_storage.BlockStorage
	sfcLib            *sfclib100.ContractCaller
	eventChannel      chan *block_storage.BlockStorageNewHash
	conn              *ethclient.Client
	sub               event.Subscription
	validatorCountLRU *lru.Cache
}

func NewVerifier(nodeCfg node.Config, validatorId idx.ValidatorID) *Verifier {
	ipcPath := fmt.Sprintf("%s/%s", nodeCfg.DataDir, nodeCfg.IPCPath)

	return &Verifier{
		enabled:      false,
		ipcPath:      ipcPath,
		numOfWorkers: numOfWorkers,
		backlog:      backlog,
		validatorId:  uint32(validatorId),
	}
}

func (x *Verifier) Start() {
	log.Info("Starting Block storage watcher")
	time.Sleep(5 * time.Second)
	x.enabled = true
	for {
		if _, err := os.Stat(x.ipcPath); errors.Is(err, os.ErrNotExist) {
			log.Warn("IPC file not found, waiting 10 seconds")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	var err error

	x.validatorCountLRU, err = lru.New(100)

	x.conn, err = ethclient.Dial(x.ipcPath)
	if err != nil {
		panic(err)
	}

	x.bs, err = block_storage.NewBlockStorage(common.HexToAddress(blockStorageAddr), x.conn)
	if err != nil {
		panic(err)
	}

	x.sfcLib, err = sfclib100.NewContractCaller(sfc.ContractAddress, x.conn)
	if err != nil {
		panic(err)
	}

	x.eventChannel = make(chan *block_storage.BlockStorageNewHash, backlog)
	for w := 1; w <= numOfWorkers; w++ {
		go x.worker(x.eventChannel)
	}

	// Start a goroutine which watches new events
	go func() {
		x.sub, err = x.bs.WatchNewHash(nil, x.eventChannel, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-x.sub.Err():
				log.Error("Error in BlockStorage watcher", "err", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
}

func (x *Verifier) worker(events <-chan *block_storage.BlockStorageNewHash) {
	for e := range events {
		x.handleEvent(e)
	}
}

func (x *Verifier) handleEvent(event *block_storage.BlockStorageNewHash) {
	log.Debug("NewHash event received", "hashId", event.HashId, "epoch", event.Epoch, "account", event.Account)

	if !x.shouldVote(event.Raw.BlockNumber) {
		log.Debug("Not voting for this hash", "hash", event.Raw.BlockHash)
		return
	}

	dr, err := x.bs.DecodeRecordBytes0(nil, event.HashId)
	if err != nil {
		return
	}

	hashToVerify, err := argon2Hash(dr.C, dr.M, dr.T, dr.S, dr.K)
	if err != nil {
		panic(err)
	}

	if validateSalt(dr.S) {
		log.Info("Hash verified", "hash", hashToVerify)
	} else {
		log.Warn("Hash verification failed", "hash", hashToVerify)
	}

	// TODO: parse XNM, XUNI or XNM+X.BLK
	// TODO: vote for each hash
}

func argon2Hash(parallelism uint8, memory uint32, iterations uint8, salt []byte, key [32]byte) (string, error) {
	ctx := argon2.NewContext()
	ctx.Iterations = int(iterations)
	ctx.Memory = int(memory)
	ctx.Parallelism = int(parallelism)
	ctx.HashLen = hashLen
	ctx.Mode = argon2.ModeArgon2i
	ctx.Version = argon2.Version13

	return argon2.HashEncoded(ctx, key[:], salt)
}

func (x *Verifier) getValidatorCount(blockNumber uint64) (int, error) {
	bn := new(big.Int).SetUint64(blockNumber)
	epoch, err := x.sfcLib.CurrentSealedEpoch(&bind.CallOpts{BlockNumber: bn})
	log.Trace("CurrentSealedEpoch", "epoch", epoch)
	if err != nil {
		return 0, err
	}

	// store the validator count in cache
	r, ok := x.validatorCountLRU.Get(epoch.Int64())
	if ok {
		return r.(int), nil
	}

	validators, err := x.sfcLib.GetEpochValidatorIDs(&bind.CallOpts{BlockNumber: bn}, epoch)
	log.Trace("GetEpochValidatorIDs", "validators", validators)

	count := len(validators)
	x.validatorCountLRU.Add(epoch.Int64(), count)

	return count, err
}

func (x *Verifier) shouldVote(blockNumber uint64) bool {
	validatorsCount, err := x.getValidatorCount(blockNumber)
	if err != nil {
		log.Error("Error getting validator count", "err", err)
		return false
	}

	seed := sha256.Sum256([]byte(fmt.Sprintf("%d-%d", blockNumber, x.validatorId)))
	randomState := int(binary.LittleEndian.Uint64(seed[:])) % validatorsCount

	// Calculate the number of validators to vote (percentage of validators)
	numVotingValidators := max(1, int(math.Round(validatorsFactor*float64(validatorsCount))))

	// Calculate the position of the validator within the selected validators
	positionInSelectedValidators := randomState % numVotingValidators

	// Check if the given validator index matches the calculated position
	return positionInSelectedValidators == int(x.validatorId)%numVotingValidators
}

func (x *Verifier) Close() {
	if x.enabled {
		log.Info("Closing Block storage watcher")
		x.sub.Unsubscribe()
		time.Sleep(time.Second)
		close(x.eventChannel)
		x.conn.Close()
	}
}
