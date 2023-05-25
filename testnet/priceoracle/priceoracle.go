// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package priceoracle

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

// PriceoracleMetaData contains all meta data concerning the Priceoracle contract.
var PriceoracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_postageStamp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_UPDATER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redundancy\",\"type\":\"uint256\"}],\"name\":\"adjustPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"increaseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postageStamp\",\"outputs\":[{\"internalType\":\"contractPostageStamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PriceoracleABI is the input ABI used to generate the binding from.
// Deprecated: Use PriceoracleMetaData.ABI instead.
var PriceoracleABI = PriceoracleMetaData.ABI

// Priceoracle is an auto generated Go binding around an Ethereum contract.
type Priceoracle struct {
	PriceoracleCaller     // Read-only binding to the contract
	PriceoracleTransactor // Write-only binding to the contract
	PriceoracleFilterer   // Log filterer for contract events
}

// PriceoracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceoracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceoracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceoracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceoracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceoracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceoracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceoracleSession struct {
	Contract     *Priceoracle      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriceoracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceoracleCallerSession struct {
	Contract *PriceoracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PriceoracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceoracleTransactorSession struct {
	Contract     *PriceoracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PriceoracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceoracleRaw struct {
	Contract *Priceoracle // Generic contract binding to access the raw methods on
}

// PriceoracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceoracleCallerRaw struct {
	Contract *PriceoracleCaller // Generic read-only contract binding to access the raw methods on
}

// PriceoracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceoracleTransactorRaw struct {
	Contract *PriceoracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceoracle creates a new instance of Priceoracle, bound to a specific deployed contract.
func NewPriceoracle(address common.Address, backend bind.ContractBackend) (*Priceoracle, error) {
	contract, err := bindPriceoracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Priceoracle{PriceoracleCaller: PriceoracleCaller{contract: contract}, PriceoracleTransactor: PriceoracleTransactor{contract: contract}, PriceoracleFilterer: PriceoracleFilterer{contract: contract}}, nil
}

// NewPriceoracleCaller creates a new read-only instance of Priceoracle, bound to a specific deployed contract.
func NewPriceoracleCaller(address common.Address, caller bind.ContractCaller) (*PriceoracleCaller, error) {
	contract, err := bindPriceoracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceoracleCaller{contract: contract}, nil
}

// NewPriceoracleTransactor creates a new write-only instance of Priceoracle, bound to a specific deployed contract.
func NewPriceoracleTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceoracleTransactor, error) {
	contract, err := bindPriceoracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceoracleTransactor{contract: contract}, nil
}

// NewPriceoracleFilterer creates a new log filterer instance of Priceoracle, bound to a specific deployed contract.
func NewPriceoracleFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceoracleFilterer, error) {
	contract, err := bindPriceoracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceoracleFilterer{contract: contract}, nil
}

// bindPriceoracle binds a generic wrapper to an already deployed contract.
func bindPriceoracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PriceoracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Priceoracle *PriceoracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Priceoracle.Contract.PriceoracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Priceoracle *PriceoracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoracle.Contract.PriceoracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Priceoracle *PriceoracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Priceoracle.Contract.PriceoracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Priceoracle *PriceoracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Priceoracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Priceoracle *PriceoracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Priceoracle *PriceoracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Priceoracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Priceoracle.Contract.DEFAULTADMINROLE(&_Priceoracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Priceoracle.Contract.DEFAULTADMINROLE(&_Priceoracle.CallOpts)
}

// PRICEUPDATERROLE is a free data retrieval call binding the contract method 0xfb8d8101.
//
// Solidity: function PRICE_UPDATER_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleCaller) PRICEUPDATERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "PRICE_UPDATER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PRICEUPDATERROLE is a free data retrieval call binding the contract method 0xfb8d8101.
//
// Solidity: function PRICE_UPDATER_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleSession) PRICEUPDATERROLE() ([32]byte, error) {
	return _Priceoracle.Contract.PRICEUPDATERROLE(&_Priceoracle.CallOpts)
}

// PRICEUPDATERROLE is a free data retrieval call binding the contract method 0xfb8d8101.
//
// Solidity: function PRICE_UPDATER_ROLE() view returns(bytes32)
func (_Priceoracle *PriceoracleCallerSession) PRICEUPDATERROLE() ([32]byte, error) {
	return _Priceoracle.Contract.PRICEUPDATERROLE(&_Priceoracle.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Priceoracle *PriceoracleCaller) CurrentPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "currentPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Priceoracle *PriceoracleSession) CurrentPrice() (*big.Int, error) {
	return _Priceoracle.Contract.CurrentPrice(&_Priceoracle.CallOpts)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x9d1b464a.
//
// Solidity: function currentPrice() view returns(uint256)
func (_Priceoracle *PriceoracleCallerSession) CurrentPrice() (*big.Int, error) {
	return _Priceoracle.Contract.CurrentPrice(&_Priceoracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Priceoracle *PriceoracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Priceoracle *PriceoracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Priceoracle.Contract.GetRoleAdmin(&_Priceoracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Priceoracle *PriceoracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Priceoracle.Contract.GetRoleAdmin(&_Priceoracle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Priceoracle *PriceoracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Priceoracle *PriceoracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Priceoracle.Contract.HasRole(&_Priceoracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Priceoracle *PriceoracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Priceoracle.Contract.HasRole(&_Priceoracle.CallOpts, role, account)
}

// IncreaseRate is a free data retrieval call binding the contract method 0x0080248f.
//
// Solidity: function increaseRate(uint256 ) view returns(uint256)
func (_Priceoracle *PriceoracleCaller) IncreaseRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "increaseRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreaseRate is a free data retrieval call binding the contract method 0x0080248f.
//
// Solidity: function increaseRate(uint256 ) view returns(uint256)
func (_Priceoracle *PriceoracleSession) IncreaseRate(arg0 *big.Int) (*big.Int, error) {
	return _Priceoracle.Contract.IncreaseRate(&_Priceoracle.CallOpts, arg0)
}

// IncreaseRate is a free data retrieval call binding the contract method 0x0080248f.
//
// Solidity: function increaseRate(uint256 ) view returns(uint256)
func (_Priceoracle *PriceoracleCallerSession) IncreaseRate(arg0 *big.Int) (*big.Int, error) {
	return _Priceoracle.Contract.IncreaseRate(&_Priceoracle.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Priceoracle *PriceoracleCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Priceoracle *PriceoracleSession) IsPaused() (bool, error) {
	return _Priceoracle.Contract.IsPaused(&_Priceoracle.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Priceoracle *PriceoracleCallerSession) IsPaused() (bool, error) {
	return _Priceoracle.Contract.IsPaused(&_Priceoracle.CallOpts)
}

// MinimumPrice is a free data retrieval call binding the contract method 0x7f386b6c.
//
// Solidity: function minimumPrice() view returns(uint256)
func (_Priceoracle *PriceoracleCaller) MinimumPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "minimumPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumPrice is a free data retrieval call binding the contract method 0x7f386b6c.
//
// Solidity: function minimumPrice() view returns(uint256)
func (_Priceoracle *PriceoracleSession) MinimumPrice() (*big.Int, error) {
	return _Priceoracle.Contract.MinimumPrice(&_Priceoracle.CallOpts)
}

// MinimumPrice is a free data retrieval call binding the contract method 0x7f386b6c.
//
// Solidity: function minimumPrice() view returns(uint256)
func (_Priceoracle *PriceoracleCallerSession) MinimumPrice() (*big.Int, error) {
	return _Priceoracle.Contract.MinimumPrice(&_Priceoracle.CallOpts)
}

// PostageStamp is a free data retrieval call binding the contract method 0xe0632c64.
//
// Solidity: function postageStamp() view returns(address)
func (_Priceoracle *PriceoracleCaller) PostageStamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "postageStamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PostageStamp is a free data retrieval call binding the contract method 0xe0632c64.
//
// Solidity: function postageStamp() view returns(address)
func (_Priceoracle *PriceoracleSession) PostageStamp() (common.Address, error) {
	return _Priceoracle.Contract.PostageStamp(&_Priceoracle.CallOpts)
}

// PostageStamp is a free data retrieval call binding the contract method 0xe0632c64.
//
// Solidity: function postageStamp() view returns(address)
func (_Priceoracle *PriceoracleCallerSession) PostageStamp() (common.Address, error) {
	return _Priceoracle.Contract.PostageStamp(&_Priceoracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Priceoracle *PriceoracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Priceoracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Priceoracle *PriceoracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Priceoracle.Contract.SupportsInterface(&_Priceoracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Priceoracle *PriceoracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Priceoracle.Contract.SupportsInterface(&_Priceoracle.CallOpts, interfaceId)
}

// AdjustPrice is a paid mutator transaction binding the contract method 0x72bf079e.
//
// Solidity: function adjustPrice(uint256 redundancy) returns()
func (_Priceoracle *PriceoracleTransactor) AdjustPrice(opts *bind.TransactOpts, redundancy *big.Int) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "adjustPrice", redundancy)
}

// AdjustPrice is a paid mutator transaction binding the contract method 0x72bf079e.
//
// Solidity: function adjustPrice(uint256 redundancy) returns()
func (_Priceoracle *PriceoracleSession) AdjustPrice(redundancy *big.Int) (*types.Transaction, error) {
	return _Priceoracle.Contract.AdjustPrice(&_Priceoracle.TransactOpts, redundancy)
}

// AdjustPrice is a paid mutator transaction binding the contract method 0x72bf079e.
//
// Solidity: function adjustPrice(uint256 redundancy) returns()
func (_Priceoracle *PriceoracleTransactorSession) AdjustPrice(redundancy *big.Int) (*types.Transaction, error) {
	return _Priceoracle.Contract.AdjustPrice(&_Priceoracle.TransactOpts, redundancy)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.GrantRole(&_Priceoracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.GrantRole(&_Priceoracle.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Priceoracle *PriceoracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Priceoracle *PriceoracleSession) Pause() (*types.Transaction, error) {
	return _Priceoracle.Contract.Pause(&_Priceoracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Priceoracle *PriceoracleTransactorSession) Pause() (*types.Transaction, error) {
	return _Priceoracle.Contract.Pause(&_Priceoracle.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.RenounceRole(&_Priceoracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.RenounceRole(&_Priceoracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.RevokeRole(&_Priceoracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Priceoracle *PriceoracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Priceoracle.Contract.RevokeRole(&_Priceoracle.TransactOpts, role, account)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Priceoracle *PriceoracleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Priceoracle *PriceoracleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Priceoracle.Contract.SetPrice(&_Priceoracle.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns()
func (_Priceoracle *PriceoracleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _Priceoracle.Contract.SetPrice(&_Priceoracle.TransactOpts, _price)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Priceoracle *PriceoracleTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Priceoracle.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Priceoracle *PriceoracleSession) UnPause() (*types.Transaction, error) {
	return _Priceoracle.Contract.UnPause(&_Priceoracle.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Priceoracle *PriceoracleTransactorSession) UnPause() (*types.Transaction, error) {
	return _Priceoracle.Contract.UnPause(&_Priceoracle.TransactOpts)
}

// PriceoraclePriceUpdateIterator is returned from FilterPriceUpdate and is used to iterate over the raw logs and unpacked data for PriceUpdate events raised by the Priceoracle contract.
type PriceoraclePriceUpdateIterator struct {
	Event *PriceoraclePriceUpdate // Event containing the contract specifics and raw log

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
func (it *PriceoraclePriceUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceoraclePriceUpdate)
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
		it.Event = new(PriceoraclePriceUpdate)
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
func (it *PriceoraclePriceUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceoraclePriceUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceoraclePriceUpdate represents a PriceUpdate event raised by the Priceoracle contract.
type PriceoraclePriceUpdate struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdate is a free log retrieval operation binding the contract event 0xae46785019700e30375a5d7b4f91e32f8060ef085111f896ebf889450aa2ab5a.
//
// Solidity: event PriceUpdate(uint256 price)
func (_Priceoracle *PriceoracleFilterer) FilterPriceUpdate(opts *bind.FilterOpts) (*PriceoraclePriceUpdateIterator, error) {

	logs, sub, err := _Priceoracle.contract.FilterLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return &PriceoraclePriceUpdateIterator{contract: _Priceoracle.contract, event: "PriceUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceUpdate is a free log subscription operation binding the contract event 0xae46785019700e30375a5d7b4f91e32f8060ef085111f896ebf889450aa2ab5a.
//
// Solidity: event PriceUpdate(uint256 price)
func (_Priceoracle *PriceoracleFilterer) WatchPriceUpdate(opts *bind.WatchOpts, sink chan<- *PriceoraclePriceUpdate) (event.Subscription, error) {

	logs, sub, err := _Priceoracle.contract.WatchLogs(opts, "PriceUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceoraclePriceUpdate)
				if err := _Priceoracle.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
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
func (_Priceoracle *PriceoracleFilterer) ParsePriceUpdate(log types.Log) (*PriceoraclePriceUpdate, error) {
	event := new(PriceoraclePriceUpdate)
	if err := _Priceoracle.contract.UnpackLog(event, "PriceUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceoracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Priceoracle contract.
type PriceoracleRoleAdminChangedIterator struct {
	Event *PriceoracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PriceoracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceoracleRoleAdminChanged)
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
		it.Event = new(PriceoracleRoleAdminChanged)
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
func (it *PriceoracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceoracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceoracleRoleAdminChanged represents a RoleAdminChanged event raised by the Priceoracle contract.
type PriceoracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Priceoracle *PriceoracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PriceoracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Priceoracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PriceoracleRoleAdminChangedIterator{contract: _Priceoracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Priceoracle *PriceoracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PriceoracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Priceoracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceoracleRoleAdminChanged)
				if err := _Priceoracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Priceoracle *PriceoracleFilterer) ParseRoleAdminChanged(log types.Log) (*PriceoracleRoleAdminChanged, error) {
	event := new(PriceoracleRoleAdminChanged)
	if err := _Priceoracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceoracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Priceoracle contract.
type PriceoracleRoleGrantedIterator struct {
	Event *PriceoracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *PriceoracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceoracleRoleGranted)
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
		it.Event = new(PriceoracleRoleGranted)
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
func (it *PriceoracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceoracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceoracleRoleGranted represents a RoleGranted event raised by the Priceoracle contract.
type PriceoracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Priceoracle *PriceoracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceoracleRoleGrantedIterator, error) {

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

	logs, sub, err := _Priceoracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceoracleRoleGrantedIterator{contract: _Priceoracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Priceoracle *PriceoracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PriceoracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Priceoracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceoracleRoleGranted)
				if err := _Priceoracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Priceoracle *PriceoracleFilterer) ParseRoleGranted(log types.Log) (*PriceoracleRoleGranted, error) {
	event := new(PriceoracleRoleGranted)
	if err := _Priceoracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PriceoracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Priceoracle contract.
type PriceoracleRoleRevokedIterator struct {
	Event *PriceoracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PriceoracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceoracleRoleRevoked)
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
		it.Event = new(PriceoracleRoleRevoked)
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
func (it *PriceoracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceoracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceoracleRoleRevoked represents a RoleRevoked event raised by the Priceoracle contract.
type PriceoracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Priceoracle *PriceoracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PriceoracleRoleRevokedIterator, error) {

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

	logs, sub, err := _Priceoracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PriceoracleRoleRevokedIterator{contract: _Priceoracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Priceoracle *PriceoracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PriceoracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Priceoracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceoracleRoleRevoked)
				if err := _Priceoracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Priceoracle *PriceoracleFilterer) ParseRoleRevoked(log types.Log) (*PriceoracleRoleRevoked, error) {
	event := new(PriceoracleRoleRevoked)
	if err := _Priceoracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
