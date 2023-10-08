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

// CategoryFactoryMetaData contains all meta data concerning the CategoryFactory contract.
var CategoryFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"category\",\"type\":\"address\"}],\"name\":\"CategoryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collectionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCategory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"category\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CategoryFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CategoryFactoryMetaData.ABI instead.
var CategoryFactoryABI = CategoryFactoryMetaData.ABI

// CategoryFactory is an auto generated Go binding around an Ethereum contract.
type CategoryFactory struct {
	CategoryFactoryCaller     // Read-only binding to the contract
	CategoryFactoryTransactor // Write-only binding to the contract
	CategoryFactoryFilterer   // Log filterer for contract events
}

// CategoryFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CategoryFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CategoryFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CategoryFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CategoryFactorySession struct {
	Contract     *CategoryFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CategoryFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CategoryFactoryCallerSession struct {
	Contract *CategoryFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CategoryFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CategoryFactoryTransactorSession struct {
	Contract     *CategoryFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CategoryFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CategoryFactoryRaw struct {
	Contract *CategoryFactory // Generic contract binding to access the raw methods on
}

// CategoryFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CategoryFactoryCallerRaw struct {
	Contract *CategoryFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CategoryFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CategoryFactoryTransactorRaw struct {
	Contract *CategoryFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCategoryFactory creates a new instance of CategoryFactory, bound to a specific deployed contract.
func NewCategoryFactory(address common.Address, backend bind.ContractBackend) (*CategoryFactory, error) {
	contract, err := bindCategoryFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CategoryFactory{CategoryFactoryCaller: CategoryFactoryCaller{contract: contract}, CategoryFactoryTransactor: CategoryFactoryTransactor{contract: contract}, CategoryFactoryFilterer: CategoryFactoryFilterer{contract: contract}}, nil
}

// NewCategoryFactoryCaller creates a new read-only instance of CategoryFactory, bound to a specific deployed contract.
func NewCategoryFactoryCaller(address common.Address, caller bind.ContractCaller) (*CategoryFactoryCaller, error) {
	contract, err := bindCategoryFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CategoryFactoryCaller{contract: contract}, nil
}

// NewCategoryFactoryTransactor creates a new write-only instance of CategoryFactory, bound to a specific deployed contract.
func NewCategoryFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CategoryFactoryTransactor, error) {
	contract, err := bindCategoryFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CategoryFactoryTransactor{contract: contract}, nil
}

// NewCategoryFactoryFilterer creates a new log filterer instance of CategoryFactory, bound to a specific deployed contract.
func NewCategoryFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CategoryFactoryFilterer, error) {
	contract, err := bindCategoryFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CategoryFactoryFilterer{contract: contract}, nil
}

// bindCategoryFactory binds a generic wrapper to an already deployed contract.
func bindCategoryFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CategoryFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CategoryFactory *CategoryFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CategoryFactory.Contract.CategoryFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CategoryFactory *CategoryFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryFactory.Contract.CategoryFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CategoryFactory *CategoryFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CategoryFactory.Contract.CategoryFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CategoryFactory *CategoryFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CategoryFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CategoryFactory *CategoryFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CategoryFactory *CategoryFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CategoryFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CategoryFactory *CategoryFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CategoryFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CategoryFactory *CategoryFactorySession) AdminWallet() (common.Address, error) {
	return _CategoryFactory.Contract.AdminWallet(&_CategoryFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CategoryFactory *CategoryFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _CategoryFactory.Contract.AdminWallet(&_CategoryFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CategoryFactory *CategoryFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CategoryFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CategoryFactory *CategoryFactorySession) Owner() (common.Address, error) {
	return _CategoryFactory.Contract.Owner(&_CategoryFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CategoryFactory *CategoryFactoryCallerSession) Owner() (common.Address, error) {
	return _CategoryFactory.Contract.Owner(&_CategoryFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CategoryFactory *CategoryFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CategoryFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CategoryFactory *CategoryFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.ChangeAdmin(&_CategoryFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CategoryFactory *CategoryFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.ChangeAdmin(&_CategoryFactory.TransactOpts, _newAdmin)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address category)
func (_CategoryFactory *CategoryFactoryTransactor) CreateCategory(opts *bind.TransactOpts, collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.contract.Transact(opts, "createCategory", collectionId, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address category)
func (_CategoryFactory *CategoryFactorySession) CreateCategory(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.CreateCategory(&_CategoryFactory.TransactOpts, collectionId, owner)
}

// CreateCategory is a paid mutator transaction binding the contract method 0x2dca7b1c.
//
// Solidity: function createCategory(uint256 collectionId, address owner) returns(address category)
func (_CategoryFactory *CategoryFactoryTransactorSession) CreateCategory(collectionId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.CreateCategory(&_CategoryFactory.TransactOpts, collectionId, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CategoryFactory *CategoryFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CategoryFactory *CategoryFactorySession) Initialize() (*types.Transaction, error) {
	return _CategoryFactory.Contract.Initialize(&_CategoryFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CategoryFactory *CategoryFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _CategoryFactory.Contract.Initialize(&_CategoryFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CategoryFactory *CategoryFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CategoryFactory *CategoryFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CategoryFactory.Contract.RenounceOwnership(&_CategoryFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CategoryFactory *CategoryFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CategoryFactory.Contract.RenounceOwnership(&_CategoryFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CategoryFactory *CategoryFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CategoryFactory *CategoryFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.TransferOwnership(&_CategoryFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CategoryFactory *CategoryFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CategoryFactory.Contract.TransferOwnership(&_CategoryFactory.TransactOpts, newOwner)
}

// CategoryFactoryCategoryCreatedIterator is returned from FilterCategoryCreated and is used to iterate over the raw logs and unpacked data for CategoryCreated events raised by the CategoryFactory contract.
type CategoryFactoryCategoryCreatedIterator struct {
	Event *CategoryFactoryCategoryCreated // Event containing the contract specifics and raw log

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
func (it *CategoryFactoryCategoryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CategoryFactoryCategoryCreated)
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
		it.Event = new(CategoryFactoryCategoryCreated)
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
func (it *CategoryFactoryCategoryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CategoryFactoryCategoryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CategoryFactoryCategoryCreated represents a CategoryCreated event raised by the CategoryFactory contract.
type CategoryFactoryCategoryCreated struct {
	Category common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCategoryCreated is a free log retrieval operation binding the contract event 0xaa1f4b4dd901b2af8a5bcd9bc027b6f06a8f2de93695e8c1294d31814a05de3d.
//
// Solidity: event CategoryCreated(address category)
func (_CategoryFactory *CategoryFactoryFilterer) FilterCategoryCreated(opts *bind.FilterOpts) (*CategoryFactoryCategoryCreatedIterator, error) {

	logs, sub, err := _CategoryFactory.contract.FilterLogs(opts, "CategoryCreated")
	if err != nil {
		return nil, err
	}
	return &CategoryFactoryCategoryCreatedIterator{contract: _CategoryFactory.contract, event: "CategoryCreated", logs: logs, sub: sub}, nil
}

// WatchCategoryCreated is a free log subscription operation binding the contract event 0xaa1f4b4dd901b2af8a5bcd9bc027b6f06a8f2de93695e8c1294d31814a05de3d.
//
// Solidity: event CategoryCreated(address category)
func (_CategoryFactory *CategoryFactoryFilterer) WatchCategoryCreated(opts *bind.WatchOpts, sink chan<- *CategoryFactoryCategoryCreated) (event.Subscription, error) {

	logs, sub, err := _CategoryFactory.contract.WatchLogs(opts, "CategoryCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CategoryFactoryCategoryCreated)
				if err := _CategoryFactory.contract.UnpackLog(event, "CategoryCreated", log); err != nil {
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

// ParseCategoryCreated is a log parse operation binding the contract event 0xaa1f4b4dd901b2af8a5bcd9bc027b6f06a8f2de93695e8c1294d31814a05de3d.
//
// Solidity: event CategoryCreated(address category)
func (_CategoryFactory *CategoryFactoryFilterer) ParseCategoryCreated(log types.Log) (*CategoryFactoryCategoryCreated, error) {
	event := new(CategoryFactoryCategoryCreated)
	if err := _CategoryFactory.contract.UnpackLog(event, "CategoryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CategoryFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CategoryFactory contract.
type CategoryFactoryOwnershipTransferredIterator struct {
	Event *CategoryFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CategoryFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CategoryFactoryOwnershipTransferred)
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
		it.Event = new(CategoryFactoryOwnershipTransferred)
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
func (it *CategoryFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CategoryFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CategoryFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CategoryFactory contract.
type CategoryFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CategoryFactory *CategoryFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CategoryFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CategoryFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CategoryFactoryOwnershipTransferredIterator{contract: _CategoryFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CategoryFactory *CategoryFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CategoryFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CategoryFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CategoryFactoryOwnershipTransferred)
				if err := _CategoryFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CategoryFactory *CategoryFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CategoryFactoryOwnershipTransferred, error) {
	event := new(CategoryFactoryOwnershipTransferred)
	if err := _CategoryFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
