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

// IDayMonthFactoryMetaData contains all meta data concerning the IDayMonthFactory contract.
var IDayMonthFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createDayMonth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDayMonthFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IDayMonthFactoryMetaData.ABI instead.
var IDayMonthFactoryABI = IDayMonthFactoryMetaData.ABI

// IDayMonthFactory is an auto generated Go binding around an Ethereum contract.
type IDayMonthFactory struct {
	IDayMonthFactoryCaller     // Read-only binding to the contract
	IDayMonthFactoryTransactor // Write-only binding to the contract
	IDayMonthFactoryFilterer   // Log filterer for contract events
}

// IDayMonthFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDayMonthFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDayMonthFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDayMonthFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDayMonthFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDayMonthFactorySession struct {
	Contract     *IDayMonthFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDayMonthFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDayMonthFactoryCallerSession struct {
	Contract *IDayMonthFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IDayMonthFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDayMonthFactoryTransactorSession struct {
	Contract     *IDayMonthFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IDayMonthFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDayMonthFactoryRaw struct {
	Contract *IDayMonthFactory // Generic contract binding to access the raw methods on
}

// IDayMonthFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDayMonthFactoryCallerRaw struct {
	Contract *IDayMonthFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IDayMonthFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDayMonthFactoryTransactorRaw struct {
	Contract *IDayMonthFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDayMonthFactory creates a new instance of IDayMonthFactory, bound to a specific deployed contract.
func NewIDayMonthFactory(address common.Address, backend bind.ContractBackend) (*IDayMonthFactory, error) {
	contract, err := bindIDayMonthFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDayMonthFactory{IDayMonthFactoryCaller: IDayMonthFactoryCaller{contract: contract}, IDayMonthFactoryTransactor: IDayMonthFactoryTransactor{contract: contract}, IDayMonthFactoryFilterer: IDayMonthFactoryFilterer{contract: contract}}, nil
}

// NewIDayMonthFactoryCaller creates a new read-only instance of IDayMonthFactory, bound to a specific deployed contract.
func NewIDayMonthFactoryCaller(address common.Address, caller bind.ContractCaller) (*IDayMonthFactoryCaller, error) {
	contract, err := bindIDayMonthFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDayMonthFactoryCaller{contract: contract}, nil
}

// NewIDayMonthFactoryTransactor creates a new write-only instance of IDayMonthFactory, bound to a specific deployed contract.
func NewIDayMonthFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IDayMonthFactoryTransactor, error) {
	contract, err := bindIDayMonthFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDayMonthFactoryTransactor{contract: contract}, nil
}

// NewIDayMonthFactoryFilterer creates a new log filterer instance of IDayMonthFactory, bound to a specific deployed contract.
func NewIDayMonthFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IDayMonthFactoryFilterer, error) {
	contract, err := bindIDayMonthFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDayMonthFactoryFilterer{contract: contract}, nil
}

// bindIDayMonthFactory binds a generic wrapper to an already deployed contract.
func bindIDayMonthFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDayMonthFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDayMonthFactory *IDayMonthFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDayMonthFactory.Contract.IDayMonthFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDayMonthFactory *IDayMonthFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.IDayMonthFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDayMonthFactory *IDayMonthFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.IDayMonthFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDayMonthFactory *IDayMonthFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDayMonthFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDayMonthFactory *IDayMonthFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDayMonthFactory *IDayMonthFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address)
func (_IDayMonthFactory *IDayMonthFactoryTransactor) CreateDayMonth(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IDayMonthFactory.contract.Transact(opts, "createDayMonth", collectionId, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address)
func (_IDayMonthFactory *IDayMonthFactorySession) CreateDayMonth(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.CreateDayMonth(&_IDayMonthFactory.TransactOpts, collectionId, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address)
func (_IDayMonthFactory *IDayMonthFactoryTransactorSession) CreateDayMonth(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IDayMonthFactory.Contract.CreateDayMonth(&_IDayMonthFactory.TransactOpts, collectionId, owner)
}
