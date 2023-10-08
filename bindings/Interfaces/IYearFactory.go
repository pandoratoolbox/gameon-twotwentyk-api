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

// IYearFactoryMetaData contains all meta data concerning the IYearFactory contract.
var IYearFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createYear\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IYearFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IYearFactoryMetaData.ABI instead.
var IYearFactoryABI = IYearFactoryMetaData.ABI

// IYearFactory is an auto generated Go binding around an Ethereum contract.
type IYearFactory struct {
	IYearFactoryCaller     // Read-only binding to the contract
	IYearFactoryTransactor // Write-only binding to the contract
	IYearFactoryFilterer   // Log filterer for contract events
}

// IYearFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IYearFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IYearFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IYearFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IYearFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IYearFactorySession struct {
	Contract     *IYearFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IYearFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IYearFactoryCallerSession struct {
	Contract *IYearFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IYearFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IYearFactoryTransactorSession struct {
	Contract     *IYearFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IYearFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IYearFactoryRaw struct {
	Contract *IYearFactory // Generic contract binding to access the raw methods on
}

// IYearFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IYearFactoryCallerRaw struct {
	Contract *IYearFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IYearFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IYearFactoryTransactorRaw struct {
	Contract *IYearFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIYearFactory creates a new instance of IYearFactory, bound to a specific deployed contract.
func NewIYearFactory(address common.Address, backend bind.ContractBackend) (*IYearFactory, error) {
	contract, err := bindIYearFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IYearFactory{IYearFactoryCaller: IYearFactoryCaller{contract: contract}, IYearFactoryTransactor: IYearFactoryTransactor{contract: contract}, IYearFactoryFilterer: IYearFactoryFilterer{contract: contract}}, nil
}

// NewIYearFactoryCaller creates a new read-only instance of IYearFactory, bound to a specific deployed contract.
func NewIYearFactoryCaller(address common.Address, caller bind.ContractCaller) (*IYearFactoryCaller, error) {
	contract, err := bindIYearFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IYearFactoryCaller{contract: contract}, nil
}

// NewIYearFactoryTransactor creates a new write-only instance of IYearFactory, bound to a specific deployed contract.
func NewIYearFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IYearFactoryTransactor, error) {
	contract, err := bindIYearFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IYearFactoryTransactor{contract: contract}, nil
}

// NewIYearFactoryFilterer creates a new log filterer instance of IYearFactory, bound to a specific deployed contract.
func NewIYearFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IYearFactoryFilterer, error) {
	contract, err := bindIYearFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IYearFactoryFilterer{contract: contract}, nil
}

// bindIYearFactory binds a generic wrapper to an already deployed contract.
func bindIYearFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IYearFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYearFactory *IYearFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYearFactory.Contract.IYearFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYearFactory *IYearFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYearFactory.Contract.IYearFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYearFactory *IYearFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYearFactory.Contract.IYearFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IYearFactory *IYearFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IYearFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IYearFactory *IYearFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IYearFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IYearFactory *IYearFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IYearFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address)
func (_IYearFactory *IYearFactoryTransactor) CreateYear(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IYearFactory.contract.Transact(opts, "createYear", collectionId, owner)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address)
func (_IYearFactory *IYearFactorySession) CreateYear(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IYearFactory.Contract.CreateYear(&_IYearFactory.TransactOpts, collectionId, owner)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address)
func (_IYearFactory *IYearFactoryTransactorSession) CreateYear(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IYearFactory.Contract.CreateYear(&_IYearFactory.TransactOpts, collectionId, owner)
}
