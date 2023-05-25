// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bzzToken\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_NetworkId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"multisig\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"slashed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StakeFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"slashed\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"name\":\"StakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDISTRIBUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bzzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"freezeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"}],\"name\":\"lastUpdatedBlockNumberOfOverlay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"}],\"name\":\"ownerOfOverlay\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"}],\"name\":\"stakeOfOverlay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"}],\"name\":\"usableStakeOfOverlay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"overlay\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Staking.Contract.DEFAULTADMINROLE(&_Staking.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingSession) PAUSERROLE() ([32]byte, error) {
	return _Staking.Contract.PAUSERROLE(&_Staking.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Staking.Contract.PAUSERROLE(&_Staking.CallOpts)
}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Staking *StakingCaller) REDISTRIBUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "REDISTRIBUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Staking *StakingSession) REDISTRIBUTORROLE() ([32]byte, error) {
	return _Staking.Contract.REDISTRIBUTORROLE(&_Staking.CallOpts)
}

// REDISTRIBUTORROLE is a free data retrieval call binding the contract method 0xa6471a1d.
//
// Solidity: function REDISTRIBUTOR_ROLE() view returns(bytes32)
func (_Staking *StakingCallerSession) REDISTRIBUTORROLE() ([32]byte, error) {
	return _Staking.Contract.REDISTRIBUTORROLE(&_Staking.CallOpts)
}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Staking *StakingCaller) BzzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "bzzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Staking *StakingSession) BzzToken() (common.Address, error) {
	return _Staking.Contract.BzzToken(&_Staking.CallOpts)
}

// BzzToken is a free data retrieval call binding the contract method 0x420fc4db.
//
// Solidity: function bzzToken() view returns(address)
func (_Staking *StakingCallerSession) BzzToken() (common.Address, error) {
	return _Staking.Contract.BzzToken(&_Staking.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Staking *StakingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Staking.Contract.GetRoleAdmin(&_Staking.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Staking *StakingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Staking.Contract.HasRole(&_Staking.CallOpts, role, account)
}

// LastUpdatedBlockNumberOfOverlay is a free data retrieval call binding the contract method 0xede41302.
//
// Solidity: function lastUpdatedBlockNumberOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCaller) LastUpdatedBlockNumberOfOverlay(opts *bind.CallOpts, overlay [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "lastUpdatedBlockNumberOfOverlay", overlay)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedBlockNumberOfOverlay is a free data retrieval call binding the contract method 0xede41302.
//
// Solidity: function lastUpdatedBlockNumberOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingSession) LastUpdatedBlockNumberOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.LastUpdatedBlockNumberOfOverlay(&_Staking.CallOpts, overlay)
}

// LastUpdatedBlockNumberOfOverlay is a free data retrieval call binding the contract method 0xede41302.
//
// Solidity: function lastUpdatedBlockNumberOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCallerSession) LastUpdatedBlockNumberOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.LastUpdatedBlockNumberOfOverlay(&_Staking.CallOpts, overlay)
}

// OwnerOfOverlay is a free data retrieval call binding the contract method 0xa0d22b21.
//
// Solidity: function ownerOfOverlay(bytes32 overlay) view returns(address)
func (_Staking *StakingCaller) OwnerOfOverlay(opts *bind.CallOpts, overlay [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "ownerOfOverlay", overlay)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOfOverlay is a free data retrieval call binding the contract method 0xa0d22b21.
//
// Solidity: function ownerOfOverlay(bytes32 overlay) view returns(address)
func (_Staking *StakingSession) OwnerOfOverlay(overlay [32]byte) (common.Address, error) {
	return _Staking.Contract.OwnerOfOverlay(&_Staking.CallOpts, overlay)
}

// OwnerOfOverlay is a free data retrieval call binding the contract method 0xa0d22b21.
//
// Solidity: function ownerOfOverlay(bytes32 overlay) view returns(address)
func (_Staking *StakingCallerSession) OwnerOfOverlay(overlay [32]byte) (common.Address, error) {
	return _Staking.Contract.OwnerOfOverlay(&_Staking.CallOpts, overlay)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// StakeOfOverlay is a free data retrieval call binding the contract method 0x48962b93.
//
// Solidity: function stakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCaller) StakeOfOverlay(opts *bind.CallOpts, overlay [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakeOfOverlay", overlay)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeOfOverlay is a free data retrieval call binding the contract method 0x48962b93.
//
// Solidity: function stakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingSession) StakeOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.StakeOfOverlay(&_Staking.CallOpts, overlay)
}

// StakeOfOverlay is a free data retrieval call binding the contract method 0x48962b93.
//
// Solidity: function stakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCallerSession) StakeOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.StakeOfOverlay(&_Staking.CallOpts, overlay)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(bytes32 overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlockNumber, bool isValue)
func (_Staking *StakingCaller) Stakes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Overlay                [32]byte
	StakeAmount            *big.Int
	Owner                  common.Address
	LastUpdatedBlockNumber *big.Int
	IsValue                bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Overlay                [32]byte
		StakeAmount            *big.Int
		Owner                  common.Address
		LastUpdatedBlockNumber *big.Int
		IsValue                bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Overlay = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.StakeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.LastUpdatedBlockNumber = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsValue = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(bytes32 overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlockNumber, bool isValue)
func (_Staking *StakingSession) Stakes(arg0 [32]byte) (struct {
	Overlay                [32]byte
	StakeAmount            *big.Int
	Owner                  common.Address
	LastUpdatedBlockNumber *big.Int
	IsValue                bool
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x8fee6407.
//
// Solidity: function stakes(bytes32 ) view returns(bytes32 overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlockNumber, bool isValue)
func (_Staking *StakingCallerSession) Stakes(arg0 [32]byte) (struct {
	Overlay                [32]byte
	StakeAmount            *big.Int
	Owner                  common.Address
	LastUpdatedBlockNumber *big.Int
	IsValue                bool
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Staking *StakingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Staking.Contract.SupportsInterface(&_Staking.CallOpts, interfaceId)
}

// UsableStakeOfOverlay is a free data retrieval call binding the contract method 0xabe38543.
//
// Solidity: function usableStakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCaller) UsableStakeOfOverlay(opts *bind.CallOpts, overlay [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "usableStakeOfOverlay", overlay)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsableStakeOfOverlay is a free data retrieval call binding the contract method 0xabe38543.
//
// Solidity: function usableStakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingSession) UsableStakeOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.UsableStakeOfOverlay(&_Staking.CallOpts, overlay)
}

// UsableStakeOfOverlay is a free data retrieval call binding the contract method 0xabe38543.
//
// Solidity: function usableStakeOfOverlay(bytes32 overlay) view returns(uint256)
func (_Staking *StakingCallerSession) UsableStakeOfOverlay(overlay [32]byte) (*big.Int, error) {
	return _Staking.Contract.UsableStakeOfOverlay(&_Staking.CallOpts, overlay)
}

// DepositStake is a paid mutator transaction binding the contract method 0x1ed2cd40.
//
// Solidity: function depositStake(address _owner, bytes32 nonce, uint256 amount) returns()
func (_Staking *StakingTransactor) DepositStake(opts *bind.TransactOpts, _owner common.Address, nonce [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "depositStake", _owner, nonce, amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0x1ed2cd40.
//
// Solidity: function depositStake(address _owner, bytes32 nonce, uint256 amount) returns()
func (_Staking *StakingSession) DepositStake(_owner common.Address, nonce [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DepositStake(&_Staking.TransactOpts, _owner, nonce, amount)
}

// DepositStake is a paid mutator transaction binding the contract method 0x1ed2cd40.
//
// Solidity: function depositStake(address _owner, bytes32 nonce, uint256 amount) returns()
func (_Staking *StakingTransactorSession) DepositStake(_owner common.Address, nonce [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DepositStake(&_Staking.TransactOpts, _owner, nonce, amount)
}

// FreezeDeposit is a paid mutator transaction binding the contract method 0x837fd16a.
//
// Solidity: function freezeDeposit(bytes32 overlay, uint256 time) returns()
func (_Staking *StakingTransactor) FreezeDeposit(opts *bind.TransactOpts, overlay [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "freezeDeposit", overlay, time)
}

// FreezeDeposit is a paid mutator transaction binding the contract method 0x837fd16a.
//
// Solidity: function freezeDeposit(bytes32 overlay, uint256 time) returns()
func (_Staking *StakingSession) FreezeDeposit(overlay [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FreezeDeposit(&_Staking.TransactOpts, overlay, time)
}

// FreezeDeposit is a paid mutator transaction binding the contract method 0x837fd16a.
//
// Solidity: function freezeDeposit(bytes32 overlay, uint256 time) returns()
func (_Staking *StakingTransactorSession) FreezeDeposit(overlay [32]byte, time *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.FreezeDeposit(&_Staking.TransactOpts, overlay, time)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GrantRole(&_Staking.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RenounceRole(&_Staking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Staking *StakingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RevokeRole(&_Staking.TransactOpts, role, account)
}

// SlashDeposit is a paid mutator transaction binding the contract method 0xa6ce31d4.
//
// Solidity: function slashDeposit(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingTransactor) SlashDeposit(opts *bind.TransactOpts, overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slashDeposit", overlay, amount)
}

// SlashDeposit is a paid mutator transaction binding the contract method 0xa6ce31d4.
//
// Solidity: function slashDeposit(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingSession) SlashDeposit(overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SlashDeposit(&_Staking.TransactOpts, overlay, amount)
}

// SlashDeposit is a paid mutator transaction binding the contract method 0xa6ce31d4.
//
// Solidity: function slashDeposit(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingTransactorSession) SlashDeposit(overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SlashDeposit(&_Staking.TransactOpts, overlay, amount)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Staking *StakingTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Staking *StakingSession) UnPause() (*types.Transaction, error) {
	return _Staking.Contract.UnPause(&_Staking.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Staking *StakingTransactorSession) UnPause() (*types.Transaction, error) {
	return _Staking.Contract.UnPause(&_Staking.TransactOpts)
}

// WithdrawFromStake is a paid mutator transaction binding the contract method 0xe34c4527.
//
// Solidity: function withdrawFromStake(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingTransactor) WithdrawFromStake(opts *bind.TransactOpts, overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdrawFromStake", overlay, amount)
}

// WithdrawFromStake is a paid mutator transaction binding the contract method 0xe34c4527.
//
// Solidity: function withdrawFromStake(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingSession) WithdrawFromStake(overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawFromStake(&_Staking.TransactOpts, overlay, amount)
}

// WithdrawFromStake is a paid mutator transaction binding the contract method 0xe34c4527.
//
// Solidity: function withdrawFromStake(bytes32 overlay, uint256 amount) returns()
func (_Staking *StakingTransactorSession) WithdrawFromStake(overlay [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.WithdrawFromStake(&_Staking.TransactOpts, overlay, amount)
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Staking contract.
type StakingRoleAdminChangedIterator struct {
	Event *StakingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleAdminChanged)
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
		it.Event = new(StakingRoleAdminChanged)
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
func (it *StakingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleAdminChanged represents a RoleAdminChanged event raised by the Staking contract.
type StakingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleAdminChangedIterator{contract: _Staking.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Staking *StakingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleAdminChanged)
				if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleAdminChanged(log types.Log) (*StakingRoleAdminChanged, error) {
	event := new(StakingRoleAdminChanged)
	if err := _Staking.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Staking contract.
type StakingRoleGrantedIterator struct {
	Event *StakingRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleGranted)
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
		it.Event = new(StakingRoleGranted)
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
func (it *StakingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleGranted represents a RoleGranted event raised by the Staking contract.
type StakingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleGrantedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleGrantedIterator{contract: _Staking.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleGranted)
				if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleGranted(log types.Log) (*StakingRoleGranted, error) {
	event := new(StakingRoleGranted)
	if err := _Staking.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Staking contract.
type StakingRoleRevokedIterator struct {
	Event *StakingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRoleRevoked)
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
		it.Event = new(StakingRoleRevoked)
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
func (it *StakingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRoleRevoked represents a RoleRevoked event raised by the Staking contract.
type StakingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingRoleRevokedIterator, error) {

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

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingRoleRevokedIterator{contract: _Staking.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Staking *StakingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRoleRevoked)
				if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Staking *StakingFilterer) ParseRoleRevoked(log types.Log) (*StakingRoleRevoked, error) {
	event := new(StakingRoleRevoked)
	if err := _Staking.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeFrozenIterator is returned from FilterStakeFrozen and is used to iterate over the raw logs and unpacked data for StakeFrozen events raised by the Staking contract.
type StakingStakeFrozenIterator struct {
	Event *StakingStakeFrozen // Event containing the contract specifics and raw log

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
func (it *StakingStakeFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeFrozen)
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
		it.Event = new(StakingStakeFrozen)
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
func (it *StakingStakeFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeFrozen represents a StakeFrozen event raised by the Staking contract.
type StakingStakeFrozen struct {
	Slashed [32]byte
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeFrozen is a free log retrieval operation binding the contract event 0x340439a63c1ee2404f5b7010cab559b4dcbfc28b8baab9acade354fd688ae2b9.
//
// Solidity: event StakeFrozen(bytes32 slashed, uint256 time)
func (_Staking *StakingFilterer) FilterStakeFrozen(opts *bind.FilterOpts) (*StakingStakeFrozenIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeFrozen")
	if err != nil {
		return nil, err
	}
	return &StakingStakeFrozenIterator{contract: _Staking.contract, event: "StakeFrozen", logs: logs, sub: sub}, nil
}

// WatchStakeFrozen is a free log subscription operation binding the contract event 0x340439a63c1ee2404f5b7010cab559b4dcbfc28b8baab9acade354fd688ae2b9.
//
// Solidity: event StakeFrozen(bytes32 slashed, uint256 time)
func (_Staking *StakingFilterer) WatchStakeFrozen(opts *bind.WatchOpts, sink chan<- *StakingStakeFrozen) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeFrozen)
				if err := _Staking.contract.UnpackLog(event, "StakeFrozen", log); err != nil {
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

// ParseStakeFrozen is a log parse operation binding the contract event 0x340439a63c1ee2404f5b7010cab559b4dcbfc28b8baab9acade354fd688ae2b9.
//
// Solidity: event StakeFrozen(bytes32 slashed, uint256 time)
func (_Staking *StakingFilterer) ParseStakeFrozen(log types.Log) (*StakingStakeFrozen, error) {
	event := new(StakingStakeFrozen)
	if err := _Staking.contract.UnpackLog(event, "StakeFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeSlashedIterator is returned from FilterStakeSlashed and is used to iterate over the raw logs and unpacked data for StakeSlashed events raised by the Staking contract.
type StakingStakeSlashedIterator struct {
	Event *StakingStakeSlashed // Event containing the contract specifics and raw log

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
func (it *StakingStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeSlashed)
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
		it.Event = new(StakingStakeSlashed)
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
func (it *StakingStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeSlashed represents a StakeSlashed event raised by the Staking contract.
type StakingStakeSlashed struct {
	Slashed [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeSlashed is a free log retrieval operation binding the contract event 0x0956b50d4d586f6b9c90800d4e713bd2b866e044decd047e6d70ea20314ab308.
//
// Solidity: event StakeSlashed(bytes32 slashed, uint256 amount)
func (_Staking *StakingFilterer) FilterStakeSlashed(opts *bind.FilterOpts) (*StakingStakeSlashedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeSlashed")
	if err != nil {
		return nil, err
	}
	return &StakingStakeSlashedIterator{contract: _Staking.contract, event: "StakeSlashed", logs: logs, sub: sub}, nil
}

// WatchStakeSlashed is a free log subscription operation binding the contract event 0x0956b50d4d586f6b9c90800d4e713bd2b866e044decd047e6d70ea20314ab308.
//
// Solidity: event StakeSlashed(bytes32 slashed, uint256 amount)
func (_Staking *StakingFilterer) WatchStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakingStakeSlashed) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeSlashed)
				if err := _Staking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
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

// ParseStakeSlashed is a log parse operation binding the contract event 0x0956b50d4d586f6b9c90800d4e713bd2b866e044decd047e6d70ea20314ab308.
//
// Solidity: event StakeSlashed(bytes32 slashed, uint256 amount)
func (_Staking *StakingFilterer) ParseStakeSlashed(log types.Log) (*StakingStakeSlashed, error) {
	event := new(StakingStakeSlashed)
	if err := _Staking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakeUpdatedIterator is returned from FilterStakeUpdated and is used to iterate over the raw logs and unpacked data for StakeUpdated events raised by the Staking contract.
type StakingStakeUpdatedIterator struct {
	Event *StakingStakeUpdated // Event containing the contract specifics and raw log

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
func (it *StakingStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakeUpdated)
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
		it.Event = new(StakingStakeUpdated)
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
func (it *StakingStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakeUpdated represents a StakeUpdated event raised by the Staking contract.
type StakingStakeUpdated struct {
	Overlay          [32]byte
	StakeAmount      *big.Int
	Owner            common.Address
	LastUpdatedBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakeUpdated is a free log retrieval operation binding the contract event 0x61e979698346a2aa868a3a9f08d30c846174841dc9b074bbf2a82d20554bc682.
//
// Solidity: event StakeUpdated(bytes32 indexed overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlock)
func (_Staking *StakingFilterer) FilterStakeUpdated(opts *bind.FilterOpts, overlay [][32]byte) (*StakingStakeUpdatedIterator, error) {

	var overlayRule []interface{}
	for _, overlayItem := range overlay {
		overlayRule = append(overlayRule, overlayItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakeUpdated", overlayRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakeUpdatedIterator{contract: _Staking.contract, event: "StakeUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeUpdated is a free log subscription operation binding the contract event 0x61e979698346a2aa868a3a9f08d30c846174841dc9b074bbf2a82d20554bc682.
//
// Solidity: event StakeUpdated(bytes32 indexed overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlock)
func (_Staking *StakingFilterer) WatchStakeUpdated(opts *bind.WatchOpts, sink chan<- *StakingStakeUpdated, overlay [][32]byte) (event.Subscription, error) {

	var overlayRule []interface{}
	for _, overlayItem := range overlay {
		overlayRule = append(overlayRule, overlayItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakeUpdated", overlayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakeUpdated)
				if err := _Staking.contract.UnpackLog(event, "StakeUpdated", log); err != nil {
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

// ParseStakeUpdated is a log parse operation binding the contract event 0x61e979698346a2aa868a3a9f08d30c846174841dc9b074bbf2a82d20554bc682.
//
// Solidity: event StakeUpdated(bytes32 indexed overlay, uint256 stakeAmount, address owner, uint256 lastUpdatedBlock)
func (_Staking *StakingFilterer) ParseStakeUpdated(log types.Log) (*StakingStakeUpdated, error) {
	event := new(StakingStakeUpdated)
	if err := _Staking.contract.UnpackLog(event, "StakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
