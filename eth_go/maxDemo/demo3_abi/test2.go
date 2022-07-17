// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package demo3_abi

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

// TestAbi2MetaData contains all meta data concerning the TestAbi2 contract.
var TestAbi2MetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"get_money\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bla\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestAbi2ABI is the input ABI used to generate the binding from.
// Deprecated: Use TestAbi2MetaData.ABI instead.
var TestAbi2ABI = TestAbi2MetaData.ABI

// TestAbi2 is an auto generated Go binding around an Ethereum contract.
type TestAbi2 struct {
	TestAbi2Caller     // Read-only binding to the contract
	TestAbi2Transactor // Write-only binding to the contract
	TestAbi2Filterer   // Log filterer for contract events
}

// TestAbi2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TestAbi2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbi2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TestAbi2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbi2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestAbi2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestAbi2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestAbi2Session struct {
	Contract     *TestAbi2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestAbi2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestAbi2CallerSession struct {
	Contract *TestAbi2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestAbi2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestAbi2TransactorSession struct {
	Contract     *TestAbi2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestAbi2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TestAbi2Raw struct {
	Contract *TestAbi2 // Generic contract binding to access the raw methods on
}

// TestAbi2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestAbi2CallerRaw struct {
	Contract *TestAbi2Caller // Generic read-only contract binding to access the raw methods on
}

// TestAbi2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestAbi2TransactorRaw struct {
	Contract *TestAbi2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTestAbi2 creates a new instance of TestAbi2, bound to a specific deployed contract.
func NewTestAbi2(address common.Address, backend bind.ContractBackend) (*TestAbi2, error) {
	contract, err := bindTestAbi2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestAbi2{TestAbi2Caller: TestAbi2Caller{contract: contract}, TestAbi2Transactor: TestAbi2Transactor{contract: contract}, TestAbi2Filterer: TestAbi2Filterer{contract: contract}}, nil
}

// NewTestAbi2Caller creates a new read-only instance of TestAbi2, bound to a specific deployed contract.
func NewTestAbi2Caller(address common.Address, caller bind.ContractCaller) (*TestAbi2Caller, error) {
	contract, err := bindTestAbi2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbi2Caller{contract: contract}, nil
}

// NewTestAbi2Transactor creates a new write-only instance of TestAbi2, bound to a specific deployed contract.
func NewTestAbi2Transactor(address common.Address, transactor bind.ContractTransactor) (*TestAbi2Transactor, error) {
	contract, err := bindTestAbi2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestAbi2Transactor{contract: contract}, nil
}

// NewTestAbi2Filterer creates a new log filterer instance of TestAbi2, bound to a specific deployed contract.
func NewTestAbi2Filterer(address common.Address, filterer bind.ContractFilterer) (*TestAbi2Filterer, error) {
	contract, err := bindTestAbi2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestAbi2Filterer{contract: contract}, nil
}

// bindTestAbi2 binds a generic wrapper to an already deployed contract.
func bindTestAbi2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestAbi2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbi2 *TestAbi2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbi2.Contract.TestAbi2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbi2 *TestAbi2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbi2.Contract.TestAbi2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbi2 *TestAbi2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbi2.Contract.TestAbi2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestAbi2 *TestAbi2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestAbi2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestAbi2 *TestAbi2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbi2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestAbi2 *TestAbi2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestAbi2.Contract.contract.Transact(opts, method, params...)
}

// Bla is a free data retrieval call binding the contract method 0xf07494a0.
//
// Solidity: function bla() view returns(uint256)
func (_TestAbi2 *TestAbi2Caller) Bla(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestAbi2.contract.Call(opts, &out, "bla")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bla is a free data retrieval call binding the contract method 0xf07494a0.
//
// Solidity: function bla() view returns(uint256)
func (_TestAbi2 *TestAbi2Session) Bla() (*big.Int, error) {
	return _TestAbi2.Contract.Bla(&_TestAbi2.CallOpts)
}

// Bla is a free data retrieval call binding the contract method 0xf07494a0.
//
// Solidity: function bla() view returns(uint256)
func (_TestAbi2 *TestAbi2CallerSession) Bla() (*big.Int, error) {
	return _TestAbi2.Contract.Bla(&_TestAbi2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAbi2 *TestAbi2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestAbi2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAbi2 *TestAbi2Session) Owner() (common.Address, error) {
	return _TestAbi2.Contract.Owner(&_TestAbi2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestAbi2 *TestAbi2CallerSession) Owner() (common.Address, error) {
	return _TestAbi2.Contract.Owner(&_TestAbi2.CallOpts)
}

// GetMoney is a paid mutator transaction binding the contract method 0xb8029269.
//
// Solidity: function get_money() payable returns()
func (_TestAbi2 *TestAbi2Transactor) GetMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbi2.contract.Transact(opts, "get_money")
}

// GetMoney is a paid mutator transaction binding the contract method 0xb8029269.
//
// Solidity: function get_money() payable returns()
func (_TestAbi2 *TestAbi2Session) GetMoney() (*types.Transaction, error) {
	return _TestAbi2.Contract.GetMoney(&_TestAbi2.TransactOpts)
}

// GetMoney is a paid mutator transaction binding the contract method 0xb8029269.
//
// Solidity: function get_money() payable returns()
func (_TestAbi2 *TestAbi2TransactorSession) GetMoney() (*types.Transaction, error) {
	return _TestAbi2.Contract.GetMoney(&_TestAbi2.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_TestAbi2 *TestAbi2Transactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbi2.contract.Transact(opts, "get_owner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_TestAbi2 *TestAbi2Session) GetOwner() (*types.Transaction, error) {
	return _TestAbi2.Contract.GetOwner(&_TestAbi2.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_TestAbi2 *TestAbi2TransactorSession) GetOwner() (*types.Transaction, error) {
	return _TestAbi2.Contract.GetOwner(&_TestAbi2.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_TestAbi2 *TestAbi2Transactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestAbi2.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_TestAbi2 *TestAbi2Session) Test() (*types.Transaction, error) {
	return _TestAbi2.Contract.Test(&_TestAbi2.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_TestAbi2 *TestAbi2TransactorSession) Test() (*types.Transaction, error) {
	return _TestAbi2.Contract.Test(&_TestAbi2.TransactOpts)
}
