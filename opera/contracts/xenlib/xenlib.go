package xenlib

import (
	"bytes"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/crypto/argon2"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

var (
	ContractAddress = common.HexToAddress("0xd100ec0000000000000000000000000000000005")
	// ContractABI is the input ABI used to generate the binding from

	ContractABI string = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"argon2Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
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

func (_ PreCompiledContract) Run(stateDB vm.StateDB, _ vm.BlockContext, txCtx vm.TxContext, caller common.Address, input []byte, suppliedGas uint64) ([]byte, uint64, error) {
	log.Warn("hello you are in xenlib")
	if len(input) < 1 {
		return nil, 0, vm.ErrExecutionReverted
	}
	if bytes.Equal(input[:4], argon2HashIdMethodID) {
		hash := bytes.Trim(input[68:100], "\x00")
		key := argon2.IDKey(hash, []byte(""), 1, 100*1024, 10, 32)
		log.Info("argon2 hash", "password", string(hash), "key", string(key))

		// don't know how to return the key, yet
		return nil, suppliedGas, nil

	} else {
		return nil, 0, vm.ErrExecutionReverted
	}
}
