// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openpack

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

// PackOpenerMetaData contains all meta data concerning the PackOpener contract.
var PackOpenerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cardPackAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dayMonthAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggerContract\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCategories\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenYears\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenDayMonths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCraftingCards\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenTriggers\",\"type\":\"string[]\"}],\"name\":\"openPack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PackOpenerABI is the input ABI used to generate the binding from.
// Deprecated: Use PackOpenerMetaData.ABI instead.
var PackOpenerABI = PackOpenerMetaData.ABI

// PackOpener is an auto generated Go binding around an Ethereum contract.
type PackOpener struct {
	PackOpenerCaller     // Read-only binding to the contract
	PackOpenerTransactor // Write-only binding to the contract
	PackOpenerFilterer   // Log filterer for contract events
}

// PackOpenerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackOpenerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackOpenerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackOpenerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackOpenerSession struct {
	Contract     *PackOpener       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackOpenerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackOpenerCallerSession struct {
	Contract *PackOpenerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PackOpenerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackOpenerTransactorSession struct {
	Contract     *PackOpenerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PackOpenerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackOpenerRaw struct {
	Contract *PackOpener // Generic contract binding to access the raw methods on
}

// PackOpenerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackOpenerCallerRaw struct {
	Contract *PackOpenerCaller // Generic read-only contract binding to access the raw methods on
}

// PackOpenerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackOpenerTransactorRaw struct {
	Contract *PackOpenerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPackOpener creates a new instance of PackOpener, bound to a specific deployed contract.
func NewPackOpener(address common.Address, backend bind.ContractBackend) (*PackOpener, error) {
	contract, err := bindPackOpener(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PackOpener{PackOpenerCaller: PackOpenerCaller{contract: contract}, PackOpenerTransactor: PackOpenerTransactor{contract: contract}, PackOpenerFilterer: PackOpenerFilterer{contract: contract}}, nil
}

// NewPackOpenerCaller creates a new read-only instance of PackOpener, bound to a specific deployed contract.
func NewPackOpenerCaller(address common.Address, caller bind.ContractCaller) (*PackOpenerCaller, error) {
	contract, err := bindPackOpener(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackOpenerCaller{contract: contract}, nil
}

// NewPackOpenerTransactor creates a new write-only instance of PackOpener, bound to a specific deployed contract.
func NewPackOpenerTransactor(address common.Address, transactor bind.ContractTransactor) (*PackOpenerTransactor, error) {
	contract, err := bindPackOpener(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackOpenerTransactor{contract: contract}, nil
}

// NewPackOpenerFilterer creates a new log filterer instance of PackOpener, bound to a specific deployed contract.
func NewPackOpenerFilterer(address common.Address, filterer bind.ContractFilterer) (*PackOpenerFilterer, error) {
	contract, err := bindPackOpener(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackOpenerFilterer{contract: contract}, nil
}

// bindPackOpener binds a generic wrapper to an already deployed contract.
func bindPackOpener(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PackOpenerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackOpener *PackOpenerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackOpener.Contract.PackOpenerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackOpener *PackOpenerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpener.Contract.PackOpenerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackOpener *PackOpenerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackOpener.Contract.PackOpenerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackOpener *PackOpenerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackOpener.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackOpener *PackOpenerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpener.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackOpener *PackOpenerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackOpener.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpener *PackOpenerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PackOpener.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpener *PackOpenerSession) Owner() (common.Address, error) {
	return _PackOpener.Contract.Owner(&_PackOpener.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpener *PackOpenerCallerSession) Owner() (common.Address, error) {
	return _PackOpener.Contract.Owner(&_PackOpener.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpener *PackOpenerTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpener.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpener *PackOpenerSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpener.Contract.ChangeAdmin(&_PackOpener.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpener *PackOpenerTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpener.Contract.ChangeAdmin(&_PackOpener.TransactOpts, _newAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PackOpener *PackOpenerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpener.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PackOpener *PackOpenerSession) Initialize() (*types.Transaction, error) {
	return _PackOpener.Contract.Initialize(&_PackOpener.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_PackOpener *PackOpenerTransactorSession) Initialize() (*types.Transaction, error) {
	return _PackOpener.Contract.Initialize(&_PackOpener.TransactOpts)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_PackOpener *PackOpenerTransactor) OpenPack(opts *bind.TransactOpts, _tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _PackOpener.contract.Transact(opts, "openPack", _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_PackOpener *PackOpenerSession) OpenPack(_tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _PackOpener.Contract.OpenPack(&_PackOpener.TransactOpts, _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// OpenPack is a paid mutator transaction binding the contract method 0x3874820b.
//
// Solidity: function openPack(uint256 _tokenId, address _cardPackAddress, address _categoryAddress, address _yearAddress, address _dayMonthAddress, address _craftingCardAddress, address _triggerContract, string[] _tokenCategories, string[] _tokenYears, string[] _tokenDayMonths, string[] _tokenCraftingCards, string[] _tokenTriggers) returns()
func (_PackOpener *PackOpenerTransactorSession) OpenPack(_tokenId *big.Int, _cardPackAddress common.Address, _categoryAddress common.Address, _yearAddress common.Address, _dayMonthAddress common.Address, _craftingCardAddress common.Address, _triggerContract common.Address, _tokenCategories []string, _tokenYears []string, _tokenDayMonths []string, _tokenCraftingCards []string, _tokenTriggers []string) (*types.Transaction, error) {
	return _PackOpener.Contract.OpenPack(&_PackOpener.TransactOpts, _tokenId, _cardPackAddress, _categoryAddress, _yearAddress, _dayMonthAddress, _craftingCardAddress, _triggerContract, _tokenCategories, _tokenYears, _tokenDayMonths, _tokenCraftingCards, _tokenTriggers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpener *PackOpenerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpener.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpener *PackOpenerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PackOpener.Contract.RenounceOwnership(&_PackOpener.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpener *PackOpenerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PackOpener.Contract.RenounceOwnership(&_PackOpener.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpener *PackOpenerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PackOpener.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpener *PackOpenerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PackOpener.Contract.TransferOwnership(&_PackOpener.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpener *PackOpenerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PackOpener.Contract.TransferOwnership(&_PackOpener.TransactOpts, newOwner)
}

// PackOpenerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PackOpener contract.
type PackOpenerOwnershipTransferredIterator struct {
	Event *PackOpenerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PackOpenerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackOpenerOwnershipTransferred)
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
		it.Event = new(PackOpenerOwnershipTransferred)
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
func (it *PackOpenerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackOpenerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackOpenerOwnershipTransferred represents a OwnershipTransferred event raised by the PackOpener contract.
type PackOpenerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PackOpener *PackOpenerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PackOpenerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PackOpener.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PackOpenerOwnershipTransferredIterator{contract: _PackOpener.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PackOpener *PackOpenerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PackOpenerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PackOpener.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackOpenerOwnershipTransferred)
				if err := _PackOpener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PackOpener *PackOpenerFilterer) ParseOwnershipTransferred(log types.Log) (*PackOpenerOwnershipTransferred, error) {
	event := new(PackOpenerOwnershipTransferred)
	if err := _PackOpener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
