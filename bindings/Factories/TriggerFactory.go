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

// TriggerFactoryMetaData contains all meta data concerning the TriggerFactory contract.
var TriggerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daymonth\",\"type\":\"address\"}],\"name\":\"DayMonthCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createTrigger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trigger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TriggerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TriggerFactoryMetaData.ABI instead.
var TriggerFactoryABI = TriggerFactoryMetaData.ABI

// TriggerFactory is an auto generated Go binding around an Ethereum contract.
type TriggerFactory struct {
	TriggerFactoryCaller     // Read-only binding to the contract
	TriggerFactoryTransactor // Write-only binding to the contract
	TriggerFactoryFilterer   // Log filterer for contract events
}

// TriggerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriggerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriggerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriggerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriggerFactorySession struct {
	Contract     *TriggerFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TriggerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriggerFactoryCallerSession struct {
	Contract *TriggerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TriggerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriggerFactoryTransactorSession struct {
	Contract     *TriggerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TriggerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriggerFactoryRaw struct {
	Contract *TriggerFactory // Generic contract binding to access the raw methods on
}

// TriggerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriggerFactoryCallerRaw struct {
	Contract *TriggerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TriggerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriggerFactoryTransactorRaw struct {
	Contract *TriggerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTriggerFactory creates a new instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactory(address common.Address, backend bind.ContractBackend) (*TriggerFactory, error) {
	contract, err := bindTriggerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TriggerFactory{TriggerFactoryCaller: TriggerFactoryCaller{contract: contract}, TriggerFactoryTransactor: TriggerFactoryTransactor{contract: contract}, TriggerFactoryFilterer: TriggerFactoryFilterer{contract: contract}}, nil
}

// NewTriggerFactoryCaller creates a new read-only instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryCaller(address common.Address, caller bind.ContractCaller) (*TriggerFactoryCaller, error) {
	contract, err := bindTriggerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryCaller{contract: contract}, nil
}

// NewTriggerFactoryTransactor creates a new write-only instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TriggerFactoryTransactor, error) {
	contract, err := bindTriggerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryTransactor{contract: contract}, nil
}

// NewTriggerFactoryFilterer creates a new log filterer instance of TriggerFactory, bound to a specific deployed contract.
func NewTriggerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TriggerFactoryFilterer, error) {
	contract, err := bindTriggerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryFilterer{contract: contract}, nil
}

// bindTriggerFactory binds a generic wrapper to an already deployed contract.
func bindTriggerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TriggerFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerFactory *TriggerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerFactory.Contract.TriggerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerFactory *TriggerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TriggerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerFactory *TriggerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TriggerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TriggerFactory *TriggerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TriggerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TriggerFactory *TriggerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TriggerFactory *TriggerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TriggerFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_TriggerFactory *TriggerFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_TriggerFactory *TriggerFactorySession) AdminWallet() (common.Address, error) {
	return _TriggerFactory.Contract.AdminWallet(&_TriggerFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_TriggerFactory *TriggerFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _TriggerFactory.Contract.AdminWallet(&_TriggerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TriggerFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactorySession) Owner() (common.Address, error) {
	return _TriggerFactory.Contract.Owner(&_TriggerFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TriggerFactory *TriggerFactoryCallerSession) Owner() (common.Address, error) {
	return _TriggerFactory.Contract.Owner(&_TriggerFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_TriggerFactory *TriggerFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_TriggerFactory *TriggerFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeAdmin(&_TriggerFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.ChangeAdmin(&_TriggerFactory.TransactOpts, _newAdmin)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address trigger)
func (_TriggerFactory *TriggerFactoryTransactor) CreateTrigger(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "createTrigger", collectionId, owner)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address trigger)
func (_TriggerFactory *TriggerFactorySession) CreateTrigger(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.CreateTrigger(&_TriggerFactory.TransactOpts, collectionId, owner)
}

// CreateTrigger is a paid mutator transaction binding the contract method 0xf91d8464.
//
// Solidity: function createTrigger(uint256 collectionId, address owner) returns(address trigger)
func (_TriggerFactory *TriggerFactoryTransactorSession) CreateTrigger(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.CreateTrigger(&_TriggerFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TriggerFactory *TriggerFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TriggerFactory *TriggerFactorySession) Initialize() (*types.Transaction, error) {
	return _TriggerFactory.Contract.Initialize(&_TriggerFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _TriggerFactory.Contract.Initialize(&_TriggerFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerFactory.Contract.RenounceOwnership(&_TriggerFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TriggerFactory.Contract.RenounceOwnership(&_TriggerFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferOwnership(&_TriggerFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TriggerFactory *TriggerFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TriggerFactory.Contract.TransferOwnership(&_TriggerFactory.TransactOpts, newOwner)
}

// TriggerFactoryDayMonthCreatedIterator is returned from FilterDayMonthCreated and is used to iterate over the raw logs and unpacked data for DayMonthCreated events raised by the TriggerFactory contract.
type TriggerFactoryDayMonthCreatedIterator struct {
	Event *TriggerFactoryDayMonthCreated // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryDayMonthCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryDayMonthCreated)
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
		it.Event = new(TriggerFactoryDayMonthCreated)
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
func (it *TriggerFactoryDayMonthCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryDayMonthCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryDayMonthCreated represents a DayMonthCreated event raised by the TriggerFactory contract.
type TriggerFactoryDayMonthCreated struct {
	Daymonth common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDayMonthCreated is a free log retrieval operation binding the contract event 0xf1346e9b554c01082fa5038330a28b7871c1bd39d2cb899a7b7735a8bce7f216.
//
// Solidity: event DayMonthCreated(address daymonth)
func (_TriggerFactory *TriggerFactoryFilterer) FilterDayMonthCreated(opts *bind.FilterOpts) (*TriggerFactoryDayMonthCreatedIterator, error) {

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "DayMonthCreated")
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryDayMonthCreatedIterator{contract: _TriggerFactory.contract, event: "DayMonthCreated", logs: logs, sub: sub}, nil
}

// WatchDayMonthCreated is a free log subscription operation binding the contract event 0xf1346e9b554c01082fa5038330a28b7871c1bd39d2cb899a7b7735a8bce7f216.
//
// Solidity: event DayMonthCreated(address daymonth)
func (_TriggerFactory *TriggerFactoryFilterer) WatchDayMonthCreated(opts *bind.WatchOpts, sink chan<- *TriggerFactoryDayMonthCreated) (event.Subscription, error) {

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "DayMonthCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryDayMonthCreated)
				if err := _TriggerFactory.contract.UnpackLog(event, "DayMonthCreated", log); err != nil {
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

// ParseDayMonthCreated is a log parse operation binding the contract event 0xf1346e9b554c01082fa5038330a28b7871c1bd39d2cb899a7b7735a8bce7f216.
//
// Solidity: event DayMonthCreated(address daymonth)
func (_TriggerFactory *TriggerFactoryFilterer) ParseDayMonthCreated(log types.Log) (*TriggerFactoryDayMonthCreated, error) {
	event := new(TriggerFactoryDayMonthCreated)
	if err := _TriggerFactory.contract.UnpackLog(event, "DayMonthCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TriggerFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TriggerFactory contract.
type TriggerFactoryOwnershipTransferredIterator struct {
	Event *TriggerFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TriggerFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerFactoryOwnershipTransferred)
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
		it.Event = new(TriggerFactoryOwnershipTransferred)
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
func (it *TriggerFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the TriggerFactory contract.
type TriggerFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerFactory *TriggerFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TriggerFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TriggerFactoryOwnershipTransferredIterator{contract: _TriggerFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TriggerFactory *TriggerFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TriggerFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TriggerFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerFactoryOwnershipTransferred)
				if err := _TriggerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TriggerFactory *TriggerFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TriggerFactoryOwnershipTransferred, error) {
	event := new(TriggerFactoryOwnershipTransferred)
	if err := _TriggerFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
