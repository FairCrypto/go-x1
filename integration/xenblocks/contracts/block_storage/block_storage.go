// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package block_storage

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BlockStorageHashRecord is an auto generated low-level Go binding around an user-defined struct.
type BlockStorageHashRecord struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}

// BlockStorageMetaData contains all meta data concerning the BlockStorage contract.
var BlockStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"EpochDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"}],\"name\":\"NewEpoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hashId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"NewHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"TargetDifficultySet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TARGET_BLOCK_TIME_S\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressesByHashId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpochTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentHashId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDurationSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochHashCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashIdsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recordCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochDurationSec\",\"type\":\"uint256\"}],\"name\":\"setEpochDurationSec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_targetDifficulty\",\"type\":\"uint256\"}],\"name\":\"setTargetDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doHousekeeping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"internalType\":\"structBlockStorage.HashRecord\",\"name\":\"r\",\"type\":\"tuple\"}],\"name\":\"storeNewRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"name\":\"storeNewRecordData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_bytes\",\"type\":\"bytes\"}],\"name\":\"storeNewRecordBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"internalType\":\"structBlockStorage.HashRecord[]\",\"name\":\"hh\",\"type\":\"tuple[]\"}],\"name\":\"bulkStoreRecordsInc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"bb\",\"type\":\"bytes[]\"}],\"name\":\"bulkStoreRecordBytesInc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_hasIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"internalType\":\"structBlockStorage.HashRecord[]\",\"name\":\"hh\",\"type\":\"tuple[]\"}],\"name\":\"bulkStoreRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_hasIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"bb\",\"type\":\"bytes[]\"}],\"name\":\"bulkStoreRecordBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getBytesLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"decodeRecordBytes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hashId\",\"type\":\"uint256\"}],\"name\":\"decodeRecordBytes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getHashIdsByAddress\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getRecordsByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"c\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"m\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"}],\"internalType\":\"structBlockStorage.HashRecord[]\",\"name\":\"recs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockStorageMetaData.ABI instead.
var BlockStorageABI = BlockStorageMetaData.ABI

// BlockStorage is an auto generated Go binding around an Ethereum contract.
type BlockStorage struct {
	BlockStorageCaller     // Read-only binding to the contract
	BlockStorageTransactor // Write-only binding to the contract
	BlockStorageFilterer   // Log filterer for contract events
}

// BlockStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockStorageSession struct {
	Contract     *BlockStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockStorageCallerSession struct {
	Contract *BlockStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlockStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockStorageTransactorSession struct {
	Contract     *BlockStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlockStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockStorageRaw struct {
	Contract *BlockStorage // Generic contract binding to access the raw methods on
}

// BlockStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockStorageCallerRaw struct {
	Contract *BlockStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BlockStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockStorageTransactorRaw struct {
	Contract *BlockStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockStorage creates a new instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorage(address common.Address, backend bind.ContractBackend) (*BlockStorage, error) {
	contract, err := bindBlockStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockStorage{BlockStorageCaller: BlockStorageCaller{contract: contract}, BlockStorageTransactor: BlockStorageTransactor{contract: contract}, BlockStorageFilterer: BlockStorageFilterer{contract: contract}}, nil
}

// NewBlockStorageCaller creates a new read-only instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageCaller(address common.Address, caller bind.ContractCaller) (*BlockStorageCaller, error) {
	contract, err := bindBlockStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockStorageCaller{contract: contract}, nil
}

// NewBlockStorageTransactor creates a new write-only instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockStorageTransactor, error) {
	contract, err := bindBlockStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockStorageTransactor{contract: contract}, nil
}

// NewBlockStorageFilterer creates a new log filterer instance of BlockStorage, bound to a specific deployed contract.
func NewBlockStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockStorageFilterer, error) {
	contract, err := bindBlockStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockStorageFilterer{contract: contract}, nil
}

// bindBlockStorage binds a generic wrapper to an already deployed contract.
func bindBlockStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockStorage *BlockStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockStorage.Contract.BlockStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockStorage *BlockStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.Contract.BlockStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockStorage *BlockStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockStorage.Contract.BlockStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockStorage *BlockStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockStorage *BlockStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockStorage *BlockStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockStorage.Contract.contract.Transact(opts, method, params...)
}

// TARGETBLOCKTIMES is a free data retrieval call binding the contract method 0x0ff99d59.
//
// Solidity: function TARGET_BLOCK_TIME_S() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) TARGETBLOCKTIMES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "TARGET_BLOCK_TIME_S")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETBLOCKTIMES is a free data retrieval call binding the contract method 0x0ff99d59.
//
// Solidity: function TARGET_BLOCK_TIME_S() view returns(uint256)
func (_BlockStorage *BlockStorageSession) TARGETBLOCKTIMES() (*big.Int, error) {
	return _BlockStorage.Contract.TARGETBLOCKTIMES(&_BlockStorage.CallOpts)
}

// TARGETBLOCKTIMES is a free data retrieval call binding the contract method 0x0ff99d59.
//
// Solidity: function TARGET_BLOCK_TIME_S() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) TARGETBLOCKTIMES() (*big.Int, error) {
	return _BlockStorage.Contract.TARGETBLOCKTIMES(&_BlockStorage.CallOpts)
}

// AddressesByHashId is a free data retrieval call binding the contract method 0xa6def311.
//
// Solidity: function addressesByHashId(uint256 ) view returns(address)
func (_BlockStorage *BlockStorageCaller) AddressesByHashId(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "addressesByHashId", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesByHashId is a free data retrieval call binding the contract method 0xa6def311.
//
// Solidity: function addressesByHashId(uint256 ) view returns(address)
func (_BlockStorage *BlockStorageSession) AddressesByHashId(arg0 *big.Int) (common.Address, error) {
	return _BlockStorage.Contract.AddressesByHashId(&_BlockStorage.CallOpts, arg0)
}

// AddressesByHashId is a free data retrieval call binding the contract method 0xa6def311.
//
// Solidity: function addressesByHashId(uint256 ) view returns(address)
func (_BlockStorage *BlockStorageCallerSession) AddressesByHashId(arg0 *big.Int) (common.Address, error) {
	return _BlockStorage.Contract.AddressesByHashId(&_BlockStorage.CallOpts, arg0)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_BlockStorage *BlockStorageSession) CurrentEpoch() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentEpoch(&_BlockStorage.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) CurrentEpoch() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentEpoch(&_BlockStorage.CallOpts)
}

// CurrentEpochTs is a free data retrieval call binding the contract method 0xa9abec5b.
//
// Solidity: function currentEpochTs() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) CurrentEpochTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "currentEpochTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochTs is a free data retrieval call binding the contract method 0xa9abec5b.
//
// Solidity: function currentEpochTs() view returns(uint256)
func (_BlockStorage *BlockStorageSession) CurrentEpochTs() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentEpochTs(&_BlockStorage.CallOpts)
}

// CurrentEpochTs is a free data retrieval call binding the contract method 0xa9abec5b.
//
// Solidity: function currentEpochTs() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) CurrentEpochTs() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentEpochTs(&_BlockStorage.CallOpts)
}

// CurrentHashId is a free data retrieval call binding the contract method 0x06ffc5ea.
//
// Solidity: function currentHashId() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) CurrentHashId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "currentHashId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentHashId is a free data retrieval call binding the contract method 0x06ffc5ea.
//
// Solidity: function currentHashId() view returns(uint256)
func (_BlockStorage *BlockStorageSession) CurrentHashId() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentHashId(&_BlockStorage.CallOpts)
}

// CurrentHashId is a free data retrieval call binding the contract method 0x06ffc5ea.
//
// Solidity: function currentHashId() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) CurrentHashId() (*big.Int, error) {
	return _BlockStorage.Contract.CurrentHashId(&_BlockStorage.CallOpts)
}

// DecodeRecordBytes is a free data retrieval call binding the contract method 0x0747b9c1.
//
// Solidity: function decodeRecordBytes(address _address, uint256 _idx) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageCaller) DecodeRecordBytes(opts *bind.CallOpts, _address common.Address, _idx *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "decodeRecordBytes", _address, _idx)

	outstruct := new(struct {
		C uint8
		M uint32
		T uint8
		V uint8
		K [32]byte
		S []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.C = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.M = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.T = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.V = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.K = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DecodeRecordBytes is a free data retrieval call binding the contract method 0x0747b9c1.
//
// Solidity: function decodeRecordBytes(address _address, uint256 _idx) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageSession) DecodeRecordBytes(_address common.Address, _idx *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	return _BlockStorage.Contract.DecodeRecordBytes(&_BlockStorage.CallOpts, _address, _idx)
}

// DecodeRecordBytes is a free data retrieval call binding the contract method 0x0747b9c1.
//
// Solidity: function decodeRecordBytes(address _address, uint256 _idx) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageCallerSession) DecodeRecordBytes(_address common.Address, _idx *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	return _BlockStorage.Contract.DecodeRecordBytes(&_BlockStorage.CallOpts, _address, _idx)
}

// DecodeRecordBytes0 is a free data retrieval call binding the contract method 0xcf81e899.
//
// Solidity: function decodeRecordBytes(uint256 _hashId) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageCaller) DecodeRecordBytes0(opts *bind.CallOpts, _hashId *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "decodeRecordBytes0", _hashId)

	outstruct := new(struct {
		C uint8
		M uint32
		T uint8
		V uint8
		K [32]byte
		S []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.C = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.M = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.T = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.V = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.K = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DecodeRecordBytes0 is a free data retrieval call binding the contract method 0xcf81e899.
//
// Solidity: function decodeRecordBytes(uint256 _hashId) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageSession) DecodeRecordBytes0(_hashId *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	return _BlockStorage.Contract.DecodeRecordBytes0(&_BlockStorage.CallOpts, _hashId)
}

// DecodeRecordBytes0 is a free data retrieval call binding the contract method 0xcf81e899.
//
// Solidity: function decodeRecordBytes(uint256 _hashId) view returns(uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s)
func (_BlockStorage *BlockStorageCallerSession) DecodeRecordBytes0(_hashId *big.Int) (struct {
	C uint8
	M uint32
	T uint8
	V uint8
	K [32]byte
	S []byte
}, error) {
	return _BlockStorage.Contract.DecodeRecordBytes0(&_BlockStorage.CallOpts, _hashId)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_BlockStorage *BlockStorageSession) Difficulty() (*big.Int, error) {
	return _BlockStorage.Contract.Difficulty(&_BlockStorage.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) Difficulty() (*big.Int, error) {
	return _BlockStorage.Contract.Difficulty(&_BlockStorage.CallOpts)
}

// EpochDurationSec is a free data retrieval call binding the contract method 0x151b8f4e.
//
// Solidity: function epochDurationSec() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) EpochDurationSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "epochDurationSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDurationSec is a free data retrieval call binding the contract method 0x151b8f4e.
//
// Solidity: function epochDurationSec() view returns(uint256)
func (_BlockStorage *BlockStorageSession) EpochDurationSec() (*big.Int, error) {
	return _BlockStorage.Contract.EpochDurationSec(&_BlockStorage.CallOpts)
}

// EpochDurationSec is a free data retrieval call binding the contract method 0x151b8f4e.
//
// Solidity: function epochDurationSec() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) EpochDurationSec() (*big.Int, error) {
	return _BlockStorage.Contract.EpochDurationSec(&_BlockStorage.CallOpts)
}

// EpochHashCounter is a free data retrieval call binding the contract method 0x47efe9ab.
//
// Solidity: function epochHashCounter() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) EpochHashCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "epochHashCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochHashCounter is a free data retrieval call binding the contract method 0x47efe9ab.
//
// Solidity: function epochHashCounter() view returns(uint256)
func (_BlockStorage *BlockStorageSession) EpochHashCounter() (*big.Int, error) {
	return _BlockStorage.Contract.EpochHashCounter(&_BlockStorage.CallOpts)
}

// EpochHashCounter is a free data retrieval call binding the contract method 0x47efe9ab.
//
// Solidity: function epochHashCounter() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) EpochHashCounter() (*big.Int, error) {
	return _BlockStorage.Contract.EpochHashCounter(&_BlockStorage.CallOpts)
}

// GetBytesLen is a free data retrieval call binding the contract method 0x27b50bfa.
//
// Solidity: function getBytesLen(address _address, uint256 _index) view returns(uint256 len)
func (_BlockStorage *BlockStorageCaller) GetBytesLen(opts *bind.CallOpts, _address common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "getBytesLen", _address, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBytesLen is a free data retrieval call binding the contract method 0x27b50bfa.
//
// Solidity: function getBytesLen(address _address, uint256 _index) view returns(uint256 len)
func (_BlockStorage *BlockStorageSession) GetBytesLen(_address common.Address, _index *big.Int) (*big.Int, error) {
	return _BlockStorage.Contract.GetBytesLen(&_BlockStorage.CallOpts, _address, _index)
}

// GetBytesLen is a free data retrieval call binding the contract method 0x27b50bfa.
//
// Solidity: function getBytesLen(address _address, uint256 _index) view returns(uint256 len)
func (_BlockStorage *BlockStorageCallerSession) GetBytesLen(_address common.Address, _index *big.Int) (*big.Int, error) {
	return _BlockStorage.Contract.GetBytesLen(&_BlockStorage.CallOpts, _address, _index)
}

// GetHashIdsByAddress is a free data retrieval call binding the contract method 0x45fbbb8c.
//
// Solidity: function getHashIdsByAddress(address _address) view returns(uint256[] ids)
func (_BlockStorage *BlockStorageCaller) GetHashIdsByAddress(opts *bind.CallOpts, _address common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "getHashIdsByAddress", _address)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetHashIdsByAddress is a free data retrieval call binding the contract method 0x45fbbb8c.
//
// Solidity: function getHashIdsByAddress(address _address) view returns(uint256[] ids)
func (_BlockStorage *BlockStorageSession) GetHashIdsByAddress(_address common.Address) ([]*big.Int, error) {
	return _BlockStorage.Contract.GetHashIdsByAddress(&_BlockStorage.CallOpts, _address)
}

// GetHashIdsByAddress is a free data retrieval call binding the contract method 0x45fbbb8c.
//
// Solidity: function getHashIdsByAddress(address _address) view returns(uint256[] ids)
func (_BlockStorage *BlockStorageCallerSession) GetHashIdsByAddress(_address common.Address) ([]*big.Int, error) {
	return _BlockStorage.Contract.GetHashIdsByAddress(&_BlockStorage.CallOpts, _address)
}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _address) view returns((uint8,uint32,uint8,uint8,bytes32,bytes)[] recs)
func (_BlockStorage *BlockStorageCaller) GetRecordsByAddress(opts *bind.CallOpts, _address common.Address) ([]BlockStorageHashRecord, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "getRecordsByAddress", _address)

	if err != nil {
		return *new([]BlockStorageHashRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlockStorageHashRecord)).(*[]BlockStorageHashRecord)

	return out0, err

}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _address) view returns((uint8,uint32,uint8,uint8,bytes32,bytes)[] recs)
func (_BlockStorage *BlockStorageSession) GetRecordsByAddress(_address common.Address) ([]BlockStorageHashRecord, error) {
	return _BlockStorage.Contract.GetRecordsByAddress(&_BlockStorage.CallOpts, _address)
}

// GetRecordsByAddress is a free data retrieval call binding the contract method 0xad88bb27.
//
// Solidity: function getRecordsByAddress(address _address) view returns((uint8,uint32,uint8,uint8,bytes32,bytes)[] recs)
func (_BlockStorage *BlockStorageCallerSession) GetRecordsByAddress(_address common.Address) ([]BlockStorageHashRecord, error) {
	return _BlockStorage.Contract.GetRecordsByAddress(&_BlockStorage.CallOpts, _address)
}

// HashIdsByAddress is a free data retrieval call binding the contract method 0xa3d2216d.
//
// Solidity: function hashIdsByAddress(address , uint256 ) view returns(uint256)
func (_BlockStorage *BlockStorageCaller) HashIdsByAddress(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "hashIdsByAddress", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashIdsByAddress is a free data retrieval call binding the contract method 0xa3d2216d.
//
// Solidity: function hashIdsByAddress(address , uint256 ) view returns(uint256)
func (_BlockStorage *BlockStorageSession) HashIdsByAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BlockStorage.Contract.HashIdsByAddress(&_BlockStorage.CallOpts, arg0, arg1)
}

// HashIdsByAddress is a free data retrieval call binding the contract method 0xa3d2216d.
//
// Solidity: function hashIdsByAddress(address , uint256 ) view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) HashIdsByAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BlockStorage.Contract.HashIdsByAddress(&_BlockStorage.CallOpts, arg0, arg1)
}

// HashRecords is a free data retrieval call binding the contract method 0xdb658b0f.
//
// Solidity: function hashRecords(uint256 ) view returns(bytes)
func (_BlockStorage *BlockStorageCaller) HashRecords(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "hashRecords", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// HashRecords is a free data retrieval call binding the contract method 0xdb658b0f.
//
// Solidity: function hashRecords(uint256 ) view returns(bytes)
func (_BlockStorage *BlockStorageSession) HashRecords(arg0 *big.Int) ([]byte, error) {
	return _BlockStorage.Contract.HashRecords(&_BlockStorage.CallOpts, arg0)
}

// HashRecords is a free data retrieval call binding the contract method 0xdb658b0f.
//
// Solidity: function hashRecords(uint256 ) view returns(bytes)
func (_BlockStorage *BlockStorageCallerSession) HashRecords(arg0 *big.Int) ([]byte, error) {
	return _BlockStorage.Contract.HashRecords(&_BlockStorage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockStorage *BlockStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockStorage *BlockStorageSession) Owner() (common.Address, error) {
	return _BlockStorage.Contract.Owner(&_BlockStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockStorage *BlockStorageCallerSession) Owner() (common.Address, error) {
	return _BlockStorage.Contract.Owner(&_BlockStorage.CallOpts)
}

// RecordCount is a free data retrieval call binding the contract method 0xba63245d.
//
// Solidity: function recordCount(address ) view returns(uint256)
func (_BlockStorage *BlockStorageCaller) RecordCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "recordCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecordCount is a free data retrieval call binding the contract method 0xba63245d.
//
// Solidity: function recordCount(address ) view returns(uint256)
func (_BlockStorage *BlockStorageSession) RecordCount(arg0 common.Address) (*big.Int, error) {
	return _BlockStorage.Contract.RecordCount(&_BlockStorage.CallOpts, arg0)
}

// RecordCount is a free data retrieval call binding the contract method 0xba63245d.
//
// Solidity: function recordCount(address ) view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) RecordCount(arg0 common.Address) (*big.Int, error) {
	return _BlockStorage.Contract.RecordCount(&_BlockStorage.CallOpts, arg0)
}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_BlockStorage *BlockStorageCaller) TargetDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockStorage.contract.Call(opts, &out, "targetDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_BlockStorage *BlockStorageSession) TargetDifficulty() (*big.Int, error) {
	return _BlockStorage.Contract.TargetDifficulty(&_BlockStorage.CallOpts)
}

// TargetDifficulty is a free data retrieval call binding the contract method 0x8b2db16e.
//
// Solidity: function targetDifficulty() view returns(uint256)
func (_BlockStorage *BlockStorageCallerSession) TargetDifficulty() (*big.Int, error) {
	return _BlockStorage.Contract.TargetDifficulty(&_BlockStorage.CallOpts)
}

// BulkStoreRecordBytes is a paid mutator transaction binding the contract method 0xee5f8832.
//
// Solidity: function bulkStoreRecordBytes(address _address, uint256[] _hasIds, bytes[] bb) returns()
func (_BlockStorage *BlockStorageTransactor) BulkStoreRecordBytes(opts *bind.TransactOpts, _address common.Address, _hasIds []*big.Int, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "bulkStoreRecordBytes", _address, _hasIds, bb)
}

// BulkStoreRecordBytes is a paid mutator transaction binding the contract method 0xee5f8832.
//
// Solidity: function bulkStoreRecordBytes(address _address, uint256[] _hasIds, bytes[] bb) returns()
func (_BlockStorage *BlockStorageSession) BulkStoreRecordBytes(_address common.Address, _hasIds []*big.Int, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordBytes(&_BlockStorage.TransactOpts, _address, _hasIds, bb)
}

// BulkStoreRecordBytes is a paid mutator transaction binding the contract method 0xee5f8832.
//
// Solidity: function bulkStoreRecordBytes(address _address, uint256[] _hasIds, bytes[] bb) returns()
func (_BlockStorage *BlockStorageTransactorSession) BulkStoreRecordBytes(_address common.Address, _hasIds []*big.Int, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordBytes(&_BlockStorage.TransactOpts, _address, _hasIds, bb)
}

// BulkStoreRecordBytesInc is a paid mutator transaction binding the contract method 0xcb48c1b6.
//
// Solidity: function bulkStoreRecordBytesInc(address _address, bytes[] bb) returns()
func (_BlockStorage *BlockStorageTransactor) BulkStoreRecordBytesInc(opts *bind.TransactOpts, _address common.Address, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "bulkStoreRecordBytesInc", _address, bb)
}

// BulkStoreRecordBytesInc is a paid mutator transaction binding the contract method 0xcb48c1b6.
//
// Solidity: function bulkStoreRecordBytesInc(address _address, bytes[] bb) returns()
func (_BlockStorage *BlockStorageSession) BulkStoreRecordBytesInc(_address common.Address, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordBytesInc(&_BlockStorage.TransactOpts, _address, bb)
}

// BulkStoreRecordBytesInc is a paid mutator transaction binding the contract method 0xcb48c1b6.
//
// Solidity: function bulkStoreRecordBytesInc(address _address, bytes[] bb) returns()
func (_BlockStorage *BlockStorageTransactorSession) BulkStoreRecordBytesInc(_address common.Address, bb [][]byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordBytesInc(&_BlockStorage.TransactOpts, _address, bb)
}

// BulkStoreRecords is a paid mutator transaction binding the contract method 0x2c292d7a.
//
// Solidity: function bulkStoreRecords(address _address, uint256[] _hasIds, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageTransactor) BulkStoreRecords(opts *bind.TransactOpts, _address common.Address, _hasIds []*big.Int, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "bulkStoreRecords", _address, _hasIds, hh)
}

// BulkStoreRecords is a paid mutator transaction binding the contract method 0x2c292d7a.
//
// Solidity: function bulkStoreRecords(address _address, uint256[] _hasIds, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageSession) BulkStoreRecords(_address common.Address, _hasIds []*big.Int, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecords(&_BlockStorage.TransactOpts, _address, _hasIds, hh)
}

// BulkStoreRecords is a paid mutator transaction binding the contract method 0x2c292d7a.
//
// Solidity: function bulkStoreRecords(address _address, uint256[] _hasIds, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageTransactorSession) BulkStoreRecords(_address common.Address, _hasIds []*big.Int, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecords(&_BlockStorage.TransactOpts, _address, _hasIds, hh)
}

// BulkStoreRecordsInc is a paid mutator transaction binding the contract method 0xc78d90d4.
//
// Solidity: function bulkStoreRecordsInc(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageTransactor) BulkStoreRecordsInc(opts *bind.TransactOpts, _address common.Address, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "bulkStoreRecordsInc", _address, hh)
}

// BulkStoreRecordsInc is a paid mutator transaction binding the contract method 0xc78d90d4.
//
// Solidity: function bulkStoreRecordsInc(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageSession) BulkStoreRecordsInc(_address common.Address, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordsInc(&_BlockStorage.TransactOpts, _address, hh)
}

// BulkStoreRecordsInc is a paid mutator transaction binding the contract method 0xc78d90d4.
//
// Solidity: function bulkStoreRecordsInc(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes)[] hh) returns()
func (_BlockStorage *BlockStorageTransactorSession) BulkStoreRecordsInc(_address common.Address, hh []BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.BulkStoreRecordsInc(&_BlockStorage.TransactOpts, _address, hh)
}

// DoHousekeeping is a paid mutator transaction binding the contract method 0x85ae2fb7.
//
// Solidity: function doHousekeeping() returns()
func (_BlockStorage *BlockStorageTransactor) DoHousekeeping(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "doHousekeeping")
}

// DoHousekeeping is a paid mutator transaction binding the contract method 0x85ae2fb7.
//
// Solidity: function doHousekeeping() returns()
func (_BlockStorage *BlockStorageSession) DoHousekeeping() (*types.Transaction, error) {
	return _BlockStorage.Contract.DoHousekeeping(&_BlockStorage.TransactOpts)
}

// DoHousekeeping is a paid mutator transaction binding the contract method 0x85ae2fb7.
//
// Solidity: function doHousekeeping() returns()
func (_BlockStorage *BlockStorageTransactorSession) DoHousekeeping() (*types.Transaction, error) {
	return _BlockStorage.Contract.DoHousekeeping(&_BlockStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockStorage *BlockStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockStorage *BlockStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockStorage.Contract.RenounceOwnership(&_BlockStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockStorage *BlockStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockStorage.Contract.RenounceOwnership(&_BlockStorage.TransactOpts)
}

// SetEpochDurationSec is a paid mutator transaction binding the contract method 0x799f6335.
//
// Solidity: function setEpochDurationSec(uint256 _epochDurationSec) returns()
func (_BlockStorage *BlockStorageTransactor) SetEpochDurationSec(opts *bind.TransactOpts, _epochDurationSec *big.Int) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "setEpochDurationSec", _epochDurationSec)
}

// SetEpochDurationSec is a paid mutator transaction binding the contract method 0x799f6335.
//
// Solidity: function setEpochDurationSec(uint256 _epochDurationSec) returns()
func (_BlockStorage *BlockStorageSession) SetEpochDurationSec(_epochDurationSec *big.Int) (*types.Transaction, error) {
	return _BlockStorage.Contract.SetEpochDurationSec(&_BlockStorage.TransactOpts, _epochDurationSec)
}

// SetEpochDurationSec is a paid mutator transaction binding the contract method 0x799f6335.
//
// Solidity: function setEpochDurationSec(uint256 _epochDurationSec) returns()
func (_BlockStorage *BlockStorageTransactorSession) SetEpochDurationSec(_epochDurationSec *big.Int) (*types.Transaction, error) {
	return _BlockStorage.Contract.SetEpochDurationSec(&_BlockStorage.TransactOpts, _epochDurationSec)
}

// SetTargetDifficulty is a paid mutator transaction binding the contract method 0x74811f6f.
//
// Solidity: function setTargetDifficulty(uint256 _targetDifficulty) returns()
func (_BlockStorage *BlockStorageTransactor) SetTargetDifficulty(opts *bind.TransactOpts, _targetDifficulty *big.Int) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "setTargetDifficulty", _targetDifficulty)
}

// SetTargetDifficulty is a paid mutator transaction binding the contract method 0x74811f6f.
//
// Solidity: function setTargetDifficulty(uint256 _targetDifficulty) returns()
func (_BlockStorage *BlockStorageSession) SetTargetDifficulty(_targetDifficulty *big.Int) (*types.Transaction, error) {
	return _BlockStorage.Contract.SetTargetDifficulty(&_BlockStorage.TransactOpts, _targetDifficulty)
}

// SetTargetDifficulty is a paid mutator transaction binding the contract method 0x74811f6f.
//
// Solidity: function setTargetDifficulty(uint256 _targetDifficulty) returns()
func (_BlockStorage *BlockStorageTransactorSession) SetTargetDifficulty(_targetDifficulty *big.Int) (*types.Transaction, error) {
	return _BlockStorage.Contract.SetTargetDifficulty(&_BlockStorage.TransactOpts, _targetDifficulty)
}

// StoreNewRecord is a paid mutator transaction binding the contract method 0xe1616531.
//
// Solidity: function storeNewRecord(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes) r) returns()
func (_BlockStorage *BlockStorageTransactor) StoreNewRecord(opts *bind.TransactOpts, _address common.Address, r BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "storeNewRecord", _address, r)
}

// StoreNewRecord is a paid mutator transaction binding the contract method 0xe1616531.
//
// Solidity: function storeNewRecord(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes) r) returns()
func (_BlockStorage *BlockStorageSession) StoreNewRecord(_address common.Address, r BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecord(&_BlockStorage.TransactOpts, _address, r)
}

// StoreNewRecord is a paid mutator transaction binding the contract method 0xe1616531.
//
// Solidity: function storeNewRecord(address _address, (uint8,uint32,uint8,uint8,bytes32,bytes) r) returns()
func (_BlockStorage *BlockStorageTransactorSession) StoreNewRecord(_address common.Address, r BlockStorageHashRecord) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecord(&_BlockStorage.TransactOpts, _address, r)
}

// StoreNewRecordBytes is a paid mutator transaction binding the contract method 0xa5cdb7e5.
//
// Solidity: function storeNewRecordBytes(address _address, bytes _bytes) returns()
func (_BlockStorage *BlockStorageTransactor) StoreNewRecordBytes(opts *bind.TransactOpts, _address common.Address, _bytes []byte) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "storeNewRecordBytes", _address, _bytes)
}

// StoreNewRecordBytes is a paid mutator transaction binding the contract method 0xa5cdb7e5.
//
// Solidity: function storeNewRecordBytes(address _address, bytes _bytes) returns()
func (_BlockStorage *BlockStorageSession) StoreNewRecordBytes(_address common.Address, _bytes []byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecordBytes(&_BlockStorage.TransactOpts, _address, _bytes)
}

// StoreNewRecordBytes is a paid mutator transaction binding the contract method 0xa5cdb7e5.
//
// Solidity: function storeNewRecordBytes(address _address, bytes _bytes) returns()
func (_BlockStorage *BlockStorageTransactorSession) StoreNewRecordBytes(_address common.Address, _bytes []byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecordBytes(&_BlockStorage.TransactOpts, _address, _bytes)
}

// StoreNewRecordData is a paid mutator transaction binding the contract method 0x0e86b75e.
//
// Solidity: function storeNewRecordData(address _address, uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s) returns()
func (_BlockStorage *BlockStorageTransactor) StoreNewRecordData(opts *bind.TransactOpts, _address common.Address, c uint8, m uint32, t uint8, v uint8, k [32]byte, s []byte) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "storeNewRecordData", _address, c, m, t, v, k, s)
}

// StoreNewRecordData is a paid mutator transaction binding the contract method 0x0e86b75e.
//
// Solidity: function storeNewRecordData(address _address, uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s) returns()
func (_BlockStorage *BlockStorageSession) StoreNewRecordData(_address common.Address, c uint8, m uint32, t uint8, v uint8, k [32]byte, s []byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecordData(&_BlockStorage.TransactOpts, _address, c, m, t, v, k, s)
}

// StoreNewRecordData is a paid mutator transaction binding the contract method 0x0e86b75e.
//
// Solidity: function storeNewRecordData(address _address, uint8 c, uint32 m, uint8 t, uint8 v, bytes32 k, bytes s) returns()
func (_BlockStorage *BlockStorageTransactorSession) StoreNewRecordData(_address common.Address, c uint8, m uint32, t uint8, v uint8, k [32]byte, s []byte) (*types.Transaction, error) {
	return _BlockStorage.Contract.StoreNewRecordData(&_BlockStorage.TransactOpts, _address, c, m, t, v, k, s)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockStorage *BlockStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockStorage *BlockStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockStorage.Contract.TransferOwnership(&_BlockStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockStorage *BlockStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockStorage.Contract.TransferOwnership(&_BlockStorage.TransactOpts, newOwner)
}

// BlockStorageEpochDurationSetIterator is returned from FilterEpochDurationSet and is used to iterate over the raw logs and unpacked data for EpochDurationSet events raised by the BlockStorage contract.
type BlockStorageEpochDurationSetIterator struct {
	Event *BlockStorageEpochDurationSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockStorageEpochDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockStorageEpochDurationSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockStorageEpochDurationSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockStorageEpochDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockStorageEpochDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockStorageEpochDurationSet represents a EpochDurationSet event raised by the BlockStorage contract.
type BlockStorageEpochDurationSet struct {
	Epoch    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEpochDurationSet is a free log retrieval operation binding the contract event 0x361d876894ff3854d2e192162a59a0ce8d9e27637af1302a2b8af18e3b007b24.
//
// Solidity: event EpochDurationSet(uint256 indexed epoch, uint256 indexed duration)
func (_BlockStorage *BlockStorageFilterer) FilterEpochDurationSet(opts *bind.FilterOpts, epoch []*big.Int, duration []*big.Int) (*BlockStorageEpochDurationSetIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _BlockStorage.contract.FilterLogs(opts, "EpochDurationSet", epochRule, durationRule)
	if err != nil {
		return nil, err
	}
	return &BlockStorageEpochDurationSetIterator{contract: _BlockStorage.contract, event: "EpochDurationSet", logs: logs, sub: sub}, nil
}

// WatchEpochDurationSet is a free log subscription operation binding the contract event 0x361d876894ff3854d2e192162a59a0ce8d9e27637af1302a2b8af18e3b007b24.
//
// Solidity: event EpochDurationSet(uint256 indexed epoch, uint256 indexed duration)
func (_BlockStorage *BlockStorageFilterer) WatchEpochDurationSet(opts *bind.WatchOpts, sink chan<- *BlockStorageEpochDurationSet, epoch []*big.Int, duration []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _BlockStorage.contract.WatchLogs(opts, "EpochDurationSet", epochRule, durationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockStorageEpochDurationSet)
				if err := _BlockStorage.contract.UnpackLog(event, "EpochDurationSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochDurationSet is a log parse operation binding the contract event 0x361d876894ff3854d2e192162a59a0ce8d9e27637af1302a2b8af18e3b007b24.
//
// Solidity: event EpochDurationSet(uint256 indexed epoch, uint256 indexed duration)
func (_BlockStorage *BlockStorageFilterer) ParseEpochDurationSet(log types.Log) (*BlockStorageEpochDurationSet, error) {
	event := new(BlockStorageEpochDurationSet)
	if err := _BlockStorage.contract.UnpackLog(event, "EpochDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockStorageNewEpochIterator is returned from FilterNewEpoch and is used to iterate over the raw logs and unpacked data for NewEpoch events raised by the BlockStorage contract.
type BlockStorageNewEpochIterator struct {
	Event *BlockStorageNewEpoch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockStorageNewEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockStorageNewEpoch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockStorageNewEpoch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockStorageNewEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockStorageNewEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockStorageNewEpoch represents a NewEpoch event raised by the BlockStorage contract.
type BlockStorageNewEpoch struct {
	Epoch      *big.Int
	Difficulty *big.Int
	Ts         *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewEpoch is a free log retrieval operation binding the contract event 0x3bb7b347508b7c148ec2094ac60d2e3d8b7595421025643f08b45cb78b326b58.
//
// Solidity: event NewEpoch(uint256 indexed epoch, uint256 difficulty, uint256 ts)
func (_BlockStorage *BlockStorageFilterer) FilterNewEpoch(opts *bind.FilterOpts, epoch []*big.Int) (*BlockStorageNewEpochIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BlockStorage.contract.FilterLogs(opts, "NewEpoch", epochRule)
	if err != nil {
		return nil, err
	}
	return &BlockStorageNewEpochIterator{contract: _BlockStorage.contract, event: "NewEpoch", logs: logs, sub: sub}, nil
}

// WatchNewEpoch is a free log subscription operation binding the contract event 0x3bb7b347508b7c148ec2094ac60d2e3d8b7595421025643f08b45cb78b326b58.
//
// Solidity: event NewEpoch(uint256 indexed epoch, uint256 difficulty, uint256 ts)
func (_BlockStorage *BlockStorageFilterer) WatchNewEpoch(opts *bind.WatchOpts, sink chan<- *BlockStorageNewEpoch, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _BlockStorage.contract.WatchLogs(opts, "NewEpoch", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockStorageNewEpoch)
				if err := _BlockStorage.contract.UnpackLog(event, "NewEpoch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewEpoch is a log parse operation binding the contract event 0x3bb7b347508b7c148ec2094ac60d2e3d8b7595421025643f08b45cb78b326b58.
//
// Solidity: event NewEpoch(uint256 indexed epoch, uint256 difficulty, uint256 ts)
func (_BlockStorage *BlockStorageFilterer) ParseNewEpoch(log types.Log) (*BlockStorageNewEpoch, error) {
	event := new(BlockStorageNewEpoch)
	if err := _BlockStorage.contract.UnpackLog(event, "NewEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockStorageNewHashIterator is returned from FilterNewHash and is used to iterate over the raw logs and unpacked data for NewHash events raised by the BlockStorage contract.
type BlockStorageNewHashIterator struct {
	Event *BlockStorageNewHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockStorageNewHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockStorageNewHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockStorageNewHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockStorageNewHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockStorageNewHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockStorageNewHash represents a NewHash event raised by the BlockStorage contract.
type BlockStorageNewHash struct {
	HashId  *big.Int
	Epoch   *big.Int
	Account common.Address
	Bytes   []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewHash is a free log retrieval operation binding the contract event 0x30e5fb28e00aab05ae396f77b088b8627fa5f78c42fbd55fdb593caa2e3f2a2e.
//
// Solidity: event NewHash(uint256 indexed hashId, uint256 indexed epoch, address indexed account, bytes _bytes)
func (_BlockStorage *BlockStorageFilterer) FilterNewHash(opts *bind.FilterOpts, hashId []*big.Int, epoch []*big.Int, account []common.Address) (*BlockStorageNewHashIterator, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BlockStorage.contract.FilterLogs(opts, "NewHash", hashIdRule, epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BlockStorageNewHashIterator{contract: _BlockStorage.contract, event: "NewHash", logs: logs, sub: sub}, nil
}

// WatchNewHash is a free log subscription operation binding the contract event 0x30e5fb28e00aab05ae396f77b088b8627fa5f78c42fbd55fdb593caa2e3f2a2e.
//
// Solidity: event NewHash(uint256 indexed hashId, uint256 indexed epoch, address indexed account, bytes _bytes)
func (_BlockStorage *BlockStorageFilterer) WatchNewHash(opts *bind.WatchOpts, sink chan<- *BlockStorageNewHash, hashId []*big.Int, epoch []*big.Int, account []common.Address) (event.Subscription, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BlockStorage.contract.WatchLogs(opts, "NewHash", hashIdRule, epochRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockStorageNewHash)
				if err := _BlockStorage.contract.UnpackLog(event, "NewHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewHash is a log parse operation binding the contract event 0x30e5fb28e00aab05ae396f77b088b8627fa5f78c42fbd55fdb593caa2e3f2a2e.
//
// Solidity: event NewHash(uint256 indexed hashId, uint256 indexed epoch, address indexed account, bytes _bytes)
func (_BlockStorage *BlockStorageFilterer) ParseNewHash(log types.Log) (*BlockStorageNewHash, error) {
	event := new(BlockStorageNewHash)
	if err := _BlockStorage.contract.UnpackLog(event, "NewHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockStorage contract.
type BlockStorageOwnershipTransferredIterator struct {
	Event *BlockStorageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockStorageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockStorageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockStorageOwnershipTransferred represents a OwnershipTransferred event raised by the BlockStorage contract.
type BlockStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockStorage *BlockStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockStorageOwnershipTransferredIterator{contract: _BlockStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockStorage *BlockStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockStorageOwnershipTransferred)
				if err := _BlockStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockStorage *BlockStorageFilterer) ParseOwnershipTransferred(log types.Log) (*BlockStorageOwnershipTransferred, error) {
	event := new(BlockStorageOwnershipTransferred)
	if err := _BlockStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockStorageTargetDifficultySetIterator is returned from FilterTargetDifficultySet and is used to iterate over the raw logs and unpacked data for TargetDifficultySet events raised by the BlockStorage contract.
type BlockStorageTargetDifficultySetIterator struct {
	Event *BlockStorageTargetDifficultySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BlockStorageTargetDifficultySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockStorageTargetDifficultySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BlockStorageTargetDifficultySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BlockStorageTargetDifficultySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockStorageTargetDifficultySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockStorageTargetDifficultySet represents a TargetDifficultySet event raised by the BlockStorage contract.
type BlockStorageTargetDifficultySet struct {
	Epoch      *big.Int
	Difficulty *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTargetDifficultySet is a free log retrieval operation binding the contract event 0xa0c1b6fbc0703071b511f72bd3575b2013f2af47bf123f0d84602f476488799c.
//
// Solidity: event TargetDifficultySet(uint256 indexed epoch, uint256 indexed difficulty)
func (_BlockStorage *BlockStorageFilterer) FilterTargetDifficultySet(opts *bind.FilterOpts, epoch []*big.Int, difficulty []*big.Int) (*BlockStorageTargetDifficultySetIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var difficultyRule []interface{}
	for _, difficultyItem := range difficulty {
		difficultyRule = append(difficultyRule, difficultyItem)
	}

	logs, sub, err := _BlockStorage.contract.FilterLogs(opts, "TargetDifficultySet", epochRule, difficultyRule)
	if err != nil {
		return nil, err
	}
	return &BlockStorageTargetDifficultySetIterator{contract: _BlockStorage.contract, event: "TargetDifficultySet", logs: logs, sub: sub}, nil
}

// WatchTargetDifficultySet is a free log subscription operation binding the contract event 0xa0c1b6fbc0703071b511f72bd3575b2013f2af47bf123f0d84602f476488799c.
//
// Solidity: event TargetDifficultySet(uint256 indexed epoch, uint256 indexed difficulty)
func (_BlockStorage *BlockStorageFilterer) WatchTargetDifficultySet(opts *bind.WatchOpts, sink chan<- *BlockStorageTargetDifficultySet, epoch []*big.Int, difficulty []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}
	var difficultyRule []interface{}
	for _, difficultyItem := range difficulty {
		difficultyRule = append(difficultyRule, difficultyItem)
	}

	logs, sub, err := _BlockStorage.contract.WatchLogs(opts, "TargetDifficultySet", epochRule, difficultyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockStorageTargetDifficultySet)
				if err := _BlockStorage.contract.UnpackLog(event, "TargetDifficultySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTargetDifficultySet is a log parse operation binding the contract event 0xa0c1b6fbc0703071b511f72bd3575b2013f2af47bf123f0d84602f476488799c.
//
// Solidity: event TargetDifficultySet(uint256 indexed epoch, uint256 indexed difficulty)
func (_BlockStorage *BlockStorageFilterer) ParseTargetDifficultySet(log types.Log) (*BlockStorageTargetDifficultySet, error) {
	event := new(BlockStorageTargetDifficultySet)
	if err := _BlockStorage.contract.UnpackLog(event, "TargetDifficultySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
