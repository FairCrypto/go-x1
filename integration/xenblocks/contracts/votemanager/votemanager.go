// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package votemanager

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
)

// VoteManagerVotePayload is an auto generated low-level Go binding around an user-defined struct.
type VoteManagerVotePayload struct {
	HashId            *big.Int
	CurrencyType      *big.Int
	MintedBlockNumber *big.Int
	Version           uint16
}

// VotemanagerMetaData contains all meta data concerning the Votemanager contract.
var VotemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hashId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"}],\"name\":\"MintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"hashId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"votes\",\"type\":\"uint16\"}],\"name\":\"VoteToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockCheckpointByValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockStorage\",\"outputs\":[{\"internalType\":\"contractBlockStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialSfcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialBlockStorageAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"initialVotePercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"initialVoteBufferPercent\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mintedByHashIdAndCurrencyType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredNumOfValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredNumOfVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sfc\",\"outputs\":[{\"internalType\":\"contractSFC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hashId\",\"type\":\"uint256\"}],\"name\":\"shouldVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractTokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockStorageAddress\",\"type\":\"address\"}],\"name\":\"updateBlockStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"updateCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sfcAddress\",\"type\":\"address\"}],\"name\":\"updateSfcAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenRegistryAddress_\",\"type\":\"address\"}],\"name\":\"updateTokenRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"voteBufferPercent_\",\"type\":\"uint8\"}],\"name\":\"updateVoteBufferPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"votePercentage_\",\"type\":\"uint8\"}],\"name\":\"updateVotePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"hashId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"internalType\":\"structVoteManager.VotePayload[]\",\"name\":\"payload\",\"type\":\"tuple[]\"}],\"name\":\"voteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteBufferPercent\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votePercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votesByHashIdAndCurrencyType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votesByHashIdAndValidatorIdAndCurrencyType\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VotemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VotemanagerMetaData.ABI instead.
var VotemanagerABI = VotemanagerMetaData.ABI

// Votemanager is an auto generated Go binding around an Ethereum contract.
type Votemanager struct {
	VotemanagerCaller     // Read-only binding to the contract
	VotemanagerTransactor // Write-only binding to the contract
	VotemanagerFilterer   // Log filterer for contract events
}

// VotemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotemanagerSession struct {
	Contract     *Votemanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotemanagerCallerSession struct {
	Contract *VotemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VotemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotemanagerTransactorSession struct {
	Contract     *VotemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VotemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotemanagerRaw struct {
	Contract *Votemanager // Generic contract binding to access the raw methods on
}

// VotemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotemanagerCallerRaw struct {
	Contract *VotemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// VotemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotemanagerTransactorRaw struct {
	Contract *VotemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotemanager creates a new instance of Votemanager, bound to a specific deployed contract.
func NewVotemanager(address common.Address, backend bind.ContractBackend) (*Votemanager, error) {
	contract, err := bindVotemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Votemanager{VotemanagerCaller: VotemanagerCaller{contract: contract}, VotemanagerTransactor: VotemanagerTransactor{contract: contract}, VotemanagerFilterer: VotemanagerFilterer{contract: contract}}, nil
}

// NewVotemanagerCaller creates a new read-only instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerCaller(address common.Address, caller bind.ContractCaller) (*VotemanagerCaller, error) {
	contract, err := bindVotemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotemanagerCaller{contract: contract}, nil
}

// NewVotemanagerTransactor creates a new write-only instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VotemanagerTransactor, error) {
	contract, err := bindVotemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotemanagerTransactor{contract: contract}, nil
}

// NewVotemanagerFilterer creates a new log filterer instance of Votemanager, bound to a specific deployed contract.
func NewVotemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VotemanagerFilterer, error) {
	contract, err := bindVotemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotemanagerFilterer{contract: contract}, nil
}

// bindVotemanager binds a generic wrapper to an already deployed contract.
func bindVotemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votemanager *VotemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votemanager.Contract.VotemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votemanager *VotemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.Contract.VotemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votemanager *VotemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votemanager.Contract.VotemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Votemanager *VotemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Votemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Votemanager *VotemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Votemanager *VotemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Votemanager.Contract.contract.Transact(opts, method, params...)
}

// BlockCheckpointByValidatorId is a free data retrieval call binding the contract method 0x2b3b2c0a.
//
// Solidity: function blockCheckpointByValidatorId(uint256 ) view returns(uint256)
func (_Votemanager *VotemanagerCaller) BlockCheckpointByValidatorId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "blockCheckpointByValidatorId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockCheckpointByValidatorId is a free data retrieval call binding the contract method 0x2b3b2c0a.
//
// Solidity: function blockCheckpointByValidatorId(uint256 ) view returns(uint256)
func (_Votemanager *VotemanagerSession) BlockCheckpointByValidatorId(arg0 *big.Int) (*big.Int, error) {
	return _Votemanager.Contract.BlockCheckpointByValidatorId(&_Votemanager.CallOpts, arg0)
}

// BlockCheckpointByValidatorId is a free data retrieval call binding the contract method 0x2b3b2c0a.
//
// Solidity: function blockCheckpointByValidatorId(uint256 ) view returns(uint256)
func (_Votemanager *VotemanagerCallerSession) BlockCheckpointByValidatorId(arg0 *big.Int) (*big.Int, error) {
	return _Votemanager.Contract.BlockCheckpointByValidatorId(&_Votemanager.CallOpts, arg0)
}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() view returns(address)
func (_Votemanager *VotemanagerCaller) BlockStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "blockStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() view returns(address)
func (_Votemanager *VotemanagerSession) BlockStorage() (common.Address, error) {
	return _Votemanager.Contract.BlockStorage(&_Votemanager.CallOpts)
}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() view returns(address)
func (_Votemanager *VotemanagerCallerSession) BlockStorage() (common.Address, error) {
	return _Votemanager.Contract.BlockStorage(&_Votemanager.CallOpts)
}

// MintedByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x47ce239d.
//
// Solidity: function mintedByHashIdAndCurrencyType(uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerCaller) MintedByHashIdAndCurrencyType(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "mintedByHashIdAndCurrencyType", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintedByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x47ce239d.
//
// Solidity: function mintedByHashIdAndCurrencyType(uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerSession) MintedByHashIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Votemanager.Contract.MintedByHashIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1)
}

// MintedByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x47ce239d.
//
// Solidity: function mintedByHashIdAndCurrencyType(uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) MintedByHashIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _Votemanager.Contract.MintedByHashIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerSession) Owner() (common.Address, error) {
	return _Votemanager.Contract.Owner(&_Votemanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Votemanager *VotemanagerCallerSession) Owner() (common.Address, error) {
	return _Votemanager.Contract.Owner(&_Votemanager.CallOpts)
}

// RequiredNumOfValidators is a free data retrieval call binding the contract method 0x5370c435.
//
// Solidity: function requiredNumOfValidators() view returns(uint256)
func (_Votemanager *VotemanagerCaller) RequiredNumOfValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "requiredNumOfValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredNumOfValidators is a free data retrieval call binding the contract method 0x5370c435.
//
// Solidity: function requiredNumOfValidators() view returns(uint256)
func (_Votemanager *VotemanagerSession) RequiredNumOfValidators() (*big.Int, error) {
	return _Votemanager.Contract.RequiredNumOfValidators(&_Votemanager.CallOpts)
}

// RequiredNumOfValidators is a free data retrieval call binding the contract method 0x5370c435.
//
// Solidity: function requiredNumOfValidators() view returns(uint256)
func (_Votemanager *VotemanagerCallerSession) RequiredNumOfValidators() (*big.Int, error) {
	return _Votemanager.Contract.RequiredNumOfValidators(&_Votemanager.CallOpts)
}

// RequiredNumOfVotes is a free data retrieval call binding the contract method 0x27b26533.
//
// Solidity: function requiredNumOfVotes() view returns(uint256)
func (_Votemanager *VotemanagerCaller) RequiredNumOfVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "requiredNumOfVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredNumOfVotes is a free data retrieval call binding the contract method 0x27b26533.
//
// Solidity: function requiredNumOfVotes() view returns(uint256)
func (_Votemanager *VotemanagerSession) RequiredNumOfVotes() (*big.Int, error) {
	return _Votemanager.Contract.RequiredNumOfVotes(&_Votemanager.CallOpts)
}

// RequiredNumOfVotes is a free data retrieval call binding the contract method 0x27b26533.
//
// Solidity: function requiredNumOfVotes() view returns(uint256)
func (_Votemanager *VotemanagerCallerSession) RequiredNumOfVotes() (*big.Int, error) {
	return _Votemanager.Contract.RequiredNumOfVotes(&_Votemanager.CallOpts)
}

// Sfc is a free data retrieval call binding the contract method 0x8992229f.
//
// Solidity: function sfc() view returns(address)
func (_Votemanager *VotemanagerCaller) Sfc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "sfc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sfc is a free data retrieval call binding the contract method 0x8992229f.
//
// Solidity: function sfc() view returns(address)
func (_Votemanager *VotemanagerSession) Sfc() (common.Address, error) {
	return _Votemanager.Contract.Sfc(&_Votemanager.CallOpts)
}

// Sfc is a free data retrieval call binding the contract method 0x8992229f.
//
// Solidity: function sfc() view returns(address)
func (_Votemanager *VotemanagerCallerSession) Sfc() (common.Address, error) {
	return _Votemanager.Contract.Sfc(&_Votemanager.CallOpts)
}

// ShouldVote is a free data retrieval call binding the contract method 0xb1faefbb.
//
// Solidity: function shouldVote(uint256 mintedBlockNumber, uint256 validatorId, uint256 hashId) view returns(bool)
func (_Votemanager *VotemanagerCaller) ShouldVote(opts *bind.CallOpts, mintedBlockNumber *big.Int, validatorId *big.Int, hashId *big.Int) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "shouldVote", mintedBlockNumber, validatorId, hashId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldVote is a free data retrieval call binding the contract method 0xb1faefbb.
//
// Solidity: function shouldVote(uint256 mintedBlockNumber, uint256 validatorId, uint256 hashId) view returns(bool)
func (_Votemanager *VotemanagerSession) ShouldVote(mintedBlockNumber *big.Int, validatorId *big.Int, hashId *big.Int) (bool, error) {
	return _Votemanager.Contract.ShouldVote(&_Votemanager.CallOpts, mintedBlockNumber, validatorId, hashId)
}

// ShouldVote is a free data retrieval call binding the contract method 0xb1faefbb.
//
// Solidity: function shouldVote(uint256 mintedBlockNumber, uint256 validatorId, uint256 hashId) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) ShouldVote(mintedBlockNumber *big.Int, validatorId *big.Int, hashId *big.Int) (bool, error) {
	return _Votemanager.Contract.ShouldVote(&_Votemanager.CallOpts, mintedBlockNumber, validatorId, hashId)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Votemanager *VotemanagerCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "tokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Votemanager *VotemanagerSession) TokenRegistry() (common.Address, error) {
	return _Votemanager.Contract.TokenRegistry(&_Votemanager.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Votemanager *VotemanagerCallerSession) TokenRegistry() (common.Address, error) {
	return _Votemanager.Contract.TokenRegistry(&_Votemanager.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Votemanager *VotemanagerCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Votemanager *VotemanagerSession) ValidatorCount() (*big.Int, error) {
	return _Votemanager.Contract.ValidatorCount(&_Votemanager.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Votemanager *VotemanagerCallerSession) ValidatorCount() (*big.Int, error) {
	return _Votemanager.Contract.ValidatorCount(&_Votemanager.CallOpts)
}

// VoteBufferPercent is a free data retrieval call binding the contract method 0xa2379199.
//
// Solidity: function voteBufferPercent() view returns(uint8)
func (_Votemanager *VotemanagerCaller) VoteBufferPercent(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "voteBufferPercent")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VoteBufferPercent is a free data retrieval call binding the contract method 0xa2379199.
//
// Solidity: function voteBufferPercent() view returns(uint8)
func (_Votemanager *VotemanagerSession) VoteBufferPercent() (uint8, error) {
	return _Votemanager.Contract.VoteBufferPercent(&_Votemanager.CallOpts)
}

// VoteBufferPercent is a free data retrieval call binding the contract method 0xa2379199.
//
// Solidity: function voteBufferPercent() view returns(uint8)
func (_Votemanager *VotemanagerCallerSession) VoteBufferPercent() (uint8, error) {
	return _Votemanager.Contract.VoteBufferPercent(&_Votemanager.CallOpts)
}

// VotePercentage is a free data retrieval call binding the contract method 0x34b502c0.
//
// Solidity: function votePercentage() view returns(uint8)
func (_Votemanager *VotemanagerCaller) VotePercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "votePercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VotePercentage is a free data retrieval call binding the contract method 0x34b502c0.
//
// Solidity: function votePercentage() view returns(uint8)
func (_Votemanager *VotemanagerSession) VotePercentage() (uint8, error) {
	return _Votemanager.Contract.VotePercentage(&_Votemanager.CallOpts)
}

// VotePercentage is a free data retrieval call binding the contract method 0x34b502c0.
//
// Solidity: function votePercentage() view returns(uint8)
func (_Votemanager *VotemanagerCallerSession) VotePercentage() (uint8, error) {
	return _Votemanager.Contract.VotePercentage(&_Votemanager.CallOpts)
}

// VotesByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x348fbd64.
//
// Solidity: function votesByHashIdAndCurrencyType(uint256 , uint256 ) view returns(uint16)
func (_Votemanager *VotemanagerCaller) VotesByHashIdAndCurrencyType(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "votesByHashIdAndCurrencyType", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// VotesByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x348fbd64.
//
// Solidity: function votesByHashIdAndCurrencyType(uint256 , uint256 ) view returns(uint16)
func (_Votemanager *VotemanagerSession) VotesByHashIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int) (uint16, error) {
	return _Votemanager.Contract.VotesByHashIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1)
}

// VotesByHashIdAndCurrencyType is a free data retrieval call binding the contract method 0x348fbd64.
//
// Solidity: function votesByHashIdAndCurrencyType(uint256 , uint256 ) view returns(uint16)
func (_Votemanager *VotemanagerCallerSession) VotesByHashIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int) (uint16, error) {
	return _Votemanager.Contract.VotesByHashIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1)
}

// VotesByHashIdAndValidatorIdAndCurrencyType is a free data retrieval call binding the contract method 0x95c6286c.
//
// Solidity: function votesByHashIdAndValidatorIdAndCurrencyType(uint256 , uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerCaller) VotesByHashIdAndValidatorIdAndCurrencyType(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _Votemanager.contract.Call(opts, &out, "votesByHashIdAndValidatorIdAndCurrencyType", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VotesByHashIdAndValidatorIdAndCurrencyType is a free data retrieval call binding the contract method 0x95c6286c.
//
// Solidity: function votesByHashIdAndValidatorIdAndCurrencyType(uint256 , uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerSession) VotesByHashIdAndValidatorIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Votemanager.Contract.VotesByHashIdAndValidatorIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1, arg2)
}

// VotesByHashIdAndValidatorIdAndCurrencyType is a free data retrieval call binding the contract method 0x95c6286c.
//
// Solidity: function votesByHashIdAndValidatorIdAndCurrencyType(uint256 , uint256 , uint256 ) view returns(bool)
func (_Votemanager *VotemanagerCallerSession) VotesByHashIdAndValidatorIdAndCurrencyType(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int) (bool, error) {
	return _Votemanager.Contract.VotesByHashIdAndValidatorIdAndCurrencyType(&_Votemanager.CallOpts, arg0, arg1, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0x903cd3e3.
//
// Solidity: function initialize(address initialOwner, address initialSfcAddress, address initialBlockStorageAddress, uint8 initialVotePercentage, uint8 initialVoteBufferPercent) returns()
func (_Votemanager *VotemanagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialSfcAddress common.Address, initialBlockStorageAddress common.Address, initialVotePercentage uint8, initialVoteBufferPercent uint8) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "initialize", initialOwner, initialSfcAddress, initialBlockStorageAddress, initialVotePercentage, initialVoteBufferPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x903cd3e3.
//
// Solidity: function initialize(address initialOwner, address initialSfcAddress, address initialBlockStorageAddress, uint8 initialVotePercentage, uint8 initialVoteBufferPercent) returns()
func (_Votemanager *VotemanagerSession) Initialize(initialOwner common.Address, initialSfcAddress common.Address, initialBlockStorageAddress common.Address, initialVotePercentage uint8, initialVoteBufferPercent uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.Initialize(&_Votemanager.TransactOpts, initialOwner, initialSfcAddress, initialBlockStorageAddress, initialVotePercentage, initialVoteBufferPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x903cd3e3.
//
// Solidity: function initialize(address initialOwner, address initialSfcAddress, address initialBlockStorageAddress, uint8 initialVotePercentage, uint8 initialVoteBufferPercent) returns()
func (_Votemanager *VotemanagerTransactorSession) Initialize(initialOwner common.Address, initialSfcAddress common.Address, initialBlockStorageAddress common.Address, initialVotePercentage uint8, initialVoteBufferPercent uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.Initialize(&_Votemanager.TransactOpts, initialOwner, initialSfcAddress, initialBlockStorageAddress, initialVotePercentage, initialVoteBufferPercent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Votemanager.Contract.RenounceOwnership(&_Votemanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Votemanager *VotemanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Votemanager.Contract.RenounceOwnership(&_Votemanager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.TransferOwnership(&_Votemanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Votemanager *VotemanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.TransferOwnership(&_Votemanager.TransactOpts, newOwner)
}

// UpdateBlockStorageAddress is a paid mutator transaction binding the contract method 0x3cc66f46.
//
// Solidity: function updateBlockStorageAddress(address blockStorageAddress) returns()
func (_Votemanager *VotemanagerTransactor) UpdateBlockStorageAddress(opts *bind.TransactOpts, blockStorageAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateBlockStorageAddress", blockStorageAddress)
}

// UpdateBlockStorageAddress is a paid mutator transaction binding the contract method 0x3cc66f46.
//
// Solidity: function updateBlockStorageAddress(address blockStorageAddress) returns()
func (_Votemanager *VotemanagerSession) UpdateBlockStorageAddress(blockStorageAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateBlockStorageAddress(&_Votemanager.TransactOpts, blockStorageAddress)
}

// UpdateBlockStorageAddress is a paid mutator transaction binding the contract method 0x3cc66f46.
//
// Solidity: function updateBlockStorageAddress(address blockStorageAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateBlockStorageAddress(blockStorageAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateBlockStorageAddress(&_Votemanager.TransactOpts, blockStorageAddress)
}

// UpdateCheckpoint is a paid mutator transaction binding the contract method 0x3e92aebe.
//
// Solidity: function updateCheckpoint(uint256 blockNumber) returns()
func (_Votemanager *VotemanagerTransactor) UpdateCheckpoint(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateCheckpoint", blockNumber)
}

// UpdateCheckpoint is a paid mutator transaction binding the contract method 0x3e92aebe.
//
// Solidity: function updateCheckpoint(uint256 blockNumber) returns()
func (_Votemanager *VotemanagerSession) UpdateCheckpoint(blockNumber *big.Int) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateCheckpoint(&_Votemanager.TransactOpts, blockNumber)
}

// UpdateCheckpoint is a paid mutator transaction binding the contract method 0x3e92aebe.
//
// Solidity: function updateCheckpoint(uint256 blockNumber) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateCheckpoint(blockNumber *big.Int) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateCheckpoint(&_Votemanager.TransactOpts, blockNumber)
}

// UpdateSfcAddress is a paid mutator transaction binding the contract method 0xcbb9d8c9.
//
// Solidity: function updateSfcAddress(address sfcAddress) returns()
func (_Votemanager *VotemanagerTransactor) UpdateSfcAddress(opts *bind.TransactOpts, sfcAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateSfcAddress", sfcAddress)
}

// UpdateSfcAddress is a paid mutator transaction binding the contract method 0xcbb9d8c9.
//
// Solidity: function updateSfcAddress(address sfcAddress) returns()
func (_Votemanager *VotemanagerSession) UpdateSfcAddress(sfcAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateSfcAddress(&_Votemanager.TransactOpts, sfcAddress)
}

// UpdateSfcAddress is a paid mutator transaction binding the contract method 0xcbb9d8c9.
//
// Solidity: function updateSfcAddress(address sfcAddress) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateSfcAddress(sfcAddress common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateSfcAddress(&_Votemanager.TransactOpts, sfcAddress)
}

// UpdateTokenRegistryAddress is a paid mutator transaction binding the contract method 0x520051d7.
//
// Solidity: function updateTokenRegistryAddress(address tokenRegistryAddress_) returns()
func (_Votemanager *VotemanagerTransactor) UpdateTokenRegistryAddress(opts *bind.TransactOpts, tokenRegistryAddress_ common.Address) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateTokenRegistryAddress", tokenRegistryAddress_)
}

// UpdateTokenRegistryAddress is a paid mutator transaction binding the contract method 0x520051d7.
//
// Solidity: function updateTokenRegistryAddress(address tokenRegistryAddress_) returns()
func (_Votemanager *VotemanagerSession) UpdateTokenRegistryAddress(tokenRegistryAddress_ common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateTokenRegistryAddress(&_Votemanager.TransactOpts, tokenRegistryAddress_)
}

// UpdateTokenRegistryAddress is a paid mutator transaction binding the contract method 0x520051d7.
//
// Solidity: function updateTokenRegistryAddress(address tokenRegistryAddress_) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateTokenRegistryAddress(tokenRegistryAddress_ common.Address) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateTokenRegistryAddress(&_Votemanager.TransactOpts, tokenRegistryAddress_)
}

// UpdateVoteBufferPercentage is a paid mutator transaction binding the contract method 0xf3a3a784.
//
// Solidity: function updateVoteBufferPercentage(uint8 voteBufferPercent_) returns()
func (_Votemanager *VotemanagerTransactor) UpdateVoteBufferPercentage(opts *bind.TransactOpts, voteBufferPercent_ uint8) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateVoteBufferPercentage", voteBufferPercent_)
}

// UpdateVoteBufferPercentage is a paid mutator transaction binding the contract method 0xf3a3a784.
//
// Solidity: function updateVoteBufferPercentage(uint8 voteBufferPercent_) returns()
func (_Votemanager *VotemanagerSession) UpdateVoteBufferPercentage(voteBufferPercent_ uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateVoteBufferPercentage(&_Votemanager.TransactOpts, voteBufferPercent_)
}

// UpdateVoteBufferPercentage is a paid mutator transaction binding the contract method 0xf3a3a784.
//
// Solidity: function updateVoteBufferPercentage(uint8 voteBufferPercent_) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateVoteBufferPercentage(voteBufferPercent_ uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateVoteBufferPercentage(&_Votemanager.TransactOpts, voteBufferPercent_)
}

// UpdateVotePercentage is a paid mutator transaction binding the contract method 0x8ba8405c.
//
// Solidity: function updateVotePercentage(uint8 votePercentage_) returns()
func (_Votemanager *VotemanagerTransactor) UpdateVotePercentage(opts *bind.TransactOpts, votePercentage_ uint8) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "updateVotePercentage", votePercentage_)
}

// UpdateVotePercentage is a paid mutator transaction binding the contract method 0x8ba8405c.
//
// Solidity: function updateVotePercentage(uint8 votePercentage_) returns()
func (_Votemanager *VotemanagerSession) UpdateVotePercentage(votePercentage_ uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateVotePercentage(&_Votemanager.TransactOpts, votePercentage_)
}

// UpdateVotePercentage is a paid mutator transaction binding the contract method 0x8ba8405c.
//
// Solidity: function updateVotePercentage(uint8 votePercentage_) returns()
func (_Votemanager *VotemanagerTransactorSession) UpdateVotePercentage(votePercentage_ uint8) (*types.Transaction, error) {
	return _Votemanager.Contract.UpdateVotePercentage(&_Votemanager.TransactOpts, votePercentage_)
}

// VoteBatch is a paid mutator transaction binding the contract method 0xe53edd93.
//
// Solidity: function voteBatch((uint256,uint256,uint256,uint16)[] payload) returns()
func (_Votemanager *VotemanagerTransactor) VoteBatch(opts *bind.TransactOpts, payload []VoteManagerVotePayload) (*types.Transaction, error) {
	return _Votemanager.contract.Transact(opts, "voteBatch", payload)
}

// VoteBatch is a paid mutator transaction binding the contract method 0xe53edd93.
//
// Solidity: function voteBatch((uint256,uint256,uint256,uint16)[] payload) returns()
func (_Votemanager *VotemanagerSession) VoteBatch(payload []VoteManagerVotePayload) (*types.Transaction, error) {
	return _Votemanager.Contract.VoteBatch(&_Votemanager.TransactOpts, payload)
}

// VoteBatch is a paid mutator transaction binding the contract method 0xe53edd93.
//
// Solidity: function voteBatch((uint256,uint256,uint256,uint16)[] payload) returns()
func (_Votemanager *VotemanagerTransactorSession) VoteBatch(payload []VoteManagerVotePayload) (*types.Transaction, error) {
	return _Votemanager.Contract.VoteBatch(&_Votemanager.TransactOpts, payload)
}

// VotemanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Votemanager contract.
type VotemanagerInitializedIterator struct {
	Event *VotemanagerInitialized // Event containing the contract specifics and raw log

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
func (it *VotemanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerInitialized)
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
		it.Event = new(VotemanagerInitialized)
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
func (it *VotemanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerInitialized represents a Initialized event raised by the Votemanager contract.
type VotemanagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Votemanager *VotemanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*VotemanagerInitializedIterator, error) {

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VotemanagerInitializedIterator{contract: _Votemanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Votemanager *VotemanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VotemanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerInitialized)
				if err := _Votemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Votemanager *VotemanagerFilterer) ParseInitialized(log types.Log) (*VotemanagerInitialized, error) {
	event := new(VotemanagerInitialized)
	if err := _Votemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemanagerMintTokenIterator is returned from FilterMintToken and is used to iterate over the raw logs and unpacked data for MintToken events raised by the Votemanager contract.
type VotemanagerMintTokenIterator struct {
	Event *VotemanagerMintToken // Event containing the contract specifics and raw log

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
func (it *VotemanagerMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerMintToken)
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
		it.Event = new(VotemanagerMintToken)
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
func (it *VotemanagerMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerMintToken represents a MintToken event raised by the Votemanager contract.
type VotemanagerMintToken struct {
	HashId       *big.Int
	Account      common.Address
	CurrencyType *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintToken is a free log retrieval operation binding the contract event 0x207f55ba01c4294fc34d3fa8c3c37cdc1863b4056044e4bb590d6d1aa9387136.
//
// Solidity: event MintToken(uint256 indexed hashId, address indexed account, uint256 indexed currencyType)
func (_Votemanager *VotemanagerFilterer) FilterMintToken(opts *bind.FilterOpts, hashId []*big.Int, account []common.Address, currencyType []*big.Int) (*VotemanagerMintTokenIterator, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var currencyTypeRule []interface{}
	for _, currencyTypeItem := range currencyType {
		currencyTypeRule = append(currencyTypeRule, currencyTypeItem)
	}

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "MintToken", hashIdRule, accountRule, currencyTypeRule)
	if err != nil {
		return nil, err
	}
	return &VotemanagerMintTokenIterator{contract: _Votemanager.contract, event: "MintToken", logs: logs, sub: sub}, nil
}

// WatchMintToken is a free log subscription operation binding the contract event 0x207f55ba01c4294fc34d3fa8c3c37cdc1863b4056044e4bb590d6d1aa9387136.
//
// Solidity: event MintToken(uint256 indexed hashId, address indexed account, uint256 indexed currencyType)
func (_Votemanager *VotemanagerFilterer) WatchMintToken(opts *bind.WatchOpts, sink chan<- *VotemanagerMintToken, hashId []*big.Int, account []common.Address, currencyType []*big.Int) (event.Subscription, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var currencyTypeRule []interface{}
	for _, currencyTypeItem := range currencyType {
		currencyTypeRule = append(currencyTypeRule, currencyTypeItem)
	}

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "MintToken", hashIdRule, accountRule, currencyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerMintToken)
				if err := _Votemanager.contract.UnpackLog(event, "MintToken", log); err != nil {
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

// ParseMintToken is a log parse operation binding the contract event 0x207f55ba01c4294fc34d3fa8c3c37cdc1863b4056044e4bb590d6d1aa9387136.
//
// Solidity: event MintToken(uint256 indexed hashId, address indexed account, uint256 indexed currencyType)
func (_Votemanager *VotemanagerFilterer) ParseMintToken(log types.Log) (*VotemanagerMintToken, error) {
	event := new(VotemanagerMintToken)
	if err := _Votemanager.contract.UnpackLog(event, "MintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Votemanager contract.
type VotemanagerOwnershipTransferredIterator struct {
	Event *VotemanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VotemanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerOwnershipTransferred)
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
		it.Event = new(VotemanagerOwnershipTransferred)
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
func (it *VotemanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Votemanager contract.
type VotemanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Votemanager *VotemanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VotemanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VotemanagerOwnershipTransferredIterator{contract: _Votemanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Votemanager *VotemanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VotemanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerOwnershipTransferred)
				if err := _Votemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Votemanager *VotemanagerFilterer) ParseOwnershipTransferred(log types.Log) (*VotemanagerOwnershipTransferred, error) {
	event := new(VotemanagerOwnershipTransferred)
	if err := _Votemanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotemanagerVoteTokenIterator is returned from FilterVoteToken and is used to iterate over the raw logs and unpacked data for VoteToken events raised by the Votemanager contract.
type VotemanagerVoteTokenIterator struct {
	Event *VotemanagerVoteToken // Event containing the contract specifics and raw log

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
func (it *VotemanagerVoteTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotemanagerVoteToken)
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
		it.Event = new(VotemanagerVoteToken)
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
func (it *VotemanagerVoteTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotemanagerVoteTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotemanagerVoteToken represents a VoteToken event raised by the Votemanager contract.
type VotemanagerVoteToken struct {
	HashId       *big.Int
	ValidatorId  *big.Int
	CurrencyType *big.Int
	Votes        uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteToken is a free log retrieval operation binding the contract event 0x964a6dd8f4d3c3f266a1a580e00a7eacf0f3701286d408d2bb1583449f1943a1.
//
// Solidity: event VoteToken(uint256 indexed hashId, uint256 indexed validatorId, uint256 indexed currencyType, uint16 votes)
func (_Votemanager *VotemanagerFilterer) FilterVoteToken(opts *bind.FilterOpts, hashId []*big.Int, validatorId []*big.Int, currencyType []*big.Int) (*VotemanagerVoteTokenIterator, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var currencyTypeRule []interface{}
	for _, currencyTypeItem := range currencyType {
		currencyTypeRule = append(currencyTypeRule, currencyTypeItem)
	}

	logs, sub, err := _Votemanager.contract.FilterLogs(opts, "VoteToken", hashIdRule, validatorIdRule, currencyTypeRule)
	if err != nil {
		return nil, err
	}
	return &VotemanagerVoteTokenIterator{contract: _Votemanager.contract, event: "VoteToken", logs: logs, sub: sub}, nil
}

// WatchVoteToken is a free log subscription operation binding the contract event 0x964a6dd8f4d3c3f266a1a580e00a7eacf0f3701286d408d2bb1583449f1943a1.
//
// Solidity: event VoteToken(uint256 indexed hashId, uint256 indexed validatorId, uint256 indexed currencyType, uint16 votes)
func (_Votemanager *VotemanagerFilterer) WatchVoteToken(opts *bind.WatchOpts, sink chan<- *VotemanagerVoteToken, hashId []*big.Int, validatorId []*big.Int, currencyType []*big.Int) (event.Subscription, error) {

	var hashIdRule []interface{}
	for _, hashIdItem := range hashId {
		hashIdRule = append(hashIdRule, hashIdItem)
	}
	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}
	var currencyTypeRule []interface{}
	for _, currencyTypeItem := range currencyType {
		currencyTypeRule = append(currencyTypeRule, currencyTypeItem)
	}

	logs, sub, err := _Votemanager.contract.WatchLogs(opts, "VoteToken", hashIdRule, validatorIdRule, currencyTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotemanagerVoteToken)
				if err := _Votemanager.contract.UnpackLog(event, "VoteToken", log); err != nil {
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

// ParseVoteToken is a log parse operation binding the contract event 0x964a6dd8f4d3c3f266a1a580e00a7eacf0f3701286d408d2bb1583449f1943a1.
//
// Solidity: event VoteToken(uint256 indexed hashId, uint256 indexed validatorId, uint256 indexed currencyType, uint16 votes)
func (_Votemanager *VotemanagerFilterer) ParseVoteToken(log types.Log) (*VotemanagerVoteToken, error) {
	event := new(VotemanagerVoteToken)
	if err := _Votemanager.contract.UnpackLog(event, "VoteToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

