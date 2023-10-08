// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cardpack

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

// CardPackMetaData contains all meta data concerning the CardPack contract.
var CardPackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"CardPackCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CardPackOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"typeOf\",\"type\":\"string\"}],\"name\":\"LimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousMinter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMinter\",\"type\":\"address\"}],\"name\":\"MinterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOpener\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOpener\",\"type\":\"address\"}],\"name\":\"OpenerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"PayProxyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMinter\",\"type\":\"address\"}],\"name\":\"changeMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOpener\",\"type\":\"address\"}],\"name\":\"changeOpener\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"changeToOpened\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"changeToTotalLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tier\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"opened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opener\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setPayProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CardPackABI is the input ABI used to generate the binding from.
// Deprecated: Use CardPackMetaData.ABI instead.
var CardPackABI = CardPackMetaData.ABI

// CardPack is an auto generated Go binding around an Ethereum contract.
type CardPack struct {
	CardPackCaller     // Read-only binding to the contract
	CardPackTransactor // Write-only binding to the contract
	CardPackFilterer   // Log filterer for contract events
}

// CardPackCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardPackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardPackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardPackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardPackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardPackSession struct {
	Contract     *CardPack         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardPackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardPackCallerSession struct {
	Contract *CardPackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CardPackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardPackTransactorSession struct {
	Contract     *CardPackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CardPackRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardPackRaw struct {
	Contract *CardPack // Generic contract binding to access the raw methods on
}

// CardPackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardPackCallerRaw struct {
	Contract *CardPackCaller // Generic read-only contract binding to access the raw methods on
}

// CardPackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardPackTransactorRaw struct {
	Contract *CardPackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardPack creates a new instance of CardPack, bound to a specific deployed contract.
func NewCardPack(address common.Address, backend bind.ContractBackend) (*CardPack, error) {
	contract, err := bindCardPack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardPack{CardPackCaller: CardPackCaller{contract: contract}, CardPackTransactor: CardPackTransactor{contract: contract}, CardPackFilterer: CardPackFilterer{contract: contract}}, nil
}

// NewCardPackCaller creates a new read-only instance of CardPack, bound to a specific deployed contract.
func NewCardPackCaller(address common.Address, caller bind.ContractCaller) (*CardPackCaller, error) {
	contract, err := bindCardPack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardPackCaller{contract: contract}, nil
}

// NewCardPackTransactor creates a new write-only instance of CardPack, bound to a specific deployed contract.
func NewCardPackTransactor(address common.Address, transactor bind.ContractTransactor) (*CardPackTransactor, error) {
	contract, err := bindCardPack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardPackTransactor{contract: contract}, nil
}

// NewCardPackFilterer creates a new log filterer instance of CardPack, bound to a specific deployed contract.
func NewCardPackFilterer(address common.Address, filterer bind.ContractFilterer) (*CardPackFilterer, error) {
	contract, err := bindCardPack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardPackFilterer{contract: contract}, nil
}

// bindCardPack binds a generic wrapper to an already deployed contract.
func bindCardPack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CardPackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardPack *CardPackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardPack.Contract.CardPackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardPack *CardPackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPack.Contract.CardPackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardPack *CardPackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardPack.Contract.CardPackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardPack *CardPackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CardPack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardPack *CardPackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardPack *CardPackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardPack.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CardPack *CardPackCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CardPack *CardPackSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CardPack.Contract.BalanceOf(&_CardPack.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CardPack *CardPackCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CardPack.Contract.BalanceOf(&_CardPack.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CardPack *CardPackCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CardPack *CardPackSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CardPack.Contract.GetApproved(&_CardPack.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CardPack *CardPackCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CardPack.Contract.GetApproved(&_CardPack.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CardPack *CardPackCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CardPack *CardPackSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CardPack.Contract.IsApprovedForAll(&_CardPack.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CardPack *CardPackCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CardPack.Contract.IsApprovedForAll(&_CardPack.CallOpts, owner, operator)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_CardPack *CardPackCaller) IsOpened(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "isOpened", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_CardPack *CardPackSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _CardPack.Contract.IsOpened(&_CardPack.CallOpts, tokenId)
}

// IsOpened is a free data retrieval call binding the contract method 0x5faf46bb.
//
// Solidity: function isOpened(uint256 tokenId) view returns(bool)
func (_CardPack *CardPackCallerSession) IsOpened(tokenId *big.Int) (bool, error) {
	return _CardPack.Contract.IsOpened(&_CardPack.CallOpts, tokenId)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CardPack *CardPackCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CardPack *CardPackSession) Minter() (common.Address, error) {
	return _CardPack.Contract.Minter(&_CardPack.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CardPack *CardPackCallerSession) Minter() (common.Address, error) {
	return _CardPack.Contract.Minter(&_CardPack.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CardPack *CardPackCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CardPack *CardPackSession) Name() (string, error) {
	return _CardPack.Contract.Name(&_CardPack.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CardPack *CardPackCallerSession) Name() (string, error) {
	return _CardPack.Contract.Name(&_CardPack.CallOpts)
}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_CardPack *CardPackCaller) Opened(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "opened", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_CardPack *CardPackSession) Opened(arg0 *big.Int) (bool, error) {
	return _CardPack.Contract.Opened(&_CardPack.CallOpts, arg0)
}

// Opened is a free data retrieval call binding the contract method 0xf1ea5cd3.
//
// Solidity: function opened(uint256 ) view returns(bool)
func (_CardPack *CardPackCallerSession) Opened(arg0 *big.Int) (bool, error) {
	return _CardPack.Contract.Opened(&_CardPack.CallOpts, arg0)
}

// Opener is a free data retrieval call binding the contract method 0x9c9b4cdb.
//
// Solidity: function opener() view returns(address)
func (_CardPack *CardPackCaller) Opener(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "opener")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Opener is a free data retrieval call binding the contract method 0x9c9b4cdb.
//
// Solidity: function opener() view returns(address)
func (_CardPack *CardPackSession) Opener() (common.Address, error) {
	return _CardPack.Contract.Opener(&_CardPack.CallOpts)
}

// Opener is a free data retrieval call binding the contract method 0x9c9b4cdb.
//
// Solidity: function opener() view returns(address)
func (_CardPack *CardPackCallerSession) Opener() (common.Address, error) {
	return _CardPack.Contract.Opener(&_CardPack.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPack *CardPackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPack *CardPackSession) Owner() (common.Address, error) {
	return _CardPack.Contract.Owner(&_CardPack.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CardPack *CardPackCallerSession) Owner() (common.Address, error) {
	return _CardPack.Contract.Owner(&_CardPack.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CardPack *CardPackCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CardPack *CardPackSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CardPack.Contract.OwnerOf(&_CardPack.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CardPack *CardPackCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CardPack.Contract.OwnerOf(&_CardPack.CallOpts, tokenId)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_CardPack *CardPackCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_CardPack *CardPackSession) Proxy() (common.Address, error) {
	return _CardPack.Contract.Proxy(&_CardPack.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_CardPack *CardPackCallerSession) Proxy() (common.Address, error) {
	return _CardPack.Contract.Proxy(&_CardPack.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CardPack *CardPackCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CardPack *CardPackSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CardPack.Contract.SupportsInterface(&_CardPack.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CardPack *CardPackCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CardPack.Contract.SupportsInterface(&_CardPack.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CardPack *CardPackCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CardPack *CardPackSession) Symbol() (string, error) {
	return _CardPack.Contract.Symbol(&_CardPack.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CardPack *CardPackCallerSession) Symbol() (string, error) {
	return _CardPack.Contract.Symbol(&_CardPack.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CardPack *CardPackCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CardPack *CardPackSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CardPack.Contract.TokenURI(&_CardPack.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CardPack *CardPackCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CardPack.Contract.TokenURI(&_CardPack.CallOpts, tokenId)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_CardPack *CardPackCaller) TotalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CardPack.contract.Call(opts, &out, "totalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_CardPack *CardPackSession) TotalLimit() (*big.Int, error) {
	return _CardPack.Contract.TotalLimit(&_CardPack.CallOpts)
}

// TotalLimit is a free data retrieval call binding the contract method 0xa36298c7.
//
// Solidity: function totalLimit() view returns(uint256)
func (_CardPack *CardPackCallerSession) TotalLimit() (*big.Int, error) {
	return _CardPack.Contract.TotalLimit(&_CardPack.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CardPack *CardPackSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.Approve(&_CardPack.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.Approve(&_CardPack.TransactOpts, to, tokenId)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_CardPack *CardPackTransactor) ChangeMinter(opts *bind.TransactOpts, _newMinter common.Address) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "changeMinter", _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_CardPack *CardPackSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeMinter(&_CardPack.TransactOpts, _newMinter)
}

// ChangeMinter is a paid mutator transaction binding the contract method 0x2c4d4d18.
//
// Solidity: function changeMinter(address _newMinter) returns()
func (_CardPack *CardPackTransactorSession) ChangeMinter(_newMinter common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeMinter(&_CardPack.TransactOpts, _newMinter)
}

// ChangeOpener is a paid mutator transaction binding the contract method 0xc9d39c4d.
//
// Solidity: function changeOpener(address _newOpener) returns()
func (_CardPack *CardPackTransactor) ChangeOpener(opts *bind.TransactOpts, _newOpener common.Address) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "changeOpener", _newOpener)
}

// ChangeOpener is a paid mutator transaction binding the contract method 0xc9d39c4d.
//
// Solidity: function changeOpener(address _newOpener) returns()
func (_CardPack *CardPackSession) ChangeOpener(_newOpener common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeOpener(&_CardPack.TransactOpts, _newOpener)
}

// ChangeOpener is a paid mutator transaction binding the contract method 0xc9d39c4d.
//
// Solidity: function changeOpener(address _newOpener) returns()
func (_CardPack *CardPackTransactorSession) ChangeOpener(_newOpener common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeOpener(&_CardPack.TransactOpts, _newOpener)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_CardPack *CardPackTransactor) ChangeToOpened(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "changeToOpened", tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_CardPack *CardPackSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeToOpened(&_CardPack.TransactOpts, tokenId)
}

// ChangeToOpened is a paid mutator transaction binding the contract method 0xd9a99ac5.
//
// Solidity: function changeToOpened(uint256 tokenId) returns()
func (_CardPack *CardPackTransactorSession) ChangeToOpened(tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeToOpened(&_CardPack.TransactOpts, tokenId)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_CardPack *CardPackTransactor) ChangeToTotalLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "changeToTotalLimit", newLimit)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_CardPack *CardPackSession) ChangeToTotalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeToTotalLimit(&_CardPack.TransactOpts, newLimit)
}

// ChangeToTotalLimit is a paid mutator transaction binding the contract method 0x4bc8355a.
//
// Solidity: function changeToTotalLimit(uint256 newLimit) returns()
func (_CardPack *CardPackTransactorSession) ChangeToTotalLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.ChangeToTotalLimit(&_CardPack.TransactOpts, newLimit)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address _to, uint256 _quantity, uint256 _tier, string _tokenURI) payable returns(uint256)
func (_CardPack *CardPackTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _quantity *big.Int, _tier *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "mint", _to, _quantity, _tier, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address _to, uint256 _quantity, uint256 _tier, string _tokenURI) payable returns(uint256)
func (_CardPack *CardPackSession) Mint(_to common.Address, _quantity *big.Int, _tier *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CardPack.Contract.Mint(&_CardPack.TransactOpts, _to, _quantity, _tier, _tokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xbb7fde71.
//
// Solidity: function mint(address _to, uint256 _quantity, uint256 _tier, string _tokenURI) payable returns(uint256)
func (_CardPack *CardPackTransactorSession) Mint(_to common.Address, _quantity *big.Int, _tier *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _CardPack.Contract.Mint(&_CardPack.TransactOpts, _to, _quantity, _tier, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPack *CardPackTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPack *CardPackSession) RenounceOwnership() (*types.Transaction, error) {
	return _CardPack.Contract.RenounceOwnership(&_CardPack.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CardPack *CardPackTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CardPack.Contract.RenounceOwnership(&_CardPack.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.SafeTransferFrom(&_CardPack.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.SafeTransferFrom(&_CardPack.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CardPack *CardPackTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CardPack *CardPackSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CardPack.Contract.SafeTransferFrom0(&_CardPack.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CardPack *CardPackTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CardPack.Contract.SafeTransferFrom0(&_CardPack.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CardPack *CardPackTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CardPack *CardPackSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CardPack.Contract.SetApprovalForAll(&_CardPack.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CardPack *CardPackTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CardPack.Contract.SetApprovalForAll(&_CardPack.TransactOpts, operator, approved)
}

// SetPayProxy is a paid mutator transaction binding the contract method 0x285d9a28.
//
// Solidity: function setPayProxy(address _proxy) returns()
func (_CardPack *CardPackTransactor) SetPayProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "setPayProxy", _proxy)
}

// SetPayProxy is a paid mutator transaction binding the contract method 0x285d9a28.
//
// Solidity: function setPayProxy(address _proxy) returns()
func (_CardPack *CardPackSession) SetPayProxy(_proxy common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.SetPayProxy(&_CardPack.TransactOpts, _proxy)
}

// SetPayProxy is a paid mutator transaction binding the contract method 0x285d9a28.
//
// Solidity: function setPayProxy(address _proxy) returns()
func (_CardPack *CardPackTransactorSession) SetPayProxy(_proxy common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.SetPayProxy(&_CardPack.TransactOpts, _proxy)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.TransferFrom(&_CardPack.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CardPack *CardPackTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CardPack.Contract.TransferFrom(&_CardPack.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPack *CardPackTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CardPack.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPack *CardPackSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.TransferOwnership(&_CardPack.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CardPack *CardPackTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CardPack.Contract.TransferOwnership(&_CardPack.TransactOpts, newOwner)
}

// CardPackApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CardPack contract.
type CardPackApprovalIterator struct {
	Event *CardPackApproval // Event containing the contract specifics and raw log

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
func (it *CardPackApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackApproval)
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
		it.Event = new(CardPackApproval)
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
func (it *CardPackApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackApproval represents a Approval event raised by the CardPack contract.
type CardPackApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CardPackApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardPackApprovalIterator{contract: _CardPack.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CardPackApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackApproval)
				if err := _CardPack.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) ParseApproval(log types.Log) (*CardPackApproval, error) {
	event := new(CardPackApproval)
	if err := _CardPack.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CardPack contract.
type CardPackApprovalForAllIterator struct {
	Event *CardPackApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CardPackApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackApprovalForAll)
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
		it.Event = new(CardPackApprovalForAll)
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
func (it *CardPackApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackApprovalForAll represents a ApprovalForAll event raised by the CardPack contract.
type CardPackApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CardPack *CardPackFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CardPackApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CardPackApprovalForAllIterator{contract: _CardPack.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CardPack *CardPackFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CardPackApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackApprovalForAll)
				if err := _CardPack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CardPack *CardPackFilterer) ParseApprovalForAll(log types.Log) (*CardPackApprovalForAll, error) {
	event := new(CardPackApprovalForAll)
	if err := _CardPack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackCardPackCreatedIterator is returned from FilterCardPackCreated and is used to iterate over the raw logs and unpacked data for CardPackCreated events raised by the CardPack contract.
type CardPackCardPackCreatedIterator struct {
	Event *CardPackCardPackCreated // Event containing the contract specifics and raw log

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
func (it *CardPackCardPackCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackCardPackCreated)
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
		it.Event = new(CardPackCardPackCreated)
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
func (it *CardPackCardPackCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackCardPackCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackCardPackCreated represents a CardPackCreated event raised by the CardPack contract.
type CardPackCardPackCreated struct {
	TokenId  *big.Int
	Tier     *big.Int
	TokenURI string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCardPackCreated is a free log retrieval operation binding the contract event 0xe6439b927659f474e7c41df5081b02830e23b5627bd7da5652da84e0f7308057.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 tier, string tokenURI)
func (_CardPack *CardPackFilterer) FilterCardPackCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*CardPackCardPackCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "CardPackCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardPackCardPackCreatedIterator{contract: _CardPack.contract, event: "CardPackCreated", logs: logs, sub: sub}, nil
}

// WatchCardPackCreated is a free log subscription operation binding the contract event 0xe6439b927659f474e7c41df5081b02830e23b5627bd7da5652da84e0f7308057.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 tier, string tokenURI)
func (_CardPack *CardPackFilterer) WatchCardPackCreated(opts *bind.WatchOpts, sink chan<- *CardPackCardPackCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "CardPackCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackCardPackCreated)
				if err := _CardPack.contract.UnpackLog(event, "CardPackCreated", log); err != nil {
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

// ParseCardPackCreated is a log parse operation binding the contract event 0xe6439b927659f474e7c41df5081b02830e23b5627bd7da5652da84e0f7308057.
//
// Solidity: event CardPackCreated(uint256 indexed tokenId, uint256 tier, string tokenURI)
func (_CardPack *CardPackFilterer) ParseCardPackCreated(log types.Log) (*CardPackCardPackCreated, error) {
	event := new(CardPackCardPackCreated)
	if err := _CardPack.contract.UnpackLog(event, "CardPackCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackCardPackOpenedIterator is returned from FilterCardPackOpened and is used to iterate over the raw logs and unpacked data for CardPackOpened events raised by the CardPack contract.
type CardPackCardPackOpenedIterator struct {
	Event *CardPackCardPackOpened // Event containing the contract specifics and raw log

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
func (it *CardPackCardPackOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackCardPackOpened)
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
		it.Event = new(CardPackCardPackOpened)
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
func (it *CardPackCardPackOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackCardPackOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackCardPackOpened represents a CardPackOpened event raised by the CardPack contract.
type CardPackCardPackOpened struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCardPackOpened is a free log retrieval operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) FilterCardPackOpened(opts *bind.FilterOpts, tokenId []*big.Int) (*CardPackCardPackOpenedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "CardPackOpened", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardPackCardPackOpenedIterator{contract: _CardPack.contract, event: "CardPackOpened", logs: logs, sub: sub}, nil
}

// WatchCardPackOpened is a free log subscription operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) WatchCardPackOpened(opts *bind.WatchOpts, sink chan<- *CardPackCardPackOpened, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "CardPackOpened", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackCardPackOpened)
				if err := _CardPack.contract.UnpackLog(event, "CardPackOpened", log); err != nil {
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

// ParseCardPackOpened is a log parse operation binding the contract event 0xf3e8995f49b67b9b40d19d0ae915aff4ef00ecd661b45eaacdc47b3684279a34.
//
// Solidity: event CardPackOpened(uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) ParseCardPackOpened(log types.Log) (*CardPackCardPackOpened, error) {
	event := new(CardPackCardPackOpened)
	if err := _CardPack.contract.UnpackLog(event, "CardPackOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackLimitChangedIterator is returned from FilterLimitChanged and is used to iterate over the raw logs and unpacked data for LimitChanged events raised by the CardPack contract.
type CardPackLimitChangedIterator struct {
	Event *CardPackLimitChanged // Event containing the contract specifics and raw log

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
func (it *CardPackLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackLimitChanged)
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
		it.Event = new(CardPackLimitChanged)
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
func (it *CardPackLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackLimitChanged represents a LimitChanged event raised by the CardPack contract.
type CardPackLimitChanged struct {
	NewLimit *big.Int
	TypeOf   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLimitChanged is a free log retrieval operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_CardPack *CardPackFilterer) FilterLimitChanged(opts *bind.FilterOpts) (*CardPackLimitChangedIterator, error) {

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return &CardPackLimitChangedIterator{contract: _CardPack.contract, event: "LimitChanged", logs: logs, sub: sub}, nil
}

// WatchLimitChanged is a free log subscription operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_CardPack *CardPackFilterer) WatchLimitChanged(opts *bind.WatchOpts, sink chan<- *CardPackLimitChanged) (event.Subscription, error) {

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "LimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackLimitChanged)
				if err := _CardPack.contract.UnpackLog(event, "LimitChanged", log); err != nil {
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

// ParseLimitChanged is a log parse operation binding the contract event 0x9fc072c98b6eb02f368bc5c6765f963c699b663e23b0bbf06712102ce81d0b65.
//
// Solidity: event LimitChanged(uint256 newLimit, string typeOf)
func (_CardPack *CardPackFilterer) ParseLimitChanged(log types.Log) (*CardPackLimitChanged, error) {
	event := new(CardPackLimitChanged)
	if err := _CardPack.contract.UnpackLog(event, "LimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackMinterChangedIterator is returned from FilterMinterChanged and is used to iterate over the raw logs and unpacked data for MinterChanged events raised by the CardPack contract.
type CardPackMinterChangedIterator struct {
	Event *CardPackMinterChanged // Event containing the contract specifics and raw log

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
func (it *CardPackMinterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackMinterChanged)
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
		it.Event = new(CardPackMinterChanged)
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
func (it *CardPackMinterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackMinterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackMinterChanged represents a MinterChanged event raised by the CardPack contract.
type CardPackMinterChanged struct {
	PreviousMinter common.Address
	NewMinter      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinterChanged is a free log retrieval operation binding the contract event 0x3b0007eb941cf645526cbb3a4fdaecda9d28ce4843167d9263b536a1f1edc0f6.
//
// Solidity: event MinterChanged(address indexed previousMinter, address indexed newMinter)
func (_CardPack *CardPackFilterer) FilterMinterChanged(opts *bind.FilterOpts, previousMinter []common.Address, newMinter []common.Address) (*CardPackMinterChangedIterator, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "MinterChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &CardPackMinterChangedIterator{contract: _CardPack.contract, event: "MinterChanged", logs: logs, sub: sub}, nil
}

// WatchMinterChanged is a free log subscription operation binding the contract event 0x3b0007eb941cf645526cbb3a4fdaecda9d28ce4843167d9263b536a1f1edc0f6.
//
// Solidity: event MinterChanged(address indexed previousMinter, address indexed newMinter)
func (_CardPack *CardPackFilterer) WatchMinterChanged(opts *bind.WatchOpts, sink chan<- *CardPackMinterChanged, previousMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var previousMinterRule []interface{}
	for _, previousMinterItem := range previousMinter {
		previousMinterRule = append(previousMinterRule, previousMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "MinterChanged", previousMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackMinterChanged)
				if err := _CardPack.contract.UnpackLog(event, "MinterChanged", log); err != nil {
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

// ParseMinterChanged is a log parse operation binding the contract event 0x3b0007eb941cf645526cbb3a4fdaecda9d28ce4843167d9263b536a1f1edc0f6.
//
// Solidity: event MinterChanged(address indexed previousMinter, address indexed newMinter)
func (_CardPack *CardPackFilterer) ParseMinterChanged(log types.Log) (*CardPackMinterChanged, error) {
	event := new(CardPackMinterChanged)
	if err := _CardPack.contract.UnpackLog(event, "MinterChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackOpenerChangedIterator is returned from FilterOpenerChanged and is used to iterate over the raw logs and unpacked data for OpenerChanged events raised by the CardPack contract.
type CardPackOpenerChangedIterator struct {
	Event *CardPackOpenerChanged // Event containing the contract specifics and raw log

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
func (it *CardPackOpenerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackOpenerChanged)
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
		it.Event = new(CardPackOpenerChanged)
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
func (it *CardPackOpenerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackOpenerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackOpenerChanged represents a OpenerChanged event raised by the CardPack contract.
type CardPackOpenerChanged struct {
	PreviousOpener common.Address
	NewOpener      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOpenerChanged is a free log retrieval operation binding the contract event 0x990db547fbdba970bfd1d591a8721559e6b2ab12cf06967606db7e26f9d19d0b.
//
// Solidity: event OpenerChanged(address indexed previousOpener, address indexed newOpener)
func (_CardPack *CardPackFilterer) FilterOpenerChanged(opts *bind.FilterOpts, previousOpener []common.Address, newOpener []common.Address) (*CardPackOpenerChangedIterator, error) {

	var previousOpenerRule []interface{}
	for _, previousOpenerItem := range previousOpener {
		previousOpenerRule = append(previousOpenerRule, previousOpenerItem)
	}
	var newOpenerRule []interface{}
	for _, newOpenerItem := range newOpener {
		newOpenerRule = append(newOpenerRule, newOpenerItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "OpenerChanged", previousOpenerRule, newOpenerRule)
	if err != nil {
		return nil, err
	}
	return &CardPackOpenerChangedIterator{contract: _CardPack.contract, event: "OpenerChanged", logs: logs, sub: sub}, nil
}

// WatchOpenerChanged is a free log subscription operation binding the contract event 0x990db547fbdba970bfd1d591a8721559e6b2ab12cf06967606db7e26f9d19d0b.
//
// Solidity: event OpenerChanged(address indexed previousOpener, address indexed newOpener)
func (_CardPack *CardPackFilterer) WatchOpenerChanged(opts *bind.WatchOpts, sink chan<- *CardPackOpenerChanged, previousOpener []common.Address, newOpener []common.Address) (event.Subscription, error) {

	var previousOpenerRule []interface{}
	for _, previousOpenerItem := range previousOpener {
		previousOpenerRule = append(previousOpenerRule, previousOpenerItem)
	}
	var newOpenerRule []interface{}
	for _, newOpenerItem := range newOpener {
		newOpenerRule = append(newOpenerRule, newOpenerItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "OpenerChanged", previousOpenerRule, newOpenerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackOpenerChanged)
				if err := _CardPack.contract.UnpackLog(event, "OpenerChanged", log); err != nil {
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

// ParseOpenerChanged is a log parse operation binding the contract event 0x990db547fbdba970bfd1d591a8721559e6b2ab12cf06967606db7e26f9d19d0b.
//
// Solidity: event OpenerChanged(address indexed previousOpener, address indexed newOpener)
func (_CardPack *CardPackFilterer) ParseOpenerChanged(log types.Log) (*CardPackOpenerChanged, error) {
	event := new(CardPackOpenerChanged)
	if err := _CardPack.contract.UnpackLog(event, "OpenerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CardPack contract.
type CardPackOwnershipTransferredIterator struct {
	Event *CardPackOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CardPackOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackOwnershipTransferred)
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
		it.Event = new(CardPackOwnershipTransferred)
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
func (it *CardPackOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackOwnershipTransferred represents a OwnershipTransferred event raised by the CardPack contract.
type CardPackOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CardPack *CardPackFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CardPackOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CardPackOwnershipTransferredIterator{contract: _CardPack.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CardPack *CardPackFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CardPackOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackOwnershipTransferred)
				if err := _CardPack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CardPack *CardPackFilterer) ParseOwnershipTransferred(log types.Log) (*CardPackOwnershipTransferred, error) {
	event := new(CardPackOwnershipTransferred)
	if err := _CardPack.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackPayProxyChangedIterator is returned from FilterPayProxyChanged and is used to iterate over the raw logs and unpacked data for PayProxyChanged events raised by the CardPack contract.
type CardPackPayProxyChangedIterator struct {
	Event *CardPackPayProxyChanged // Event containing the contract specifics and raw log

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
func (it *CardPackPayProxyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackPayProxyChanged)
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
		it.Event = new(CardPackPayProxyChanged)
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
func (it *CardPackPayProxyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackPayProxyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackPayProxyChanged represents a PayProxyChanged event raised by the CardPack contract.
type CardPackPayProxyChanged struct {
	PreviousProxy common.Address
	NewProxy      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPayProxyChanged is a free log retrieval operation binding the contract event 0x1b2b8822d1235c1b20a8ba5fa879672ca8dacd95a234479ba71eef6d29303a2b.
//
// Solidity: event PayProxyChanged(address indexed previousProxy, address indexed newProxy)
func (_CardPack *CardPackFilterer) FilterPayProxyChanged(opts *bind.FilterOpts, previousProxy []common.Address, newProxy []common.Address) (*CardPackPayProxyChangedIterator, error) {

	var previousProxyRule []interface{}
	for _, previousProxyItem := range previousProxy {
		previousProxyRule = append(previousProxyRule, previousProxyItem)
	}
	var newProxyRule []interface{}
	for _, newProxyItem := range newProxy {
		newProxyRule = append(newProxyRule, newProxyItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "PayProxyChanged", previousProxyRule, newProxyRule)
	if err != nil {
		return nil, err
	}
	return &CardPackPayProxyChangedIterator{contract: _CardPack.contract, event: "PayProxyChanged", logs: logs, sub: sub}, nil
}

// WatchPayProxyChanged is a free log subscription operation binding the contract event 0x1b2b8822d1235c1b20a8ba5fa879672ca8dacd95a234479ba71eef6d29303a2b.
//
// Solidity: event PayProxyChanged(address indexed previousProxy, address indexed newProxy)
func (_CardPack *CardPackFilterer) WatchPayProxyChanged(opts *bind.WatchOpts, sink chan<- *CardPackPayProxyChanged, previousProxy []common.Address, newProxy []common.Address) (event.Subscription, error) {

	var previousProxyRule []interface{}
	for _, previousProxyItem := range previousProxy {
		previousProxyRule = append(previousProxyRule, previousProxyItem)
	}
	var newProxyRule []interface{}
	for _, newProxyItem := range newProxy {
		newProxyRule = append(newProxyRule, newProxyItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "PayProxyChanged", previousProxyRule, newProxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackPayProxyChanged)
				if err := _CardPack.contract.UnpackLog(event, "PayProxyChanged", log); err != nil {
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

// ParsePayProxyChanged is a log parse operation binding the contract event 0x1b2b8822d1235c1b20a8ba5fa879672ca8dacd95a234479ba71eef6d29303a2b.
//
// Solidity: event PayProxyChanged(address indexed previousProxy, address indexed newProxy)
func (_CardPack *CardPackFilterer) ParsePayProxyChanged(log types.Log) (*CardPackPayProxyChanged, error) {
	event := new(CardPackPayProxyChanged)
	if err := _CardPack.contract.UnpackLog(event, "PayProxyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CardPackTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CardPack contract.
type CardPackTransferIterator struct {
	Event *CardPackTransfer // Event containing the contract specifics and raw log

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
func (it *CardPackTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardPackTransfer)
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
		it.Event = new(CardPackTransfer)
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
func (it *CardPackTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardPackTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardPackTransfer represents a Transfer event raised by the CardPack contract.
type CardPackTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CardPackTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CardPackTransferIterator{contract: _CardPack.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardPackTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CardPack.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardPackTransfer)
				if err := _CardPack.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CardPack *CardPackFilterer) ParseTransfer(log types.Log) (*CardPackTransfer, error) {
	event := new(CardPackTransfer)
	if err := _CardPack.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
