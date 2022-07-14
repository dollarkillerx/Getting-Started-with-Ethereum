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

// AB2T2MetaData contains all meta data concerning the AB2T2 contract.
var AB2T2MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AB2T2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AB2T2MetaData.ABI instead.
var AB2T2ABI = AB2T2MetaData.ABI

// AB2T2 is an auto generated Go binding around an Ethereum contract.
type AB2T2 struct {
	AB2T2Caller     // Read-only binding to the contract
	AB2T2Transactor // Write-only binding to the contract
	AB2T2Filterer   // Log filterer for contract events
}

// AB2T2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AB2T2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AB2T2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AB2T2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AB2T2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AB2T2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AB2T2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AB2T2Session struct {
	Contract     *AB2T2            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AB2T2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AB2T2CallerSession struct {
	Contract *AB2T2Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AB2T2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AB2T2TransactorSession struct {
	Contract     *AB2T2Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AB2T2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AB2T2Raw struct {
	Contract *AB2T2 // Generic contract binding to access the raw methods on
}

// AB2T2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AB2T2CallerRaw struct {
	Contract *AB2T2Caller // Generic read-only contract binding to access the raw methods on
}

// AB2T2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AB2T2TransactorRaw struct {
	Contract *AB2T2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAB2T2 creates a new instance of AB2T2, bound to a specific deployed contract.
func NewAB2T2(address common.Address, backend bind.ContractBackend) (*AB2T2, error) {
	contract, err := bindAB2T2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AB2T2{AB2T2Caller: AB2T2Caller{contract: contract}, AB2T2Transactor: AB2T2Transactor{contract: contract}, AB2T2Filterer: AB2T2Filterer{contract: contract}}, nil
}

// NewAB2T2Caller creates a new read-only instance of AB2T2, bound to a specific deployed contract.
func NewAB2T2Caller(address common.Address, caller bind.ContractCaller) (*AB2T2Caller, error) {
	contract, err := bindAB2T2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AB2T2Caller{contract: contract}, nil
}

// NewAB2T2Transactor creates a new write-only instance of AB2T2, bound to a specific deployed contract.
func NewAB2T2Transactor(address common.Address, transactor bind.ContractTransactor) (*AB2T2Transactor, error) {
	contract, err := bindAB2T2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AB2T2Transactor{contract: contract}, nil
}

// NewAB2T2Filterer creates a new log filterer instance of AB2T2, bound to a specific deployed contract.
func NewAB2T2Filterer(address common.Address, filterer bind.ContractFilterer) (*AB2T2Filterer, error) {
	contract, err := bindAB2T2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AB2T2Filterer{contract: contract}, nil
}

// bindAB2T2 binds a generic wrapper to an already deployed contract.
func bindAB2T2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AB2T2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AB2T2 *AB2T2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AB2T2.Contract.AB2T2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AB2T2 *AB2T2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AB2T2.Contract.AB2T2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AB2T2 *AB2T2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AB2T2.Contract.AB2T2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AB2T2 *AB2T2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AB2T2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AB2T2 *AB2T2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AB2T2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AB2T2 *AB2T2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AB2T2.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AB2T2 *AB2T2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AB2T2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AB2T2 *AB2T2Session) Name() (string, error) {
	return _AB2T2.Contract.Name(&_AB2T2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AB2T2 *AB2T2CallerSession) Name() (string, error) {
	return _AB2T2.Contract.Name(&_AB2T2.CallOpts)
}
