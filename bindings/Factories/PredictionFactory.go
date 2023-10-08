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

// PredictionFactoryMetaData contains all meta data concerning the PredictionFactory contract.
var PredictionFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prediction\",\"type\":\"address\"}],\"name\":\"PredictionCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCraftingPredection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"prediction\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PredictionFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionFactoryMetaData.ABI instead.
var PredictionFactoryABI = PredictionFactoryMetaData.ABI

// PredictionFactory is an auto generated Go binding around an Ethereum contract.
type PredictionFactory struct {
	PredictionFactoryCaller     // Read-only binding to the contract
	PredictionFactoryTransactor // Write-only binding to the contract
	PredictionFactoryFilterer   // Log filterer for contract events
}

// PredictionFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionFactorySession struct {
	Contract     *PredictionFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PredictionFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionFactoryCallerSession struct {
	Contract *PredictionFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PredictionFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionFactoryTransactorSession struct {
	Contract     *PredictionFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PredictionFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionFactoryRaw struct {
	Contract *PredictionFactory // Generic contract binding to access the raw methods on
}

// PredictionFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionFactoryCallerRaw struct {
	Contract *PredictionFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionFactoryTransactorRaw struct {
	Contract *PredictionFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredictionFactory creates a new instance of PredictionFactory, bound to a specific deployed contract.
func NewPredictionFactory(address common.Address, backend bind.ContractBackend) (*PredictionFactory, error) {
	contract, err := bindPredictionFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PredictionFactory{PredictionFactoryCaller: PredictionFactoryCaller{contract: contract}, PredictionFactoryTransactor: PredictionFactoryTransactor{contract: contract}, PredictionFactoryFilterer: PredictionFactoryFilterer{contract: contract}}, nil
}

// NewPredictionFactoryCaller creates a new read-only instance of PredictionFactory, bound to a specific deployed contract.
func NewPredictionFactoryCaller(address common.Address, caller bind.ContractCaller) (*PredictionFactoryCaller, error) {
	contract, err := bindPredictionFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionFactoryCaller{contract: contract}, nil
}

// NewPredictionFactoryTransactor creates a new write-only instance of PredictionFactory, bound to a specific deployed contract.
func NewPredictionFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionFactoryTransactor, error) {
	contract, err := bindPredictionFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionFactoryTransactor{contract: contract}, nil
}

// NewPredictionFactoryFilterer creates a new log filterer instance of PredictionFactory, bound to a specific deployed contract.
func NewPredictionFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionFactoryFilterer, error) {
	contract, err := bindPredictionFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionFactoryFilterer{contract: contract}, nil
}

// bindPredictionFactory binds a generic wrapper to an already deployed contract.
func bindPredictionFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PredictionFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionFactory *PredictionFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionFactory.Contract.PredictionFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionFactory *PredictionFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionFactory.Contract.PredictionFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionFactory *PredictionFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionFactory.Contract.PredictionFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionFactory *PredictionFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionFactory *PredictionFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionFactory *PredictionFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_PredictionFactory *PredictionFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_PredictionFactory *PredictionFactorySession) AdminWallet() (common.Address, error) {
	return _PredictionFactory.Contract.AdminWallet(&_PredictionFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_PredictionFactory *PredictionFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _PredictionFactory.Contract.AdminWallet(&_PredictionFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionFactory *PredictionFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionFactory *PredictionFactorySession) Owner() (common.Address, error) {
	return _PredictionFactory.Contract.Owner(&_PredictionFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PredictionFactory *PredictionFactoryCallerSession) Owner() (common.Address, error) {
	return _PredictionFactory.Contract.Owner(&_PredictionFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionFactory *PredictionFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionFactory *PredictionFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.ChangeAdmin(&_PredictionFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PredictionFactory *PredictionFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.ChangeAdmin(&_PredictionFactory.TransactOpts, _newAdmin)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address prediction)
func (_PredictionFactory *PredictionFactoryTransactor) CreateCraftingPredection(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.contract.Transact(opts, "createCraftingPredection", collectionId, owner)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address prediction)
func (_PredictionFactory *PredictionFactorySession) CreateCraftingPredection(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.CreateCraftingPredection(&_PredictionFactory.TransactOpts, collectionId, owner)
}

// CreateCraftingPredection is a paid mutator transaction binding the contract method 0x88fcc6f1.
//
// Solidity: function createCraftingPredection(uint256 collectionId, address owner) returns(address prediction)
func (_PredictionFactory *PredictionFactoryTransactorSession) CreateCraftingPredection(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.CreateCraftingPredection(&_PredictionFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PredictionFactory *PredictionFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PredictionFactory *PredictionFactorySession) Initialize() (*types.Transaction, error) {
	return _PredictionFactory.Contract.Initialize(&_PredictionFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PredictionFactory *PredictionFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _PredictionFactory.Contract.Initialize(&_PredictionFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionFactory *PredictionFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionFactory *PredictionFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionFactory.Contract.RenounceOwnership(&_PredictionFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PredictionFactory *PredictionFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PredictionFactory.Contract.RenounceOwnership(&_PredictionFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionFactory *PredictionFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionFactory *PredictionFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.TransferOwnership(&_PredictionFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PredictionFactory *PredictionFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PredictionFactory.Contract.TransferOwnership(&_PredictionFactory.TransactOpts, newOwner)
}

// PredictionFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PredictionFactory contract.
type PredictionFactoryOwnershipTransferredIterator struct {
	Event *PredictionFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PredictionFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionFactoryOwnershipTransferred)
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
		it.Event = new(PredictionFactoryOwnershipTransferred)
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
func (it *PredictionFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the PredictionFactory contract.
type PredictionFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionFactory *PredictionFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PredictionFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PredictionFactoryOwnershipTransferredIterator{contract: _PredictionFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PredictionFactory *PredictionFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PredictionFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PredictionFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionFactoryOwnershipTransferred)
				if err := _PredictionFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PredictionFactory *PredictionFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*PredictionFactoryOwnershipTransferred, error) {
	event := new(PredictionFactoryOwnershipTransferred)
	if err := _PredictionFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PredictionFactoryPredictionCreatedIterator is returned from FilterPredictionCreated and is used to iterate over the raw logs and unpacked data for PredictionCreated events raised by the PredictionFactory contract.
type PredictionFactoryPredictionCreatedIterator struct {
	Event *PredictionFactoryPredictionCreated // Event containing the contract specifics and raw log

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
func (it *PredictionFactoryPredictionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PredictionFactoryPredictionCreated)
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
		it.Event = new(PredictionFactoryPredictionCreated)
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
func (it *PredictionFactoryPredictionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PredictionFactoryPredictionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PredictionFactoryPredictionCreated represents a PredictionCreated event raised by the PredictionFactory contract.
type PredictionFactoryPredictionCreated struct {
	Prediction common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPredictionCreated is a free log retrieval operation binding the contract event 0xd7b2de33199111c8a48c5ce17ca9116a7ce19a6a409a29a00edfb10ac18bdeb6.
//
// Solidity: event PredictionCreated(address prediction)
func (_PredictionFactory *PredictionFactoryFilterer) FilterPredictionCreated(opts *bind.FilterOpts) (*PredictionFactoryPredictionCreatedIterator, error) {

	logs, sub, err := _PredictionFactory.contract.FilterLogs(opts, "PredictionCreated")
	if err != nil {
		return nil, err
	}
	return &PredictionFactoryPredictionCreatedIterator{contract: _PredictionFactory.contract, event: "PredictionCreated", logs: logs, sub: sub}, nil
}

// WatchPredictionCreated is a free log subscription operation binding the contract event 0xd7b2de33199111c8a48c5ce17ca9116a7ce19a6a409a29a00edfb10ac18bdeb6.
//
// Solidity: event PredictionCreated(address prediction)
func (_PredictionFactory *PredictionFactoryFilterer) WatchPredictionCreated(opts *bind.WatchOpts, sink chan<- *PredictionFactoryPredictionCreated) (event.Subscription, error) {

	logs, sub, err := _PredictionFactory.contract.WatchLogs(opts, "PredictionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PredictionFactoryPredictionCreated)
				if err := _PredictionFactory.contract.UnpackLog(event, "PredictionCreated", log); err != nil {
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

// ParsePredictionCreated is a log parse operation binding the contract event 0xd7b2de33199111c8a48c5ce17ca9116a7ce19a6a409a29a00edfb10ac18bdeb6.
//
// Solidity: event PredictionCreated(address prediction)
func (_PredictionFactory *PredictionFactoryFilterer) ParsePredictionCreated(log types.Log) (*PredictionFactoryPredictionCreated, error) {
	event := new(PredictionFactoryPredictionCreated)
	if err := _PredictionFactory.contract.UnpackLog(event, "PredictionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
