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

// IYearMetaData contains all meta data concerning the IYear contract.
var IYearMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintYear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IYearABI is the input ABI used to generate the binding from.
// Deprecated: Use IYearMetaData.ABI instead.
var IYearABI = IYearMetaData.ABI

// IYear is an auto generated Go binding around an Ethereum contract.
type IYear struct {
	IYearCaller     // Read-only binding to the contract
	IYearTransactor // Write-only binding to the contract
	IYearFilterer   // Log filterer for contract events
}

// IYearCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYearCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYearTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYearFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYearSession struct {
	Contract     *IYear            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYearCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYearCallerSession struct {
	Contract *IYearCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IYearTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYearTransactorSession struct {
	Contract     *IYearTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYearRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYearRaw struct {
	Contract *IYear // Generic contract binding to access the raw methods on
}

// IYearCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYearCallerRaw struct {
	Contract *IYearCaller // Generic read-only contract binding to access the raw methods on
}

// IYearTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYearTransactorRaw struct {
	Contract *IYearTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYear creates a new instance of IYear, bound to a specific deployed contract.
func NewIYear(address common.Address, backend bind.ContractBackend) (*IYear, error) {
	contract, err := bindIYear(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYear{IYearCaller: IYearCaller{contract: contract}, IYearTransactor: IYearTransactor{contract: contract}, IYearFilterer: IYearFilterer{contract: contract}}, nil
}

// NewIYearCaller creates a new read-only instance of IYear, bound to a specific deployed contract.
func NewIYearCaller(address common.Address, caller bind.ContractCaller) (*IYearCaller, error) {
	contract, err := bindIYear(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYearCaller{contract: contract}, nil
}

// NewIYearTransactor creates a new write-only instance of IYear, bound to a specific deployed contract.
func NewIYearTransactor(address common.Address, transactor bind.ContractTransactor) (*IYearTransactor, error) {
	contract, err := bindIYear(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYearTransactor{contract: contract}, nil
}

// NewIYearFilterer creates a new log filterer instance of IYear, bound to a specific deployed contract.
func NewIYearFilterer(address common.Address, filterer bind.ContractFilterer) (*IYearFilterer, error) {
	contract, err := bindIYear(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYearFilterer{contract: contract}, nil
}

// bindIYear binds a generic wrapper to an already deployed contract.
func bindIYear(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IYearMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYear *IYearRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYear.Contract.IYearCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYear *IYearRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYear.Contract.IYearTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYear *IYearRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYear.Contract.IYearTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYear *IYearCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYear.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYear *IYearTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYear.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYear *IYearTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYear.Contract.contract.Transact(opts, method, params...)
}

// MintYear is a paid mutator transaction binding the contract method 0x62eea0ed.
//
// Solidity: function mintYear(string _tokenURI, address _owner) returns()
func (_IYear *IYearTransactor) MintYear(opts *bind.TransactOpts, _tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _IYear.contract.Transact(opts, "mintYear", _tokenURI, _owner)
}

// MintYear is a paid mutator transaction binding the contract method 0x62eea0ed.
//
// Solidity: function mintYear(string _tokenURI, address _owner) returns()
func (_IYear *IYearSession) MintYear(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _IYear.Contract.MintYear(&_IYear.TransactOpts, _tokenURI, _owner)
}

// MintYear is a paid mutator transaction binding the contract method 0x62eea0ed.
//
// Solidity: function mintYear(string _tokenURI, address _owner) returns()
func (_IYear *IYearTransactorSession) MintYear(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _IYear.Contract.MintYear(&_IYear.TransactOpts, _tokenURI, _owner)
}
