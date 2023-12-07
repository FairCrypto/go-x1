package verifier

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"regexp"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
