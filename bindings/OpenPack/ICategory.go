// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openpack

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

// ICategoryMetaData contains all meta data concerning the ICategory contract.
var ICategoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCategory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ICategoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ICategoryMetaData.ABI instead.
var ICategoryABI = ICategoryMetaData.ABI

// ICategory is an auto generated Go binding around an Ethereum contract.
type ICategory struct {
	ICategoryCaller     // Read-only binding to the contract
	ICategoryTransactor // Write-only binding to the contract
	ICategoryFilterer   // Log filterer for contract events
}

// ICategoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICategoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICategoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICategoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICategorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICategorySession struct {
	Contract     *ICategory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICategoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICategoryCallerSession struct {
	Contract *ICategoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ICategoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICategoryTransactorSession struct {
	Contract     *ICategoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICategoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICategoryRaw struct {
	Contract *ICategory // Generic contract binding to access the raw methods on
}

// ICategoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICategoryCallerRaw struct {
	Contract *ICategoryCaller // Generic read-only contract binding to access the raw methods on
}

// ICategoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICategoryTransactorRaw struct {
	Contract *ICategoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICategory creates a new instance of ICategory, bound to a specific deployed contract.
func NewICategory(address common.Address, backend bind.ContractBackend) (*ICategory, error) {
	contract, err := bindICategory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICategory{ICategoryCaller: ICategoryCaller{contract: contract}, ICategoryTransactor: ICategoryTransactor{contract: contract}, ICategoryFilterer: ICategoryFilterer{contract: contract}}, nil
}

// NewICategoryCaller creates a new read-only instance of ICategory, bound to a specific deployed contract.
func NewICategoryCaller(address common.Address, caller bind.ContractCaller) (*ICategoryCaller, error) {
	contract, err := bindICategory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICategoryCaller{contract: contract}, nil
}

// NewICategoryTransactor creates a new write-only instance of ICategory, bound to a specific deployed contract.
func NewICategoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ICategoryTransactor, error) {
	contract, err := bindICategory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICategoryTransactor{contract: contract}, nil
}

// NewICategoryFilterer creates a new log filterer instance of ICategory, bound to a specific deployed contract.
func NewICategoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ICategoryFilterer, error) {
	contract, err := bindICategory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICategoryFilterer{contract: contract}, nil
}

// bindICategory binds a generic wrapper to an already deployed contract.
func bindICategory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICategoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICategory *ICategoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICategory.Contract.ICategoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICategory *ICategoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICategory.Contract.ICategoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICategory *ICategoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICategory.Contract.ICategoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICategory *ICategoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICategory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICategory *ICategoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICategory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICategory *ICategoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICategory.Contract.contract.Transact(opts, method, params...)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x06e837f0.
//
// Solidity: function createCategory(string tokenURI, address owner) returns(uint256)
func (_ICategory *ICategoryTransactor) CreateCategory(opts *bind.TransactOpts, tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICategory.contract.Transact(opts, "createCategory", tokenURI, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x06e837f0.
//
// Solidity: function createCategory(string tokenURI, address owner) returns(uint256)
func (_ICategory *ICategorySession) CreateCategory(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICategory.Contract.CreateCategory(&_ICategory.TransactOpts, tokenURI, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x06e837f0.
//
// Solidity: function createCategory(string tokenURI, address owner) returns(uint256)
func (_ICategory *ICategoryTransactorSession) CreateCategory(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICategory.Contract.CreateCategory(&_ICategory.TransactOpts, tokenURI, owner)
}
