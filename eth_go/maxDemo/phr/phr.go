// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package phr

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

// PHRMetaData contains all meta data concerning the PHR contract.
var PHRMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PHRABI is the input ABI used to generate the binding from.
// Deprecated: Use PHRMetaData.ABI instead.
var PHRABI = PHRMetaData.ABI

// PHR is an auto generated Go binding around an Ethereum contract.
type PHR struct {
	PHRCaller     // Read-only binding to the contract
	PHRTransactor // Write-only binding to the contract
	PHRFilterer   // Log filterer for contract events
}

// PHRCaller is an auto generated read-only Go binding around an Ethereum contract.
type PHRCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PHRTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PHRTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PHRFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PHRFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PHRSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PHRSession struct {
	Contract     *PHR              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PHRCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PHRCallerSession struct {
	Contract *PHRCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PHRTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PHRTransactorSession struct {
	Contract     *PHRTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PHRRaw is an auto generated low-level Go binding around an Ethereum contract.
type PHRRaw struct {
	Contract *PHR // Generic contract binding to access the raw methods on
}

// PHRCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PHRCallerRaw struct {
	Contract *PHRCaller // Generic read-only contract binding to access the raw methods on
}

// PHRTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PHRTransactorRaw struct {
	Contract *PHRTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPHR creates a new instance of PHR, bound to a specific deployed contract.
func NewPHR(address common.Address, backend bind.ContractBackend) (*PHR, error) {
	contract, err := bindPHR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PHR{PHRCaller: PHRCaller{contract: contract}, PHRTransactor: PHRTransactor{contract: contract}, PHRFilterer: PHRFilterer{contract: contract}}, nil
}

// NewPHRCaller creates a new read-only instance of PHR, bound to a specific deployed contract.
func NewPHRCaller(address common.Address, caller bind.ContractCaller) (*PHRCaller, error) {
	contract, err := bindPHR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PHRCaller{contract: contract}, nil
}

// NewPHRTransactor creates a new write-only instance of PHR, bound to a specific deployed contract.
func NewPHRTransactor(address common.Address, transactor bind.ContractTransactor) (*PHRTransactor, error) {
	contract, err := bindPHR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PHRTransactor{contract: contract}, nil
}

// NewPHRFilterer creates a new log filterer instance of PHR, bound to a specific deployed contract.
func NewPHRFilterer(address common.Address, filterer bind.ContractFilterer) (*PHRFilterer, error) {
	contract, err := bindPHR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PHRFilterer{contract: contract}, nil
}

// bindPHR binds a generic wrapper to an already deployed contract.
func bindPHR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PHRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PHR *PHRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PHR.Contract.PHRCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PHR *PHRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PHR.Contract.PHRTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PHR *PHRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PHR.Contract.PHRTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PHR *PHRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PHR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PHR *PHRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PHR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PHR *PHRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PHR.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_PHR *PHRCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PHR.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_PHR *PHRSession) GetName() (string, error) {
	return _PHR.Contract.GetName(&_PHR.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_PHR *PHRCallerSession) GetName() (string, error) {
	return _PHR.Contract.GetName(&_PHR.CallOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PHR *PHRTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _PHR.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PHR *PHRSession) SetName(_name string) (*types.Transaction, error) {
	return _PHR.Contract.SetName(&_PHR.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_PHR *PHRTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _PHR.Contract.SetName(&_PHR.TransactOpts, _name)
}
