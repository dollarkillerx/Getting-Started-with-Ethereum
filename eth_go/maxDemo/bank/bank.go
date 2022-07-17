// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"internalTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"publicTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawMoney\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bankBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// BankBalances is a free data retrieval call binding the contract method 0x56981be5.
//
// Solidity: function bankBalances() view returns(uint256)
func (_Bank *BankCaller) BankBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "bankBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BankBalances is a free data retrieval call binding the contract method 0x56981be5.
//
// Solidity: function bankBalances() view returns(uint256)
func (_Bank *BankSession) BankBalances() (*big.Int, error) {
	return _Bank.Contract.BankBalances(&_Bank.CallOpts)
}

// BankBalances is a free data retrieval call binding the contract method 0x56981be5.
//
// Solidity: function bankBalances() view returns(uint256)
func (_Bank *BankCallerSession) BankBalances() (*big.Int, error) {
	return _Bank.Contract.BankBalances(&_Bank.CallOpts)
}

// CheckBalances is a free data retrieval call binding the contract method 0x9e241e85.
//
// Solidity: function checkBalances() view returns(uint256)
func (_Bank *BankCaller) CheckBalances(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "checkBalances")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBalances is a free data retrieval call binding the contract method 0x9e241e85.
//
// Solidity: function checkBalances() view returns(uint256)
func (_Bank *BankSession) CheckBalances() (*big.Int, error) {
	return _Bank.Contract.CheckBalances(&_Bank.CallOpts)
}

// CheckBalances is a free data retrieval call binding the contract method 0x9e241e85.
//
// Solidity: function checkBalances() view returns(uint256)
func (_Bank *BankCallerSession) CheckBalances() (*big.Int, error) {
	return _Bank.Contract.CheckBalances(&_Bank.CallOpts)
}

// InternalTransfer is a paid mutator transaction binding the contract method 0x3e7ba166.
//
// Solidity: function internalTransfer(address _to, uint256 amount) returns()
func (_Bank *BankTransactor) InternalTransfer(opts *bind.TransactOpts, _to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "internalTransfer", _to, amount)
}

// InternalTransfer is a paid mutator transaction binding the contract method 0x3e7ba166.
//
// Solidity: function internalTransfer(address _to, uint256 amount) returns()
func (_Bank *BankSession) InternalTransfer(_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.InternalTransfer(&_Bank.TransactOpts, _to, amount)
}

// InternalTransfer is a paid mutator transaction binding the contract method 0x3e7ba166.
//
// Solidity: function internalTransfer(address _to, uint256 amount) returns()
func (_Bank *BankTransactorSession) InternalTransfer(_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.InternalTransfer(&_Bank.TransactOpts, _to, amount)
}

// PublicTransfer is a paid mutator transaction binding the contract method 0x7e9df872.
//
// Solidity: function publicTransfer(address _to) payable returns(bool)
func (_Bank *BankTransactor) PublicTransfer(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "publicTransfer", _to)
}

// PublicTransfer is a paid mutator transaction binding the contract method 0x7e9df872.
//
// Solidity: function publicTransfer(address _to) payable returns(bool)
func (_Bank *BankSession) PublicTransfer(_to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.PublicTransfer(&_Bank.TransactOpts, _to)
}

// PublicTransfer is a paid mutator transaction binding the contract method 0x7e9df872.
//
// Solidity: function publicTransfer(address _to) payable returns(bool)
func (_Bank *BankTransactorSession) PublicTransfer(_to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.PublicTransfer(&_Bank.TransactOpts, _to)
}

// SaveMoney is a paid mutator transaction binding the contract method 0x2466de24.
//
// Solidity: function saveMoney() payable returns()
func (_Bank *BankTransactor) SaveMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "saveMoney")
}

// SaveMoney is a paid mutator transaction binding the contract method 0x2466de24.
//
// Solidity: function saveMoney() payable returns()
func (_Bank *BankSession) SaveMoney() (*types.Transaction, error) {
	return _Bank.Contract.SaveMoney(&_Bank.TransactOpts)
}

// SaveMoney is a paid mutator transaction binding the contract method 0x2466de24.
//
// Solidity: function saveMoney() payable returns()
func (_Bank *BankTransactorSession) SaveMoney() (*types.Transaction, error) {
	return _Bank.Contract.SaveMoney(&_Bank.TransactOpts)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x54876921.
//
// Solidity: function withdrawMoney(uint256 amount) returns(bool)
func (_Bank *BankTransactor) WithdrawMoney(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdrawMoney", amount)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x54876921.
//
// Solidity: function withdrawMoney(uint256 amount) returns(bool)
func (_Bank *BankSession) WithdrawMoney(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts, amount)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0x54876921.
//
// Solidity: function withdrawMoney(uint256 amount) returns(bool)
func (_Bank *BankTransactorSession) WithdrawMoney(amount *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts, amount)
}
