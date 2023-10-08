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

// IPackFactoryMetaData contains all meta data concerning the IPackFactory contract.
var IPackFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCardPack\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPackFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPackFactoryMetaData.ABI instead.
var IPackFactoryABI = IPackFactoryMetaData.ABI

// IPackFactory is an auto generated Go binding around an Ethereum contract.
type IPackFactory struct {
	IPackFactoryCaller     // Read-only binding to the contract
	IPackFactoryTransactor // Write-only binding to the contract
	IPackFactoryFilterer   // Log filterer for contract events
}

// IPackFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPackFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPackFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPackFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPackFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPackFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPackFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPackFactorySession struct {
	Contract     *IPackFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPackFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPackFactoryCallerSession struct {
	Contract *IPackFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IPackFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPackFactoryTransactorSession struct {
	Contract     *IPackFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IPackFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPackFactoryRaw struct {
	Contract *IPackFactory // Generic contract binding to access the raw methods on
}

// IPackFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPackFactoryCallerRaw struct {
	Contract *IPackFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPackFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPackFactoryTransactorRaw struct {
	Contract *IPackFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPackFactory creates a new instance of IPackFactory, bound to a specific deployed contract.
func NewIPackFactory(address common.Address, backend bind.ContractBackend) (*IPackFactory, error) {
	contract, err := bindIPackFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPackFactory{IPackFactoryCaller: IPackFactoryCaller{contract: contract}, IPackFactoryTransactor: IPackFactoryTransactor{contract: contract}, IPackFactoryFilterer: IPackFactoryFilterer{contract: contract}}, nil
}

// NewIPackFactoryCaller creates a new read-only instance of IPackFactory, bound to a specific deployed contract.
func NewIPackFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPackFactoryCaller, error) {
	contract, err := bindIPackFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPackFactoryCaller{contract: contract}, nil
}

// NewIPackFactoryTransactor creates a new write-only instance of IPackFactory, bound to a specific deployed contract.
func NewIPackFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPackFactoryTransactor, error) {
	contract, err := bindIPackFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPackFactoryTransactor{contract: contract}, nil
}

// NewIPackFactoryFilterer creates a new log filterer instance of IPackFactory, bound to a specific deployed contract.
func NewIPackFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPackFactoryFilterer, error) {
	contract, err := bindIPackFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPackFactoryFilterer{contract: contract}, nil
}

// bindIPackFactory binds a generic wrapper to an already deployed contract.
func bindIPackFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPackFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPackFactory *IPackFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPackFactory.Contract.IPackFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPackFactory *IPackFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPackFactory.Contract.IPackFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPackFactory *IPackFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPackFactory.Contract.IPackFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPackFactory *IPackFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPackFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPackFactory *IPackFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPackFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPackFactory *IPackFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPackFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address)
func (_IPackFactory *IPackFactoryTransactor) CreateCardPack(opts *bind.TransactOpts, _totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPackFactory.contract.Transact(opts, "createCardPack", _totalAmount, owner)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address)
func (_IPackFactory *IPackFactorySession) CreateCardPack(_totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPackFactory.Contract.CreateCardPack(&_IPackFactory.TransactOpts, _totalAmount, owner)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address)
func (_IPackFactory *IPackFactoryTransactorSession) CreateCardPack(_totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPackFactory.Contract.CreateCardPack(&_IPackFactory.TransactOpts, _totalAmount, owner)
}
