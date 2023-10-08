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

// DayMonthFactoryMetaData contains all meta data concerning the DayMonthFactory contract.
var DayMonthFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"daymonth\",\"type\":\"address\"}],\"name\":\"DayMonthCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createDayMonth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"daymonth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DayMonthFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DayMonthFactoryMetaData.ABI instead.
var DayMonthFactoryABI = DayMonthFactoryMetaData.ABI

// DayMonthFactory is an auto generated Go binding around an Ethereum contract.
type DayMonthFactory struct {
	DayMonthFactoryCaller     // Read-only binding to the contract
	DayMonthFactoryTransactor // Write-only binding to the contract
	DayMonthFactoryFilterer   // Log filterer for contract events
}

// DayMonthFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DayMonthFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayMonthFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DayMonthFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayMonthFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DayMonthFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DayMonthFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DayMonthFactorySession struct {
	Contract     *DayMonthFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DayMonthFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DayMonthFactoryCallerSession struct {
	Contract *DayMonthFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DayMonthFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DayMonthFactoryTransactorSession struct {
	Contract     *DayMonthFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DayMonthFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DayMonthFactoryRaw struct {
	Contract *DayMonthFactory // Generic contract binding to access the raw methods on
}

// DayMonthFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DayMonthFactoryCallerRaw struct {
	Contract *DayMonthFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DayMonthFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DayMonthFactoryTransactorRaw struct {
	Contract *DayMonthFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDayMonthFactory creates a new instance of DayMonthFactory, bound to a specific deployed contract.
func NewDayMonthFactory(address common.Address, backend bind.ContractBackend) (*DayMonthFactory, error) {
	contract, err := bindDayMonthFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DayMonthFactory{DayMonthFactoryCaller: DayMonthFactoryCaller{contract: contract}, DayMonthFactoryTransactor: DayMonthFactoryTransactor{contract: contract}, DayMonthFactoryFilterer: DayMonthFactoryFilterer{contract: contract}}, nil
}

// NewDayMonthFactoryCaller creates a new read-only instance of DayMonthFactory, bound to a specific deployed contract.
func NewDayMonthFactoryCaller(address common.Address, caller bind.ContractCaller) (*DayMonthFactoryCaller, error) {
	contract, err := bindDayMonthFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DayMonthFactoryCaller{contract: contract}, nil
}

// NewDayMonthFactoryTransactor creates a new write-only instance of DayMonthFactory, bound to a specific deployed contract.
func NewDayMonthFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DayMonthFactoryTransactor, error) {
	contract, err := bindDayMonthFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DayMonthFactoryTransactor{contract: contract}, nil
}

// NewDayMonthFactoryFilterer creates a new log filterer instance of DayMonthFactory, bound to a specific deployed contract.
func NewDayMonthFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DayMonthFactoryFilterer, error) {
	contract, err := bindDayMonthFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DayMonthFactoryFilterer{contract: contract}, nil
}

// bindDayMonthFactory binds a generic wrapper to an already deployed contract.
func bindDayMonthFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DayMonthFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayMonthFactory *DayMonthFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DayMonthFactory.Contract.DayMonthFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayMonthFactory *DayMonthFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.DayMonthFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayMonthFactory *DayMonthFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.DayMonthFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DayMonthFactory *DayMonthFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DayMonthFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DayMonthFactory *DayMonthFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DayMonthFactory *DayMonthFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_DayMonthFactory *DayMonthFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DayMonthFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_DayMonthFactory *DayMonthFactorySession) AdminWallet() (common.Address, error) {
	return _DayMonthFactory.Contract.AdminWallet(&_DayMonthFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_DayMonthFactory *DayMonthFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _DayMonthFactory.Contract.AdminWallet(&_DayMonthFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DayMonthFactory *DayMonthFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DayMonthFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DayMonthFactory *DayMonthFactorySession) Owner() (common.Address, error) {
	return _DayMonthFactory.Contract.Owner(&_DayMonthFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DayMonthFactory *DayMonthFactoryCallerSession) Owner() (common.Address, error) {
	return _DayMonthFactory.Contract.Owner(&_DayMonthFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_DayMonthFactory *DayMonthFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_DayMonthFactory *DayMonthFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.ChangeAdmin(&_DayMonthFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_DayMonthFactory *DayMonthFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.ChangeAdmin(&_DayMonthFactory.TransactOpts, _newAdmin)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address daymonth)
func (_DayMonthFactory *DayMonthFactoryTransactor) CreateDayMonth(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.contract.Transact(opts, "createDayMonth", collectionId, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address daymonth)
func (_DayMonthFactory *DayMonthFactorySession) CreateDayMonth(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.CreateDayMonth(&_DayMonthFactory.TransactOpts, collectionId, owner)
}

// CreateDayMonth is a paid mutator transaction binding the contract method 0x7a7552f7.
//
// Solidity: function createDayMonth(uint256 collectionId, address owner) returns(address daymonth)
func (_DayMonthFactory *DayMonthFactoryTransactorSession) CreateDayMonth(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.CreateDayMonth(&_DayMonthFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DayMonthFactory *DayMonthFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayMonthFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DayMonthFactory *DayMonthFactorySession) Initialize() (*types.Transaction, error) {
	return _DayMonthFactory.Contract.Initialize(&_DayMonthFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_DayMonthFactory *DayMonthFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _DayMonthFactory.Contract.Initialize(&_DayMonthFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DayMonthFactory *DayMonthFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DayMonthFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DayMonthFactory *DayMonthFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _DayMonthFactory.Contract.RenounceOwnership(&_DayMonthFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DayMonthFactory *DayMonthFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DayMonthFactory.Contract.RenounceOwnership(&_DayMonthFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DayMonthFactory *DayMonthFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DayMonthFactory *DayMonthFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.TransferOwnership(&_DayMonthFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DayMonthFactory *DayMonthFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DayMonthFactory.Contract.TransferOwnership(&_DayMonthFactory.TransactOpts, newOwner)
}

// DayMonthFactoryDayMonthCreatedIterator is returned from FilterDayMonthCreated and is used to iterate over the raw logs and unpacked data for DayMonthCreated events raised by the DayMonthFactory contract.
type DayMonthFactoryDayMonthCreatedIterator struct {
	Event *DayMonthFactoryDayMonthCreated // Event containing the contract specifics and raw log

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
func (it *DayMonthFactoryDayMonthCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DayMonthFactoryDayMonthCreated)
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
		it.Event = new(DayMonthFactoryDayMonthCreated)
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
func (it *DayMonthFactoryDayMonthCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DayMonthFactoryDayMonthCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DayMonthFactoryDayMonthCreated represents a DayMonthCreated event raised by the DayMonthFactory contract.
type DayMonthFactoryDayMonthCreated struct {
	Daymonth common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDayMonthCreated is a free log retrieval operation binding the contract event 0xf1346e9b554c01082fa5038330a28b7871c1bd39d2cb899a7b7735a8bce7f216.
//
// Solidity: event DayMonthCreated(address daymonth)
func (_DayMonthFactory *DayMonthFactoryFilterer) FilterDayMonthCreated(opts *bind.FilterOpts) (*DayMonthFactoryDayMonthCreatedIterator, error) {

	logs, sub, err := _DayMonthFactory.contract.FilterLogs(opts, "DayMonthCreated")
	if err != nil {
		return nil, err
	}
	return &DayMonthFactoryDayMonthCreatedIterator{contract: _DayMonthFactory.contract, event: "DayMonthCreated", logs: logs, sub: sub}, nil
}

// WatchDayMonthCreated is a free log subscription operation binding the contract event 0xf1346e9b554c01082fa5038330a28b7871c1bd39d2cb899a7b7735a8bce7f216.
//
// Solidity: event DayMonthCreated(address daymonth)
func (_DayMonthFactory *DayMonthFactoryFilterer) WatchDayMonthCreated(opts *bind.WatchOpts, sink chan<- *DayMonthFactoryDayMonthCreated) (event.Subscription, error) {

	logs, sub, err := _DayMonthFactory.contract.WatchLogs(opts, "DayMonthCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DayMonthFactoryDayMonthCreated)
				if err := _DayMonthFactory.contract.UnpackLog(event, "DayMonthCreated", log); err != nil {
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
func (_DayMonthFactory *DayMonthFactoryFilterer) ParseDayMonthCreated(log types.Log) (*DayMonthFactoryDayMonthCreated, error) {
	event := new(DayMonthFactoryDayMonthCreated)
	if err := _DayMonthFactory.contract.UnpackLog(event, "DayMonthCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DayMonthFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DayMonthFactory contract.
type DayMonthFactoryOwnershipTransferredIterator struct {
	Event *DayMonthFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DayMonthFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DayMonthFactoryOwnershipTransferred)
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
		it.Event = new(DayMonthFactoryOwnershipTransferred)
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
func (it *DayMonthFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DayMonthFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DayMonthFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the DayMonthFactory contract.
type DayMonthFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DayMonthFactory *DayMonthFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DayMonthFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DayMonthFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DayMonthFactoryOwnershipTransferredIterator{contract: _DayMonthFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DayMonthFactory *DayMonthFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DayMonthFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DayMonthFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DayMonthFactoryOwnershipTransferred)
				if err := _DayMonthFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DayMonthFactory *DayMonthFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*DayMonthFactoryOwnershipTransferred, error) {
	event := new(DayMonthFactoryOwnershipTransferred)
	if err := _DayMonthFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
