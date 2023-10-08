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

// IPredictionFactoryMetaData contains all meta data concerning the IPredictionFactory contract.
var IPredictionFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCraftingPredection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPredictionFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IPredictionFactoryMetaData.ABI instead.
var IPredictionFactoryABI = IPredictionFactoryMetaData.ABI

// IPredictionFactory is an auto generated Go binding around an Ethereum contract.
type IPredictionFactory struct {
	IPredictionFactoryCaller     // Read-only binding to the contract
	IPredictionFactoryTransactor // Write-only binding to the contract
	IPredictionFactoryFilterer   // Log filterer for contract events
}

// IPredictionFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPredictionFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPredictionFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPredictionFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPredictionFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPredictionFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPredictionFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPredictionFactorySession struct {
	Contract     *IPredictionFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IPredictionFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPredictionFactoryCallerSession struct {
	Contract *IPredictionFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IPredictionFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPredictionFactoryTransactorSession struct {
	Contract     *IPredictionFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IPredictionFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPredictionFactoryRaw struct {
	Contract *IPredictionFactory // Generic contract binding to access the raw methods on
}

// IPredictionFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPredictionFactoryCallerRaw struct {
	Contract *IPredictionFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IPredictionFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPredictionFactoryTransactorRaw struct {
	Contract *IPredictionFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPredictionFactory creates a new instance of IPredictionFactory, bound to a specific deployed contract.
func NewIPredictionFactory(address common.Address, backend bind.ContractBackend) (*IPredictionFactory, error) {
	contract, err := bindIPredictionFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPredictionFactory{IPredictionFactoryCaller: IPredictionFactoryCaller{contract: contract}, IPredictionFactoryTransactor: IPredictionFactoryTransactor{contract: contract}, IPredictionFactoryFilterer: IPredictionFactoryFilterer{contract: contract}}, nil
}

// NewIPredictionFactoryCaller creates a new read-only instance of IPredictionFactory, bound to a specific deployed contract.
func NewIPredictionFactoryCaller(address common.Address, caller bind.ContractCaller) (*IPredictionFactoryCaller, error) {
	contract, err := bindIPredictionFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPredictionFactoryCaller{contract: contract}, nil
}

// NewIPredictionFactoryTransactor creates a new write-only instance of IPredictionFactory, bound to a specific deployed contract.
func NewIPredictionFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IPredictionFactoryTransactor, error) {
	contract, err := bindIPredictionFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPredictionFactoryTransactor{contract: contract}, nil
}

// NewIPredictionFactoryFilterer creates a new log filterer instance of IPredictionFactory, bound to a specific deployed contract.
func NewIPredictionFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IPredictionFactoryFilterer, error) {
	contract, err := bindIPredictionFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPredictionFactoryFilterer{contract: contract}, nil
}

// bindIPredictionFactory binds a generic wrapper to an already deployed contract.
func bindIPredictionFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPredictionFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPredictionFactory *IPredictionFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPredictionFactory.Contract.IPredictionFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPredictionFactory *IPredictionFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.IPredictionFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPredictionFactory *IPredictionFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.IPredictionFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPredictionFactory *IPredictionFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPredictionFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPredictionFactory *IPredictionFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPredictionFactory *IPredictionFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address)
func (_IPredictionFactory *IPredictionFactoryTransactor) CreateCraftingPredection(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPredictionFactory.contract.Transact(opts, "createCraftingPredection", collectionId, owner)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address)
func (_IPredictionFactory *IPredictionFactorySession) CreateCraftingPredection(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.CreateCraftingPredection(&_IPredictionFactory.TransactOpts, collectionId, owner)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address)
func (_IPredictionFactory *IPredictionFactoryTransactorSession) CreateCraftingPredection(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IPredictionFactory.Contract.CreateCraftingPredection(&_IPredictionFactory.TransactOpts, collectionId, owner)
}
