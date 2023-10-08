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

// ITriggerFactoryMetaData contains all meta data concerning the ITriggerFactory contract.
var ITriggerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createTrigger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITriggerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ITriggerFactoryMetaData.ABI instead.
var ITriggerFactoryABI = ITriggerFactoryMetaData.ABI

// ITriggerFactory is an auto generated Go binding around an Ethereum contract.
type ITriggerFactory struct {
	ITriggerFactoryCaller     // Read-only binding to the contract
	ITriggerFactoryTransactor // Write-only binding to the contract
	ITriggerFactoryFilterer   // Log filterer for contract events
}

// ITriggerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITriggerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITriggerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITriggerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITriggerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITriggerFactorySession struct {
	Contract     *ITriggerFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITriggerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITriggerFactoryCallerSession struct {
	Contract *ITriggerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITriggerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITriggerFactoryTransactorSession struct {
	Contract     *ITriggerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITriggerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITriggerFactoryRaw struct {
	Contract *ITriggerFactory // Generic contract binding to access the raw methods on
}

// ITriggerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITriggerFactoryCallerRaw struct {
	Contract *ITriggerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ITriggerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITriggerFactoryTransactorRaw struct {
	Contract *ITriggerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITriggerFactory creates a new instance of ITriggerFactory, bound to a specific deployed contract.
func NewITriggerFactory(address common.Address, backend bind.ContractBackend) (*ITriggerFactory, error) {
	contract, err := bindITriggerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITriggerFactory{ITriggerFactoryCaller: ITriggerFactoryCaller{contract: contract}, ITriggerFactoryTransactor: ITriggerFactoryTransactor{contract: contract}, ITriggerFactoryFilterer: ITriggerFactoryFilterer{contract: contract}}, nil
}

// NewITriggerFactoryCaller creates a new read-only instance of ITriggerFactory, bound to a specific deployed contract.
func NewITriggerFactoryCaller(address common.Address, caller bind.ContractCaller) (*ITriggerFactoryCaller, error) {
	contract, err := bindITriggerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITriggerFactoryCaller{contract: contract}, nil
}

// NewITriggerFactoryTransactor creates a new write-only instance of ITriggerFactory, bound to a specific deployed contract.
func NewITriggerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ITriggerFactoryTransactor, error) {
	contract, err := bindITriggerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITriggerFactoryTransactor{contract: contract}, nil
}

// NewITriggerFactoryFilterer creates a new log filterer instance of ITriggerFactory, bound to a specific deployed contract.
func NewITriggerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ITriggerFactoryFilterer, error) {
	contract, err := bindITriggerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITriggerFactoryFilterer{contract: contract}, nil
}

// bindITriggerFactory binds a generic wrapper to an already deployed contract.
func bindITriggerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITriggerFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITriggerFactory *ITriggerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITriggerFactory.Contract.ITriggerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITriggerFactory *ITriggerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.ITriggerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITriggerFactory *ITriggerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.ITriggerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITriggerFactory *ITriggerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITriggerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITriggerFactory *ITriggerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITriggerFactory *ITriggerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address)
func (_ITriggerFactory *ITriggerFactoryTransactor) CreateTrigger(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ITriggerFactory.contract.Transact(opts, "createTrigger", collectionId, owner)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address)
func (_ITriggerFactory *ITriggerFactorySession) CreateTrigger(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.CreateTrigger(&_ITriggerFactory.TransactOpts, collectionId, owner)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address)
func (_ITriggerFactory *ITriggerFactoryTransactorSession) CreateTrigger(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ITriggerFactory.Contract.CreateTrigger(&_ITriggerFactory.TransactOpts, collectionId, owner)
}
