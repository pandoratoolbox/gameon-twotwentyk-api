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

// PackOpenerV1MetaData contains all meta data concerning the PackOpenerV1 contract.
var PackOpenerV1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cardPackAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dayMonthAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_yearAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_craftingCardContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_triggerContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dayMonthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_yearAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_categoryAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_craftingCardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_triggerAmount\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_tokenDayMonth\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenYear\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCategory\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenCraftingCard\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenTrigger\",\"type\":\"string[]\"}],\"name\":\"openPack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PackOpenerV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use PackOpenerV1MetaData.ABI instead.
var PackOpenerV1ABI = PackOpenerV1MetaData.ABI

// PackOpenerV1 is an auto generated Go binding around an Ethereum contract.
type PackOpenerV1 struct {
	PackOpenerV1Caller     // Read-only binding to the contract
	PackOpenerV1Transactor // Write-only binding to the contract
	PackOpenerV1Filterer   // Log filterer for contract events
}

// PackOpenerV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type PackOpenerV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type PackOpenerV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackOpenerV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackOpenerV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackOpenerV1Session struct {
	Contract     *PackOpenerV1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackOpenerV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackOpenerV1CallerSession struct {
	Contract *PackOpenerV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PackOpenerV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackOpenerV1TransactorSession struct {
	Contract     *PackOpenerV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PackOpenerV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type PackOpenerV1Raw struct {
	Contract *PackOpenerV1 // Generic contract binding to access the raw methods on
}

// PackOpenerV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackOpenerV1CallerRaw struct {
	Contract *PackOpenerV1Caller // Generic read-only contract binding to access the raw methods on
}

// PackOpenerV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackOpenerV1TransactorRaw struct {
	Contract *PackOpenerV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPackOpenerV1 creates a new instance of PackOpenerV1, bound to a specific deployed contract.
func NewPackOpenerV1(address common.Address, backend bind.ContractBackend) (*PackOpenerV1, error) {
	contract, err := bindPackOpenerV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PackOpenerV1{PackOpenerV1Caller: PackOpenerV1Caller{contract: contract}, PackOpenerV1Transactor: PackOpenerV1Transactor{contract: contract}, PackOpenerV1Filterer: PackOpenerV1Filterer{contract: contract}}, nil
}

// NewPackOpenerV1Caller creates a new read-only instance of PackOpenerV1, bound to a specific deployed contract.
func NewPackOpenerV1Caller(address common.Address, caller bind.ContractCaller) (*PackOpenerV1Caller, error) {
	contract, err := bindPackOpenerV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackOpenerV1Caller{contract: contract}, nil
}

// NewPackOpenerV1Transactor creates a new write-only instance of PackOpenerV1, bound to a specific deployed contract.
func NewPackOpenerV1Transactor(address common.Address, transactor bind.ContractTransactor) (*PackOpenerV1Transactor, error) {
	contract, err := bindPackOpenerV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackOpenerV1Transactor{contract: contract}, nil
}

// NewPackOpenerV1Filterer creates a new log filterer instance of PackOpenerV1, bound to a specific deployed contract.
func NewPackOpenerV1Filterer(address common.Address, filterer bind.ContractFilterer) (*PackOpenerV1Filterer, error) {
	contract, err := bindPackOpenerV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackOpenerV1Filterer{contract: contract}, nil
}

// bindPackOpenerV1 binds a generic wrapper to an already deployed contract.
func bindPackOpenerV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PackOpenerV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackOpenerV1 *PackOpenerV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackOpenerV1.Contract.PackOpenerV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackOpenerV1 *PackOpenerV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.PackOpenerV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackOpenerV1 *PackOpenerV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.PackOpenerV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackOpenerV1 *PackOpenerV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackOpenerV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackOpenerV1 *PackOpenerV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackOpenerV1 *PackOpenerV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpenerV1 *PackOpenerV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PackOpenerV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpenerV1 *PackOpenerV1Session) Owner() (common.Address, error) {
	return _PackOpenerV1.Contract.Owner(&_PackOpenerV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PackOpenerV1 *PackOpenerV1CallerSession) Owner() (common.Address, error) {
	return _PackOpenerV1.Contract.Owner(&_PackOpenerV1.CallOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpenerV1 *PackOpenerV1Transactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpenerV1 *PackOpenerV1Session) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.ChangeAdmin(&_PackOpenerV1.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_PackOpenerV1 *PackOpenerV1TransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.ChangeAdmin(&_PackOpenerV1.TransactOpts, _newAdmin)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _cardPackAddress, address _dayMonthAddress, address _yearAddress, address _categoryAddress, address _craftingCardContract, address _triggerContract) returns()
func (_PackOpenerV1 *PackOpenerV1Transactor) Initialize(opts *bind.TransactOpts, _cardPackAddress common.Address, _dayMonthAddress common.Address, _yearAddress common.Address, _categoryAddress common.Address, _craftingCardContract common.Address, _triggerContract common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.contract.Transact(opts, "initialize", _cardPackAddress, _dayMonthAddress, _yearAddress, _categoryAddress, _craftingCardContract, _triggerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _cardPackAddress, address _dayMonthAddress, address _yearAddress, address _categoryAddress, address _craftingCardContract, address _triggerContract) returns()
func (_PackOpenerV1 *PackOpenerV1Session) Initialize(_cardPackAddress common.Address, _dayMonthAddress common.Address, _yearAddress common.Address, _categoryAddress common.Address, _craftingCardContract common.Address, _triggerContract common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.Initialize(&_PackOpenerV1.TransactOpts, _cardPackAddress, _dayMonthAddress, _yearAddress, _categoryAddress, _craftingCardContract, _triggerContract)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address _cardPackAddress, address _dayMonthAddress, address _yearAddress, address _categoryAddress, address _craftingCardContract, address _triggerContract) returns()
func (_PackOpenerV1 *PackOpenerV1TransactorSession) Initialize(_cardPackAddress common.Address, _dayMonthAddress common.Address, _yearAddress common.Address, _categoryAddress common.Address, _craftingCardContract common.Address, _triggerContract common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.Initialize(&_PackOpenerV1.TransactOpts, _cardPackAddress, _dayMonthAddress, _yearAddress, _categoryAddress, _craftingCardContract, _triggerContract)
}

// OpenPack is a paid mutator transaction binding the contract method 0x299f2d03.
//
// Solidity: function openPack(uint256 _tokenId, uint256 _dayMonthAmount, uint256 _yearAmount, uint256 _categoryAmount, uint256 _craftingCardAmount, uint256 _triggerAmount, string[] _tokenDayMonth, string[] _tokenYear, string[] _tokenCategory, string[] _tokenCraftingCard, string[] _tokenTrigger) returns()
func (_PackOpenerV1 *PackOpenerV1Transactor) OpenPack(opts *bind.TransactOpts, _tokenId *big.Int, _dayMonthAmount *big.Int, _yearAmount *big.Int, _categoryAmount *big.Int, _craftingCardAmount *big.Int, _triggerAmount *big.Int, _tokenDayMonth []string, _tokenYear []string, _tokenCategory []string, _tokenCraftingCard []string, _tokenTrigger []string) (*types.Transaction, error) {
	return _PackOpenerV1.contract.Transact(opts, "openPack", _tokenId, _dayMonthAmount, _yearAmount, _categoryAmount, _craftingCardAmount, _triggerAmount, _tokenDayMonth, _tokenYear, _tokenCategory, _tokenCraftingCard, _tokenTrigger)
}

// OpenPack is a paid mutator transaction binding the contract method 0x299f2d03.
//
// Solidity: function openPack(uint256 _tokenId, uint256 _dayMonthAmount, uint256 _yearAmount, uint256 _categoryAmount, uint256 _craftingCardAmount, uint256 _triggerAmount, string[] _tokenDayMonth, string[] _tokenYear, string[] _tokenCategory, string[] _tokenCraftingCard, string[] _tokenTrigger) returns()
func (_PackOpenerV1 *PackOpenerV1Session) OpenPack(_tokenId *big.Int, _dayMonthAmount *big.Int, _yearAmount *big.Int, _categoryAmount *big.Int, _craftingCardAmount *big.Int, _triggerAmount *big.Int, _tokenDayMonth []string, _tokenYear []string, _tokenCategory []string, _tokenCraftingCard []string, _tokenTrigger []string) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.OpenPack(&_PackOpenerV1.TransactOpts, _tokenId, _dayMonthAmount, _yearAmount, _categoryAmount, _craftingCardAmount, _triggerAmount, _tokenDayMonth, _tokenYear, _tokenCategory, _tokenCraftingCard, _tokenTrigger)
}

// OpenPack is a paid mutator transaction binding the contract method 0x299f2d03.
//
// Solidity: function openPack(uint256 _tokenId, uint256 _dayMonthAmount, uint256 _yearAmount, uint256 _categoryAmount, uint256 _craftingCardAmount, uint256 _triggerAmount, string[] _tokenDayMonth, string[] _tokenYear, string[] _tokenCategory, string[] _tokenCraftingCard, string[] _tokenTrigger) returns()
func (_PackOpenerV1 *PackOpenerV1TransactorSession) OpenPack(_tokenId *big.Int, _dayMonthAmount *big.Int, _yearAmount *big.Int, _categoryAmount *big.Int, _craftingCardAmount *big.Int, _triggerAmount *big.Int, _tokenDayMonth []string, _tokenYear []string, _tokenCategory []string, _tokenCraftingCard []string, _tokenTrigger []string) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.OpenPack(&_PackOpenerV1.TransactOpts, _tokenId, _dayMonthAmount, _yearAmount, _categoryAmount, _craftingCardAmount, _triggerAmount, _tokenDayMonth, _tokenYear, _tokenCategory, _tokenCraftingCard, _tokenTrigger)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpenerV1 *PackOpenerV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackOpenerV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpenerV1 *PackOpenerV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _PackOpenerV1.Contract.RenounceOwnership(&_PackOpenerV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PackOpenerV1 *PackOpenerV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PackOpenerV1.Contract.RenounceOwnership(&_PackOpenerV1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpenerV1 *PackOpenerV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpenerV1 *PackOpenerV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.TransferOwnership(&_PackOpenerV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PackOpenerV1 *PackOpenerV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PackOpenerV1.Contract.TransferOwnership(&_PackOpenerV1.TransactOpts, newOwner)
}

// PackOpenerV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PackOpenerV1 contract.
type PackOpenerV1OwnershipTransferredIterator struct {
	Event *PackOpenerV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PackOpenerV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackOpenerV1OwnershipTransferred)
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
		it.Event = new(PackOpenerV1OwnershipTransferred)
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
func (it *PackOpenerV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackOpenerV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackOpenerV1OwnershipTransferred represents a OwnershipTransferred event raised by the PackOpenerV1 contract.
type PackOpenerV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PackOpenerV1 *PackOpenerV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PackOpenerV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PackOpenerV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PackOpenerV1OwnershipTransferredIterator{contract: _PackOpenerV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PackOpenerV1 *PackOpenerV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PackOpenerV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PackOpenerV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackOpenerV1OwnershipTransferred)
				if err := _PackOpenerV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PackOpenerV1 *PackOpenerV1Filterer) ParseOwnershipTransferred(log types.Log) (*PackOpenerV1OwnershipTransferred, error) {
	event := new(PackOpenerV1OwnershipTransferred)
	if err := _PackOpenerV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
