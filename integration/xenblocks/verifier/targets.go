package verifier

import (
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"regexp"
	"strings"
	"time"
)

var (
	// BLOCK_NUMBER_TIME_CHECK_REQ is the block number at which the time check starts to be required.
	BLOCK_NUMBER_TIME_CHECK_REQ = big.NewInt(0)
)

type Token struct {
	name         string
	currencyCode uint8
	value        int
}

type TokenFilter struct {
	token     string
	pattern   string                                          // regex pattern to filter hashes
	exclusive bool                                            // if true, stop searching for other targets
	timeCheck func(time time.Time, blockNumber *big.Int) bool // custom function to filter based on time
	hashCheck func(hash string) bool                          // custom function to filter based on hash
}

// Tokens These are token definitions for the verifier and their initial mint values.
var Tokens = map[string]Token{
	"XNM": {
		name:         "XNM",
		currencyCode: 1,
		value:        10,
	},
	"XUNI": {
		name:         "XUNI",
		currencyCode: 2,
		value:        1,
	},
	"XBLK": {
		name:         "XBLK",
		currencyCode: 3,
		value:        1,
	},
}

// TokenFilters These are filters to run on each verified hash.
// They are run in order as listed below.
var TokenFilters = []TokenFilter{
	{
		token:     "XUNI",
		pattern:   ".*XUNI[0-9].*",
		exclusive: true,
		timeCheck: func(time time.Time, blockNumber *big.Int) bool {
			if BLOCK_NUMBER_TIME_CHECK_REQ.Cmp(big.NewInt(0)) == 0 || blockNumber.Cmp(BLOCK_NUMBER_TIME_CHECK_REQ) == -1 {
				return true
			}
			minutes := time.Minute()
			return (0 <= minutes && minutes < 5) || (55 <= minutes && minutes < 60)
		},
	},
	{
		token:     "XNM",
		pattern:   ".*XEN11.*",
		exclusive: false,
	},
	{
		token:     "XBLK",
		pattern:   ".*XEN11.*",
		exclusive: false,
		hashCheck: func(hash string) bool {
			return countUppercase(hash) >= 50
		},
	},
}

// FindTokensFromHash returns a list of tokens found in the given hash.
func FindTokensFromHash(argon2Result string, blockTime time.Time, blockNumber uint64) []Token {
	var foundTargets []Token

	splits := strings.Split(argon2Result, "$")
	hashOnly := splits[len(splits)-1]

	for i := range TokenFilters {
		target := TokenFilters[i]

		if target.pattern != "" && !regexp.MustCompile(target.pattern).MatchString(hashOnly) {
			continue
		}

		if target.hashCheck != nil && !target.hashCheck(hashOnly) {
			log.Trace("hash check failed", "token", target.token, "hash", hashOnly)
			continue
		}

		bn := new(big.Int).SetUint64(blockNumber)
		if target.timeCheck != nil && !target.timeCheck(blockTime, bn) {
			log.Trace("time check failed", "token", target.token, "time", blockTime)
			continue
		}

		log.Trace("target found", "token", target.token, "hash", hashOnly)
		foundTargets = append(foundTargets, Tokens[target.token])

		if target.exclusive {
			log.Trace("exclusive target found. stop searching", "token", target.token)
			break
		}
	}

	return foundTargets
}
