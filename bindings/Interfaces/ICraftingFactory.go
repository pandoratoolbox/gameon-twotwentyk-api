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

// ICraftingFactoryMetaData contains all meta data concerning the ICraftingFactory contract.
var ICraftingFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCrafting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICraftingFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ICraftingFactoryMetaData.ABI instead.
var ICraftingFactoryABI = ICraftingFactoryMetaData.ABI

// ICraftingFactory is an auto generated Go binding around an Ethereum contract.
type ICraftingFactory struct {
	ICraftingFactoryCaller     // Read-only binding to the contract
	ICraftingFactoryTransactor // Write-only binding to the contract
	ICraftingFactoryFilterer   // Log filterer for contract events
}

// ICraftingFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICraftingFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICraftingFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICraftingFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICraftingFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICraftingFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICraftingFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICraftingFactorySession struct {
	Contract     *ICraftingFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICraftingFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICraftingFactoryCallerSession struct {
	Contract *ICraftingFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ICraftingFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICraftingFactoryTransactorSession struct {
	Contract     *ICraftingFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ICraftingFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICraftingFactoryRaw struct {
	Contract *ICraftingFactory // Generic contract binding to access the raw methods on
}

// ICraftingFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICraftingFactoryCallerRaw struct {
	Contract *ICraftingFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ICraftingFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICraftingFactoryTransactorRaw struct {
	Contract *ICraftingFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICraftingFactory creates a new instance of ICraftingFactory, bound to a specific deployed contract.
func NewICraftingFactory(address common.Address, backend bind.ContractBackend) (*ICraftingFactory, error) {
	contract, err := bindICraftingFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICraftingFactory{ICraftingFactoryCaller: ICraftingFactoryCaller{contract: contract}, ICraftingFactoryTransactor: ICraftingFactoryTransactor{contract: contract}, ICraftingFactoryFilterer: ICraftingFactoryFilterer{contract: contract}}, nil
}

// NewICraftingFactoryCaller creates a new read-only instance of ICraftingFactory, bound to a specific deployed contract.
func NewICraftingFactoryCaller(address common.Address, caller bind.ContractCaller) (*ICraftingFactoryCaller, error) {
	contract, err := bindICraftingFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICraftingFactoryCaller{contract: contract}, nil
}

// NewICraftingFactoryTransactor creates a new write-only instance of ICraftingFactory, bound to a specific deployed contract.
func NewICraftingFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICraftingFactoryTransactor, error) {
	contract, err := bindICraftingFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICraftingFactoryTransactor{contract: contract}, nil
}

// NewICraftingFactoryFilterer creates a new log filterer instance of ICraftingFactory, bound to a specific deployed contract.
func NewICraftingFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICraftingFactoryFilterer, error) {
	contract, err := bindICraftingFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICraftingFactoryFilterer{contract: contract}, nil
}

// bindICraftingFactory binds a generic wrapper to an already deployed contract.
func bindICraftingFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICraftingFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICraftingFactory *ICraftingFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICraftingFactory.Contract.ICraftingFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICraftingFactory *ICraftingFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.ICraftingFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICraftingFactory *ICraftingFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.ICraftingFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICraftingFactory *ICraftingFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICraftingFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICraftingFactory *ICraftingFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICraftingFactory *ICraftingFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address)
func (_ICraftingFactory *ICraftingFactoryTransactor) CreateCrafting(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICraftingFactory.contract.Transact(opts, "createCrafting", collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address)
func (_ICraftingFactory *ICraftingFactorySession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.CreateCrafting(&_ICraftingFactory.TransactOpts, collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address)
func (_ICraftingFactory *ICraftingFactoryTransactorSession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ICraftingFactory.Contract.CreateCrafting(&_ICraftingFactory.TransactOpts, collectionId, owner)
}
