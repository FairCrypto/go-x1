package verifier

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Fantom-foundation/go-opera/integration/xenblocks/contracts/block_storage"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/tvdburgt/go-argon2"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	hashId  []*big.Int
	epoch   []*big.Int
	account []common.Address

	ctx = &argon2.Context{
		HashLen: 64,
		Mode:    argon2.ModeArgon2id,
		Version: argon2.Version13,
	}
)

func Worker(cfg node.Config) {
	log.Info("Starting Block storage watcher")

	// check if file exists
	time.Sleep(10 * time.Second)
	ipcPath := fmt.Sprintf("%s/%s", cfg.DataDir, cfg.IPCPath)

	for {
		if _, err := os.Stat(ipcPath); errors.Is(err, os.ErrNotExist) {
			log.Warn("IPC file not found, waiting 10 seconds")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	conn, err := ethclient.Dial(ipcPath)
	if err != nil {
		panic(err)
	}

	bs, err := block_storage.NewBlockStorage(common.HexToAddress("0x23213196F7d13153a906c456f5a1008E23EA94bA"), conn)
	if err != nil {
		panic(err)
	}

	channel := make(chan *block_storage.BlockStorageNewHash, 5)

	// Start a goroutine which watches new events
	go func() {
		sub, err := bs.WatchNewHash(nil, channel, hashId, epoch, account)
		if err != nil {
			panic(err)
		}

		for {
			select {
			case err := <-sub.Err():
				panic(err)
			case t := <-channel:
				log.Trace("NewHash event received", "hashId", t.HashId, "epoch", t.Epoch, "account", t.Account)
				dr, err := bs.DecodeRecordBytes0(nil, t.HashId)
				if err != nil {
					return
				}

				ctx.Parallelism = int(dr.C)
				ctx.Memory = int(dr.M)
				ctx.Iterations = int(dr.T)
				hashToVerify, err := argon2.HashEncoded(ctx, dr.K[:], dr.S)
				if err != nil {
					panic(err)
				}

				if validateHash(hashToVerify) {
					log.Info("Hash verified", "hash", hashToVerify)
				} else {
					log.Warn("Hash verification failed", "hash", hashToVerify)
				}

				// then I can call a vote function

				default:
					log.Warn("Channel is full")
				}
			}
		}
	}()
}

func restoreEIP55Address(addr string) bool {
	re := regexp.MustCompile("^(0x|)[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

func extractSaltFromHash(hashToVerify string) string {
	parts := strings.Split(hashToVerify, "$")
	if len(parts) != 6 {
		log.Warn("less than 6 parts")
		return ""
	}
	return parts[4]
}

func validatePattern1(salt string) bool {
	return salt == "WEVOMTAwODIwMjJYRU4"
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
	if !restoreEIP55Address(decodedStr) {
		log.Warn("decoded string is not a valid hash", "decodedStr", decodedStr)
		return false
	}

	return true
}

func validateHash(hash string) bool {
	salt := extractSaltFromHash(hash)
	if salt == "" {
		return false
	}

	if validatePattern1(salt) {
		return true
	}

	return validatePattern2(salt)
}
