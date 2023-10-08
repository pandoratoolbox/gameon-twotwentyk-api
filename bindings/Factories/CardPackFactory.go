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

// CardPackFactoryMetaData contains all meta data concerning the CardPackFactory contract.
var CardPackFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cardpack\",\"type\":\"address\"}],\"name\":\"cardPackCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"createCardPack\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cardpack\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CardPackFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CardPackFactoryMetaData.ABI instead.
var CardPackFactoryABI = CardPackFactoryMetaData.ABI

// CardPackFactory is an auto generated Go binding around an Ethereum contract.
type CardPackFactory struct {
	CardPackFactoryCaller     // Read-only binding to the contract
	CardPackFactoryTransactor // Write-only binding to the contract
	CardPackFactoryFilterer   // Log filterer for contract events
}

// CardPackFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardPackFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardPackFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardPackFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardPackFactorySession struct {
	Contract     *CardPackFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardPackFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardPackFactoryCallerSession struct {
	Contract *CardPackFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CardPackFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardPackFactoryTransactorSession struct {
	Contract     *CardPackFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CardPackFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardPackFactoryRaw struct {
	Contract *CardPackFactory // Generic contract binding to access the raw methods on
}

// CardPackFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardPackFactoryCallerRaw struct {
	Contract *CardPackFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CardPackFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardPackFactoryTransactorRaw struct {
	Contract *CardPackFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardPackFactory creates a new instance of CardPackFactory, bound to a specific deployed contract.
func NewCardPackFactory(address common.Address, backend bind.ContractBackend) (*CardPackFactory, error) {
	contract, err := bindCardPackFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardPackFactory{CardPackFactoryCaller: CardPackFactoryCaller{contract: contract}, CardPackFactoryTransactor: CardPackFactoryTransactor{contract: contract}, CardPackFactoryFilterer: CardPackFactoryFilterer{contract: contract}}, nil
}

// NewCardPackFactoryCaller creates a new read-only instance of CardPackFactory, bound to a specific deployed contract.
func NewCardPackFactoryCaller(address common.Address, caller bind.ContractCaller) (*CardPackFactoryCaller, error) {
	contract, err := bindCardPackFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardPackFactoryCaller{contract: contract}, nil
}

// NewCardPackFactoryTransactor creates a new write-only instance of CardPackFactory, bound to a specific deployed contract.
func NewCardPackFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CardPackFactoryTransactor, error) {
	contract, err := bindCardPackFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardPackFactoryTransactor{contract: contract}, nil
}

// NewCardPackFactoryFilterer creates a new log filterer instance of CardPackFactory, bound to a specific deployed contract.
func NewCardPackFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CardPackFactoryFilterer, error) {
	contract, err := bindCardPackFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardPackFactoryFilterer{contract: contract}, nil
}

// bindCardPackFactory binds a generic wrapper to an already deployed contract.
func bindCardPackFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CardPackFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardPackFactory *CardPackFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardPackFactory.Contract.CardPackFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardPackFactory *CardPackFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPackFactory.Contract.CardPackFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardPackFactory *CardPackFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardPackFactory.Contract.CardPackFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardPackFactory *CardPackFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardPackFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardPackFactory *CardPackFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPackFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardPackFactory *CardPackFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardPackFactory.Contract.contract.Transact(opts, method, params...)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CardPackFactory *CardPackFactoryCaller) AdminWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPackFactory.contract.Call(opts, &out, "adminWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CardPackFactory *CardPackFactorySession) AdminWallet() (common.Address, error) {
	return _CardPackFactory.Contract.AdminWallet(&_CardPackFactory.CallOpts)
}

// AdminWallet is a free data retrieval call binding the contract method 0x36b19cd7.
//
// Solidity: function adminWallet() view returns(address)
func (_CardPackFactory *CardPackFactoryCallerSession) AdminWallet() (common.Address, error) {
	return _CardPackFactory.Contract.AdminWallet(&_CardPackFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPackFactory *CardPackFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPackFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPackFactory *CardPackFactorySession) Owner() (common.Address, error) {
	return _CardPackFactory.Contract.Owner(&_CardPackFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPackFactory *CardPackFactoryCallerSession) Owner() (common.Address, error) {
	return _CardPackFactory.Contract.Owner(&_CardPackFactory.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CardPackFactory *CardPackFactoryTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _CardPackFactory.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CardPackFactory *CardPackFactorySession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.ChangeAdmin(&_CardPackFactory.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_CardPackFactory *CardPackFactoryTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.ChangeAdmin(&_CardPackFactory.TransactOpts, _newAdmin)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address cardpack)
func (_CardPackFactory *CardPackFactoryTransactor) CreateCardPack(opts *bind.TransactOpts, _totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.contract.Transact(opts, "createCardPack", _totalAmount, owner)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address cardpack)
func (_CardPackFactory *CardPackFactorySession) CreateCardPack(_totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.CreateCardPack(&_CardPackFactory.TransactOpts, _totalAmount, owner)
}

// CreateCardPack is a paid mutator transaction binding the contract method 0x6e646c3c.
//
// Solidity: function createCardPack(uint256 _totalAmount, address owner) returns(address cardpack)
func (_CardPackFactory *CardPackFactoryTransactorSession) CreateCardPack(_totalAmount *big.Int, owner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.CreateCardPack(&_CardPackFactory.TransactOpts, _totalAmount, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CardPackFactory *CardPackFactoryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPackFactory.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CardPackFactory *CardPackFactorySession) Initialize() (*types.Transaction, error) {
	return _CardPackFactory.Contract.Initialize(&_CardPackFactory.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_CardPackFactory *CardPackFactoryTransactorSession) Initialize() (*types.Transaction, error) {
	return _CardPackFactory.Contract.Initialize(&_CardPackFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPackFactory *CardPackFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPackFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPackFactory *CardPackFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CardPackFactory.Contract.RenounceOwnership(&_CardPackFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPackFactory *CardPackFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CardPackFactory.Contract.RenounceOwnership(&_CardPackFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPackFactory *CardPackFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPackFactory *CardPackFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.TransferOwnership(&_CardPackFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPackFactory *CardPackFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CardPackFactory.Contract.TransferOwnership(&_CardPackFactory.TransactOpts, newOwner)
}

// CardPackFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CardPackFactory contract.
type CardPackFactoryOwnershipTransferredIterator struct {
	Event *CardPackFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CardPackFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackFactoryOwnershipTransferred)
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
		it.Event = new(CardPackFactoryOwnershipTransferred)
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
func (it *CardPackFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CardPackFactory contract.
type CardPackFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CardPackFactory *CardPackFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CardPackFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CardPackFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CardPackFactoryOwnershipTransferredIterator{contract: _CardPackFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CardPackFactory *CardPackFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CardPackFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CardPackFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackFactoryOwnershipTransferred)
				if err := _CardPackFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CardPackFactory *CardPackFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CardPackFactoryOwnershipTransferred, error) {
	event := new(CardPackFactoryOwnershipTransferred)
	if err := _CardPackFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackFactoryCardPackCreatedIterator is returned from FilterCardPackCreated and is used to iterate over the raw logs and unpacked data for CardPackCreated events raised by the CardPackFactory contract.
type CardPackFactoryCardPackCreatedIterator struct {
	Event *CardPackFactoryCardPackCreated // Event containing the contract specifics and raw log

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
func (it *CardPackFactoryCardPackCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackFactoryCardPackCreated)
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
		it.Event = new(CardPackFactoryCardPackCreated)
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
func (it *CardPackFactoryCardPackCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackFactoryCardPackCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackFactoryCardPackCreated represents a CardPackCreated event raised by the CardPackFactory contract.
type CardPackFactoryCardPackCreated struct {
	Cardpack common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCardPackCreated is a free log retrieval operation binding the contract event 0x2b61f7f1ec3a76953375c09cdf817a02154590c7b94eb7776baa2933b3a06ff5.
//
// Solidity: event cardPackCreated(address cardpack)
func (_CardPackFactory *CardPackFactoryFilterer) FilterCardPackCreated(opts *bind.FilterOpts) (*CardPackFactoryCardPackCreatedIterator, error) {

	logs, sub, err := _CardPackFactory.contract.FilterLogs(opts, "cardPackCreated")
	if err != nil {
		return nil, err
	}
	return &CardPackFactoryCardPackCreatedIterator{contract: _CardPackFactory.contract, event: "cardPackCreated", logs: logs, sub: sub}, nil
}

// WatchCardPackCreated is a free log subscription operation binding the contract event 0x2b61f7f1ec3a76953375c09cdf817a02154590c7b94eb7776baa2933b3a06ff5.
//
// Solidity: event cardPackCreated(address cardpack)
func (_CardPackFactory *CardPackFactoryFilterer) WatchCardPackCreated(opts *bind.WatchOpts, sink chan<- *CardPackFactoryCardPackCreated) (event.Subscription, error) {

	logs, sub, err := _CardPackFactory.contract.WatchLogs(opts, "cardPackCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackFactoryCardPackCreated)
				if err := _CardPackFactory.contract.UnpackLog(event, "cardPackCreated", log); err != nil {
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

// ParseCardPackCreated is a log parse operation binding the contract event 0x2b61f7f1ec3a76953375c09cdf817a02154590c7b94eb7776baa2933b3a06ff5.
//
// Solidity: event cardPackCreated(address cardpack)
func (_CardPackFactory *CardPackFactoryFilterer) ParseCardPackCreated(log types.Log) (*CardPackFactoryCardPackCreated, error) {
	event := new(CardPackFactoryCardPackCreated)
	if err := _CardPackFactory.contract.UnpackLog(event, "cardPackCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
