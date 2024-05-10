// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethwallet

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

// EthwalletMetaData contains all meta data concerning the Ethwallet contract.
var EthwalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EthwalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EthwalletMetaData.ABI instead.
var EthwalletABI = EthwalletMetaData.ABI

// Ethwallet is an auto generated Go binding around an Ethereum contract.
type Ethwallet struct {
	EthwalletCaller     // Read-only binding to the contract
	EthwalletTransactor // Write-only binding to the contract
	EthwalletFilterer   // Log filterer for contract events
}

// EthwalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthwalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthwalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthwalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthwalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthwalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthwalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthwalletSession struct {
	Contract     *Ethwallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthwalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthwalletCallerSession struct {
	Contract *EthwalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthwalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthwalletTransactorSession struct {
	Contract     *EthwalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthwalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthwalletRaw struct {
	Contract *Ethwallet // Generic contract binding to access the raw methods on
}

// EthwalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthwalletCallerRaw struct {
	Contract *EthwalletCaller // Generic read-only contract binding to access the raw methods on
}

// EthwalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthwalletTransactorRaw struct {
	Contract *EthwalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthwallet creates a new instance of Ethwallet, bound to a specific deployed contract.
func NewEthwallet(address common.Address, backend bind.ContractBackend) (*Ethwallet, error) {
	contract, err := bindEthwallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethwallet{EthwalletCaller: EthwalletCaller{contract: contract}, EthwalletTransactor: EthwalletTransactor{contract: contract}, EthwalletFilterer: EthwalletFilterer{contract: contract}}, nil
}

// NewEthwalletCaller creates a new read-only instance of Ethwallet, bound to a specific deployed contract.
func NewEthwalletCaller(address common.Address, caller bind.ContractCaller) (*EthwalletCaller, error) {
	contract, err := bindEthwallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthwalletCaller{contract: contract}, nil
}

// NewEthwalletTransactor creates a new write-only instance of Ethwallet, bound to a specific deployed contract.
func NewEthwalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EthwalletTransactor, error) {
	contract, err := bindEthwallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthwalletTransactor{contract: contract}, nil
}

// NewEthwalletFilterer creates a new log filterer instance of Ethwallet, bound to a specific deployed contract.
func NewEthwalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EthwalletFilterer, error) {
	contract, err := bindEthwallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthwalletFilterer{contract: contract}, nil
}

// bindEthwallet binds a generic wrapper to an already deployed contract.
func bindEthwallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthwalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethwallet *EthwalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethwallet.Contract.EthwalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethwallet *EthwalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethwallet.Contract.EthwalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethwallet *EthwalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethwallet.Contract.EthwalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethwallet *EthwalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethwallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethwallet *EthwalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethwallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethwallet *EthwalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethwallet.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Ethwallet *EthwalletCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethwallet.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Ethwallet *EthwalletSession) GetBalance() (*big.Int, error) {
	return _Ethwallet.Contract.GetBalance(&_Ethwallet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Ethwallet *EthwalletCallerSession) GetBalance() (*big.Int, error) {
	return _Ethwallet.Contract.GetBalance(&_Ethwallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethwallet *EthwalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethwallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethwallet *EthwalletSession) Owner() (common.Address, error) {
	return _Ethwallet.Contract.Owner(&_Ethwallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ethwallet *EthwalletCallerSession) Owner() (common.Address, error) {
	return _Ethwallet.Contract.Owner(&_Ethwallet.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Ethwallet *EthwalletTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Ethwallet.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Ethwallet *EthwalletSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Ethwallet.Contract.Withdraw(&_Ethwallet.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Ethwallet *EthwalletTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Ethwallet.Contract.Withdraw(&_Ethwallet.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethwallet *EthwalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethwallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethwallet *EthwalletSession) Receive() (*types.Transaction, error) {
	return _Ethwallet.Contract.Receive(&_Ethwallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethwallet *EthwalletTransactorSession) Receive() (*types.Transaction, error) {
	return _Ethwallet.Contract.Receive(&_Ethwallet.TransactOpts)
}
