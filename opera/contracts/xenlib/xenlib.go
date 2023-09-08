package xenlib

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/crypto/argon2"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ContractAddress = common.HexToAddress("0xd100ec0000000000000000000000000000000005")
	// ContractABI is the input ABI used to generate the binding from

	ContractABI string = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"password\",\"type\":\"bytes\"}],\"name\":\"argon2Hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

	Gas     uint64 = 30 // Once per operation.
	WordGas uint64 = 6  // Once per word of the operation's data.
)

var (
	argon2HashIdMethodID []byte
)

func init() {
	abi, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		panic(err)
	}

	for name, constID := range map[string]*[]byte{
		"argon2Hash": &argon2HashIdMethodID,
	} {
		method, exist := abi.Methods[name]
		if !exist {
			panic("unknown xenlib method")
		}

		*constID = make([]byte, len(method.ID))
		copy(*constID, method.ID)
	}
}

type PreCompiledContract struct{}

func (_ PreCompiledContract) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*WordGas + Gas
}

func (_ PreCompiledContract) Run(input []byte) ([]byte, error) {
	salt := []byte("XEN10082022XEN")
	time := uint32(1)
	memory := uint32(80)
	threads := uint8(8)
	keyLen := uint32(64)

	password := bytes.TrimSpace(input)
	hash := argon2.IDKey(password, salt, time, memory, threads, keyLen)

	log.Info("argon2Hash", "password", string(password))
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	output := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, time, threads, b64Salt, b64Hash)
	log.Info("argon2Hash", "output", string(output))

	return []byte(output), nil
}
