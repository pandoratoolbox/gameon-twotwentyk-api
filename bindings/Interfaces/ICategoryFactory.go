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

// ICategoryFactoryMetaData contains all meta data concerning the ICategoryFactory contract.
var ICategoryFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCategory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICategoryFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ICategoryFactoryMetaData.ABI instead.
var ICategoryFactoryABI = ICategoryFactoryMetaData.ABI

// ICategoryFactory is an auto generated Go binding around an Ethereum contract.
type ICategoryFactory struct {
	ICategoryFactoryCaller     // Read-only binding to the contract
	ICategoryFactoryTransactor // Write-only binding to the contract
	ICategoryFactoryFilterer   // Log filterer for contract events
}

// ICategoryFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICategoryFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategoryFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICategoryFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategoryFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICategoryFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategoryFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICategoryFactorySession struct {
	Contract     *ICategoryFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICategoryFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICategoryFactoryCallerSession struct {
	Contract *ICategoryFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ICategoryFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICategoryFactoryTransactorSession struct {
	Contract     *ICategoryFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ICategoryFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICategoryFactoryRaw struct {
	Contract *ICategoryFactory // Generic contract binding to access the raw methods on
}

// ICategoryFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICategoryFactoryCallerRaw struct {
	Contract *ICategoryFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ICategoryFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICategoryFactoryTransactorRaw struct {
	Contract *ICategoryFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICategoryFactory creates a new instance of ICategoryFactory, bound to a specific deployed contract.
func NewICategoryFactory(address common.Address, backend bind.ContractBackend) (*ICategoryFactory, error) {
	contract, err := bindICategoryFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICategoryFactory{ICategoryFactoryCaller: ICategoryFactoryCaller{contract: contract}, ICategoryFactoryTransactor: ICategoryFactoryTransactor{contract: contract}, ICategoryFactoryFilterer: ICategoryFactoryFilterer{contract: contract}}, nil
}

// NewICategoryFactoryCaller creates a new read-only instance of ICategoryFactory, bound to a specific deployed contract.
func NewICategoryFactoryCaller(address common.Address, caller bind.ContractCaller) (*ICategoryFactoryCaller, error) {
	contract, err := bindICategoryFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICategoryFactoryCaller{contract: contract}, nil
}

// NewICategoryFactoryTransactor creates a new write-only instance of ICategoryFactory, bound to a specific deployed contract.
func NewICategoryFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICategoryFactoryTransactor, error) {
	contract, err := bindICategoryFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICategoryFactoryTransactor{contract: contract}, nil
}

// NewICategoryFactoryFilterer creates a new log filterer instance of ICategoryFactory, bound to a specific deployed contract.
func NewICategoryFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICategoryFactoryFilterer, error) {
	contract, err := bindICategoryFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICategoryFactoryFilterer{contract: contract}, nil
}

// bindICategoryFactory binds a generic wrapper to an already deployed contract.
func bindICategoryFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICategoryFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICategoryFactory *ICategoryFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICategoryFactory.Contract.ICategoryFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICategoryFactory *ICategoryFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.ICategoryFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICategoryFactory *ICategoryFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.ICategoryFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICategoryFactory *ICategoryFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICategoryFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICategoryFactory *ICategoryFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICategoryFactory *ICategoryFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address)
func (_ICategoryFactory *ICategoryFactoryTransactor) CreateCategory(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICategoryFactory.contract.Transact(opts, "createCategory", collectionId, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address)
func (_ICategoryFactory *ICategoryFactorySession) CreateCategory(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.CreateCategory(&_ICategoryFactory.TransactOpts, collectionId, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address)
func (_ICategoryFactory *ICategoryFactoryTransactorSession) CreateCategory(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICategoryFactory.Contract.CreateCategory(&_ICategoryFactory.TransactOpts, collectionId, owner)
}
