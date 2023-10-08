// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crafting

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

// IPartMetaData contains all meta data concerning the IPart contract.
var IPartMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IPartABI is the input ABI used to generate the binding from.
// Deprecated: Use IPartMetaData.ABI instead.
var IPartABI = IPartMetaData.ABI

// IPart is an auto generated Go binding around an Ethereum contract.
type IPart struct {
	IPartCaller     // Read-only binding to the contract
	IPartTransactor // Write-only binding to the contract
	IPartFilterer   // Log filterer for contract events
}

// IPartCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPartCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPartTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPartTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPartFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPartFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPartSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPartSession struct {
	Contract     *IPart            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPartCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPartCallerSession struct {
	Contract *IPartCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPartTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPartTransactorSession struct {
	Contract     *IPartTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPartRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPartRaw struct {
	Contract *IPart // Generic contract binding to access the raw methods on
}

// IPartCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPartCallerRaw struct {
	Contract *IPartCaller // Generic read-only contract binding to access the raw methods on
}

// IPartTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPartTransactorRaw struct {
	Contract *IPartTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPart creates a new instance of IPart, bound to a specific deployed contract.
func NewIPart(address common.Address, backend bind.ContractBackend) (*IPart, error) {
	contract, err := bindIPart(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPart{IPartCaller: IPartCaller{contract: contract}, IPartTransactor: IPartTransactor{contract: contract}, IPartFilterer: IPartFilterer{contract: contract}}, nil
}

// NewIPartCaller creates a new read-only instance of IPart, bound to a specific deployed contract.
func NewIPartCaller(address common.Address, caller bind.ContractCaller) (*IPartCaller, error) {
	contract, err := bindIPart(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPartCaller{contract: contract}, nil
}

// NewIPartTransactor creates a new write-only instance of IPart, bound to a specific deployed contract.
func NewIPartTransactor(address common.Address, transactor bind.ContractTransactor) (*IPartTransactor, error) {
	contract, err := bindIPart(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPartTransactor{contract: contract}, nil
}

// NewIPartFilterer creates a new log filterer instance of IPart, bound to a specific deployed contract.
func NewIPartFilterer(address common.Address, filterer bind.ContractFilterer) (*IPartFilterer, error) {
	contract, err := bindIPart(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPartFilterer{contract: contract}, nil
}

// bindIPart binds a generic wrapper to an already deployed contract.
func bindIPart(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPartMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPart *IPartRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPart.Contract.IPartCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPart *IPartRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPart.Contract.IPartTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPart *IPartRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPart.Contract.IPartTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPart *IPartCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPart.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPart *IPartTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPart.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPart *IPartTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPart.Contract.contract.Transact(opts, method, params...)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_IPart *IPartCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _IPart.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_IPart *IPartSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _IPart.Contract.IsCrafted(&_IPart.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_IPart *IPartCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _IPart.Contract.IsCrafted(&_IPart.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPart *IPartCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPart.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPart *IPartSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IPart.Contract.OwnerOf(&_IPart.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPart *IPartCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IPart.Contract.OwnerOf(&_IPart.CallOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_IPart *IPartTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_IPart *IPartSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.Contract.ChangeToCrafted(&_IPart.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_IPart *IPartTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.Contract.ChangeToCrafted(&_IPart.TransactOpts, tokenId)
}

// ChangeToCrafted2 is a paid mutator transaction binding the contract method 0xe851d37b.
//
// Solidity: function changeToCrafted2(uint256 tokenId) returns()
func (_IPart *IPartTransactor) ChangeToCrafted2(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.contract.Transact(opts, "changeToCrafted2", tokenId)
}

// ChangeToCrafted2 is a paid mutator transaction binding the contract method 0xe851d37b.
//
// Solidity: function changeToCrafted2(uint256 tokenId) returns()
func (_IPart *IPartSession) ChangeToCrafted2(tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.Contract.ChangeToCrafted2(&_IPart.TransactOpts, tokenId)
}

// ChangeToCrafted2 is a paid mutator transaction binding the contract method 0xe851d37b.
//
// Solidity: function changeToCrafted2(uint256 tokenId) returns()
func (_IPart *IPartTransactorSession) ChangeToCrafted2(tokenId *big.Int) (*types.Transaction, error) {
	return _IPart.Contract.ChangeToCrafted2(&_IPart.TransactOpts, tokenId)
}
