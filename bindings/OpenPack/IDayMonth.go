// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openpack

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

// IDayMonthMetaData contains all meta data concerning the IDayMonth contract.
var IDayMonthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createDayMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDayMonthABI is the input ABI used to generate the binding from.
// Deprecated: Use IDayMonthMetaData.ABI instead.
var IDayMonthABI = IDayMonthMetaData.ABI

// IDayMonth is an auto generated Go binding around an Ethereum contract.
type IDayMonth struct {
	IDayMonthCaller     // Read-only binding to the contract
	IDayMonthTransactor // Write-only binding to the contract
	IDayMonthFilterer   // Log filterer for contract events
}

// IDayMonthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDayMonthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDayMonthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDayMonthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDayMonthSession struct {
	Contract     *IDayMonth        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDayMonthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDayMonthCallerSession struct {
	Contract *IDayMonthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IDayMonthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDayMonthTransactorSession struct {
	Contract     *IDayMonthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IDayMonthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDayMonthRaw struct {
	Contract *IDayMonth // Generic contract binding to access the raw methods on
}

// IDayMonthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDayMonthCallerRaw struct {
	Contract *IDayMonthCaller // Generic read-only contract binding to access the raw methods on
}

// IDayMonthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDayMonthTransactorRaw struct {
	Contract *IDayMonthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDayMonth creates a new instance of IDayMonth, bound to a specific deployed contract.
func NewIDayMonth(address common.Address, backend bind.ContractBackend) (*IDayMonth, error) {
	contract, err := bindIDayMonth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDayMonth{IDayMonthCaller: IDayMonthCaller{contract: contract}, IDayMonthTransactor: IDayMonthTransactor{contract: contract}, IDayMonthFilterer: IDayMonthFilterer{contract: contract}}, nil
}

// NewIDayMonthCaller creates a new read-only instance of IDayMonth, bound to a specific deployed contract.
func NewIDayMonthCaller(address common.Address, caller bind.ContractCaller) (*IDayMonthCaller, error) {
	contract, err := bindIDayMonth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDayMonthCaller{contract: contract}, nil
}

// NewIDayMonthTransactor creates a new write-only instance of IDayMonth, bound to a specific deployed contract.
func NewIDayMonthTransactor(address common.Address, transactor bind.ContractTransactor) (*IDayMonthTransactor, error) {
	contract, err := bindIDayMonth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDayMonthTransactor{contract: contract}, nil
}

// NewIDayMonthFilterer creates a new log filterer instance of IDayMonth, bound to a specific deployed contract.
func NewIDayMonthFilterer(address common.Address, filterer bind.ContractFilterer) (*IDayMonthFilterer, error) {
	contract, err := bindIDayMonth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDayMonthFilterer{contract: contract}, nil
}

// bindIDayMonth binds a generic wrapper to an already deployed contract.
func bindIDayMonth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDayMonthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDayMonth *IDayMonthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDayMonth.Contract.IDayMonthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDayMonth *IDayMonthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDayMonth.Contract.IDayMonthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDayMonth *IDayMonthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDayMonth.Contract.IDayMonthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDayMonth *IDayMonthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDayMonth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDayMonth *IDayMonthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDayMonth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDayMonth *IDayMonthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDayMonth.Contract.contract.Transact(opts, method, params...)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x523a2ffd.
//
// Solidity: function createDayMonth(string tokenURI, address owner) returns(uint256)
func (_IDayMonth *IDayMonthTransactor) CreateDayMonth(opts *bind.TransactOpts, tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _IDayMonth.contract.Transact(opts, "createDayMonth", tokenURI, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x523a2ffd.
//
// Solidity: function createDayMonth(string tokenURI, address owner) returns(uint256)
func (_IDayMonth *IDayMonthSession) CreateDayMonth(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _IDayMonth.Contract.CreateDayMonth(&_IDayMonth.TransactOpts, tokenURI, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x523a2ffd.
//
// Solidity: function createDayMonth(string tokenURI, address owner) returns(uint256)
func (_IDayMonth *IDayMonthTransactorSession) CreateDayMonth(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _IDayMonth.Contract.CreateDayMonth(&_IDayMonth.TransactOpts, tokenURI, owner)
}
