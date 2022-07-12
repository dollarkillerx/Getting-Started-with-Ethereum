// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MyStructMetaData contains all meta data concerning the MyStruct contract.
var MyStructMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"get_owner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MyStructABI is the input ABI used to generate the binding from.
// Deprecated: Use MyStructMetaData.ABI instead.
var MyStructABI = MyStructMetaData.ABI

// MyStruct is an auto generated Go binding around an Ethereum contract.
type MyStruct struct {
	MyStructCaller     // Read-only binding to the contract
	MyStructTransactor // Write-only binding to the contract
	MyStructFilterer   // Log filterer for contract events
}

// MyStructCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyStructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyStructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyStructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyStructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyStructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyStructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyStructSession struct {
	Contract     *MyStruct         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyStructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyStructCallerSession struct {
	Contract *MyStructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MyStructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyStructTransactorSession struct {
	Contract     *MyStructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MyStructRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyStructRaw struct {
	Contract *MyStruct // Generic contract binding to access the raw methods on
}

// MyStructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyStructCallerRaw struct {
	Contract *MyStructCaller // Generic read-only contract binding to access the raw methods on
}

// MyStructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyStructTransactorRaw struct {
	Contract *MyStructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyStruct creates a new instance of MyStruct, bound to a specific deployed contract.
func NewMyStruct(address common.Address, backend bind.ContractBackend) (*MyStruct, error) {
	contract, err := bindMyStruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyStruct{MyStructCaller: MyStructCaller{contract: contract}, MyStructTransactor: MyStructTransactor{contract: contract}, MyStructFilterer: MyStructFilterer{contract: contract}}, nil
}

// NewMyStructCaller creates a new read-only instance of MyStruct, bound to a specific deployed contract.
func NewMyStructCaller(address common.Address, caller bind.ContractCaller) (*MyStructCaller, error) {
	contract, err := bindMyStruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyStructCaller{contract: contract}, nil
}

// NewMyStructTransactor creates a new write-only instance of MyStruct, bound to a specific deployed contract.
func NewMyStructTransactor(address common.Address, transactor bind.ContractTransactor) (*MyStructTransactor, error) {
	contract, err := bindMyStruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyStructTransactor{contract: contract}, nil
}

// NewMyStructFilterer creates a new log filterer instance of MyStruct, bound to a specific deployed contract.
func NewMyStructFilterer(address common.Address, filterer bind.ContractFilterer) (*MyStructFilterer, error) {
	contract, err := bindMyStruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyStructFilterer{contract: contract}, nil
}

// bindMyStruct binds a generic wrapper to an already deployed contract.
func bindMyStruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyStructABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyStruct *MyStructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyStruct.Contract.MyStructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyStruct *MyStructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyStruct.Contract.MyStructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyStruct *MyStructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyStruct.Contract.MyStructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyStruct *MyStructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyStruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyStruct *MyStructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyStruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyStruct *MyStructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyStruct.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyStruct *MyStructCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyStruct.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyStruct *MyStructSession) Owner() (common.Address, error) {
	return _MyStruct.Contract.Owner(&_MyStruct.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MyStruct *MyStructCallerSession) Owner() (common.Address, error) {
	return _MyStruct.Contract.Owner(&_MyStruct.CallOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_MyStruct *MyStructTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyStruct.contract.Transact(opts, "get_owner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_MyStruct *MyStructSession) GetOwner() (*types.Transaction, error) {
	return _MyStruct.Contract.GetOwner(&_MyStruct.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x0ac298dc.
//
// Solidity: function get_owner() returns()
func (_MyStruct *MyStructTransactorSession) GetOwner() (*types.Transaction, error) {
	return _MyStruct.Contract.GetOwner(&_MyStruct.TransactOpts)
}
