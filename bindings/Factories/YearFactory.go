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

// YearFactoryMetaData contains all meta data concerning the YearFactory contract.
var YearFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"year\",\"type\":\"address\"}],\"name\":\"YearCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createYear\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"year\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YearFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use YearFactoryMetaData.ABI instead.
var YearFactoryABI = YearFactoryMetaData.ABI

// YearFactory is an auto generated Go binding around an Ethereum contract.
type YearFactory struct {
	YearFactoryCaller     // Read-only binding to the contract
	YearFactoryTransactor // Write-only binding to the contract
	YearFactoryFilterer   // Log filterer for contract events
}

// YearFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearFactorySession struct {
	Contract     *YearFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearFactoryCallerSession struct {
	Contract *YearFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YearFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearFactoryTransactorSession struct {
	Contract     *YearFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YearFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearFactoryRaw struct {
	Contract *YearFactory // Generic contract binding to access the raw methods on
}

// YearFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearFactoryCallerRaw struct {
	Contract *YearFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// YearFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearFactoryTransactorRaw struct {
	Contract *YearFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearFactory creates a new instance of YearFactory, bound to a specific deployed contract.
func NewYearFactory(address common.Address, backend bind.ContractBackend) (*YearFactory, error) {
	contract, err := bindYearFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearFactory{YearFactoryCaller: YearFactoryCaller{contract: contract}, YearFactoryTransactor: YearFactoryTransactor{contract: contract}, YearFactoryFilterer: YearFactoryFilterer{contract: contract}}, nil
}

// NewYearFactoryCaller creates a new read-only instance of YearFactory, bound to a specific deployed contract.
func NewYearFactoryCaller(address common.Address, caller bind.ContractCaller) (*YearFactoryCaller, error) {
	contract, err := bindYearFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearFactoryCaller{contract: contract}, nil
}

// NewYearFactoryTransactor creates a new write-only instance of YearFactory, bound to a specific deployed contract.
func NewYearFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*YearFactoryTransactor, error) {
	contract, err := bindYearFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearFactoryTransactor{contract: contract}, nil
}

// NewYearFactoryFilterer creates a new log filterer instance of YearFactory, bound to a specific deployed contract.
func NewYearFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*YearFactoryFilterer, error) {
	contract, err := bindYearFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearFactoryFilterer{contract: contract}, nil
}

// bindYearFactory binds a generic wrapper to an already deployed contract.
func bindYearFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearFactory *YearFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearFactory.Contract.YearFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearFactory *YearFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearFactory.Contract.YearFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearFactory *YearFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearFactory.Contract.YearFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearFactory *YearFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearFactory *YearFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearFactory *YearFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_YearFactory *YearFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_YearFactory *YearFactorySession) AdminWallet() (common.Address, error) {
	return _YearFactory.Contract.AdminWallet(&_YearFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_YearFactory *YearFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _YearFactory.Contract.AdminWallet(&_YearFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearFactory *YearFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearFactory *YearFactorySession) Owner() (common.Address, error) {
	return _YearFactory.Contract.Owner(&_YearFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearFactory *YearFactoryCallerSession) Owner() (common.Address, error) {
	return _YearFactory.Contract.Owner(&_YearFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_YearFactory *YearFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _YearFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_YearFactory *YearFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.ChangeAdmin(&_YearFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_YearFactory *YearFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.ChangeAdmin(&_YearFactory.TransactOpts, _newAdmin)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address year)
func (_YearFactory *YearFactoryTransactor) CreateYear(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _YearFactory.contract.Transact(opts, "createYear", collectionId, owner)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address year)
func (_YearFactory *YearFactorySession) CreateYear(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.CreateYear(&_YearFactory.TransactOpts, collectionId, owner)
}

// CreateYear is a paid mutator transaction binding the contract method 0xdf8eea27.
//
// Solidity: function createYear(uint256 collectionId, address owner) returns(address year)
func (_YearFactory *YearFactoryTransactorSession) CreateYear(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.CreateYear(&_YearFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_YearFactory *YearFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_YearFactory *YearFactorySession) Initialize() (*types.Transaction, error) {
	return _YearFactory.Contract.Initialize(&_YearFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_YearFactory *YearFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _YearFactory.Contract.Initialize(&_YearFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearFactory *YearFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearFactory *YearFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _YearFactory.Contract.RenounceOwnership(&_YearFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearFactory *YearFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YearFactory.Contract.RenounceOwnership(&_YearFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearFactory *YearFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YearFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearFactory *YearFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.TransferOwnership(&_YearFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearFactory *YearFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YearFactory.Contract.TransferOwnership(&_YearFactory.TransactOpts, newOwner)
}

// YearFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YearFactory contract.
type YearFactoryOwnershipTransferredIterator struct {
	Event *YearFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YearFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearFactoryOwnershipTransferred)
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
		it.Event = new(YearFactoryOwnershipTransferred)
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
func (it *YearFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the YearFactory contract.
type YearFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YearFactory *YearFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YearFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YearFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YearFactoryOwnershipTransferredIterator{contract: _YearFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YearFactory *YearFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YearFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YearFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearFactoryOwnershipTransferred)
				if err := _YearFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_YearFactory *YearFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*YearFactoryOwnershipTransferred, error) {
	event := new(YearFactoryOwnershipTransferred)
	if err := _YearFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearFactoryYearCreatedIterator is returned from FilterYearCreated and is used to iterate over the raw logs and unpacked data for YearCreated events raised by the YearFactory contract.
type YearFactoryYearCreatedIterator struct {
	Event *YearFactoryYearCreated // Event containing the contract specifics and raw log

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
func (it *YearFactoryYearCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearFactoryYearCreated)
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
		it.Event = new(YearFactoryYearCreated)
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
func (it *YearFactoryYearCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearFactoryYearCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearFactoryYearCreated represents a YearCreated event raised by the YearFactory contract.
type YearFactoryYearCreated struct {
	Year common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterYearCreated is a free log retrieval operation binding the contract event 0xb9cc74879fbd894ba9703afcac4a84d6174ae23b282ac8992c4e430ffd1298ee.
//
// Solidity: event YearCreated(address year)
func (_YearFactory *YearFactoryFilterer) FilterYearCreated(opts *bind.FilterOpts) (*YearFactoryYearCreatedIterator, error) {

	logs, sub, err := _YearFactory.contract.FilterLogs(opts, "YearCreated")
	if err != nil {
		return nil, err
	}
	return &YearFactoryYearCreatedIterator{contract: _YearFactory.contract, event: "YearCreated", logs: logs, sub: sub}, nil
}

// WatchYearCreated is a free log subscription operation binding the contract event 0xb9cc74879fbd894ba9703afcac4a84d6174ae23b282ac8992c4e430ffd1298ee.
//
// Solidity: event YearCreated(address year)
func (_YearFactory *YearFactoryFilterer) WatchYearCreated(opts *bind.WatchOpts, sink chan<- *YearFactoryYearCreated) (event.Subscription, error) {

	logs, sub, err := _YearFactory.contract.WatchLogs(opts, "YearCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearFactoryYearCreated)
				if err := _YearFactory.contract.UnpackLog(event, "YearCreated", log); err != nil {
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

// ParseYearCreated is a log parse operation binding the contract event 0xb9cc74879fbd894ba9703afcac4a84d6174ae23b282ac8992c4e430ffd1298ee.
//
// Solidity: event YearCreated(address year)
func (_YearFactory *YearFactoryFilterer) ParseYearCreated(log types.Log) (*YearFactoryYearCreated, error) {
	event := new(YearFactoryYearCreated)
	if err := _YearFactory.contract.UnpackLog(event, "YearCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
