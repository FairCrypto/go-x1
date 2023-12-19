package verifier

import (
	"context"
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
	"time"
)

const (
	hashLen          = 64
	validatorsFactor = 0.2
	pattern1Salt     = "WEVOMTAwODIwMjJYRU4"
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

func (v *Verifier) validateHashEvent(event *block_storage.BlockStorageNewHash) []Token {
	log.Debug("NewHash event received", "hashId", event.HashId, "epoch", event.Epoch, "account", event.Account)

	if !v.shouldVote(event.Raw.BlockNumber) {
		log.Debug("Not voting for this hash", "hash", event.Raw.BlockHash)
		return nil
	}

	blockTime := v.getBlockTime(event.Raw.BlockHash)

	parallelism, memory, iterations, _, salt, key := decodeRecordBytes(event.Bytes)
	argon2Result, err := argon2Hash(parallelism, memory, iterations, salt, key)
	if err != nil {
		log.Warn("Argon2 hash failed", "err", err)
		return nil
	}

	if len(argon2Result) > 150 {
		log.Warn("Hash too long", "hash", argon2Result, "hashId", event.HashId)
		return nil
	}

	if !validateSalt(salt) {
		log.Warn("Salt fails verification", "hash", argon2Result, "hashId", event.HashId)
		return nil
	}

	//if !v.verifyDifficultly(memory, event.Raw.BlockNumber) {
	//	log.Warn("Difficulty too low", "hash", argon2Result, "hashId", event.HashId)
	//	return nil
	//}

	tokens := FindTokensFromHash(argon2Result, blockTime)
	if len(tokens) == 0 {
		log.Warn("No tokens found", "hash", argon2Result, "hashId", event.HashId)
		return nil
	}

	log.Info("hash verified", "hash", argon2Result, "tokens", tokens, "hashId", event.HashId)
	return tokens
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

func (v *Verifier) verifyDifficultly(difficulty uint32, blockNumber uint64) bool {
	bn := new(big.Int).SetUint64(blockNumber)
	d, err := v.bs.Difficulty(&bind.CallOpts{BlockNumber: bn})
	if err != nil {
		panic(err)
	}

	log.Trace("Difficulty", "difficulty", d.Uint64(), "d", d.Uint64())
	return d.Uint64() > uint64(difficulty)
}

func (v *Verifier) getBlockTime(blockHash common.Hash) time.Time {
	header, err := v.conn.HeaderByHash(context.TODO(), blockHash)
	if err != nil {
		panic(err)
	}

	return time.Unix(int64(header.Time), 0)
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
		log.Warn("Base64 decode error", "err", err, "salt", salt)
		return false
	}

	decodedStr := hex.EncodeToString(rawDecodedText)
	if !common.IsHexAddress(decodedStr) {
		log.Warn("Decoded string is not a valid hash", "decodedStr", decodedStr)
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

func decodeRecordBytes(data []byte) (parallelism uint8, memory uint32, iterations uint8, version byte, salt []byte, key []byte) {
	lenData := len(data)

	// Read values using binary.BigEndian
	parallelism = data[0]
	memory = binary.BigEndian.Uint32(data[1:5])
	iterations = data[5]
	version = data[6]
	keyOut := []byte(common.Bytes2Hex(data[7:39]))

	// Copy the remaining bytes to s
	salt = make([]byte, lenData-39)
	copy(salt, data[39:])

	return parallelism, memory, iterations, version, salt, keyOut
}
