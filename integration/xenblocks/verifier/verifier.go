package verifier

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/Fantom-foundation/go-opera/gossip/contract/sfclib100"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/Fantom-foundation/go-opera/opera/contracts/sfc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/golang-lru"
	"github.com/tvdburgt/go-argon2"
	"math"
	"math/big"
	"regexp"
	"strings"
)

var (
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

func NewVerifier(validatorId uint32, conn *ethclient.Client, bs *block_storage.BlockStorage) *Verifier {
	validatorCountLRU, err := lru.New(100)
	if err != nil {
		panic(err)
	}

	sfcLib, err := sfclib100.NewContractCaller(sfc.ContractAddress, conn)
	if err != nil {
		panic(err)
	}

	return &Verifier{
		enabled:           false,
		numOfWorkers:      numOfWorkers,
		backlog:           backlog,
		validatorId:       validatorId,
		conn:              conn,
		validatorCountLRU: validatorCountLRU,
		sfcLib:            sfcLib,
		bs:                bs,
	}
}

func (v *Verifier) handleEvent(event *block_storage.BlockStorageNewHash) {
	log.Debug("NewHash event received", "hashId", event.HashId, "epoch", event.Epoch, "account", event.Account)

	if !v.shouldVote(event.Raw.BlockNumber) {
		log.Debug("Not voting for this hash", "hash", event.Raw.BlockHash)
		return
	}

	dr, err := v.bs.DecodeRecordBytes0(nil, event.HashId)
	if err != nil {
		return
	}

	key := []byte(common.Bytes2Hex(dr.K[:]))
	argon2Result, err := argon2Hash(dr.C, dr.M, dr.T, dr.S, key)
	if err != nil {
		panic(err)
	}

	if len(argon2Result) > 150 {
		log.Warn("Hash too long", "hash", argon2Result)
		return
	}

	if !validateSalt(dr.S) {
		log.Warn("Salt fails verification", "hash", argon2Result)
		return
	}

	if !v.verifyDifficultly(argon2Result) {
		log.Warn("Difficulty too low", "hash", argon2Result)
		return
	}

	hash := getHashPortion(argon2Result)
	isXuniPresent := xuniPresent(hash)
	isXenPresent := xenPresent(hash)
	superBlock := isSuperBlock(hash, isXenPresent)

	log.Info("hash verified", "hash", argon2Result, "isXuniPresent",
		isXuniPresent, "isXenPresent", isXenPresent, "superBlock", superBlock)
	// TODO: vote for each hash
}

func argon2Hash(parallelism uint8, memory uint32, iterations uint8, salt []byte, key []byte) (string, error) {
	ctx := argon2.NewContext()
	ctx.Iterations = int(iterations)
	ctx.Memory = int(memory)
	ctx.Parallelism = int(parallelism)
	ctx.HashLen = hashLen
	ctx.Mode = argon2.ModeArgon2id
	ctx.Version = argon2.Version13

	return argon2.HashEncoded(ctx, key, salt)
}

func (v *Verifier) getValidatorCount(blockNumber uint64) (int, error) {
	bn := new(big.Int).SetUint64(blockNumber)
	epoch, err := v.sfcLib.CurrentSealedEpoch(&bind.CallOpts{BlockNumber: bn})
	log.Trace("CurrentSealedEpoch", "epoch", epoch)
	if err != nil {
		return 0, err
	}

	// store the validator count in cache
	r, ok := v.validatorCountLRU.Get(epoch.Int64())
	if ok {
		return r.(int), nil
	}

	validators, err := v.sfcLib.GetEpochValidatorIDs(&bind.CallOpts{BlockNumber: bn}, epoch)
	log.Trace("GetEpochValidatorIDs", "validators", validators)

	count := len(validators)
	v.validatorCountLRU.Add(epoch.Int64(), count)

	return count, err
}

func (v *Verifier) shouldVote(blockNumber uint64) bool {
	validatorsCount, err := v.getValidatorCount(blockNumber)
	if err != nil {
		log.Error("Error getting validator count", "err", err)
		return false
	}

	seed := sha256.Sum256([]byte(fmt.Sprintf("%d-%d", blockNumber, v.validatorId)))
	randomState := int(binary.LittleEndian.Uint64(seed[:])) % validatorsCount

	// Calculate the number of validators to vote (percentage of validators)
	numVotingValidators := max(1, int(math.Round(validatorsFactor*float64(validatorsCount))))

	// Calculate the position of the validator within the selected validators
	positionInSelectedValidators := randomState % numVotingValidators

	// Check if the given validator index matches the calculated position
	return positionInSelectedValidators == int(v.validatorId)%numVotingValidators
}

func (v *Verifier) verifyDifficultly(hashToVerify string) bool {
	// TODO: verify difficulty
	return true
}

func validatePattern1(salt string) bool {
	return salt == pattern1Salt
}

func validatePattern2(salt string) bool {
	r := regexp.MustCompile(`^[A-Za-z0-9+/]{27}$`)

	if !r.MatchString(salt) {
		log.Warn("pattern 2 match failed")
		return false
	}

	missingPadding := len(salt) % 4
	if missingPadding != 0 {
		salt += strings.Repeat("=", 4-missingPadding)
	}

	rawDecodedText, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		log.Warn("base64 decode error", "err", err, "salt", salt)
		return false
	}

	decodedStr := hex.EncodeToString(rawDecodedText)
	if !common.IsHexAddress(decodedStr) {
		log.Warn("decoded string is not a valid hash", "decodedStr", decodedStr)
		return false
	}

	return true
}

func validateSalt(s []byte) bool {
	salt := base64.RawStdEncoding.EncodeToString(s)
	if validatePattern1(salt) {
		return true
	}

	return validatePattern2(salt)
}

func xuniPresent(hash string) bool {
	match := regexp.MustCompile(`XUNI[0-9]`).MatchString(hash)
	return match
}

func xenPresent(hash string) bool {
	return strings.Contains(hash, "XEN11")
}

func isSuperBlock(hash string, isXenPresent bool) bool {
	if !isXenPresent {
		return false
	}
	return countUppercase(hash) > 50
}

func getHashPortion(hashToVerify string) string {
	splits := strings.Split(hashToVerify, "$")
	return splits[len(splits)-1]
}
