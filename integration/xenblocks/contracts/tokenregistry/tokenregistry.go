// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenregistry

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

// TokenRegistryFoundToken is an auto generated low-level Go binding around an user-defined struct.
type TokenRegistryFoundToken struct {
	CurrencyType *big.Int
	Version      uint16
}

// TokenRegistryTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type TokenRegistryTokenConfig struct {
	CurrencyType *big.Int
	Token        common.Address
	Exclusive    bool
}

// TokenregistryMetaData contains all meta data concerning the Tokenregistry contract.
var TokenregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TokensUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minedTimestamp\",\"type\":\"uint256\"}],\"name\":\"findTokensFromArgon2Hash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"internalType\":\"structTokenRegistry.FoundToken[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractVoterToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTokenConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"internalType\":\"contractVoterToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exclusive\",\"type\":\"bool\"}],\"internalType\":\"structTokenRegistry.TokenConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenConfigs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"internalType\":\"contractVoterToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exclusive\",\"type\":\"bool\"}],\"internalType\":\"structTokenRegistry.TokenConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenVersions\",\"outputs\":[{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exclusive\",\"type\":\"bool\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokenIdBySymbol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokensById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currencyType\",\"type\":\"uint256\"},{\"internalType\":\"contractVoterToken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exclusive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"exclusive\",\"type\":\"bool\"}],\"name\":\"updateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voteManagerAddr\",\"type\":\"address\"}],\"name\":\"updateVoteManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenregistryMetaData.ABI instead.
var TokenregistryABI = TokenregistryMetaData.ABI

// Tokenregistry is an auto generated Go binding around an Ethereum contract.
type Tokenregistry struct {
	TokenregistryCaller     // Read-only binding to the contract
	TokenregistryTransactor // Write-only binding to the contract
	TokenregistryFilterer   // Log filterer for contract events
}

// TokenregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenregistrySession struct {
	Contract     *Tokenregistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenregistryCallerSession struct {
	Contract *TokenregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenregistryTransactorSession struct {
	Contract     *TokenregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenregistryRaw struct {
	Contract *Tokenregistry // Generic contract binding to access the raw methods on
}

// TokenregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenregistryCallerRaw struct {
	Contract *TokenregistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenregistryTransactorRaw struct {
	Contract *TokenregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenregistry creates a new instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistry(address common.Address, backend bind.ContractBackend) (*Tokenregistry, error) {
	contract, err := bindTokenregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenregistry{TokenregistryCaller: TokenregistryCaller{contract: contract}, TokenregistryTransactor: TokenregistryTransactor{contract: contract}, TokenregistryFilterer: TokenregistryFilterer{contract: contract}}, nil
}

// NewTokenregistryCaller creates a new read-only instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryCaller(address common.Address, caller bind.ContractCaller) (*TokenregistryCaller, error) {
	contract, err := bindTokenregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenregistryCaller{contract: contract}, nil
}

// NewTokenregistryTransactor creates a new write-only instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenregistryTransactor, error) {
	contract, err := bindTokenregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenregistryTransactor{contract: contract}, nil
}

// NewTokenregistryFilterer creates a new log filterer instance of Tokenregistry, bound to a specific deployed contract.
func NewTokenregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenregistryFilterer, error) {
	contract, err := bindTokenregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenregistryFilterer{contract: contract}, nil
}

// bindTokenregistry binds a generic wrapper to an already deployed contract.
func bindTokenregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenregistry *TokenregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenregistry.Contract.TokenregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenregistry *TokenregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TokenregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenregistry *TokenregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TokenregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenregistry *TokenregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenregistry *TokenregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenregistry *TokenregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenregistry.Contract.contract.Transact(opts, method, params...)
}

// FindTokensFromArgon2Hash is a free data retrieval call binding the contract method 0x7bbbe12d.
//
// Solidity: function findTokensFromArgon2Hash(string hash, uint256 minedTimestamp) view returns((uint256,uint16)[])
func (_Tokenregistry *TokenregistryCaller) FindTokensFromArgon2Hash(opts *bind.CallOpts, hash string, minedTimestamp *big.Int) ([]TokenRegistryFoundToken, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "findTokensFromArgon2Hash", hash, minedTimestamp)

	if err != nil {
		return *new([]TokenRegistryFoundToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenRegistryFoundToken)).(*[]TokenRegistryFoundToken)

	return out0, err

}

// FindTokensFromArgon2Hash is a free data retrieval call binding the contract method 0x7bbbe12d.
//
// Solidity: function findTokensFromArgon2Hash(string hash, uint256 minedTimestamp) view returns((uint256,uint16)[])
func (_Tokenregistry *TokenregistrySession) FindTokensFromArgon2Hash(hash string, minedTimestamp *big.Int) ([]TokenRegistryFoundToken, error) {
	return _Tokenregistry.Contract.FindTokensFromArgon2Hash(&_Tokenregistry.CallOpts, hash, minedTimestamp)
}

// FindTokensFromArgon2Hash is a free data retrieval call binding the contract method 0x7bbbe12d.
//
// Solidity: function findTokensFromArgon2Hash(string hash, uint256 minedTimestamp) view returns((uint256,uint16)[])
func (_Tokenregistry *TokenregistryCallerSession) FindTokensFromArgon2Hash(hash string, minedTimestamp *big.Int) ([]TokenRegistryFoundToken, error) {
	return _Tokenregistry.Contract.FindTokensFromArgon2Hash(&_Tokenregistry.CallOpts, hash, minedTimestamp)
}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 id) view returns(address)
func (_Tokenregistry *TokenregistryCaller) GetToken(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "getToken", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 id) view returns(address)
func (_Tokenregistry *TokenregistrySession) GetToken(id *big.Int) (common.Address, error) {
	return _Tokenregistry.Contract.GetToken(&_Tokenregistry.CallOpts, id)
}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 id) view returns(address)
func (_Tokenregistry *TokenregistryCallerSession) GetToken(id *big.Int) (common.Address, error) {
	return _Tokenregistry.Contract.GetToken(&_Tokenregistry.CallOpts, id)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 id) view returns((uint256,address,bool))
func (_Tokenregistry *TokenregistryCaller) GetTokenConfig(opts *bind.CallOpts, id *big.Int) (TokenRegistryTokenConfig, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "getTokenConfig", id)

	if err != nil {
		return *new(TokenRegistryTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenRegistryTokenConfig)).(*TokenRegistryTokenConfig)

	return out0, err

}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 id) view returns((uint256,address,bool))
func (_Tokenregistry *TokenregistrySession) GetTokenConfig(id *big.Int) (TokenRegistryTokenConfig, error) {
	return _Tokenregistry.Contract.GetTokenConfig(&_Tokenregistry.CallOpts, id)
}

// GetTokenConfig is a free data retrieval call binding the contract method 0x8a003888.
//
// Solidity: function getTokenConfig(uint256 id) view returns((uint256,address,bool))
func (_Tokenregistry *TokenregistryCallerSession) GetTokenConfig(id *big.Int) (TokenRegistryTokenConfig, error) {
	return _Tokenregistry.Contract.GetTokenConfig(&_Tokenregistry.CallOpts, id)
}

// GetTokenConfigs is a free data retrieval call binding the contract method 0x5547171d.
//
// Solidity: function getTokenConfigs() view returns((uint256,address,bool)[])
func (_Tokenregistry *TokenregistryCaller) GetTokenConfigs(opts *bind.CallOpts) ([]TokenRegistryTokenConfig, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "getTokenConfigs")

	if err != nil {
		return *new([]TokenRegistryTokenConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]TokenRegistryTokenConfig)).(*[]TokenRegistryTokenConfig)

	return out0, err

}

// GetTokenConfigs is a free data retrieval call binding the contract method 0x5547171d.
//
// Solidity: function getTokenConfigs() view returns((uint256,address,bool)[])
func (_Tokenregistry *TokenregistrySession) GetTokenConfigs() ([]TokenRegistryTokenConfig, error) {
	return _Tokenregistry.Contract.GetTokenConfigs(&_Tokenregistry.CallOpts)
}

// GetTokenConfigs is a free data retrieval call binding the contract method 0x5547171d.
//
// Solidity: function getTokenConfigs() view returns((uint256,address,bool)[])
func (_Tokenregistry *TokenregistryCallerSession) GetTokenConfigs() ([]TokenRegistryTokenConfig, error) {
	return _Tokenregistry.Contract.GetTokenConfigs(&_Tokenregistry.CallOpts)
}

// GetTokenVersions is a free data retrieval call binding the contract method 0x7e519bd0.
//
// Solidity: function getTokenVersions() view returns(uint16[])
func (_Tokenregistry *TokenregistryCaller) GetTokenVersions(opts *bind.CallOpts) ([]uint16, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "getTokenVersions")

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetTokenVersions is a free data retrieval call binding the contract method 0x7e519bd0.
//
// Solidity: function getTokenVersions() view returns(uint16[])
func (_Tokenregistry *TokenregistrySession) GetTokenVersions() ([]uint16, error) {
	return _Tokenregistry.Contract.GetTokenVersions(&_Tokenregistry.CallOpts)
}

// GetTokenVersions is a free data retrieval call binding the contract method 0x7e519bd0.
//
// Solidity: function getTokenVersions() view returns(uint16[])
func (_Tokenregistry *TokenregistryCallerSession) GetTokenVersions() ([]uint16, error) {
	return _Tokenregistry.Contract.GetTokenVersions(&_Tokenregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenregistry *TokenregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenregistry *TokenregistrySession) Owner() (common.Address, error) {
	return _Tokenregistry.Contract.Owner(&_Tokenregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenregistry *TokenregistryCallerSession) Owner() (common.Address, error) {
	return _Tokenregistry.Contract.Owner(&_Tokenregistry.CallOpts)
}

// TokenIdBySymbol is a free data retrieval call binding the contract method 0xb4c33d3f.
//
// Solidity: function tokenIdBySymbol(string ) view returns(uint256)
func (_Tokenregistry *TokenregistryCaller) TokenIdBySymbol(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "tokenIdBySymbol", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdBySymbol is a free data retrieval call binding the contract method 0xb4c33d3f.
//
// Solidity: function tokenIdBySymbol(string ) view returns(uint256)
func (_Tokenregistry *TokenregistrySession) TokenIdBySymbol(arg0 string) (*big.Int, error) {
	return _Tokenregistry.Contract.TokenIdBySymbol(&_Tokenregistry.CallOpts, arg0)
}

// TokenIdBySymbol is a free data retrieval call binding the contract method 0xb4c33d3f.
//
// Solidity: function tokenIdBySymbol(string ) view returns(uint256)
func (_Tokenregistry *TokenregistryCallerSession) TokenIdBySymbol(arg0 string) (*big.Int, error) {
	return _Tokenregistry.Contract.TokenIdBySymbol(&_Tokenregistry.CallOpts, arg0)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Tokenregistry *TokenregistryCaller) TokenIdCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "tokenIdCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Tokenregistry *TokenregistrySession) TokenIdCounter() (*big.Int, error) {
	return _Tokenregistry.Contract.TokenIdCounter(&_Tokenregistry.CallOpts)
}

// TokenIdCounter is a free data retrieval call binding the contract method 0x98bdf6f5.
//
// Solidity: function tokenIdCounter() view returns(uint256)
func (_Tokenregistry *TokenregistryCallerSession) TokenIdCounter() (*big.Int, error) {
	return _Tokenregistry.Contract.TokenIdCounter(&_Tokenregistry.CallOpts)
}

// TokensById is a free data retrieval call binding the contract method 0x346fc98b.
//
// Solidity: function tokensById(uint256 ) view returns(uint256 currencyType, address token, bool exclusive)
func (_Tokenregistry *TokenregistryCaller) TokensById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CurrencyType *big.Int
	Token        common.Address
	Exclusive    bool
}, error) {
	var out []interface{}
	err := _Tokenregistry.contract.Call(opts, &out, "tokensById", arg0)

	outstruct := new(struct {
		CurrencyType *big.Int
		Token        common.Address
		Exclusive    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrencyType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Exclusive = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// TokensById is a free data retrieval call binding the contract method 0x346fc98b.
//
// Solidity: function tokensById(uint256 ) view returns(uint256 currencyType, address token, bool exclusive)
func (_Tokenregistry *TokenregistrySession) TokensById(arg0 *big.Int) (struct {
	CurrencyType *big.Int
	Token        common.Address
	Exclusive    bool
}, error) {
	return _Tokenregistry.Contract.TokensById(&_Tokenregistry.CallOpts, arg0)
}

// TokensById is a free data retrieval call binding the contract method 0x346fc98b.
//
// Solidity: function tokensById(uint256 ) view returns(uint256 currencyType, address token, bool exclusive)
func (_Tokenregistry *TokenregistryCallerSession) TokensById(arg0 *big.Int) (struct {
	CurrencyType *big.Int
	Token        common.Address
	Exclusive    bool
}, error) {
	return _Tokenregistry.Contract.TokensById(&_Tokenregistry.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Tokenregistry *TokenregistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Tokenregistry *TokenregistrySession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.Initialize(&_Tokenregistry.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Tokenregistry *TokenregistryTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.Initialize(&_Tokenregistry.TransactOpts, initialOwner)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x6710b83f.
//
// Solidity: function registerToken(address addr, bool exclusive) returns()
func (_Tokenregistry *TokenregistryTransactor) RegisterToken(opts *bind.TransactOpts, addr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "registerToken", addr, exclusive)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x6710b83f.
//
// Solidity: function registerToken(address addr, bool exclusive) returns()
func (_Tokenregistry *TokenregistrySession) RegisterToken(addr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.Contract.RegisterToken(&_Tokenregistry.TransactOpts, addr, exclusive)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x6710b83f.
//
// Solidity: function registerToken(address addr, bool exclusive) returns()
func (_Tokenregistry *TokenregistryTransactorSession) RegisterToken(addr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.Contract.RegisterToken(&_Tokenregistry.TransactOpts, addr, exclusive)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenregistry *TokenregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenregistry *TokenregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenregistry.Contract.RenounceOwnership(&_Tokenregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenregistry *TokenregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenregistry.Contract.RenounceOwnership(&_Tokenregistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenregistry *TokenregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenregistry *TokenregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TransferOwnership(&_Tokenregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenregistry *TokenregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.TransferOwnership(&_Tokenregistry.TransactOpts, newOwner)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xda8d73fa.
//
// Solidity: function updateToken(uint256 tokenId, address tokenAddr, bool exclusive) returns()
func (_Tokenregistry *TokenregistryTransactor) UpdateToken(opts *bind.TransactOpts, tokenId *big.Int, tokenAddr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "updateToken", tokenId, tokenAddr, exclusive)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xda8d73fa.
//
// Solidity: function updateToken(uint256 tokenId, address tokenAddr, bool exclusive) returns()
func (_Tokenregistry *TokenregistrySession) UpdateToken(tokenId *big.Int, tokenAddr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.Contract.UpdateToken(&_Tokenregistry.TransactOpts, tokenId, tokenAddr, exclusive)
}

// UpdateToken is a paid mutator transaction binding the contract method 0xda8d73fa.
//
// Solidity: function updateToken(uint256 tokenId, address tokenAddr, bool exclusive) returns()
func (_Tokenregistry *TokenregistryTransactorSession) UpdateToken(tokenId *big.Int, tokenAddr common.Address, exclusive bool) (*types.Transaction, error) {
	return _Tokenregistry.Contract.UpdateToken(&_Tokenregistry.TransactOpts, tokenId, tokenAddr, exclusive)
}

// UpdateVoteManagerAddress is a paid mutator transaction binding the contract method 0xf5b4751d.
//
// Solidity: function updateVoteManagerAddress(address voteManagerAddr) returns()
func (_Tokenregistry *TokenregistryTransactor) UpdateVoteManagerAddress(opts *bind.TransactOpts, voteManagerAddr common.Address) (*types.Transaction, error) {
	return _Tokenregistry.contract.Transact(opts, "updateVoteManagerAddress", voteManagerAddr)
}

// UpdateVoteManagerAddress is a paid mutator transaction binding the contract method 0xf5b4751d.
//
// Solidity: function updateVoteManagerAddress(address voteManagerAddr) returns()
func (_Tokenregistry *TokenregistrySession) UpdateVoteManagerAddress(voteManagerAddr common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.UpdateVoteManagerAddress(&_Tokenregistry.TransactOpts, voteManagerAddr)
}

// UpdateVoteManagerAddress is a paid mutator transaction binding the contract method 0xf5b4751d.
//
// Solidity: function updateVoteManagerAddress(address voteManagerAddr) returns()
func (_Tokenregistry *TokenregistryTransactorSession) UpdateVoteManagerAddress(voteManagerAddr common.Address) (*types.Transaction, error) {
	return _Tokenregistry.Contract.UpdateVoteManagerAddress(&_Tokenregistry.TransactOpts, voteManagerAddr)
}

// TokenregistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenregistry contract.
type TokenregistryInitializedIterator struct {
	Event *TokenregistryInitialized // Event containing the contract specifics and raw log

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
func (it *TokenregistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenregistryInitialized)
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
		it.Event = new(TokenregistryInitialized)
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
func (it *TokenregistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenregistryInitialized represents a Initialized event raised by the Tokenregistry contract.
type TokenregistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tokenregistry *TokenregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenregistryInitializedIterator, error) {

	logs, sub, err := _Tokenregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenregistryInitializedIterator{contract: _Tokenregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tokenregistry *TokenregistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenregistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenregistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenregistryInitialized)
				if err := _Tokenregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Tokenregistry *TokenregistryFilterer) ParseInitialized(log types.Log) (*TokenregistryInitialized, error) {
	event := new(TokenregistryInitialized)
	if err := _Tokenregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenregistry contract.
type TokenregistryOwnershipTransferredIterator struct {
	Event *TokenregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenregistryOwnershipTransferred)
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
		it.Event = new(TokenregistryOwnershipTransferred)
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
func (it *TokenregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenregistry contract.
type TokenregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenregistry *TokenregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenregistryOwnershipTransferredIterator{contract: _Tokenregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenregistry *TokenregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenregistryOwnershipTransferred)
				if err := _Tokenregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokenregistry *TokenregistryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenregistryOwnershipTransferred, error) {
	event := new(TokenregistryOwnershipTransferred)
	if err := _Tokenregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenregistryTokensUpdatedIterator is returned from FilterTokensUpdated and is used to iterate over the raw logs and unpacked data for TokensUpdated events raised by the Tokenregistry contract.
type TokenregistryTokensUpdatedIterator struct {
	Event *TokenregistryTokensUpdated // Event containing the contract specifics and raw log

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
func (it *TokenregistryTokensUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenregistryTokensUpdated)
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
		it.Event = new(TokenregistryTokensUpdated)
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
func (it *TokenregistryTokensUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenregistryTokensUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenregistryTokensUpdated represents a TokensUpdated event raised by the Tokenregistry contract.
type TokenregistryTokensUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTokensUpdated is a free log retrieval operation binding the contract event 0xc7cce9e0ee378bcf7fd342cae3d8eb17caaf3e9be70c370256ba6517189638dc.
//
// Solidity: event TokensUpdated()
func (_Tokenregistry *TokenregistryFilterer) FilterTokensUpdated(opts *bind.FilterOpts) (*TokenregistryTokensUpdatedIterator, error) {

	logs, sub, err := _Tokenregistry.contract.FilterLogs(opts, "TokensUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenregistryTokensUpdatedIterator{contract: _Tokenregistry.contract, event: "TokensUpdated", logs: logs, sub: sub}, nil
}

// WatchTokensUpdated is a free log subscription operation binding the contract event 0xc7cce9e0ee378bcf7fd342cae3d8eb17caaf3e9be70c370256ba6517189638dc.
//
// Solidity: event TokensUpdated()
func (_Tokenregistry *TokenregistryFilterer) WatchTokensUpdated(opts *bind.WatchOpts, sink chan<- *TokenregistryTokensUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenregistry.contract.WatchLogs(opts, "TokensUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenregistryTokensUpdated)
				if err := _Tokenregistry.contract.UnpackLog(event, "TokensUpdated", log); err != nil {
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

// ParseTokensUpdated is a log parse operation binding the contract event 0xc7cce9e0ee378bcf7fd342cae3d8eb17caaf3e9be70c370256ba6517189638dc.
//
// Solidity: event TokensUpdated()
func (_Tokenregistry *TokenregistryFilterer) ParseTokensUpdated(log types.Log) (*TokenregistryTokensUpdated, error) {
	event := new(TokenregistryTokensUpdated)
	if err := _Tokenregistry.contract.UnpackLog(event, "TokensUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

