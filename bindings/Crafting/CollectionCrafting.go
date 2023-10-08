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

// CollectionCraftingMetaData contains all meta data concerning the CollectionCrafting contract.
var CollectionCraftingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CollectionCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categoryContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dayMonthContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dayMonthTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yearTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"categoryTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"craftingCardTokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"craftCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"crafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"craftingCardContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dayMonthContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"}],\"name\":\"getCollectionTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yearContract\",\"outputs\":[{\"internalType\":\"contractIPart\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CollectionCraftingABI is the input ABI used to generate the binding from.
// Deprecated: Use CollectionCraftingMetaData.ABI instead.
var CollectionCraftingABI = CollectionCraftingMetaData.ABI

// CollectionCrafting is an auto generated Go binding around an Ethereum contract.
type CollectionCrafting struct {
	CollectionCraftingCaller     // Read-only binding to the contract
	CollectionCraftingTransactor // Write-only binding to the contract
	CollectionCraftingFilterer   // Log filterer for contract events
}

// CollectionCraftingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionCraftingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionCraftingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionCraftingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionCraftingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionCraftingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionCraftingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionCraftingSession struct {
	Contract     *CollectionCrafting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CollectionCraftingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionCraftingCallerSession struct {
	Contract *CollectionCraftingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CollectionCraftingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionCraftingTransactorSession struct {
	Contract     *CollectionCraftingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CollectionCraftingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionCraftingRaw struct {
	Contract *CollectionCrafting // Generic contract binding to access the raw methods on
}

// CollectionCraftingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionCraftingCallerRaw struct {
	Contract *CollectionCraftingCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionCraftingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionCraftingTransactorRaw struct {
	Contract *CollectionCraftingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionCrafting creates a new instance of CollectionCrafting, bound to a specific deployed contract.
func NewCollectionCrafting(address common.Address, backend bind.ContractBackend) (*CollectionCrafting, error) {
	contract, err := bindCollectionCrafting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionCrafting{CollectionCraftingCaller: CollectionCraftingCaller{contract: contract}, CollectionCraftingTransactor: CollectionCraftingTransactor{contract: contract}, CollectionCraftingFilterer: CollectionCraftingFilterer{contract: contract}}, nil
}

// NewCollectionCraftingCaller creates a new read-only instance of CollectionCrafting, bound to a specific deployed contract.
func NewCollectionCraftingCaller(address common.Address, caller bind.ContractCaller) (*CollectionCraftingCaller, error) {
	contract, err := bindCollectionCrafting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingCaller{contract: contract}, nil
}

// NewCollectionCraftingTransactor creates a new write-only instance of CollectionCrafting, bound to a specific deployed contract.
func NewCollectionCraftingTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionCraftingTransactor, error) {
	contract, err := bindCollectionCrafting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingTransactor{contract: contract}, nil
}

// NewCollectionCraftingFilterer creates a new log filterer instance of CollectionCrafting, bound to a specific deployed contract.
func NewCollectionCraftingFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionCraftingFilterer, error) {
	contract, err := bindCollectionCrafting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingFilterer{contract: contract}, nil
}

// bindCollectionCrafting binds a generic wrapper to an already deployed contract.
func bindCollectionCrafting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollectionCraftingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionCrafting *CollectionCraftingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionCrafting.Contract.CollectionCraftingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionCrafting *CollectionCraftingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.CollectionCraftingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionCrafting *CollectionCraftingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.CollectionCraftingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionCrafting *CollectionCraftingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionCrafting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionCrafting *CollectionCraftingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionCrafting *CollectionCraftingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CollectionCrafting *CollectionCraftingCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CollectionCrafting *CollectionCraftingSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CollectionCrafting.Contract.BalanceOf(&_CollectionCrafting.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CollectionCrafting *CollectionCraftingCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CollectionCrafting.Contract.BalanceOf(&_CollectionCrafting.CallOpts, owner)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) CategoryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "categoryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) CategoryContract() (common.Address, error) {
	return _CollectionCrafting.Contract.CategoryContract(&_CollectionCrafting.CallOpts)
}

// CategoryContract is a free data retrieval call binding the contract method 0xe4f0bdcc.
//
// Solidity: function categoryContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) CategoryContract() (common.Address, error) {
	return _CollectionCrafting.Contract.CategoryContract(&_CollectionCrafting.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_CollectionCrafting *CollectionCraftingCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_CollectionCrafting *CollectionCraftingSession) CollectionId() (*big.Int, error) {
	return _CollectionCrafting.Contract.CollectionId(&_CollectionCrafting.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_CollectionCrafting *CollectionCraftingCallerSession) CollectionId() (*big.Int, error) {
	return _CollectionCrafting.Contract.CollectionId(&_CollectionCrafting.CallOpts)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCaller) Crafted(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "crafted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_CollectionCrafting *CollectionCraftingSession) Crafted(arg0 *big.Int) (bool, error) {
	return _CollectionCrafting.Contract.Crafted(&_CollectionCrafting.CallOpts, arg0)
}

// Crafted is a free data retrieval call binding the contract method 0xcac4cc29.
//
// Solidity: function crafted(uint256 ) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCallerSession) Crafted(arg0 *big.Int) (bool, error) {
	return _CollectionCrafting.Contract.Crafted(&_CollectionCrafting.CallOpts, arg0)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) CraftingCardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "craftingCardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) CraftingCardContract() (common.Address, error) {
	return _CollectionCrafting.Contract.CraftingCardContract(&_CollectionCrafting.CallOpts)
}

// CraftingCardContract is a free data retrieval call binding the contract method 0x2c948296.
//
// Solidity: function craftingCardContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) CraftingCardContract() (common.Address, error) {
	return _CollectionCrafting.Contract.CraftingCardContract(&_CollectionCrafting.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) DayMonthContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "dayMonthContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) DayMonthContract() (common.Address, error) {
	return _CollectionCrafting.Contract.DayMonthContract(&_CollectionCrafting.CallOpts)
}

// DayMonthContract is a free data retrieval call binding the contract method 0x9829ba1c.
//
// Solidity: function dayMonthContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) DayMonthContract() (common.Address, error) {
	return _CollectionCrafting.Contract.DayMonthContract(&_CollectionCrafting.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CollectionCrafting.Contract.GetApproved(&_CollectionCrafting.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CollectionCrafting.Contract.GetApproved(&_CollectionCrafting.CallOpts, tokenId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_CollectionCrafting *CollectionCraftingCaller) GetCollectionTokens(opts *bind.CallOpts, collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "getCollectionTokens", collectionId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_CollectionCrafting *CollectionCraftingSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CollectionCrafting.Contract.GetCollectionTokens(&_CollectionCrafting.CallOpts, collectionId)
}

// GetCollectionTokens is a free data retrieval call binding the contract method 0x2bb88056.
//
// Solidity: function getCollectionTokens(uint256 collectionId) view returns(uint256, uint256, uint256, uint256)
func (_CollectionCrafting *CollectionCraftingCallerSession) GetCollectionTokens(collectionId *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CollectionCrafting.Contract.GetCollectionTokens(&_CollectionCrafting.CallOpts, collectionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CollectionCrafting *CollectionCraftingSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CollectionCrafting.Contract.IsApprovedForAll(&_CollectionCrafting.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CollectionCrafting.Contract.IsApprovedForAll(&_CollectionCrafting.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _CollectionCrafting.Contract.IsCrafted(&_CollectionCrafting.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _CollectionCrafting.Contract.IsCrafted(&_CollectionCrafting.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionCrafting *CollectionCraftingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionCrafting *CollectionCraftingSession) Name() (string, error) {
	return _CollectionCrafting.Contract.Name(&_CollectionCrafting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CollectionCrafting *CollectionCraftingCallerSession) Name() (string, error) {
	return _CollectionCrafting.Contract.Name(&_CollectionCrafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) Owner() (common.Address, error) {
	return _CollectionCrafting.Contract.Owner(&_CollectionCrafting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) Owner() (common.Address, error) {
	return _CollectionCrafting.Contract.Owner(&_CollectionCrafting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CollectionCrafting.Contract.OwnerOf(&_CollectionCrafting.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CollectionCrafting.Contract.OwnerOf(&_CollectionCrafting.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CollectionCrafting.Contract.SupportsInterface(&_CollectionCrafting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CollectionCrafting *CollectionCraftingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CollectionCrafting.Contract.SupportsInterface(&_CollectionCrafting.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionCrafting *CollectionCraftingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionCrafting *CollectionCraftingSession) Symbol() (string, error) {
	return _CollectionCrafting.Contract.Symbol(&_CollectionCrafting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CollectionCrafting *CollectionCraftingCallerSession) Symbol() (string, error) {
	return _CollectionCrafting.Contract.Symbol(&_CollectionCrafting.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CollectionCrafting *CollectionCraftingCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CollectionCrafting *CollectionCraftingSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CollectionCrafting.Contract.TokenURI(&_CollectionCrafting.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CollectionCrafting *CollectionCraftingCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CollectionCrafting.Contract.TokenURI(&_CollectionCrafting.CallOpts, tokenId)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCaller) YearContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionCrafting.contract.Call(opts, &out, "yearContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingSession) YearContract() (common.Address, error) {
	return _CollectionCrafting.Contract.YearContract(&_CollectionCrafting.CallOpts)
}

// YearContract is a free data retrieval call binding the contract method 0x7230f8d4.
//
// Solidity: function yearContract() view returns(address)
func (_CollectionCrafting *CollectionCraftingCallerSession) YearContract() (common.Address, error) {
	return _CollectionCrafting.Contract.YearContract(&_CollectionCrafting.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.Approve(&_CollectionCrafting.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.Approve(&_CollectionCrafting.TransactOpts, to, tokenId)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CollectionCrafting *CollectionCraftingSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeAdmin(&_CollectionCrafting.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeAdmin(&_CollectionCrafting.TransactOpts, _newAdmin)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_CollectionCrafting *CollectionCraftingSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeCrafter(&_CollectionCrafting.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeCrafter(&_CollectionCrafting.TransactOpts, _newCrafter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeToCrafted(&_CollectionCrafting.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.ChangeToCrafted(&_CollectionCrafting.TransactOpts, tokenId)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) CraftCollection(opts *bind.TransactOpts, _dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "craftCollection", _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_CollectionCrafting *CollectionCraftingSession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.CraftCollection(&_CollectionCrafting.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// CraftCollection is a paid mutator transaction binding the contract method 0xbf0837fe.
//
// Solidity: function craftCollection(address _dayMonthContract, address _yearContract, address _categoryContract, address _craftingCardContract, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string _tokenURI) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) CraftCollection(_dayMonthContract common.Address, _yearContract common.Address, _categoryContract common.Address, _craftingCardContract common.Address, dayMonthTokenId *big.Int, yearTokenId *big.Int, categoryTokenId *big.Int, craftingCardTokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.CraftCollection(&_CollectionCrafting.TransactOpts, _dayMonthContract, _yearContract, _categoryContract, _craftingCardContract, dayMonthTokenId, yearTokenId, categoryTokenId, craftingCardTokenId, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionCrafting *CollectionCraftingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionCrafting *CollectionCraftingSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollectionCrafting.Contract.RenounceOwnership(&_CollectionCrafting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollectionCrafting.Contract.RenounceOwnership(&_CollectionCrafting.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SafeTransferFrom(&_CollectionCrafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SafeTransferFrom(&_CollectionCrafting.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CollectionCrafting *CollectionCraftingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SafeTransferFrom0(&_CollectionCrafting.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SafeTransferFrom0(&_CollectionCrafting.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CollectionCrafting *CollectionCraftingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SetApprovalForAll(&_CollectionCrafting.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.SetApprovalForAll(&_CollectionCrafting.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.TransferFrom(&_CollectionCrafting.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.TransferFrom(&_CollectionCrafting.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionCrafting *CollectionCraftingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionCrafting *CollectionCraftingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.TransferOwnership(&_CollectionCrafting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionCrafting *CollectionCraftingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollectionCrafting.Contract.TransferOwnership(&_CollectionCrafting.TransactOpts, newOwner)
}

// CollectionCraftingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CollectionCrafting contract.
type CollectionCraftingApprovalIterator struct {
	Event *CollectionCraftingApproval // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingApproval)
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
		it.Event = new(CollectionCraftingApproval)
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
func (it *CollectionCraftingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingApproval represents a Approval event raised by the CollectionCrafting contract.
type CollectionCraftingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CollectionCraftingApprovalIterator, error) {

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

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingApprovalIterator{contract: _CollectionCrafting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CollectionCraftingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingApproval)
				if err := _CollectionCrafting.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_CollectionCrafting *CollectionCraftingFilterer) ParseApproval(log types.Log) (*CollectionCraftingApproval, error) {
	event := new(CollectionCraftingApproval)
	if err := _CollectionCrafting.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CollectionCrafting contract.
type CollectionCraftingApprovalForAllIterator struct {
	Event *CollectionCraftingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingApprovalForAll)
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
		it.Event = new(CollectionCraftingApprovalForAll)
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
func (it *CollectionCraftingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingApprovalForAll represents a ApprovalForAll event raised by the CollectionCrafting contract.
type CollectionCraftingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CollectionCraftingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingApprovalForAllIterator{contract: _CollectionCrafting.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CollectionCraftingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingApprovalForAll)
				if err := _CollectionCrafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_CollectionCrafting *CollectionCraftingFilterer) ParseApprovalForAll(log types.Log) (*CollectionCraftingApprovalForAll, error) {
	event := new(CollectionCraftingApprovalForAll)
	if err := _CollectionCrafting.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingCollectionCraftedIterator is returned from FilterCollectionCrafted and is used to iterate over the raw logs and unpacked data for CollectionCrafted events raised by the CollectionCrafting contract.
type CollectionCraftingCollectionCraftedIterator struct {
	Event *CollectionCraftingCollectionCrafted // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingCollectionCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingCollectionCrafted)
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
		it.Event = new(CollectionCraftingCollectionCrafted)
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
func (it *CollectionCraftingCollectionCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingCollectionCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingCollectionCrafted represents a CollectionCrafted event raised by the CollectionCrafting contract.
type CollectionCraftingCollectionCrafted struct {
	CollectionId        *big.Int
	DayMonthTokenId     *big.Int
	YearTokenId         *big.Int
	CategoryTokenId     *big.Int
	CraftingCardTokenId *big.Int
	TokenURI            string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollectionCrafted is a free log retrieval operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterCollectionCrafted(opts *bind.FilterOpts, collectionId []*big.Int) (*CollectionCraftingCollectionCraftedIterator, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingCollectionCraftedIterator{contract: _CollectionCrafting.contract, event: "CollectionCrafted", logs: logs, sub: sub}, nil
}

// WatchCollectionCrafted is a free log subscription operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchCollectionCrafted(opts *bind.WatchOpts, sink chan<- *CollectionCraftingCollectionCrafted, collectionId []*big.Int) (event.Subscription, error) {

	var collectionIdRule []interface{}
	for _, collectionIdItem := range collectionId {
		collectionIdRule = append(collectionIdRule, collectionIdItem)
	}

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "CollectionCrafted", collectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingCollectionCrafted)
				if err := _CollectionCrafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
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

// ParseCollectionCrafted is a log parse operation binding the contract event 0xe5979afd3a6266ac0cf01741228fe2b459458306a8d7a2c3152da07ee0c2f2be.
//
// Solidity: event CollectionCrafted(uint256 indexed collectionId, uint256 dayMonthTokenId, uint256 yearTokenId, uint256 categoryTokenId, uint256 craftingCardTokenId, string tokenURI)
func (_CollectionCrafting *CollectionCraftingFilterer) ParseCollectionCrafted(log types.Log) (*CollectionCraftingCollectionCrafted, error) {
	event := new(CollectionCraftingCollectionCrafted)
	if err := _CollectionCrafting.contract.UnpackLog(event, "CollectionCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the CollectionCrafting contract.
type CollectionCraftingCraftingContractChangedIterator struct {
	Event *CollectionCraftingCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingCraftingContractChanged)
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
		it.Event = new(CollectionCraftingCraftingContractChanged)
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
func (it *CollectionCraftingCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingCraftingContractChanged represents a CraftingContractChanged event raised by the CollectionCrafting contract.
type CollectionCraftingCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*CollectionCraftingCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingCraftingContractChangedIterator{contract: _CollectionCrafting.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *CollectionCraftingCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingCraftingContractChanged)
				if err := _CollectionCrafting.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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

// ParseCraftingContractChanged is a log parse operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_CollectionCrafting *CollectionCraftingFilterer) ParseCraftingContractChanged(log types.Log) (*CollectionCraftingCraftingContractChanged, error) {
	event := new(CollectionCraftingCraftingContractChanged)
	if err := _CollectionCrafting.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CollectionCrafting contract.
type CollectionCraftingOwnershipTransferredIterator struct {
	Event *CollectionCraftingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingOwnershipTransferred)
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
		it.Event = new(CollectionCraftingOwnershipTransferred)
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
func (it *CollectionCraftingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingOwnershipTransferred represents a OwnershipTransferred event raised by the CollectionCrafting contract.
type CollectionCraftingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollectionCraftingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingOwnershipTransferredIterator{contract: _CollectionCrafting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollectionCraftingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingOwnershipTransferred)
				if err := _CollectionCrafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CollectionCrafting *CollectionCraftingFilterer) ParseOwnershipTransferred(log types.Log) (*CollectionCraftingOwnershipTransferred, error) {
	event := new(CollectionCraftingOwnershipTransferred)
	if err := _CollectionCrafting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the CollectionCrafting contract.
type CollectionCraftingPartCraftedIterator struct {
	Event *CollectionCraftingPartCrafted // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingPartCrafted)
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
		it.Event = new(CollectionCraftingPartCrafted)
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
func (it *CollectionCraftingPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingPartCrafted represents a PartCrafted event raised by the CollectionCrafting contract.
type CollectionCraftingPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*CollectionCraftingPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingPartCraftedIterator{contract: _CollectionCrafting.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *CollectionCraftingPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingPartCrafted)
				if err := _CollectionCrafting.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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

// ParsePartCrafted is a log parse operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) ParsePartCrafted(log types.Log) (*CollectionCraftingPartCrafted, error) {
	event := new(CollectionCraftingPartCrafted)
	if err := _CollectionCrafting.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionCraftingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CollectionCrafting contract.
type CollectionCraftingTransferIterator struct {
	Event *CollectionCraftingTransfer // Event containing the contract specifics and raw log

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
func (it *CollectionCraftingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionCraftingTransfer)
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
		it.Event = new(CollectionCraftingTransfer)
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
func (it *CollectionCraftingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionCraftingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionCraftingTransfer represents a Transfer event raised by the CollectionCrafting contract.
type CollectionCraftingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CollectionCraftingTransferIterator, error) {

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

	logs, sub, err := _CollectionCrafting.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CollectionCraftingTransferIterator{contract: _CollectionCrafting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CollectionCrafting *CollectionCraftingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CollectionCraftingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _CollectionCrafting.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionCraftingTransfer)
				if err := _CollectionCrafting.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CollectionCrafting *CollectionCraftingFilterer) ParseTransfer(log types.Log) (*CollectionCraftingTransfer, error) {
	event := new(CollectionCraftingTransfer)
	if err := _CollectionCrafting.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
