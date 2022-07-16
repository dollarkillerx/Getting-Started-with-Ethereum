// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ab2_t2

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

// Test2AbbMetaData contains all meta data concerning the Test2Abb contract.
var Test2AbbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Test2AbbABI is the input ABI used to generate the binding from.
// Deprecated: Use Test2AbbMetaData.ABI instead.
var Test2AbbABI = Test2AbbMetaData.ABI

// Test2Abb is an auto generated Go binding around an Ethereum contract.
type Test2Abb struct {
	Test2AbbCaller     // Read-only binding to the contract
	Test2AbbTransactor // Write-only binding to the contract
	Test2AbbFilterer   // Log filterer for contract events
}

// Test2AbbCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test2AbbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2AbbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test2AbbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2AbbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test2AbbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2AbbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test2AbbSession struct {
	Contract     *Test2Abb         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Test2AbbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test2AbbCallerSession struct {
	Contract *Test2AbbCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Test2AbbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test2AbbTransactorSession struct {
	Contract     *Test2AbbTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Test2AbbRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test2AbbRaw struct {
	Contract *Test2Abb // Generic contract binding to access the raw methods on
}

// Test2AbbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test2AbbCallerRaw struct {
	Contract *Test2AbbCaller // Generic read-only contract binding to access the raw methods on
}

// Test2AbbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test2AbbTransactorRaw struct {
	Contract *Test2AbbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest2Abb creates a new instance of Test2Abb, bound to a specific deployed contract.
func NewTest2Abb(address common.Address, backend bind.ContractBackend) (*Test2Abb, error) {
	contract, err := bindTest2Abb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test2Abb{Test2AbbCaller: Test2AbbCaller{contract: contract}, Test2AbbTransactor: Test2AbbTransactor{contract: contract}, Test2AbbFilterer: Test2AbbFilterer{contract: contract}}, nil
}

// NewTest2AbbCaller creates a new read-only instance of Test2Abb, bound to a specific deployed contract.
func NewTest2AbbCaller(address common.Address, caller bind.ContractCaller) (*Test2AbbCaller, error) {
	contract, err := bindTest2Abb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test2AbbCaller{contract: contract}, nil
}

// NewTest2AbbTransactor creates a new write-only instance of Test2Abb, bound to a specific deployed contract.
func NewTest2AbbTransactor(address common.Address, transactor bind.ContractTransactor) (*Test2AbbTransactor, error) {
	contract, err := bindTest2Abb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test2AbbTransactor{contract: contract}, nil
}

// NewTest2AbbFilterer creates a new log filterer instance of Test2Abb, bound to a specific deployed contract.
func NewTest2AbbFilterer(address common.Address, filterer bind.ContractFilterer) (*Test2AbbFilterer, error) {
	contract, err := bindTest2Abb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test2AbbFilterer{contract: contract}, nil
}

// bindTest2Abb binds a generic wrapper to an already deployed contract.
func bindTest2Abb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Test2AbbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Abb *Test2AbbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Abb.Contract.Test2AbbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Abb *Test2AbbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Abb.Contract.Test2AbbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Abb *Test2AbbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Abb.Contract.Test2AbbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Abb *Test2AbbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Abb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Abb *Test2AbbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Abb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Abb *Test2AbbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Abb.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x6932cf81.
//
// Solidity: function getName(string _name) pure returns(string)
func (_Test2Abb *Test2AbbCaller) GetName(opts *bind.CallOpts, _name string) (string, error) {
	var out []interface{}
	err := _Test2Abb.contract.Call(opts, &out, "getName", _name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x6932cf81.
//
// Solidity: function getName(string _name) pure returns(string)
func (_Test2Abb *Test2AbbSession) GetName(_name string) (string, error) {
	return _Test2Abb.Contract.GetName(&_Test2Abb.CallOpts, _name)
}

// GetName is a free data retrieval call binding the contract method 0x6932cf81.
//
// Solidity: function getName(string _name) pure returns(string)
func (_Test2Abb *Test2AbbCallerSession) GetName(_name string) (string, error) {
	return _Test2Abb.Contract.GetName(&_Test2Abb.CallOpts, _name)
}
