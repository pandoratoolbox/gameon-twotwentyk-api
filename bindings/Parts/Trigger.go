// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package parts

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

// TriggerMetaData contains all meta data concerning the Trigger contract.
var TriggerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collectionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCrafter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCrafter\",\"type\":\"address\"}],\"name\":\"CraftingContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterContractChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PartCrafted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"TriggerMinted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newCrafter\",\"type\":\"address\"}],\"name\":\"changeCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToCrafted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCrafted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"mintTrigger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TriggerABI is the input ABI used to generate the binding from.
// Deprecated: Use TriggerMetaData.ABI instead.
var TriggerABI = TriggerMetaData.ABI

// Trigger is an auto generated Go binding around an Ethereum contract.
type Trigger struct {
	TriggerCaller     // Read-only binding to the contract
	TriggerTransactor // Write-only binding to the contract
	TriggerFilterer   // Log filterer for contract events
}

// TriggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriggerSession struct {
	Contract     *Trigger          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TriggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriggerCallerSession struct {
	Contract *TriggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TriggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriggerTransactorSession struct {
	Contract     *TriggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TriggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriggerRaw struct {
	Contract *Trigger // Generic contract binding to access the raw methods on
}

// TriggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriggerCallerRaw struct {
	Contract *TriggerCaller // Generic read-only contract binding to access the raw methods on
}

// TriggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriggerTransactorRaw struct {
	Contract *TriggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrigger creates a new instance of Trigger, bound to a specific deployed contract.
func NewTrigger(address common.Address, backend bind.ContractBackend) (*Trigger, error) {
	contract, err := bindTrigger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trigger{TriggerCaller: TriggerCaller{contract: contract}, TriggerTransactor: TriggerTransactor{contract: contract}, TriggerFilterer: TriggerFilterer{contract: contract}}, nil
}

// NewTriggerCaller creates a new read-only instance of Trigger, bound to a specific deployed contract.
func NewTriggerCaller(address common.Address, caller bind.ContractCaller) (*TriggerCaller, error) {
	contract, err := bindTrigger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerCaller{contract: contract}, nil
}

// NewTriggerTransactor creates a new write-only instance of Trigger, bound to a specific deployed contract.
func NewTriggerTransactor(address common.Address, transactor bind.ContractTransactor) (*TriggerTransactor, error) {
	contract, err := bindTrigger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerTransactor{contract: contract}, nil
}

// NewTriggerFilterer creates a new log filterer instance of Trigger, bound to a specific deployed contract.
func NewTriggerFilterer(address common.Address, filterer bind.ContractFilterer) (*TriggerFilterer, error) {
	contract, err := bindTrigger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriggerFilterer{contract: contract}, nil
}

// bindTrigger binds a generic wrapper to an already deployed contract.
func bindTrigger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TriggerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trigger *TriggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trigger.Contract.TriggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trigger *TriggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.Contract.TriggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trigger *TriggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trigger.Contract.TriggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trigger *TriggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trigger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trigger *TriggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trigger *TriggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trigger.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trigger *TriggerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trigger *TriggerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Trigger.Contract.BalanceOf(&_Trigger.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Trigger *TriggerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Trigger.Contract.BalanceOf(&_Trigger.CallOpts, owner)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Trigger *TriggerCaller) CollectionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "collectionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Trigger *TriggerSession) CollectionId() (*big.Int, error) {
	return _Trigger.Contract.CollectionId(&_Trigger.CallOpts)
}

// CollectionId is a free data retrieval call binding the contract method 0x3d26bb67.
//
// Solidity: function collectionId() view returns(uint256)
func (_Trigger *TriggerCallerSession) CollectionId() (*big.Int, error) {
	return _Trigger.Contract.CollectionId(&_Trigger.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trigger *TriggerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trigger *TriggerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Trigger.Contract.GetApproved(&_Trigger.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Trigger *TriggerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Trigger.Contract.GetApproved(&_Trigger.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trigger *TriggerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trigger *TriggerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Trigger.Contract.IsApprovedForAll(&_Trigger.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Trigger *TriggerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Trigger.Contract.IsApprovedForAll(&_Trigger.CallOpts, owner, operator)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Trigger *TriggerCaller) IsCrafted(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "isCrafted", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Trigger *TriggerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Trigger.Contract.IsCrafted(&_Trigger.CallOpts, tokenId)
}

// IsCrafted is a free data retrieval call binding the contract method 0xfd96527d.
//
// Solidity: function isCrafted(uint256 tokenId) view returns(bool)
func (_Trigger *TriggerCallerSession) IsCrafted(tokenId *big.Int) (bool, error) {
	return _Trigger.Contract.IsCrafted(&_Trigger.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trigger *TriggerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trigger *TriggerSession) Name() (string, error) {
	return _Trigger.Contract.Name(&_Trigger.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Trigger *TriggerCallerSession) Name() (string, error) {
	return _Trigger.Contract.Name(&_Trigger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trigger *TriggerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trigger *TriggerSession) Owner() (common.Address, error) {
	return _Trigger.Contract.Owner(&_Trigger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trigger *TriggerCallerSession) Owner() (common.Address, error) {
	return _Trigger.Contract.Owner(&_Trigger.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trigger *TriggerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trigger *TriggerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Trigger.Contract.OwnerOf(&_Trigger.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Trigger *TriggerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Trigger.Contract.OwnerOf(&_Trigger.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trigger *TriggerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trigger *TriggerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trigger.Contract.SupportsInterface(&_Trigger.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trigger *TriggerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trigger.Contract.SupportsInterface(&_Trigger.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trigger *TriggerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trigger *TriggerSession) Symbol() (string, error) {
	return _Trigger.Contract.Symbol(&_Trigger.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Trigger *TriggerCallerSession) Symbol() (string, error) {
	return _Trigger.Contract.Symbol(&_Trigger.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Trigger *TriggerCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "tokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Trigger *TriggerSession) TokenCounter() (*big.Int, error) {
	return _Trigger.Contract.TokenCounter(&_Trigger.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_Trigger *TriggerCallerSession) TokenCounter() (*big.Int, error) {
	return _Trigger.Contract.TokenCounter(&_Trigger.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trigger *TriggerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Trigger.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trigger *TriggerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Trigger.Contract.TokenURI(&_Trigger.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Trigger *TriggerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Trigger.Contract.TokenURI(&_Trigger.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trigger *TriggerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.Approve(&_Trigger.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.Approve(&_Trigger.TransactOpts, to, tokenId)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Trigger *TriggerTransactor) ChangeCrafter(opts *bind.TransactOpts, _newCrafter common.Address) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "changeCrafter", _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Trigger *TriggerSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeCrafter(&_Trigger.TransactOpts, _newCrafter)
}

// ChangeCrafter is a paid mutator transaction binding the contract method 0xe7d3b069.
//
// Solidity: function changeCrafter(address _newCrafter) returns()
func (_Trigger *TriggerTransactorSession) ChangeCrafter(_newCrafter common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeCrafter(&_Trigger.TransactOpts, _newCrafter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Trigger *TriggerTransactor) ChangeMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "changeMinter", _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Trigger *TriggerSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeMinter(&_Trigger.TransactOpts, _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_Trigger *TriggerTransactorSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeMinter(&_Trigger.TransactOpts, _newMinter)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Trigger *TriggerTransactor) ChangeToCrafted(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "changeToCrafted", tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Trigger *TriggerSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeToCrafted(&_Trigger.TransactOpts, tokenId)
}

// ChangeToCrafted is a paid mutator transaction binding the contract method 0xe6b60d9b.
//
// Solidity: function changeToCrafted(uint256 tokenId) returns()
func (_Trigger *TriggerTransactorSession) ChangeToCrafted(tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.ChangeToCrafted(&_Trigger.TransactOpts, tokenId)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns(uint256)
func (_Trigger *TriggerTransactor) MintTrigger(opts *bind.TransactOpts, _tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "mintTrigger", _tokenURI, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns(uint256)
func (_Trigger *TriggerSession) MintTrigger(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.MintTrigger(&_Trigger.TransactOpts, _tokenURI, _owner)
}

// MintTrigger is a paid mutator transaction binding the contract method 0x35dd69bb.
//
// Solidity: function mintTrigger(string _tokenURI, address _owner) returns(uint256)
func (_Trigger *TriggerTransactorSession) MintTrigger(_tokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.MintTrigger(&_Trigger.TransactOpts, _tokenURI, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trigger *TriggerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trigger *TriggerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trigger.Contract.RenounceOwnership(&_Trigger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trigger *TriggerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trigger.Contract.RenounceOwnership(&_Trigger.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.SafeTransferFrom(&_Trigger.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.SafeTransferFrom(&_Trigger.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trigger *TriggerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trigger *TriggerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trigger.Contract.SafeTransferFrom0(&_Trigger.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Trigger *TriggerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Trigger.Contract.SafeTransferFrom0(&_Trigger.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trigger *TriggerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trigger *TriggerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trigger.Contract.SetApprovalForAll(&_Trigger.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Trigger *TriggerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Trigger.Contract.SetApprovalForAll(&_Trigger.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.TransferFrom(&_Trigger.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Trigger *TriggerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.TransferFrom(&_Trigger.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trigger *TriggerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trigger *TriggerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.TransferOwnership(&_Trigger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trigger *TriggerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trigger.Contract.TransferOwnership(&_Trigger.TransactOpts, newOwner)
}

// TriggerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Trigger contract.
type TriggerApprovalIterator struct {
	Event *TriggerApproval // Event containing the contract specifics and raw log

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
func (it *TriggerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerApproval)
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
		it.Event = new(TriggerApproval)
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
func (it *TriggerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerApproval represents a Approval event raised by the Trigger contract.
type TriggerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TriggerApprovalIterator, error) {

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

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerApprovalIterator{contract: _Trigger.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TriggerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerApproval)
				if err := _Trigger.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParseApproval(log types.Log) (*TriggerApproval, error) {
	event := new(TriggerApproval)
	if err := _Trigger.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Trigger contract.
type TriggerApprovalForAllIterator struct {
	Event *TriggerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TriggerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerApprovalForAll)
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
		it.Event = new(TriggerApprovalForAll)
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
func (it *TriggerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerApprovalForAll represents a ApprovalForAll event raised by the Trigger contract.
type TriggerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Trigger *TriggerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TriggerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TriggerApprovalForAllIterator{contract: _Trigger.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Trigger *TriggerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TriggerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerApprovalForAll)
				if err := _Trigger.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParseApprovalForAll(log types.Log) (*TriggerApprovalForAll, error) {
	event := new(TriggerApprovalForAll)
	if err := _Trigger.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerCraftingContractChangedIterator is returned from FilterCraftingContractChanged and is used to iterate over the raw logs and unpacked data for CraftingContractChanged events raised by the Trigger contract.
type TriggerCraftingContractChangedIterator struct {
	Event *TriggerCraftingContractChanged // Event containing the contract specifics and raw log

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
func (it *TriggerCraftingContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerCraftingContractChanged)
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
		it.Event = new(TriggerCraftingContractChanged)
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
func (it *TriggerCraftingContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerCraftingContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerCraftingContractChanged represents a CraftingContractChanged event raised by the Trigger contract.
type TriggerCraftingContractChanged struct {
	PreviousCrafter common.Address
	NewCrafter      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCraftingContractChanged is a free log retrieval operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Trigger *TriggerFilterer) FilterCraftingContractChanged(opts *bind.FilterOpts, previousCrafter []common.Address, newCrafter []common.Address) (*TriggerCraftingContractChangedIterator, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return &TriggerCraftingContractChangedIterator{contract: _Trigger.contract, event: "CraftingContractChanged", logs: logs, sub: sub}, nil
}

// WatchCraftingContractChanged is a free log subscription operation binding the contract event 0x83149e90b909cde0e8099cb6998203f8d8dce5fa29151688382650f5637d7d88.
//
// Solidity: event CraftingContractChanged(address indexed previousCrafter, address indexed newCrafter)
func (_Trigger *TriggerFilterer) WatchCraftingContractChanged(opts *bind.WatchOpts, sink chan<- *TriggerCraftingContractChanged, previousCrafter []common.Address, newCrafter []common.Address) (event.Subscription, error) {

	var previousCrafterRule []interface{}
	for _, previousCrafterItem := range previousCrafter {
		previousCrafterRule = append(previousCrafterRule, previousCrafterItem)
	}
	var newCrafterRule []interface{}
	for _, newCrafterItem := range newCrafter {
		newCrafterRule = append(newCrafterRule, newCrafterItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "CraftingContractChanged", previousCrafterRule, newCrafterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerCraftingContractChanged)
				if err := _Trigger.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParseCraftingContractChanged(log types.Log) (*TriggerCraftingContractChanged, error) {
	event := new(TriggerCraftingContractChanged)
	if err := _Trigger.contract.UnpackLog(event, "CraftingContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerMinterContractChangedIterator is returned from FilterMinterContractChanged and is used to iterate over the raw logs and unpacked data for MinterContractChanged events raised by the Trigger contract.
type TriggerMinterContractChangedIterator struct {
	Event *TriggerMinterContractChanged // Event containing the contract specifics and raw log

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
func (it *TriggerMinterContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerMinterContractChanged)
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
		it.Event = new(TriggerMinterContractChanged)
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
func (it *TriggerMinterContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerMinterContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerMinterContractChanged represents a MinterContractChanged event raised by the Trigger contract.
type TriggerMinterContractChanged struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterContractChanged is a free log retrieval operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Trigger *TriggerFilterer) FilterMinterContractChanged(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*TriggerMinterContractChangedIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &TriggerMinterContractChangedIterator{contract: _Trigger.contract, event: "MinterContractChanged", logs: logs, sub: sub}, nil
}

// WatchMinterContractChanged is a free log subscription operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Trigger *TriggerFilterer) WatchMinterContractChanged(opts *bind.WatchOpts, sink chan<- *TriggerMinterContractChanged, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "MinterContractChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerMinterContractChanged)
				if err := _Trigger.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
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

// ParseMinterContractChanged is a log parse operation binding the contract event 0x114d06280373bf8a65ca81a7d5ec748e7747747d23dbaa3e41f3c4ea46223079.
//
// Solidity: event MinterContractChanged(address indexed previousMinter, address indexed newMinter)
func (_Trigger *TriggerFilterer) ParseMinterContractChanged(log types.Log) (*TriggerMinterContractChanged, error) {
	event := new(TriggerMinterContractChanged)
	if err := _Trigger.contract.UnpackLog(event, "MinterContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trigger contract.
type TriggerOwnershipTransferredIterator struct {
	Event *TriggerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TriggerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerOwnershipTransferred)
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
		it.Event = new(TriggerOwnershipTransferred)
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
func (it *TriggerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerOwnershipTransferred represents a OwnershipTransferred event raised by the Trigger contract.
type TriggerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trigger *TriggerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TriggerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerOwnershipTransferredIterator{contract: _Trigger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trigger *TriggerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TriggerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerOwnershipTransferred)
				if err := _Trigger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParseOwnershipTransferred(log types.Log) (*TriggerOwnershipTransferred, error) {
	event := new(TriggerOwnershipTransferred)
	if err := _Trigger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerPartCraftedIterator is returned from FilterPartCrafted and is used to iterate over the raw logs and unpacked data for PartCrafted events raised by the Trigger contract.
type TriggerPartCraftedIterator struct {
	Event *TriggerPartCrafted // Event containing the contract specifics and raw log

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
func (it *TriggerPartCraftedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerPartCrafted)
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
		it.Event = new(TriggerPartCrafted)
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
func (it *TriggerPartCraftedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerPartCraftedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerPartCrafted represents a PartCrafted event raised by the Trigger contract.
type TriggerPartCrafted struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartCrafted is a free log retrieval operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) FilterPartCrafted(opts *bind.FilterOpts, tokenId []*big.Int) (*TriggerPartCraftedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerPartCraftedIterator{contract: _Trigger.contract, event: "PartCrafted", logs: logs, sub: sub}, nil
}

// WatchPartCrafted is a free log subscription operation binding the contract event 0x71094059908fc1a5d676895c3ab6ec0e8957a09dc7a09914b791a2be4dc7fb15.
//
// Solidity: event PartCrafted(uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) WatchPartCrafted(opts *bind.WatchOpts, sink chan<- *TriggerPartCrafted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "PartCrafted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerPartCrafted)
				if err := _Trigger.contract.UnpackLog(event, "PartCrafted", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParsePartCrafted(log types.Log) (*TriggerPartCrafted, error) {
	event := new(TriggerPartCrafted)
	if err := _Trigger.contract.UnpackLog(event, "PartCrafted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Trigger contract.
type TriggerTransferIterator struct {
	Event *TriggerTransfer // Event containing the contract specifics and raw log

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
func (it *TriggerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerTransfer)
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
		it.Event = new(TriggerTransfer)
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
func (it *TriggerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerTransfer represents a Transfer event raised by the Trigger contract.
type TriggerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TriggerTransferIterator, error) {

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

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerTransferIterator{contract: _Trigger.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Trigger *TriggerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TriggerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerTransfer)
				if err := _Trigger.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Trigger *TriggerFilterer) ParseTransfer(log types.Log) (*TriggerTransfer, error) {
	event := new(TriggerTransfer)
	if err := _Trigger.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerTriggerMintedIterator is returned from FilterTriggerMinted and is used to iterate over the raw logs and unpacked data for TriggerMinted events raised by the Trigger contract.
type TriggerTriggerMintedIterator struct {
	Event *TriggerTriggerMinted // Event containing the contract specifics and raw log

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
func (it *TriggerTriggerMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerTriggerMinted)
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
		it.Event = new(TriggerTriggerMinted)
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
func (it *TriggerTriggerMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerTriggerMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerTriggerMinted represents a TriggerMinted event raised by the Trigger contract.
type TriggerTriggerMinted struct {
	CollectionId *big.Int
	TokenId      *big.Int
	TokenURI     string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTriggerMinted is a free log retrieval operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Trigger *TriggerFilterer) FilterTriggerMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*TriggerTriggerMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "TriggerMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TriggerTriggerMintedIterator{contract: _Trigger.contract, event: "TriggerMinted", logs: logs, sub: sub}, nil
}

// WatchTriggerMinted is a free log subscription operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Trigger *TriggerFilterer) WatchTriggerMinted(opts *bind.WatchOpts, sink chan<- *TriggerTriggerMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "TriggerMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerTriggerMinted)
				if err := _Trigger.contract.UnpackLog(event, "TriggerMinted", log); err != nil {
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

// ParseTriggerMinted is a log parse operation binding the contract event 0x61372f552903c29239005dd638aa00ac6da17b54079387e652d9209fc0b59052.
//
// Solidity: event TriggerMinted(uint256 collectionId, uint256 indexed tokenId, string tokenURI)
func (_Trigger *TriggerFilterer) ParseTriggerMinted(log types.Log) (*TriggerTriggerMinted, error) {
	event := new(TriggerTriggerMinted)
	if err := _Trigger.contract.UnpackLog(event, "TriggerMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
