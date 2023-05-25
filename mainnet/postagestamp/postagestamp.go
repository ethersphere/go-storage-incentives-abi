// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package postagestamp

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

// PostagestampMetaData contains all meta data concerning the Postagestamp contract.
var PostagestampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bzzToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_minimumBucketDepth\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"normalisedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"depth\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"bucketDepth\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"immutableFlag\",\"type\":\"bool\"}],\"name\":\"BatchCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newDepth\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"normalisedBalance\",\"type\":\"uint256\"}],\"name\":\"BatchDepthIncrease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"topupAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"normalisedBalance\",\"type\":\"uint256\"}],\"name\":\"BatchTopUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDISTRIBUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"depth\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"immutableFlag\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"normalisedBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bzzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialBalancePerChunk\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_depth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_bucketDepth\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_batchId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_immutable\",\"type\":\"bool\"}],\"name\":\"copyBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialBalancePerChunk\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_depth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_bucketDepth\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_immutable\",\"type\":\"bool\"}],\"name\":\"createBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTotalOutPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"empty\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"expireLimited\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiredBatchesExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstBatchId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_batchId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_newDepth\",\"type\":\"uint8\"}],\"name\":\"increaseDepth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastExpiryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumBucketDepth\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_batchId\",\"type\":\"bytes32\"}],\"name\":\"remainingBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_batchId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_topupAmountPerChunk\",\"type\":\"uint256\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"topupPot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validChunkCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PostagestampABI is the input ABI used to generate the binding from.
// Deprecated: Use PostagestampMetaData.ABI instead.
var PostagestampABI = PostagestampMetaData.ABI

// Postagestamp is an auto generated Go binding around an Ethereum contract.
type Postagestamp struct {
	PostagestampCaller     // Read-only binding to the contract
	PostagestampTransactor // Write-only binding to the contract
	PostagestampFilterer   // Log filterer for contract events
}

// PostagestampCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostagestampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostagestampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostagestampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostagestampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostagestampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostagestampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostagestampSession struct {
	Contract     *Postagestamp     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PostagestampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostagestampCallerSession struct {
	Contract *PostagestampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PostagestampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostagestampTransactorSession struct {
	Contract     *PostagestampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PostagestampRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostagestampRaw struct {
	Contract *Postagestamp // Generic contract binding to access the raw methods on
}

// PostagestampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostagestampCallerRaw struct {
	Contract *PostagestampCaller // Generic read-only contract binding to access the raw methods on
}

// PostagestampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostagestampTransactorRaw struct {
	Contract *PostagestampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPostagestamp creates a new instance of Postagestamp, bound to a specific deployed contract.
func NewPostagestamp(address common.Address, backend bind.ContractBackend) (*Postagestamp, error) {
	contract, err := bindPostagestamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Postagestamp{PostagestampCaller: PostagestampCaller{contract: contract}, PostagestampTransactor: PostagestampTransactor{contract: contract}, PostagestampFilterer: PostagestampFilterer{contract: contract}}, nil
}

// NewPostagestampCaller creates a new read-only instance of Postagestamp, bound to a specific deployed contract.
func NewPostagestampCaller(address common.Address, caller bind.ContractCaller) (*PostagestampCaller, error) {
	contract, err := bindPostagestamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostagestampCaller{contract: contract}, nil
}

// NewPostagestampTransactor creates a new write-only instance of Postagestamp, bound to a specific deployed contract.
func NewPostagestampTransactor(address common.Address, transactor bind.ContractTransactor) (*PostagestampTransactor, error) {
	contract, err := bindPostagestamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostagestampTransactor{contract: contract}, nil
}

// NewPostagestampFilterer creates a new log filterer instance of Postagestamp, bound to a specific deployed contract.
func NewPostagestampFilterer(address common.Address, filterer bind.ContractFilterer) (*PostagestampFilterer, error) {
	contract, err := bindPostagestamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostagestampFilterer{contract: contract}, nil
}

// bindPostagestamp binds a generic wrapper to an already deployed contract.
func bindPostagestamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PostagestampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Postagestamp *PostagestampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Postagestamp.Contract.PostagestampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Postagestamp *PostagestampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Postagestamp.Contract.PostagestampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Postagestamp *PostagestampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Postagestamp.Contract.PostagestampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Postagestamp *PostagestampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Postagestamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Postagestamp *PostagestampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Postagestamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Postagestamp *PostagestampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Postagestamp.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Postagestamp.Contract.DEFAULTADMINROLE(&_Postagestamp.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Postagestamp.Contract.DEFAULTADMINROLE(&_Postagestamp.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampSession) PAUSERROLE() ([32]byte, error) {
	return _Postagestamp.Contract.PAUSERROLE(&_Postagestamp.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Postagestamp.Contract.PAUSERROLE(&_Postagestamp.CallOpts)
}

// PRICEORACLEROLE is a free data retrieval call binding the contract method 0xb998902f.
//
// Solidity: function PRICE_ORACLE_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCaller) PRICEORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "PRICE_ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEORACLEROLE is a free data retrieval call binding the contract method 0xb998902f.
//
// Solidity: function PRICE_ORACLE_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampSession) PRICEORACLEROLE() ([32]byte, error) {
	return _Postagestamp.Contract.PRICEORACLEROLE(&_Postagestamp.CallOpts)
}

// PRICEORACLEROLE is a free data retrieval call binding the contract method 0xb998902f.
//
// Solidity: function PRICE_ORACLE_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) PRICEORACLEROLE() ([32]byte, error) {
	return _Postagestamp.Contract.PRICEORACLEROLE(&_Postagestamp.CallOpts)
}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCaller) REDISTRIBUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "REDISTRIBUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampSession) REDISTRIBUTORROLE() ([32]byte, error) {
	return _Postagestamp.Contract.REDISTRIBUTORROLE(&_Postagestamp.CallOpts)
}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) REDISTRIBUTORROLE() ([32]byte, error) {
	return _Postagestamp.Contract.REDISTRIBUTORROLE(&_Postagestamp.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xc81e25ab.
//
// Solidity: function batches(bytes32 ) view returns(address owner, uint8 depth, bool immutableFlag, uint256 normalisedBalance)
func (_Postagestamp *PostagestampCaller) Batches(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner             common.Address
	Depth             uint8
	ImmutableFlag     bool
	NormalisedBalance *big.Int
}, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "batches", arg0)

	outstruct := new(struct {
		Owner             common.Address
		Depth             uint8
		ImmutableFlag     bool
		NormalisedBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Depth = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ImmutableFlag = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.NormalisedBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Batches is a free data retrieval call binding the contract method 0xc81e25ab.
//
// Solidity: function batches(bytes32 ) view returns(address owner, uint8 depth, bool immutableFlag, uint256 normalisedBalance)
func (_Postagestamp *PostagestampSession) Batches(arg0 [32]byte) (struct {
	Owner             common.Address
	Depth             uint8
	ImmutableFlag     bool
	NormalisedBalance *big.Int
}, error) {
	return _Postagestamp.Contract.Batches(&_Postagestamp.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xc81e25ab.
//
// Solidity: function batches(bytes32 ) view returns(address owner, uint8 depth, bool immutableFlag, uint256 normalisedBalance)
func (_Postagestamp *PostagestampCallerSession) Batches(arg0 [32]byte) (struct {
	Owner             common.Address
	Depth             uint8
	ImmutableFlag     bool
	NormalisedBalance *big.Int
}, error) {
	return _Postagestamp.Contract.Batches(&_Postagestamp.CallOpts, arg0)
}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Postagestamp *PostagestampCaller) BzzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "bzzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Postagestamp *PostagestampSession) BzzToken() (common.Address, error) {
	return _Postagestamp.Contract.BzzToken(&_Postagestamp.CallOpts)
}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Postagestamp *PostagestampCallerSession) BzzToken() (common.Address, error) {
	return _Postagestamp.Contract.BzzToken(&_Postagestamp.CallOpts)
}

// CurrentTotalOutPayment is a free data retrieval call binding the contract method 0x51b17cd0.
//
// Solidity: function currentTotalOutPayment() view returns(uint256)
func (_Postagestamp *PostagestampCaller) CurrentTotalOutPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "currentTotalOutPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTotalOutPayment is a free data retrieval call binding the contract method 0x51b17cd0.
//
// Solidity: function currentTotalOutPayment() view returns(uint256)
func (_Postagestamp *PostagestampSession) CurrentTotalOutPayment() (*big.Int, error) {
	return _Postagestamp.Contract.CurrentTotalOutPayment(&_Postagestamp.CallOpts)
}

// CurrentTotalOutPayment is a free data retrieval call binding the contract method 0x51b17cd0.
//
// Solidity: function currentTotalOutPayment() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) CurrentTotalOutPayment() (*big.Int, error) {
	return _Postagestamp.Contract.CurrentTotalOutPayment(&_Postagestamp.CallOpts)
}

// Empty is a free data retrieval call binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() view returns(bool)
func (_Postagestamp *PostagestampCaller) Empty(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "empty")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Empty is a free data retrieval call binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() view returns(bool)
func (_Postagestamp *PostagestampSession) Empty() (bool, error) {
	return _Postagestamp.Contract.Empty(&_Postagestamp.CallOpts)
}

// Empty is a free data retrieval call binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() view returns(bool)
func (_Postagestamp *PostagestampCallerSession) Empty() (bool, error) {
	return _Postagestamp.Contract.Empty(&_Postagestamp.CallOpts)
}

// ExpiredBatchesExist is a free data retrieval call binding the contract method 0x711bfa2b.
//
// Solidity: function expiredBatchesExist() view returns(bool)
func (_Postagestamp *PostagestampCaller) ExpiredBatchesExist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "expiredBatchesExist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExpiredBatchesExist is a free data retrieval call binding the contract method 0x711bfa2b.
//
// Solidity: function expiredBatchesExist() view returns(bool)
func (_Postagestamp *PostagestampSession) ExpiredBatchesExist() (bool, error) {
	return _Postagestamp.Contract.ExpiredBatchesExist(&_Postagestamp.CallOpts)
}

// ExpiredBatchesExist is a free data retrieval call binding the contract method 0x711bfa2b.
//
// Solidity: function expiredBatchesExist() view returns(bool)
func (_Postagestamp *PostagestampCallerSession) ExpiredBatchesExist() (bool, error) {
	return _Postagestamp.Contract.ExpiredBatchesExist(&_Postagestamp.CallOpts)
}

// FirstBatchId is a free data retrieval call binding the contract method 0x8b82547f.
//
// Solidity: function firstBatchId() view returns(bytes32)
func (_Postagestamp *PostagestampCaller) FirstBatchId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "firstBatchId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FirstBatchId is a free data retrieval call binding the contract method 0x8b82547f.
//
// Solidity: function firstBatchId() view returns(bytes32)
func (_Postagestamp *PostagestampSession) FirstBatchId() ([32]byte, error) {
	return _Postagestamp.Contract.FirstBatchId(&_Postagestamp.CallOpts)
}

// FirstBatchId is a free data retrieval call binding the contract method 0x8b82547f.
//
// Solidity: function firstBatchId() view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) FirstBatchId() ([32]byte, error) {
	return _Postagestamp.Contract.FirstBatchId(&_Postagestamp.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Postagestamp *PostagestampCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Postagestamp *PostagestampSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Postagestamp.Contract.GetRoleAdmin(&_Postagestamp.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Postagestamp *PostagestampCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Postagestamp.Contract.GetRoleAdmin(&_Postagestamp.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Postagestamp *PostagestampCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Postagestamp *PostagestampSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Postagestamp.Contract.HasRole(&_Postagestamp.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Postagestamp *PostagestampCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Postagestamp.Contract.HasRole(&_Postagestamp.CallOpts, role, account)
}

// LastExpiryBalance is a free data retrieval call binding the contract method 0xea612e1f.
//
// Solidity: function lastExpiryBalance() view returns(uint256)
func (_Postagestamp *PostagestampCaller) LastExpiryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "lastExpiryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastExpiryBalance is a free data retrieval call binding the contract method 0xea612e1f.
//
// Solidity: function lastExpiryBalance() view returns(uint256)
func (_Postagestamp *PostagestampSession) LastExpiryBalance() (*big.Int, error) {
	return _Postagestamp.Contract.LastExpiryBalance(&_Postagestamp.CallOpts)
}

// LastExpiryBalance is a free data retrieval call binding the contract method 0xea612e1f.
//
// Solidity: function lastExpiryBalance() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) LastExpiryBalance() (*big.Int, error) {
	return _Postagestamp.Contract.LastExpiryBalance(&_Postagestamp.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Postagestamp *PostagestampCaller) LastPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "lastPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Postagestamp *PostagestampSession) LastPrice() (*big.Int, error) {
	return _Postagestamp.Contract.LastPrice(&_Postagestamp.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) LastPrice() (*big.Int, error) {
	return _Postagestamp.Contract.LastPrice(&_Postagestamp.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_Postagestamp *PostagestampCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "lastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_Postagestamp *PostagestampSession) LastUpdatedBlock() (*big.Int, error) {
	return _Postagestamp.Contract.LastUpdatedBlock(&_Postagestamp.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _Postagestamp.Contract.LastUpdatedBlock(&_Postagestamp.CallOpts)
}

// MinimumBucketDepth is a free data retrieval call binding the contract method 0xa81064ee.
//
// Solidity: function minimumBucketDepth() view returns(uint8)
func (_Postagestamp *PostagestampCaller) MinimumBucketDepth(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "minimumBucketDepth")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinimumBucketDepth is a free data retrieval call binding the contract method 0xa81064ee.
//
// Solidity: function minimumBucketDepth() view returns(uint8)
func (_Postagestamp *PostagestampSession) MinimumBucketDepth() (uint8, error) {
	return _Postagestamp.Contract.MinimumBucketDepth(&_Postagestamp.CallOpts)
}

// MinimumBucketDepth is a free data retrieval call binding the contract method 0xa81064ee.
//
// Solidity: function minimumBucketDepth() view returns(uint8)
func (_Postagestamp *PostagestampCallerSession) MinimumBucketDepth() (uint8, error) {
	return _Postagestamp.Contract.MinimumBucketDepth(&_Postagestamp.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Postagestamp *PostagestampCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Postagestamp *PostagestampSession) Paused() (bool, error) {
	return _Postagestamp.Contract.Paused(&_Postagestamp.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Postagestamp *PostagestampCallerSession) Paused() (bool, error) {
	return _Postagestamp.Contract.Paused(&_Postagestamp.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(uint256)
func (_Postagestamp *PostagestampCaller) Pot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "pot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(uint256)
func (_Postagestamp *PostagestampSession) Pot() (*big.Int, error) {
	return _Postagestamp.Contract.Pot(&_Postagestamp.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) Pot() (*big.Int, error) {
	return _Postagestamp.Contract.Pot(&_Postagestamp.CallOpts)
}

// RemainingBalance is a free data retrieval call binding the contract method 0xd71ba7c4.
//
// Solidity: function remainingBalance(bytes32 _batchId) view returns(uint256)
func (_Postagestamp *PostagestampCaller) RemainingBalance(opts *bind.CallOpts, _batchId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "remainingBalance", _batchId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingBalance is a free data retrieval call binding the contract method 0xd71ba7c4.
//
// Solidity: function remainingBalance(bytes32 _batchId) view returns(uint256)
func (_Postagestamp *PostagestampSession) RemainingBalance(_batchId [32]byte) (*big.Int, error) {
	return _Postagestamp.Contract.RemainingBalance(&_Postagestamp.CallOpts, _batchId)
}

// RemainingBalance is a free data retrieval call binding the contract method 0xd71ba7c4.
//
// Solidity: function remainingBalance(bytes32 _batchId) view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) RemainingBalance(_batchId [32]byte) (*big.Int, error) {
	return _Postagestamp.Contract.RemainingBalance(&_Postagestamp.CallOpts, _batchId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Postagestamp *PostagestampCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Postagestamp *PostagestampSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Postagestamp.Contract.SupportsInterface(&_Postagestamp.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Postagestamp *PostagestampCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Postagestamp.Contract.SupportsInterface(&_Postagestamp.CallOpts, interfaceId)
}

// ValidChunkCount is a free data retrieval call binding the contract method 0x8a5e8e32.
//
// Solidity: function validChunkCount() view returns(uint256)
func (_Postagestamp *PostagestampCaller) ValidChunkCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Postagestamp.contract.Call(opts, &out, "validChunkCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidChunkCount is a free data retrieval call binding the contract method 0x8a5e8e32.
//
// Solidity: function validChunkCount() view returns(uint256)
func (_Postagestamp *PostagestampSession) ValidChunkCount() (*big.Int, error) {
	return _Postagestamp.Contract.ValidChunkCount(&_Postagestamp.CallOpts)
}

// ValidChunkCount is a free data retrieval call binding the contract method 0x8a5e8e32.
//
// Solidity: function validChunkCount() view returns(uint256)
func (_Postagestamp *PostagestampCallerSession) ValidChunkCount() (*big.Int, error) {
	return _Postagestamp.Contract.ValidChunkCount(&_Postagestamp.CallOpts)
}

// CopyBatch is a paid mutator transaction binding the contract method 0x18c8572f.
//
// Solidity: function copyBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _batchId, bool _immutable) returns()
func (_Postagestamp *PostagestampTransactor) CopyBatch(opts *bind.TransactOpts, _owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _batchId [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "copyBatch", _owner, _initialBalancePerChunk, _depth, _bucketDepth, _batchId, _immutable)
}

// CopyBatch is a paid mutator transaction binding the contract method 0x18c8572f.
//
// Solidity: function copyBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _batchId, bool _immutable) returns()
func (_Postagestamp *PostagestampSession) CopyBatch(_owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _batchId [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.Contract.CopyBatch(&_Postagestamp.TransactOpts, _owner, _initialBalancePerChunk, _depth, _bucketDepth, _batchId, _immutable)
}

// CopyBatch is a paid mutator transaction binding the contract method 0x18c8572f.
//
// Solidity: function copyBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _batchId, bool _immutable) returns()
func (_Postagestamp *PostagestampTransactorSession) CopyBatch(_owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _batchId [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.Contract.CopyBatch(&_Postagestamp.TransactOpts, _owner, _initialBalancePerChunk, _depth, _bucketDepth, _batchId, _immutable)
}

// CreateBatch is a paid mutator transaction binding the contract method 0x5239af71.
//
// Solidity: function createBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _nonce, bool _immutable) returns()
func (_Postagestamp *PostagestampTransactor) CreateBatch(opts *bind.TransactOpts, _owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _nonce [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "createBatch", _owner, _initialBalancePerChunk, _depth, _bucketDepth, _nonce, _immutable)
}

// CreateBatch is a paid mutator transaction binding the contract method 0x5239af71.
//
// Solidity: function createBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _nonce, bool _immutable) returns()
func (_Postagestamp *PostagestampSession) CreateBatch(_owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _nonce [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.Contract.CreateBatch(&_Postagestamp.TransactOpts, _owner, _initialBalancePerChunk, _depth, _bucketDepth, _nonce, _immutable)
}

// CreateBatch is a paid mutator transaction binding the contract method 0x5239af71.
//
// Solidity: function createBatch(address _owner, uint256 _initialBalancePerChunk, uint8 _depth, uint8 _bucketDepth, bytes32 _nonce, bool _immutable) returns()
func (_Postagestamp *PostagestampTransactorSession) CreateBatch(_owner common.Address, _initialBalancePerChunk *big.Int, _depth uint8, _bucketDepth uint8, _nonce [32]byte, _immutable bool) (*types.Transaction, error) {
	return _Postagestamp.Contract.CreateBatch(&_Postagestamp.TransactOpts, _owner, _initialBalancePerChunk, _depth, _bucketDepth, _nonce, _immutable)
}

// ExpireLimited is a paid mutator transaction binding the contract method 0x628de877.
//
// Solidity: function expireLimited(uint256 limit) returns()
func (_Postagestamp *PostagestampTransactor) ExpireLimited(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "expireLimited", limit)
}

// ExpireLimited is a paid mutator transaction binding the contract method 0x628de877.
//
// Solidity: function expireLimited(uint256 limit) returns()
func (_Postagestamp *PostagestampSession) ExpireLimited(limit *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.ExpireLimited(&_Postagestamp.TransactOpts, limit)
}

// ExpireLimited is a paid mutator transaction binding the contract method 0x628de877.
//
// Solidity: function expireLimited(uint256 limit) returns()
func (_Postagestamp *PostagestampTransactorSession) ExpireLimited(limit *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.ExpireLimited(&_Postagestamp.TransactOpts, limit)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.GrantRole(&_Postagestamp.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.GrantRole(&_Postagestamp.TransactOpts, role, account)
}

// IncreaseDepth is a paid mutator transaction binding the contract method 0x47aab79b.
//
// Solidity: function increaseDepth(bytes32 _batchId, uint8 _newDepth) returns()
func (_Postagestamp *PostagestampTransactor) IncreaseDepth(opts *bind.TransactOpts, _batchId [32]byte, _newDepth uint8) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "increaseDepth", _batchId, _newDepth)
}

// IncreaseDepth is a paid mutator transaction binding the contract method 0x47aab79b.
//
// Solidity: function increaseDepth(bytes32 _batchId, uint8 _newDepth) returns()
func (_Postagestamp *PostagestampSession) IncreaseDepth(_batchId [32]byte, _newDepth uint8) (*types.Transaction, error) {
	return _Postagestamp.Contract.IncreaseDepth(&_Postagestamp.TransactOpts, _batchId, _newDepth)
}

// IncreaseDepth is a paid mutator transaction binding the contract method 0x47aab79b.
//
// Solidity: function increaseDepth(bytes32 _batchId, uint8 _newDepth) returns()
func (_Postagestamp *PostagestampTransactorSession) IncreaseDepth(_batchId [32]byte, _newDepth uint8) (*types.Transaction, error) {
	return _Postagestamp.Contract.IncreaseDepth(&_Postagestamp.TransactOpts, _batchId, _newDepth)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Postagestamp *PostagestampTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Postagestamp *PostagestampSession) Pause() (*types.Transaction, error) {
	return _Postagestamp.Contract.Pause(&_Postagestamp.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Postagestamp *PostagestampTransactorSession) Pause() (*types.Transaction, error) {
	return _Postagestamp.Contract.Pause(&_Postagestamp.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.RenounceRole(&_Postagestamp.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.RenounceRole(&_Postagestamp.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.RevokeRole(&_Postagestamp.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Postagestamp *PostagestampTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.RevokeRole(&_Postagestamp.TransactOpts, role, account)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Postagestamp *PostagestampTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Postagestamp *PostagestampSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.SetPrice(&_Postagestamp.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Postagestamp *PostagestampTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.SetPrice(&_Postagestamp.TransactOpts, _price)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _batchId, uint256 _topupAmountPerChunk) returns()
func (_Postagestamp *PostagestampTransactor) TopUp(opts *bind.TransactOpts, _batchId [32]byte, _topupAmountPerChunk *big.Int) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "topUp", _batchId, _topupAmountPerChunk)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _batchId, uint256 _topupAmountPerChunk) returns()
func (_Postagestamp *PostagestampSession) TopUp(_batchId [32]byte, _topupAmountPerChunk *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.TopUp(&_Postagestamp.TransactOpts, _batchId, _topupAmountPerChunk)
}

// TopUp is a paid mutator transaction binding the contract method 0xb67644b9.
//
// Solidity: function topUp(bytes32 _batchId, uint256 _topupAmountPerChunk) returns()
func (_Postagestamp *PostagestampTransactorSession) TopUp(_batchId [32]byte, _topupAmountPerChunk *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.TopUp(&_Postagestamp.TransactOpts, _batchId, _topupAmountPerChunk)
}

// TopupPot is a paid mutator transaction binding the contract method 0x99176447.
//
// Solidity: function topupPot(uint256 amount) returns()
func (_Postagestamp *PostagestampTransactor) TopupPot(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "topupPot", amount)
}

// TopupPot is a paid mutator transaction binding the contract method 0x99176447.
//
// Solidity: function topupPot(uint256 amount) returns()
func (_Postagestamp *PostagestampSession) TopupPot(amount *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.TopupPot(&_Postagestamp.TransactOpts, amount)
}

// TopupPot is a paid mutator transaction binding the contract method 0x99176447.
//
// Solidity: function topupPot(uint256 amount) returns()
func (_Postagestamp *PostagestampTransactorSession) TopupPot(amount *big.Int) (*types.Transaction, error) {
	return _Postagestamp.Contract.TopupPot(&_Postagestamp.TransactOpts, amount)
}

// TotalPot is a paid mutator transaction binding the contract method 0x24b570a9.
//
// Solidity: function totalPot() returns(uint256)
func (_Postagestamp *PostagestampTransactor) TotalPot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "totalPot")
}

// TotalPot is a paid mutator transaction binding the contract method 0x24b570a9.
//
// Solidity: function totalPot() returns(uint256)
func (_Postagestamp *PostagestampSession) TotalPot() (*types.Transaction, error) {
	return _Postagestamp.Contract.TotalPot(&_Postagestamp.TransactOpts)
}

// TotalPot is a paid mutator transaction binding the contract method 0x24b570a9.
//
// Solidity: function totalPot() returns(uint256)
func (_Postagestamp *PostagestampTransactorSession) TotalPot() (*types.Transaction, error) {
	return _Postagestamp.Contract.TotalPot(&_Postagestamp.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Postagestamp *PostagestampTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Postagestamp *PostagestampSession) UnPause() (*types.Transaction, error) {
	return _Postagestamp.Contract.UnPause(&_Postagestamp.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Postagestamp *PostagestampTransactorSession) UnPause() (*types.Transaction, error) {
	return _Postagestamp.Contract.UnPause(&_Postagestamp.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address beneficiary) returns()
func (_Postagestamp *PostagestampTransactor) Withdraw(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Postagestamp.contract.Transact(opts, "withdraw", beneficiary)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address beneficiary) returns()
func (_Postagestamp *PostagestampSession) Withdraw(beneficiary common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.Withdraw(&_Postagestamp.TransactOpts, beneficiary)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address beneficiary) returns()
func (_Postagestamp *PostagestampTransactorSession) Withdraw(beneficiary common.Address) (*types.Transaction, error) {
	return _Postagestamp.Contract.Withdraw(&_Postagestamp.TransactOpts, beneficiary)
}

// PostagestampBatchCreatedIterator is returned from FilterBatchCreated and is used to iterate over the raw logs and unpacked data for BatchCreated events raised by the Postagestamp contract.
type PostagestampBatchCreatedIterator struct {
	Event *PostagestampBatchCreated // Event containing the contract specifics and raw log

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
func (it *PostagestampBatchCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampBatchCreated)
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
		it.Event = new(PostagestampBatchCreated)
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
func (it *PostagestampBatchCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampBatchCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampBatchCreated represents a BatchCreated event raised by the Postagestamp contract.
type PostagestampBatchCreated struct {
	BatchId           [32]byte
	TotalAmount       *big.Int
	NormalisedBalance *big.Int
	Owner             common.Address
	Depth             uint8
	BucketDepth       uint8
	ImmutableFlag     bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchCreated is a free log retrieval operation binding the contract event 0x9b088e2c89b322a3c1d81515e1c88db3d386d022926f0e2d0b9b5813b7413d58.
//
// Solidity: event BatchCreated(bytes32 indexed batchId, uint256 totalAmount, uint256 normalisedBalance, address owner, uint8 depth, uint8 bucketDepth, bool immutableFlag)
func (_Postagestamp *PostagestampFilterer) FilterBatchCreated(opts *bind.FilterOpts, batchId [][32]byte) (*PostagestampBatchCreatedIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "BatchCreated", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampBatchCreatedIterator{contract: _Postagestamp.contract, event: "BatchCreated", logs: logs, sub: sub}, nil
}

// WatchBatchCreated is a free log subscription operation binding the contract event 0x9b088e2c89b322a3c1d81515e1c88db3d386d022926f0e2d0b9b5813b7413d58.
//
// Solidity: event BatchCreated(bytes32 indexed batchId, uint256 totalAmount, uint256 normalisedBalance, address owner, uint8 depth, uint8 bucketDepth, bool immutableFlag)
func (_Postagestamp *PostagestampFilterer) WatchBatchCreated(opts *bind.WatchOpts, sink chan<- *PostagestampBatchCreated, batchId [][32]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "BatchCreated", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampBatchCreated)
				if err := _Postagestamp.contract.UnpackLog(event, "BatchCreated", log); err != nil {
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

// ParseBatchCreated is a log parse operation binding the contract event 0x9b088e2c89b322a3c1d81515e1c88db3d386d022926f0e2d0b9b5813b7413d58.
//
// Solidity: event BatchCreated(bytes32 indexed batchId, uint256 totalAmount, uint256 normalisedBalance, address owner, uint8 depth, uint8 bucketDepth, bool immutableFlag)
func (_Postagestamp *PostagestampFilterer) ParseBatchCreated(log types.Log) (*PostagestampBatchCreated, error) {
	event := new(PostagestampBatchCreated)
	if err := _Postagestamp.contract.UnpackLog(event, "BatchCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampBatchDepthIncreaseIterator is returned from FilterBatchDepthIncrease and is used to iterate over the raw logs and unpacked data for BatchDepthIncrease events raised by the Postagestamp contract.
type PostagestampBatchDepthIncreaseIterator struct {
	Event *PostagestampBatchDepthIncrease // Event containing the contract specifics and raw log

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
func (it *PostagestampBatchDepthIncreaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampBatchDepthIncrease)
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
		it.Event = new(PostagestampBatchDepthIncrease)
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
func (it *PostagestampBatchDepthIncreaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampBatchDepthIncreaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampBatchDepthIncrease represents a BatchDepthIncrease event raised by the Postagestamp contract.
type PostagestampBatchDepthIncrease struct {
	BatchId           [32]byte
	NewDepth          uint8
	NormalisedBalance *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchDepthIncrease is a free log retrieval operation binding the contract event 0xaf27998ec15e9d3809edad41aec1b5551d8412e71bd07c91611a0237ead1dc8e.
//
// Solidity: event BatchDepthIncrease(bytes32 indexed batchId, uint8 newDepth, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) FilterBatchDepthIncrease(opts *bind.FilterOpts, batchId [][32]byte) (*PostagestampBatchDepthIncreaseIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "BatchDepthIncrease", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampBatchDepthIncreaseIterator{contract: _Postagestamp.contract, event: "BatchDepthIncrease", logs: logs, sub: sub}, nil
}

// WatchBatchDepthIncrease is a free log subscription operation binding the contract event 0xaf27998ec15e9d3809edad41aec1b5551d8412e71bd07c91611a0237ead1dc8e.
//
// Solidity: event BatchDepthIncrease(bytes32 indexed batchId, uint8 newDepth, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) WatchBatchDepthIncrease(opts *bind.WatchOpts, sink chan<- *PostagestampBatchDepthIncrease, batchId [][32]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "BatchDepthIncrease", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampBatchDepthIncrease)
				if err := _Postagestamp.contract.UnpackLog(event, "BatchDepthIncrease", log); err != nil {
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

// ParseBatchDepthIncrease is a log parse operation binding the contract event 0xaf27998ec15e9d3809edad41aec1b5551d8412e71bd07c91611a0237ead1dc8e.
//
// Solidity: event BatchDepthIncrease(bytes32 indexed batchId, uint8 newDepth, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) ParseBatchDepthIncrease(log types.Log) (*PostagestampBatchDepthIncrease, error) {
	event := new(PostagestampBatchDepthIncrease)
	if err := _Postagestamp.contract.UnpackLog(event, "BatchDepthIncrease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampBatchTopUpIterator is returned from FilterBatchTopUp and is used to iterate over the raw logs and unpacked data for BatchTopUp events raised by the Postagestamp contract.
type PostagestampBatchTopUpIterator struct {
	Event *PostagestampBatchTopUp // Event containing the contract specifics and raw log

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
func (it *PostagestampBatchTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampBatchTopUp)
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
		it.Event = new(PostagestampBatchTopUp)
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
func (it *PostagestampBatchTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampBatchTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampBatchTopUp represents a BatchTopUp event raised by the Postagestamp contract.
type PostagestampBatchTopUp struct {
	BatchId           [32]byte
	TopupAmount       *big.Int
	NormalisedBalance *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchTopUp is a free log retrieval operation binding the contract event 0xaf5756c62d6c0722ef9be1f82bef97ab06ea5aea7f3eb8ad348422079f01d88d.
//
// Solidity: event BatchTopUp(bytes32 indexed batchId, uint256 topupAmount, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) FilterBatchTopUp(opts *bind.FilterOpts, batchId [][32]byte) (*PostagestampBatchTopUpIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "BatchTopUp", batchIdRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampBatchTopUpIterator{contract: _Postagestamp.contract, event: "BatchTopUp", logs: logs, sub: sub}, nil
}

// WatchBatchTopUp is a free log subscription operation binding the contract event 0xaf5756c62d6c0722ef9be1f82bef97ab06ea5aea7f3eb8ad348422079f01d88d.
//
// Solidity: event BatchTopUp(bytes32 indexed batchId, uint256 topupAmount, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) WatchBatchTopUp(opts *bind.WatchOpts, sink chan<- *PostagestampBatchTopUp, batchId [][32]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "BatchTopUp", batchIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampBatchTopUp)
				if err := _Postagestamp.contract.UnpackLog(event, "BatchTopUp", log); err != nil {
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

// ParseBatchTopUp is a log parse operation binding the contract event 0xaf5756c62d6c0722ef9be1f82bef97ab06ea5aea7f3eb8ad348422079f01d88d.
//
// Solidity: event BatchTopUp(bytes32 indexed batchId, uint256 topupAmount, uint256 normalisedBalance)
func (_Postagestamp *PostagestampFilterer) ParseBatchTopUp(log types.Log) (*PostagestampBatchTopUp, error) {
	event := new(PostagestampBatchTopUp)
	if err := _Postagestamp.contract.UnpackLog(event, "BatchTopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Postagestamp contract.
type PostagestampPausedIterator struct {
	Event *PostagestampPaused // Event containing the contract specifics and raw log

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
func (it *PostagestampPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampPaused)
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
		it.Event = new(PostagestampPaused)
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
func (it *PostagestampPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampPaused represents a Paused event raised by the Postagestamp contract.
type PostagestampPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Postagestamp *PostagestampFilterer) FilterPaused(opts *bind.FilterOpts) (*PostagestampPausedIterator, error) {

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PostagestampPausedIterator{contract: _Postagestamp.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Postagestamp *PostagestampFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PostagestampPaused) (event.Subscription, error) {

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampPaused)
				if err := _Postagestamp.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Postagestamp *PostagestampFilterer) ParsePaused(log types.Log) (*PostagestampPaused, error) {
	event := new(PostagestampPaused)
	if err := _Postagestamp.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampPriceUpdateIterator is returned from FilterPriceUpdate and is used to iterate over the raw logs and unpacked data for PriceUpdate events raised by the Postagestamp contract.
type PostagestampPriceUpdateIterator struct {
	Event *PostagestampPriceUpdate // Event containing the contract specifics and raw log

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
func (it *PostagestampPriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampPriceUpdate)
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
		it.Event = new(PostagestampPriceUpdate)
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
func (it *PostagestampPriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampPriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampPriceUpdate represents a PriceUpdate event raised by the Postagestamp contract.
type PostagestampPriceUpdate struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdate is a free log retrieval operation binding the contract event 0xae46785019700e30375a5d7b4f91e32f8060ef085111f896ebf889450aa2ab5a.
//
// Solidity: event PriceUpdate(uint256 price)
func (_Postagestamp *PostagestampFilterer) FilterPriceUpdate(opts *bind.FilterOpts) (*PostagestampPriceUpdateIterator, error) {

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return &PostagestampPriceUpdateIterator{contract: _Postagestamp.contract, event: "PriceUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceUpdate is a free log subscription operation binding the contract event 0xae46785019700e30375a5d7b4f91e32f8060ef085111f896ebf889450aa2ab5a.
//
// Solidity: event PriceUpdate(uint256 price)
func (_Postagestamp *PostagestampFilterer) WatchPriceUpdate(opts *bind.WatchOpts, sink chan<- *PostagestampPriceUpdate) (event.Subscription, error) {

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampPriceUpdate)
				if err := _Postagestamp.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
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

// ParsePriceUpdate is a log parse operation binding the contract event 0xae46785019700e30375a5d7b4f91e32f8060ef085111f896ebf889450aa2ab5a.
//
// Solidity: event PriceUpdate(uint256 price)
func (_Postagestamp *PostagestampFilterer) ParsePriceUpdate(log types.Log) (*PostagestampPriceUpdate, error) {
	event := new(PostagestampPriceUpdate)
	if err := _Postagestamp.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Postagestamp contract.
type PostagestampRoleAdminChangedIterator struct {
	Event *PostagestampRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PostagestampRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampRoleAdminChanged)
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
		it.Event = new(PostagestampRoleAdminChanged)
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
func (it *PostagestampRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampRoleAdminChanged represents a RoleAdminChanged event raised by the Postagestamp contract.
type PostagestampRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Postagestamp *PostagestampFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PostagestampRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampRoleAdminChangedIterator{contract: _Postagestamp.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Postagestamp *PostagestampFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PostagestampRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampRoleAdminChanged)
				if err := _Postagestamp.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Postagestamp *PostagestampFilterer) ParseRoleAdminChanged(log types.Log) (*PostagestampRoleAdminChanged, error) {
	event := new(PostagestampRoleAdminChanged)
	if err := _Postagestamp.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Postagestamp contract.
type PostagestampRoleGrantedIterator struct {
	Event *PostagestampRoleGranted // Event containing the contract specifics and raw log

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
func (it *PostagestampRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampRoleGranted)
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
		it.Event = new(PostagestampRoleGranted)
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
func (it *PostagestampRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampRoleGranted represents a RoleGranted event raised by the Postagestamp contract.
type PostagestampRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PostagestampRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampRoleGrantedIterator{contract: _Postagestamp.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PostagestampRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampRoleGranted)
				if err := _Postagestamp.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) ParseRoleGranted(log types.Log) (*PostagestampRoleGranted, error) {
	event := new(PostagestampRoleGranted)
	if err := _Postagestamp.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Postagestamp contract.
type PostagestampRoleRevokedIterator struct {
	Event *PostagestampRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PostagestampRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampRoleRevoked)
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
		it.Event = new(PostagestampRoleRevoked)
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
func (it *PostagestampRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampRoleRevoked represents a RoleRevoked event raised by the Postagestamp contract.
type PostagestampRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PostagestampRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PostagestampRoleRevokedIterator{contract: _Postagestamp.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PostagestampRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampRoleRevoked)
				if err := _Postagestamp.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Postagestamp *PostagestampFilterer) ParseRoleRevoked(log types.Log) (*PostagestampRoleRevoked, error) {
	event := new(PostagestampRoleRevoked)
	if err := _Postagestamp.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PostagestampUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Postagestamp contract.
type PostagestampUnpausedIterator struct {
	Event *PostagestampUnpaused // Event containing the contract specifics and raw log

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
func (it *PostagestampUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostagestampUnpaused)
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
		it.Event = new(PostagestampUnpaused)
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
func (it *PostagestampUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostagestampUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostagestampUnpaused represents a Unpaused event raised by the Postagestamp contract.
type PostagestampUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Postagestamp *PostagestampFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PostagestampUnpausedIterator, error) {

	logs, sub, err := _Postagestamp.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PostagestampUnpausedIterator{contract: _Postagestamp.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Postagestamp *PostagestampFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PostagestampUnpaused) (event.Subscription, error) {

	logs, sub, err := _Postagestamp.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostagestampUnpaused)
				if err := _Postagestamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Postagestamp *PostagestampFilterer) ParseUnpaused(log types.Log) (*PostagestampUnpaused, error) {
	event := new(PostagestampUnpaused)
	if err := _Postagestamp.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
