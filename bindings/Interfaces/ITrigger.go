// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// ITriggerMetaData contains all meta data concerning the ITrigger contract.
var ITriggerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintTrigger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITriggerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITriggerMetaData.ABI instead.
var ITriggerABI = ITriggerMetaData.ABI

// ITrigger is an auto generated Go binding around an Ethereum contract.
type ITrigger struct {
	ITriggerCaller     // Read-only binding to the contract
	ITriggerTransactor // Write-only binding to the contract
	ITriggerFilterer   // Log filterer for contract events
}

// ITriggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITriggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITriggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITriggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITriggerSession struct {
	Contract     *ITrigger         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITriggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITriggerCallerSession struct {
	Contract *ITriggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ITriggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITriggerTransactorSession struct {
	Contract     *ITriggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ITriggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITriggerRaw struct {
	Contract *ITrigger // Generic contract binding to access the raw methods on
}

// ITriggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITriggerCallerRaw struct {
	Contract *ITriggerCaller // Generic read-only contract binding to access the raw methods on
}

// ITriggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITriggerTransactorRaw struct {
	Contract *ITriggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITrigger creates a new instance of ITrigger, bound to a specific deployed contract.
func NewITrigger(address common.Address, backend bind.ContractBackend) (*ITrigger, error) {
	contract, err := bindITrigger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITrigger{ITriggerCaller: ITriggerCaller{contract: contract}, ITriggerTransactor: ITriggerTransactor{contract: contract}, ITriggerFilterer: ITriggerFilterer{contract: contract}}, nil
}

// NewITriggerCaller creates a new read-only instance of ITrigger, bound to a specific deployed contract.
func NewITriggerCaller(address common.Address, caller bind.ContractCaller) (*ITriggerCaller, error) {
	contract, err := bindITrigger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITriggerCaller{contract: contract}, nil
}

// NewITriggerTransactor creates a new write-only instance of ITrigger, bound to a specific deployed contract.
func NewITriggerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITriggerTransactor, error) {
	contract, err := bindITrigger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITriggerTransactor{contract: contract}, nil
}

// NewITriggerFilterer creates a new log filterer instance of ITrigger, bound to a specific deployed contract.
func NewITriggerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITriggerFilterer, error) {
	contract, err := bindITrigger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITriggerFilterer{contract: contract}, nil
}

// bindITrigger binds a generic wrapper to an already deployed contract.
func bindITrigger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITriggerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrigger *ITriggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITrigger.Contract.ITriggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrigger *ITriggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrigger.Contract.ITriggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrigger *ITriggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrigger.Contract.ITriggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrigger *ITriggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITrigger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrigger *ITriggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrigger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrigger *ITriggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrigger.Contract.contract.Transact(opts, method, params...)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns()
func (_ITrigger *ITriggerTransactor) MintTrigger(opts *bind.TransactOpts, _tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _ITrigger.contract.Transact(opts, "mintTrigger", _tokenURI, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns()
func (_ITrigger *ITriggerSession) MintTrigger(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _ITrigger.Contract.MintTrigger(&_ITrigger.TransactOpts, _tokenURI, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns()
func (_ITrigger *ITriggerTransactorSession) MintTrigger(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _ITrigger.Contract.MintTrigger(&_ITrigger.TransactOpts, _tokenURI, _owner)
}
