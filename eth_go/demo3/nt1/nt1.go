// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nt1

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

// Nt1MetaData contains all meta data concerning the Nt1 contract.
var Nt1MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"get_name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Nt1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Nt1MetaData.ABI instead.
var Nt1ABI = Nt1MetaData.ABI

// Nt1 is an auto generated Go binding around an Ethereum contract.
type Nt1 struct {
	Nt1Caller     // Read-only binding to the contract
	Nt1Transactor // Write-only binding to the contract
	Nt1Filterer   // Log filterer for contract events
}

// Nt1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Nt1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nt1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Nt1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nt1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Nt1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Nt1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Nt1Session struct {
	Contract     *Nt1              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Nt1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Nt1CallerSession struct {
	Contract *Nt1Caller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Nt1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Nt1TransactorSession struct {
	Contract     *Nt1Transactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Nt1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Nt1Raw struct {
	Contract *Nt1 // Generic contract binding to access the raw methods on
}

// Nt1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Nt1CallerRaw struct {
	Contract *Nt1Caller // Generic read-only contract binding to access the raw methods on
}

// Nt1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Nt1TransactorRaw struct {
	Contract *Nt1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNt1 creates a new instance of Nt1, bound to a specific deployed contract.
func NewNt1(address common.Address, backend bind.ContractBackend) (*Nt1, error) {
	contract, err := bindNt1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nt1{Nt1Caller: Nt1Caller{contract: contract}, Nt1Transactor: Nt1Transactor{contract: contract}, Nt1Filterer: Nt1Filterer{contract: contract}}, nil
}

// NewNt1Caller creates a new read-only instance of Nt1, bound to a specific deployed contract.
func NewNt1Caller(address common.Address, caller bind.ContractCaller) (*Nt1Caller, error) {
	contract, err := bindNt1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Nt1Caller{contract: contract}, nil
}

// NewNt1Transactor creates a new write-only instance of Nt1, bound to a specific deployed contract.
func NewNt1Transactor(address common.Address, transactor bind.ContractTransactor) (*Nt1Transactor, error) {
	contract, err := bindNt1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Nt1Transactor{contract: contract}, nil
}

// NewNt1Filterer creates a new log filterer instance of Nt1, bound to a specific deployed contract.
func NewNt1Filterer(address common.Address, filterer bind.ContractFilterer) (*Nt1Filterer, error) {
	contract, err := bindNt1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Nt1Filterer{contract: contract}, nil
}

// bindNt1 binds a generic wrapper to an already deployed contract.
func bindNt1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Nt1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nt1 *Nt1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nt1.Contract.Nt1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nt1 *Nt1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nt1.Contract.Nt1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nt1 *Nt1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nt1.Contract.Nt1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nt1 *Nt1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nt1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nt1 *Nt1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nt1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nt1 *Nt1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nt1.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0xa8f5d7a6.
//
// Solidity: function get_name(string name) view returns(string)
func (_Nt1 *Nt1Caller) GetName(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Nt1.contract.Call(opts, &out, "get_name", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0xa8f5d7a6.
//
// Solidity: function get_name(string name) view returns(string)
func (_Nt1 *Nt1Session) GetName(name string) (string, error) {
	return _Nt1.Contract.GetName(&_Nt1.CallOpts, name)
}

// GetName is a free data retrieval call binding the contract method 0xa8f5d7a6.
//
// Solidity: function get_name(string name) view returns(string)
func (_Nt1 *Nt1CallerSession) GetName(name string) (string, error) {
	return _Nt1.Contract.GetName(&_Nt1.CallOpts, name)
}
