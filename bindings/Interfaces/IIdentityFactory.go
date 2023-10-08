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

// IIdentityFactoryMetaData contains all meta data concerning the IIdentityFactory contract.
var IIdentityFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCraftingIdentity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IIdentityFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IIdentityFactoryMetaData.ABI instead.
var IIdentityFactoryABI = IIdentityFactoryMetaData.ABI

// IIdentityFactory is an auto generated Go binding around an Ethereum contract.
type IIdentityFactory struct {
	IIdentityFactoryCaller     // Read-only binding to the contract
	IIdentityFactoryTransactor // Write-only binding to the contract
	IIdentityFactoryFilterer   // Log filterer for contract events
}

// IIdentityFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIdentityFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIdentityFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIdentityFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIdentityFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIdentityFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIdentityFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIdentityFactorySession struct {
	Contract     *IIdentityFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IIdentityFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIdentityFactoryCallerSession struct {
	Contract *IIdentityFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IIdentityFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIdentityFactoryTransactorSession struct {
	Contract     *IIdentityFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IIdentityFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIdentityFactoryRaw struct {
	Contract *IIdentityFactory // Generic contract binding to access the raw methods on
}

// IIdentityFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIdentityFactoryCallerRaw struct {
	Contract *IIdentityFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IIdentityFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIdentityFactoryTransactorRaw struct {
	Contract *IIdentityFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIdentityFactory creates a new instance of IIdentityFactory, bound to a specific deployed contract.
func NewIIdentityFactory(address common.Address, backend bind.ContractBackend) (*IIdentityFactory, error) {
	contract, err := bindIIdentityFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIdentityFactory{IIdentityFactoryCaller: IIdentityFactoryCaller{contract: contract}, IIdentityFactoryTransactor: IIdentityFactoryTransactor{contract: contract}, IIdentityFactoryFilterer: IIdentityFactoryFilterer{contract: contract}}, nil
}

// NewIIdentityFactoryCaller creates a new read-only instance of IIdentityFactory, bound to a specific deployed contract.
func NewIIdentityFactoryCaller(address common.Address, caller bind.ContractCaller) (*IIdentityFactoryCaller, error) {
	contract, err := bindIIdentityFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIdentityFactoryCaller{contract: contract}, nil
}

// NewIIdentityFactoryTransactor creates a new write-only instance of IIdentityFactory, bound to a specific deployed contract.
func NewIIdentityFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IIdentityFactoryTransactor, error) {
	contract, err := bindIIdentityFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIdentityFactoryTransactor{contract: contract}, nil
}

// NewIIdentityFactoryFilterer creates a new log filterer instance of IIdentityFactory, bound to a specific deployed contract.
func NewIIdentityFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IIdentityFactoryFilterer, error) {
	contract, err := bindIIdentityFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIdentityFactoryFilterer{contract: contract}, nil
}

// bindIIdentityFactory binds a generic wrapper to an already deployed contract.
func bindIIdentityFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IIdentityFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIdentityFactory *IIdentityFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIdentityFactory.Contract.IIdentityFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIdentityFactory *IIdentityFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.IIdentityFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIdentityFactory *IIdentityFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.IIdentityFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIdentityFactory *IIdentityFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIdentityFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIdentityFactory *IIdentityFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIdentityFactory *IIdentityFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address)
func (_IIdentityFactory *IIdentityFactoryTransactor) CreateCraftingIdentity(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IIdentityFactory.contract.Transact(opts, "createCraftingIdentity", collectionId, owner)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address)
func (_IIdentityFactory *IIdentityFactorySession) CreateCraftingIdentity(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.CreateCraftingIdentity(&_IIdentityFactory.TransactOpts, collectionId, owner)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address)
func (_IIdentityFactory *IIdentityFactoryTransactorSession) CreateCraftingIdentity(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IIdentityFactory.Contract.CreateCraftingIdentity(&_IIdentityFactory.TransactOpts, collectionId, owner)
}
