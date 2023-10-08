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

// ICardPackMetaData contains all meta data concerning the ICardPack contract.
var ICardPackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToOpened\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCard\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ICardPackABI is the input ABI used to generate the binding from.
// Deprecated: Use ICardPackMetaData.ABI instead.
var ICardPackABI = ICardPackMetaData.ABI

// ICardPack is an auto generated Go binding around an Ethereum contract.
type ICardPack struct {
	ICardPackCaller     // Read-only binding to the contract
	ICardPackTransactor // Write-only binding to the contract
	ICardPackFilterer   // Log filterer for contract events
}

// ICardPackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICardPackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICardPackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICardPackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICardPackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICardPackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICardPackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICardPackSession struct {
	Contract     *ICardPack        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICardPackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICardPackCallerSession struct {
	Contract *ICardPackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ICardPackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICardPackTransactorSession struct {
	Contract     *ICardPackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ICardPackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICardPackRaw struct {
	Contract *ICardPack // Generic contract binding to access the raw methods on
}

// ICardPackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICardPackCallerRaw struct {
	Contract *ICardPackCaller // Generic read-only contract binding to access the raw methods on
}

// ICardPackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICardPackTransactorRaw struct {
	Contract *ICardPackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICardPack creates a new instance of ICardPack, bound to a specific deployed contract.
func NewICardPack(address common.Address, backend bind.ContractBackend) (*ICardPack, error) {
	contract, err := bindICardPack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICardPack{ICardPackCaller: ICardPackCaller{contract: contract}, ICardPackTransactor: ICardPackTransactor{contract: contract}, ICardPackFilterer: ICardPackFilterer{contract: contract}}, nil
}

// NewICardPackCaller creates a new read-only instance of ICardPack, bound to a specific deployed contract.
func NewICardPackCaller(address common.Address, caller bind.ContractCaller) (*ICardPackCaller, error) {
	contract, err := bindICardPack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICardPackCaller{contract: contract}, nil
}

// NewICardPackTransactor creates a new write-only instance of ICardPack, bound to a specific deployed contract.
func NewICardPackTransactor(address common.Address, transactor bind.ContractTransactor) (*ICardPackTransactor, error) {
	contract, err := bindICardPack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICardPackTransactor{contract: contract}, nil
}

// NewICardPackFilterer creates a new log filterer instance of ICardPack, bound to a specific deployed contract.
func NewICardPackFilterer(address common.Address, filterer bind.ContractFilterer) (*ICardPackFilterer, error) {
	contract, err := bindICardPack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICardPackFilterer{contract: contract}, nil
}

// bindICardPack binds a generic wrapper to an already deployed contract.
func bindICardPack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICardPackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICardPack *ICardPackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICardPack.Contract.ICardPackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICardPack *ICardPackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICardPack.Contract.ICardPackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICardPack *ICardPackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICardPack.Contract.ICardPackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICardPack *ICardPackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICardPack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICardPack *ICardPackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICardPack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICardPack *ICardPackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICardPack.Contract.contract.Transact(opts, method, params...)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_ICardPack *ICardPackCaller) IsOpened(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _ICardPack.contract.Call(opts, &out, "isOpened", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_ICardPack *ICardPackSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _ICardPack.Contract.IsOpened(&_ICardPack.CallOpts, tokenId)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_ICardPack *ICardPackCallerSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _ICardPack.Contract.IsOpened(&_ICardPack.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ICardPack *ICardPackCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICardPack.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ICardPack *ICardPackSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ICardPack.Contract.OwnerOf(&_ICardPack.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ICardPack *ICardPackCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ICardPack.Contract.OwnerOf(&_ICardPack.CallOpts, tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_ICardPack *ICardPackTransactor) ChangeToOpened(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ICardPack.contract.Transact(opts, "changeToOpened", tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_ICardPack *ICardPackSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _ICardPack.Contract.ChangeToOpened(&_ICardPack.TransactOpts, tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_ICardPack *ICardPackTransactorSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _ICardPack.Contract.ChangeToOpened(&_ICardPack.TransactOpts, tokenId)
}

// CreateCard is a paid mutator transaction binding the contract method 0xed5d02e2.
//
// Solidity: function createCard(string tokenURI, address owner) returns(uint256)
func (_ICardPack *ICardPackTransactor) CreateCard(opts *bind.TransactOpts, tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICardPack.contract.Transact(opts, "createCard", tokenURI, owner)
}

// CreateCard is a paid mutator transaction binding the contract method 0xed5d02e2.
//
// Solidity: function createCard(string tokenURI, address owner) returns(uint256)
func (_ICardPack *ICardPackSession) CreateCard(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICardPack.Contract.CreateCard(&_ICardPack.TransactOpts, tokenURI, owner)
}

// CreateCard is a paid mutator transaction binding the contract method 0xed5d02e2.
//
// Solidity: function createCard(string tokenURI, address owner) returns(uint256)
func (_ICardPack *ICardPackTransactorSession) CreateCard(tokenURI string, owner common.Address) (*types.Transaction, error) {
	return _ICardPack.Contract.CreateCard(&_ICardPack.TransactOpts, tokenURI, owner)
}
