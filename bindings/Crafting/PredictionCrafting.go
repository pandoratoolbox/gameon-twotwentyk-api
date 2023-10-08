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

// PredictionCraftingMetaData contains all meta data concerning the PredictionCrafting contract.
var PredictionCraftingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"identityTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"triggerTokenId\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CollectionCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggersContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identityTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"triggerTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"craftCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craftingCardContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"getCollectionTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"identitiesContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggersContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"win\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"won\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PredictionCraftingABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionCraftingMetaData.ABI instead.
var PredictionCraftingABI = PredictionCraftingMetaData.ABI

// PredictionCrafting is an auto generated Go binding around an Ethereum contract.
type PredictionCrafting struct {
	PredictionCraftingCaller     // Read-only binding to the contract
	PredictionCraftingTransactor // Write-only binding to the contract
	PredictionCraftingFilterer   // Log filterer for contract events
}

// PredictionCraftingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionCraftingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionCraftingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionCraftingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionCraftingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionCraftingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionCraftingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionCraftingSession struct {
	Contract     *PredictionCrafting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PredictionCraftingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionCraftingCallerSession struct {
	Contract *PredictionCraftingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PredictionCraftingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionCraftingTransactorSession struct {
	Contract     *PredictionCraftingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PredictionCraftingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionCraftingRaw struct {
	Contract *PredictionCrafting // Generic contract binding to access the raw methods on
}

// PredictionCraftingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionCraftingCallerRaw struct {
	Contract *PredictionCraftingCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionCraftingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionCraftingTransactorRaw struct {
	Contract *PredictionCraftingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredictionCrafting creates a new instance of PredictionCrafting, bound to a specific deployed contract.
func NewPredictionCrafting(address common.Address, backend bind.ContractBackend) (*PredictionCrafting, error) {
	contract, err := bindPredictionCrafting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PredictionCrafting{PredictionCraftingCaller: PredictionCraftingCaller{contract: contract}, PredictionCraftingTransactor: PredictionCraftingTransactor{contract: contract}, PredictionCraftingFilterer: PredictionCraftingFilterer{contract: contract}}, nil
}

// NewPredictionCraftingCaller creates a new read-only instance of PredictionCrafting, bound to a specific deployed contract.
func NewPredictionCraftingCaller(address common.Address, caller bind.ContractCaller) (*PredictionCraftingCaller, error) {
	contract, err := bindPredictionCrafting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingCaller{contract: contract}, nil
}

// NewPredictionCraftingTransactor creates a new write-only instance of PredictionCrafting, bound to a specific deployed contract.
func NewPredictionCraftingTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionCraftingTransactor, error) {
	contract, err := bindPredictionCrafting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingTransactor{contract: contract}, nil
}

// NewPredictionCraftingFilterer creates a new log filterer instance of PredictionCrafting, bound to a specific deployed contract.
func NewPredictionCraftingFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionCraftingFilterer, error) {
	contract, err := bindPredictionCrafting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingFilterer{contract: contract}, nil
}

// bindPredictionCrafting binds a generic wrapper to an already deployed contract.
func bindPredictionCrafting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PredictionCraftingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionCrafting *PredictionCraftingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionCrafting.Contract.PredictionCraftingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionCrafting *PredictionCraftingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.PredictionCraftingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionCrafting *PredictionCraftingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.PredictionCraftingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionCrafting *PredictionCraftingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionCrafting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionCrafting *PredictionCraftingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionCrafting *PredictionCraftingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PredictionCrafting *PredictionCraftingCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PredictionCrafting *PredictionCraftingSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PredictionCrafting.Contract.BalanceOf(&_PredictionCrafting.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PredictionCrafting *PredictionCraftingCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PredictionCrafting.Contract.BalanceOf(&_PredictionCrafting.CallOpts, owner)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_PredictionCrafting *PredictionCraftingCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_PredictionCrafting *PredictionCraftingSession) CollectionId() (*big.Int, error) {
	return _PredictionCrafting.Contract.CollectionId(&_PredictionCrafting.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_PredictionCrafting *PredictionCraftingCallerSession) CollectionId() (*big.Int, error) {
	return _PredictionCrafting.Contract.CollectionId(&_PredictionCrafting.CallOpts)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCaller) Crafted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "crafted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingSession) Crafted(arg0 *big.Int) (bool, error) {
	return _PredictionCrafting.Contract.Crafted(&_PredictionCrafting.CallOpts, arg0)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCallerSession) Crafted(arg0 *big.Int) (bool, error) {
	return _PredictionCrafting.Contract.Crafted(&_PredictionCrafting.CallOpts, arg0)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) CraftingCardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "craftingCardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) CraftingCardContract() (common.Address, error) {
	return _PredictionCrafting.Contract.CraftingCardContract(&_PredictionCrafting.CallOpts)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) CraftingCardContract() (common.Address, error) {
	return _PredictionCrafting.Contract.CraftingCardContract(&_PredictionCrafting.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PredictionCrafting.Contract.GetApproved(&_PredictionCrafting.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _PredictionCrafting.Contract.GetApproved(&_PredictionCrafting.CallOpts, tokenId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_PredictionCrafting *PredictionCraftingCaller) GetCollectionTokens(opts *bind.CallOpts, collectionId *big.Int) (*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "getCollectionTokens", collectionId)

	if err != nil {
		return *new(*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_PredictionCrafting *PredictionCraftingSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, []*big.Int, error) {
	return _PredictionCrafting.Contract.GetCollectionTokens(&_PredictionCrafting.CallOpts, collectionId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256[])
func (_PredictionCrafting *PredictionCraftingCallerSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, []*big.Int, error) {
	return _PredictionCrafting.Contract.GetCollectionTokens(&_PredictionCrafting.CallOpts, collectionId)
}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) IdentitiesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "identitiesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) IdentitiesContract() (common.Address, error) {
	return _PredictionCrafting.Contract.IdentitiesContract(&_PredictionCrafting.CallOpts)
}

// IdentitiesContract is a free data retrieval call binding the contract method 0xbff246c2.
//
// Solidity: function identitiesContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) IdentitiesContract() (common.Address, error) {
	return _PredictionCrafting.Contract.IdentitiesContract(&_PredictionCrafting.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PredictionCrafting *PredictionCraftingSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PredictionCrafting.Contract.IsApprovedForAll(&_PredictionCrafting.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _PredictionCrafting.Contract.IsApprovedForAll(&_PredictionCrafting.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PredictionCrafting *PredictionCraftingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PredictionCrafting *PredictionCraftingSession) Name() (string, error) {
	return _PredictionCrafting.Contract.Name(&_PredictionCrafting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PredictionCrafting *PredictionCraftingCallerSession) Name() (string, error) {
	return _PredictionCrafting.Contract.Name(&_PredictionCrafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) Owner() (common.Address, error) {
	return _PredictionCrafting.Contract.Owner(&_PredictionCrafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) Owner() (common.Address, error) {
	return _PredictionCrafting.Contract.Owner(&_PredictionCrafting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PredictionCrafting.Contract.OwnerOf(&_PredictionCrafting.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _PredictionCrafting.Contract.OwnerOf(&_PredictionCrafting.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PredictionCrafting *PredictionCraftingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PredictionCrafting.Contract.SupportsInterface(&_PredictionCrafting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PredictionCrafting.Contract.SupportsInterface(&_PredictionCrafting.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PredictionCrafting *PredictionCraftingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PredictionCrafting *PredictionCraftingSession) Symbol() (string, error) {
	return _PredictionCrafting.Contract.Symbol(&_PredictionCrafting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PredictionCrafting *PredictionCraftingCallerSession) Symbol() (string, error) {
	return _PredictionCrafting.Contract.Symbol(&_PredictionCrafting.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PredictionCrafting *PredictionCraftingCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PredictionCrafting *PredictionCraftingSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PredictionCrafting.Contract.TokenURI(&_PredictionCrafting.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_PredictionCrafting *PredictionCraftingCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _PredictionCrafting.Contract.TokenURI(&_PredictionCrafting.CallOpts, tokenId)
}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCaller) TriggersContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "triggersContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingSession) TriggersContract() (common.Address, error) {
	return _PredictionCrafting.Contract.TriggersContract(&_PredictionCrafting.CallOpts)
}

// TriggersContract is a free data retrieval call binding the contract method 0x3f254c60.
//
// Solidity: function triggersContract() view returns(address)
func (_PredictionCrafting *PredictionCraftingCallerSession) TriggersContract() (common.Address, error) {
	return _PredictionCrafting.Contract.TriggersContract(&_PredictionCrafting.CallOpts)
}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCaller) Won(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PredictionCrafting.contract.Call(opts, &out, "won", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingSession) Won(arg0 common.Address) (bool, error) {
	return _PredictionCrafting.Contract.Won(&_PredictionCrafting.CallOpts, arg0)
}

// Won is a free data retrieval call binding the contract method 0x71ee1aba.
//
// Solidity: function won(address ) view returns(bool)
func (_PredictionCrafting *PredictionCraftingCallerSession) Won(arg0 common.Address) (bool, error) {
	return _PredictionCrafting.Contract.Won(&_PredictionCrafting.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.Approve(&_PredictionCrafting.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.Approve(&_PredictionCrafting.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionCrafting *PredictionCraftingSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.ChangeAdmin(&_PredictionCrafting.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.ChangeAdmin(&_PredictionCrafting.TransactOpts, _newAdmin)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x951db03f.
//
// Solidity: function craftCollection(address _identityContract, address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) CraftCollection(opts *bind.TransactOpts, _identityContract common.Address, _triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "craftCollection", _identityContract, _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x951db03f.
//
// Solidity: function craftCollection(address _identityContract, address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_PredictionCrafting *PredictionCraftingSession) CraftCollection(_identityContract common.Address, _triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.CraftCollection(&_PredictionCrafting.TransactOpts, _identityContract, _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0x951db03f.
//
// Solidity: function craftCollection(address _identityContract, address _triggersContract, address _craftingCardContract, uint256 identityTokenId, uint256[] triggerTokenIds, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) CraftCollection(_identityContract common.Address, _triggersContract common.Address, _craftingCardContract common.Address, identityTokenId *big.Int, triggerTokenIds []*big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.CraftCollection(&_PredictionCrafting.TransactOpts, _identityContract, _triggersContract, _craftingCardContract, identityTokenId, triggerTokenIds, craftingCardTokenId, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionCrafting *PredictionCraftingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionCrafting *PredictionCraftingSession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionCrafting.Contract.RenounceOwnership(&_PredictionCrafting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionCrafting.Contract.RenounceOwnership(&_PredictionCrafting.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SafeTransferFrom(&_PredictionCrafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SafeTransferFrom(&_PredictionCrafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PredictionCrafting *PredictionCraftingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SafeTransferFrom0(&_PredictionCrafting.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SafeTransferFrom0(&_PredictionCrafting.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PredictionCrafting *PredictionCraftingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SetApprovalForAll(&_PredictionCrafting.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.SetApprovalForAll(&_PredictionCrafting.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.TransferFrom(&_PredictionCrafting.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.TransferFrom(&_PredictionCrafting.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionCrafting *PredictionCraftingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.TransferOwnership(&_PredictionCrafting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.TransferOwnership(&_PredictionCrafting.TransactOpts, newOwner)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_PredictionCrafting *PredictionCraftingTransactor) Win(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.contract.Transact(opts, "win", _address)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_PredictionCrafting *PredictionCraftingSession) Win(_address common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.Win(&_PredictionCrafting.TransactOpts, _address)
}

// Win is a paid mutator transaction binding the contract method 0xa34cc845.
//
// Solidity: function win(address _address) returns()
func (_PredictionCrafting *PredictionCraftingTransactorSession) Win(_address common.Address) (*types.Transaction, error) {
	return _PredictionCrafting.Contract.Win(&_PredictionCrafting.TransactOpts, _address)
}

// PredictionCraftingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PredictionCrafting contract.
type PredictionCraftingApprovalIterator struct {
	Event *PredictionCraftingApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PredictionCraftingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCraftingApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PredictionCraftingApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PredictionCraftingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCraftingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCraftingApproval represents a Approval event raised by the PredictionCrafting contract.
type PredictionCraftingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*PredictionCraftingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingApprovalIterator{contract: _PredictionCrafting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PredictionCraftingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCraftingApproval)
				if err := _PredictionCrafting.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) ParseApproval(log types.Log) (*PredictionCraftingApproval, error) {
	event := new(PredictionCraftingApproval)
	if err := _PredictionCrafting.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionCraftingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PredictionCrafting contract.
type PredictionCraftingApprovalForAllIterator struct {
	Event *PredictionCraftingApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PredictionCraftingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCraftingApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PredictionCraftingApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PredictionCraftingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCraftingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCraftingApprovalForAll represents a ApprovalForAll event raised by the PredictionCrafting contract.
type PredictionCraftingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PredictionCrafting *PredictionCraftingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PredictionCraftingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PredictionCrafting.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingApprovalForAllIterator{contract: _PredictionCrafting.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PredictionCrafting *PredictionCraftingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PredictionCraftingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PredictionCrafting.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCraftingApprovalForAll)
				if err := _PredictionCrafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_PredictionCrafting *PredictionCraftingFilterer) ParseApprovalForAll(log types.Log) (*PredictionCraftingApprovalForAll, error) {
	event := new(PredictionCraftingApprovalForAll)
	if err := _PredictionCrafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionCraftingCollectionCraftedIterator is returned from FilterCollectionCrafted and is used to iterate over the raw logs and unpacked data for CollectionCrafted events raised by the PredictionCrafting contract.
type PredictionCraftingCollectionCraftedIterator struct {
	Event *PredictionCraftingCollectionCrafted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PredictionCraftingCollectionCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCraftingCollectionCrafted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PredictionCraftingCollectionCrafted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PredictionCraftingCollectionCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCraftingCollectionCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCraftingCollectionCrafted represents a CollectionCrafted event raised by the PredictionCrafting contract.
type PredictionCraftingCollectionCrafted struct {
	CollectionId    *big.Int
	IdentityTokenId *big.Int
	TriggerTokenId  []*big.Int
	TokenURI        string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectionCrafted is a free log retrieval operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_PredictionCrafting *PredictionCraftingFilterer) FilterCollectionCrafted(opts *bind.FilterOpts, collectionId []*big.Int) (*PredictionCraftingCollectionCraftedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.FilterLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingCollectionCraftedIterator{contract: _PredictionCrafting.contract, event: "CollectionCrafted", logs: logs, sub: sub}, nil
}

// WatchCollectionCrafted is a free log subscription operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_PredictionCrafting *PredictionCraftingFilterer) WatchCollectionCrafted(opts *bind.WatchOpts, sink chan<- *PredictionCraftingCollectionCrafted, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.WatchLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCraftingCollectionCrafted)
				if err := _PredictionCrafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectionCrafted is a log parse operation binding the contract event 0x59392a125204831e61d97876a45b9667b68def2a5e303e6068932be0fd52372d.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 identityTokenId, uint256[] triggerTokenId, string tokenURI)
func (_PredictionCrafting *PredictionCraftingFilterer) ParseCollectionCrafted(log types.Log) (*PredictionCraftingCollectionCrafted, error) {
	event := new(PredictionCraftingCollectionCrafted)
	if err := _PredictionCrafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionCraftingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PredictionCrafting contract.
type PredictionCraftingOwnershipTransferredIterator struct {
	Event *PredictionCraftingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PredictionCraftingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCraftingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PredictionCraftingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PredictionCraftingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCraftingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCraftingOwnershipTransferred represents a OwnershipTransferred event raised by the PredictionCrafting contract.
type PredictionCraftingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionCrafting *PredictionCraftingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PredictionCraftingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionCrafting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingOwnershipTransferredIterator{contract: _PredictionCrafting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionCrafting *PredictionCraftingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PredictionCraftingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionCrafting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCraftingOwnershipTransferred)
				if err := _PredictionCrafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionCrafting *PredictionCraftingFilterer) ParseOwnershipTransferred(log types.Log) (*PredictionCraftingOwnershipTransferred, error) {
	event := new(PredictionCraftingOwnershipTransferred)
	if err := _PredictionCrafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionCraftingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PredictionCrafting contract.
type PredictionCraftingTransferIterator struct {
	Event *PredictionCraftingTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PredictionCraftingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionCraftingTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PredictionCraftingTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PredictionCraftingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionCraftingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionCraftingTransfer represents a Transfer event raised by the PredictionCrafting contract.
type PredictionCraftingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*PredictionCraftingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PredictionCraftingTransferIterator{contract: _PredictionCrafting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PredictionCraftingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _PredictionCrafting.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionCraftingTransfer)
				if err := _PredictionCrafting.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_PredictionCrafting *PredictionCraftingFilterer) ParseTransfer(log types.Log) (*PredictionCraftingTransfer, error) {
	event := new(PredictionCraftingTransfer)
	if err := _PredictionCrafting.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
