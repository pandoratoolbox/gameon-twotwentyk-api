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

// IdentityFactoryMetaData contains all meta data concerning the IdentityFactory contract.
var IdentityFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"category\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCraftingIdentity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdentityFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityFactoryMetaData.ABI instead.
var IdentityFactoryABI = IdentityFactoryMetaData.ABI

// IdentityFactory is an auto generated Go binding around an Ethereum contract.
type IdentityFactory struct {
	IdentityFactoryCaller     // Read-only binding to the contract
	IdentityFactoryTransactor // Write-only binding to the contract
	IdentityFactoryFilterer   // Log filterer for contract events
}

// IdentityFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityFactorySession struct {
	Contract     *IdentityFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityFactoryCallerSession struct {
	Contract *IdentityFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentityFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityFactoryTransactorSession struct {
	Contract     *IdentityFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentityFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityFactoryRaw struct {
	Contract *IdentityFactory // Generic contract binding to access the raw methods on
}

// IdentityFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityFactoryCallerRaw struct {
	Contract *IdentityFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityFactoryTransactorRaw struct {
	Contract *IdentityFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityFactory creates a new instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactory(address common.Address, backend bind.ContractBackend) (*IdentityFactory, error) {
	contract, err := bindIdentityFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityFactory{IdentityFactoryCaller: IdentityFactoryCaller{contract: contract}, IdentityFactoryTransactor: IdentityFactoryTransactor{contract: contract}, IdentityFactoryFilterer: IdentityFactoryFilterer{contract: contract}}, nil
}

// NewIdentityFactoryCaller creates a new read-only instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryCaller(address common.Address, caller bind.ContractCaller) (*IdentityFactoryCaller, error) {
	contract, err := bindIdentityFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryCaller{contract: contract}, nil
}

// NewIdentityFactoryTransactor creates a new write-only instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityFactoryTransactor, error) {
	contract, err := bindIdentityFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryTransactor{contract: contract}, nil
}

// NewIdentityFactoryFilterer creates a new log filterer instance of IdentityFactory, bound to a specific deployed contract.
func NewIdentityFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFactoryFilterer, error) {
	contract, err := bindIdentityFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryFilterer{contract: contract}, nil
}

// bindIdentityFactory binds a generic wrapper to an already deployed contract.
func bindIdentityFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityFactory *IdentityFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityFactory.Contract.IdentityFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityFactory *IdentityFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.Contract.IdentityFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityFactory *IdentityFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityFactory.Contract.IdentityFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityFactory *IdentityFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityFactory *IdentityFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityFactory *IdentityFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_IdentityFactory *IdentityFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_IdentityFactory *IdentityFactorySession) AdminWallet() (common.Address, error) {
	return _IdentityFactory.Contract.AdminWallet(&_IdentityFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_IdentityFactory *IdentityFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _IdentityFactory.Contract.AdminWallet(&_IdentityFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityFactory *IdentityFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IdentityFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityFactory *IdentityFactorySession) Owner() (common.Address, error) {
	return _IdentityFactory.Contract.Owner(&_IdentityFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IdentityFactory *IdentityFactoryCallerSession) Owner() (common.Address, error) {
	return _IdentityFactory.Contract.Owner(&_IdentityFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_IdentityFactory *IdentityFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_IdentityFactory *IdentityFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.ChangeAdmin(&_IdentityFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_IdentityFactory *IdentityFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.ChangeAdmin(&_IdentityFactory.TransactOpts, _newAdmin)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address identity)
func (_IdentityFactory *IdentityFactoryTransactor) CreateCraftingIdentity(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "createCraftingIdentity", collectionId, owner)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address identity)
func (_IdentityFactory *IdentityFactorySession) CreateCraftingIdentity(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.CreateCraftingIdentity(&_IdentityFactory.TransactOpts, collectionId, owner)
}

// CreateCraftingIdentity is a paid mutator transaction binding the contract method 0xf77f634b.
//
// Solidity: function createCraftingIdentity(uint256 collectionId, address owner) returns(address identity)
func (_IdentityFactory *IdentityFactoryTransactorSession) CreateCraftingIdentity(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.CreateCraftingIdentity(&_IdentityFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IdentityFactory *IdentityFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IdentityFactory *IdentityFactorySession) Initialize() (*types.Transaction, error) {
	return _IdentityFactory.Contract.Initialize(&_IdentityFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IdentityFactory *IdentityFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _IdentityFactory.Contract.Initialize(&_IdentityFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityFactory *IdentityFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityFactory *IdentityFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityFactory.Contract.RenounceOwnership(&_IdentityFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IdentityFactory *IdentityFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IdentityFactory.Contract.RenounceOwnership(&_IdentityFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityFactory *IdentityFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityFactory *IdentityFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.TransferOwnership(&_IdentityFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IdentityFactory *IdentityFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IdentityFactory.Contract.TransferOwnership(&_IdentityFactory.TransactOpts, newOwner)
}

// IdentityFactoryIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the IdentityFactory contract.
type IdentityFactoryIdentityCreatedIterator struct {
	Event *IdentityFactoryIdentityCreated // Event containing the contract specifics and raw log

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
func (it *IdentityFactoryIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityFactoryIdentityCreated)
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
		it.Event = new(IdentityFactoryIdentityCreated)
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
func (it *IdentityFactoryIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityFactoryIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityFactoryIdentityCreated represents a IdentityCreated event raised by the IdentityFactory contract.
type IdentityFactoryIdentityCreated struct {
	Category common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address category)
func (_IdentityFactory *IdentityFactoryFilterer) FilterIdentityCreated(opts *bind.FilterOpts) (*IdentityFactoryIdentityCreatedIterator, error) {

	logs, sub, err := _IdentityFactory.contract.FilterLogs(opts, "IdentityCreated")
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryIdentityCreatedIterator{contract: _IdentityFactory.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address category)
func (_IdentityFactory *IdentityFactoryFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *IdentityFactoryIdentityCreated) (event.Subscription, error) {

	logs, sub, err := _IdentityFactory.contract.WatchLogs(opts, "IdentityCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityFactoryIdentityCreated)
				if err := _IdentityFactory.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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

// ParseIdentityCreated is a log parse operation binding the contract event 0xac993fde3b9423ff59e4a23cded8e89074c9c8740920d1d870f586ba7c5c8cf0.
//
// Solidity: event IdentityCreated(address category)
func (_IdentityFactory *IdentityFactoryFilterer) ParseIdentityCreated(log types.Log) (*IdentityFactoryIdentityCreated, error) {
	event := new(IdentityFactoryIdentityCreated)
	if err := _IdentityFactory.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IdentityFactory contract.
type IdentityFactoryOwnershipTransferredIterator struct {
	Event *IdentityFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentityFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityFactoryOwnershipTransferred)
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
		it.Event = new(IdentityFactoryOwnershipTransferred)
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
func (it *IdentityFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the IdentityFactory contract.
type IdentityFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityFactory *IdentityFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityFactoryOwnershipTransferredIterator{contract: _IdentityFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IdentityFactory *IdentityFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IdentityFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityFactoryOwnershipTransferred)
				if err := _IdentityFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IdentityFactory *IdentityFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityFactoryOwnershipTransferred, error) {
	event := new(IdentityFactoryOwnershipTransferred)
	if err := _IdentityFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
