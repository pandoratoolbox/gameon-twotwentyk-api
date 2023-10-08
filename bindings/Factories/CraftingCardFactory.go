// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factories

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

// CraftingCardFactoryMetaData contains all meta data concerning the CraftingCardFactory contract.
var CraftingCardFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"crafting\",\"type\":\"address\"}],\"name\":\"CraftingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCrafting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"crafting\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CraftingCardFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CraftingCardFactoryMetaData.ABI instead.
var CraftingCardFactoryABI = CraftingCardFactoryMetaData.ABI

// CraftingCardFactory is an auto generated Go binding around an Ethereum contract.
type CraftingCardFactory struct {
	CraftingCardFactoryCaller     // Read-only binding to the contract
	CraftingCardFactoryTransactor // Write-only binding to the contract
	CraftingCardFactoryFilterer   // Log filterer for contract events
}

// CraftingCardFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CraftingCardFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingCardFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CraftingCardFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingCardFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CraftingCardFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CraftingCardFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CraftingCardFactorySession struct {
	Contract     *CraftingCardFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CraftingCardFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CraftingCardFactoryCallerSession struct {
	Contract *CraftingCardFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CraftingCardFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CraftingCardFactoryTransactorSession struct {
	Contract     *CraftingCardFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CraftingCardFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CraftingCardFactoryRaw struct {
	Contract *CraftingCardFactory // Generic contract binding to access the raw methods on
}

// CraftingCardFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CraftingCardFactoryCallerRaw struct {
	Contract *CraftingCardFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CraftingCardFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CraftingCardFactoryTransactorRaw struct {
	Contract *CraftingCardFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCraftingCardFactory creates a new instance of CraftingCardFactory, bound to a specific deployed contract.
func NewCraftingCardFactory(address common.Address, backend bind.ContractBackend) (*CraftingCardFactory, error) {
	contract, err := bindCraftingCardFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactory{CraftingCardFactoryCaller: CraftingCardFactoryCaller{contract: contract}, CraftingCardFactoryTransactor: CraftingCardFactoryTransactor{contract: contract}, CraftingCardFactoryFilterer: CraftingCardFactoryFilterer{contract: contract}}, nil
}

// NewCraftingCardFactoryCaller creates a new read-only instance of CraftingCardFactory, bound to a specific deployed contract.
func NewCraftingCardFactoryCaller(address common.Address, caller bind.ContractCaller) (*CraftingCardFactoryCaller, error) {
	contract, err := bindCraftingCardFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactoryCaller{contract: contract}, nil
}

// NewCraftingCardFactoryTransactor creates a new write-only instance of CraftingCardFactory, bound to a specific deployed contract.
func NewCraftingCardFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CraftingCardFactoryTransactor, error) {
	contract, err := bindCraftingCardFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactoryTransactor{contract: contract}, nil
}

// NewCraftingCardFactoryFilterer creates a new log filterer instance of CraftingCardFactory, bound to a specific deployed contract.
func NewCraftingCardFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CraftingCardFactoryFilterer, error) {
	contract, err := bindCraftingCardFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactoryFilterer{contract: contract}, nil
}

// bindCraftingCardFactory binds a generic wrapper to an already deployed contract.
func bindCraftingCardFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CraftingCardFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingCardFactory *CraftingCardFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingCardFactory.Contract.CraftingCardFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingCardFactory *CraftingCardFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.CraftingCardFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingCardFactory *CraftingCardFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.CraftingCardFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CraftingCardFactory *CraftingCardFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CraftingCardFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CraftingCardFactory *CraftingCardFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CraftingCardFactory *CraftingCardFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingCardFactory *CraftingCardFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CraftingCardFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingCardFactory *CraftingCardFactorySession) AdminWallet() (common.Address, error) {
	return _CraftingCardFactory.Contract.AdminWallet(&_CraftingCardFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CraftingCardFactory *CraftingCardFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _CraftingCardFactory.Contract.AdminWallet(&_CraftingCardFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingCardFactory *CraftingCardFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CraftingCardFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingCardFactory *CraftingCardFactorySession) Owner() (common.Address, error) {
	return _CraftingCardFactory.Contract.Owner(&_CraftingCardFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CraftingCardFactory *CraftingCardFactoryCallerSession) Owner() (common.Address, error) {
	return _CraftingCardFactory.Contract.Owner(&_CraftingCardFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingCardFactory *CraftingCardFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.ChangeAdmin(&_CraftingCardFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.ChangeAdmin(&_CraftingCardFactory.TransactOpts, _newAdmin)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingCardFactory *CraftingCardFactoryTransactor) CreateCrafting(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.contract.Transact(opts, "createCrafting", collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingCardFactory *CraftingCardFactorySession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.CreateCrafting(&_CraftingCardFactory.TransactOpts, collectionId, owner)
}

// CreateCrafting is a paid mutator transaction binding the contract method 0x222c116c.
//
// Solidity: function createCrafting(uint256 collectionId, address owner) returns(address crafting)
func (_CraftingCardFactory *CraftingCardFactoryTransactorSession) CreateCrafting(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.CreateCrafting(&_CraftingCardFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingCardFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingCardFactory *CraftingCardFactorySession) Initialize() (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.Initialize(&_CraftingCardFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.Initialize(&_CraftingCardFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CraftingCardFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingCardFactory *CraftingCardFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.RenounceOwnership(&_CraftingCardFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.RenounceOwnership(&_CraftingCardFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingCardFactory *CraftingCardFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.TransferOwnership(&_CraftingCardFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CraftingCardFactory *CraftingCardFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CraftingCardFactory.Contract.TransferOwnership(&_CraftingCardFactory.TransactOpts, newOwner)
}

// CraftingCardFactoryCraftingCreatedIterator is returned from FilterCraftingCreated and is used to iterate over the raw logs and unpacked data for CraftingCreated events raised by the CraftingCardFactory contract.
type CraftingCardFactoryCraftingCreatedIterator struct {
	Event *CraftingCardFactoryCraftingCreated // Event containing the contract specifics and raw log

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
func (it *CraftingCardFactoryCraftingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingCardFactoryCraftingCreated)
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
		it.Event = new(CraftingCardFactoryCraftingCreated)
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
func (it *CraftingCardFactoryCraftingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingCardFactoryCraftingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingCardFactoryCraftingCreated represents a CraftingCreated event raised by the CraftingCardFactory contract.
type CraftingCardFactoryCraftingCreated struct {
	Crafting common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCraftingCreated is a free log retrieval operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingCardFactory *CraftingCardFactoryFilterer) FilterCraftingCreated(opts *bind.FilterOpts) (*CraftingCardFactoryCraftingCreatedIterator, error) {

	logs, sub, err := _CraftingCardFactory.contract.FilterLogs(opts, "CraftingCreated")
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactoryCraftingCreatedIterator{contract: _CraftingCardFactory.contract, event: "CraftingCreated", logs: logs, sub: sub}, nil
}

// WatchCraftingCreated is a free log subscription operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingCardFactory *CraftingCardFactoryFilterer) WatchCraftingCreated(opts *bind.WatchOpts, sink chan<- *CraftingCardFactoryCraftingCreated) (event.Subscription, error) {

	logs, sub, err := _CraftingCardFactory.contract.WatchLogs(opts, "CraftingCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingCardFactoryCraftingCreated)
				if err := _CraftingCardFactory.contract.UnpackLog(event, "CraftingCreated", log); err != nil {
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

// ParseCraftingCreated is a log parse operation binding the contract event 0x5949f717c13342b6ef8d9818d4878e982e2568e41a7372c059ddca6fa3d5a385.
//
// Solidity: event CraftingCreated(address crafting)
func (_CraftingCardFactory *CraftingCardFactoryFilterer) ParseCraftingCreated(log types.Log) (*CraftingCardFactoryCraftingCreated, error) {
	event := new(CraftingCardFactoryCraftingCreated)
	if err := _CraftingCardFactory.contract.UnpackLog(event, "CraftingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CraftingCardFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CraftingCardFactory contract.
type CraftingCardFactoryOwnershipTransferredIterator struct {
	Event *CraftingCardFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CraftingCardFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CraftingCardFactoryOwnershipTransferred)
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
		it.Event = new(CraftingCardFactoryOwnershipTransferred)
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
func (it *CraftingCardFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CraftingCardFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CraftingCardFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CraftingCardFactory contract.
type CraftingCardFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CraftingCardFactory *CraftingCardFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CraftingCardFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CraftingCardFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CraftingCardFactoryOwnershipTransferredIterator{contract: _CraftingCardFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CraftingCardFactory *CraftingCardFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CraftingCardFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CraftingCardFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CraftingCardFactoryOwnershipTransferred)
				if err := _CraftingCardFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CraftingCardFactory *CraftingCardFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CraftingCardFactoryOwnershipTransferred, error) {
	event := new(CraftingCardFactoryOwnershipTransferred)
	if err := _CraftingCardFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
