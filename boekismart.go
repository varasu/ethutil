// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BoekismartABI is the input ABI used to generate the binding from.
const BoekismartABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyShadow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_delta\",\"type\":\"address\"},{\"name\":\"_master\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// BoekismartBin is the compiled bytecode used for deploying new contracts.
const BoekismartBin = `6060604052341561000f57600080fd5b6040516040806103a6833981016040528080519190602001805160008054600160a060020a03958616600160a060020a0319918216179091556001805495909216941693909317909255505061033c8061006a6000396000f3006060604052600436106100485763ffffffff60e060020a6000350416630000000781146100585780632e1a7d4d146100a7578063d0e30db0146100bd578063fcc2b41f14610058575b341561005357600080fd5b600080fd5b341561006357600080fd5b6100a5600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e435166101043561012435610144356100c5565b005b34156100b257600080fd5b6100a5600435610264565b6100a56102bf565b600054600160a060020a0316636c86888b8c8c8c8c8c8c8c8c8c8c8c3060405160e060020a63ffffffff8f16028152600160a060020a039c8d166004820152602481019b909b52988b1660448b015260648a0197909752608489019590955260a488019390935290871660c487015260ff1660e486015261010485015261012484015261014483015290911661016482015261018401602060405180830381600087803b151561017457600080fd5b5af1151561018157600080fd5b50505060405180519050151561019657610257565b600054600160a060020a0316630a19b14a8c8c8c8c8c8c8c8c8c8c8c60405160e060020a63ffffffff8e16028152600160a060020a039b8c166004820152602481019a909a52978a1660448a01526064890196909652608488019490945260a487019290925290951660c485015260ff90941660e484015261010483019390935261012482019290925261014481019190915261016401600060405180830381600087803b151561024657600080fd5b5af1151561025357600080fd5b5050505b5050505050505050505050565b600054600160a060020a0316632e1a7d4d8260405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b15156102ac57600080fd5b5af115156102b957600080fd5b50505050565b600054600160a060020a031663d0e30db06040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156102fe57600080fd5b5af1151561030b57600080fd5b5050505600a165627a7a723058206e174ab5f5a176afd54b721673136b61210c71b9848b70732aab4526186da9090029`

// DeployBoekismart deploys a new Ethereum contract, binding an instance of Boekismart to it.
func DeployBoekismart(auth *bind.TransactOpts, backend bind.ContractBackend, _delta common.Address, _master common.Address) (common.Address, *types.Transaction, *Boekismart, error) {
	parsed, err := abi.JSON(strings.NewReader(BoekismartABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BoekismartBin), backend, _delta, _master)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Boekismart{BoekismartCaller: BoekismartCaller{contract: contract}, BoekismartTransactor: BoekismartTransactor{contract: contract}, BoekismartFilterer: BoekismartFilterer{contract: contract}}, nil
}

// Boekismart is an auto generated Go binding around an Ethereum contract.
type Boekismart struct {
	BoekismartCaller     // Read-only binding to the contract
	BoekismartTransactor // Write-only binding to the contract
	BoekismartFilterer   // Log filterer for contract events
}

// BoekismartCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoekismartCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekismartTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoekismartTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekismartFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoekismartFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekismartSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoekismartSession struct {
	Contract     *Boekismart       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoekismartCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoekismartCallerSession struct {
	Contract *BoekismartCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BoekismartTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoekismartTransactorSession struct {
	Contract     *BoekismartTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BoekismartRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoekismartRaw struct {
	Contract *Boekismart // Generic contract binding to access the raw methods on
}

// BoekismartCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoekismartCallerRaw struct {
	Contract *BoekismartCaller // Generic read-only contract binding to access the raw methods on
}

// BoekismartTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoekismartTransactorRaw struct {
	Contract *BoekismartTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoekismart creates a new instance of Boekismart, bound to a specific deployed contract.
func NewBoekismart(address common.Address, backend bind.ContractBackend) (*Boekismart, error) {
	contract, err := bindBoekismart(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Boekismart{BoekismartCaller: BoekismartCaller{contract: contract}, BoekismartTransactor: BoekismartTransactor{contract: contract}, BoekismartFilterer: BoekismartFilterer{contract: contract}}, nil
}

// NewBoekismartCaller creates a new read-only instance of Boekismart, bound to a specific deployed contract.
func NewBoekismartCaller(address common.Address, caller bind.ContractCaller) (*BoekismartCaller, error) {
	contract, err := bindBoekismart(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoekismartCaller{contract: contract}, nil
}

// NewBoekismartTransactor creates a new write-only instance of Boekismart, bound to a specific deployed contract.
func NewBoekismartTransactor(address common.Address, transactor bind.ContractTransactor) (*BoekismartTransactor, error) {
	contract, err := bindBoekismart(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoekismartTransactor{contract: contract}, nil
}

// NewBoekismartFilterer creates a new log filterer instance of Boekismart, bound to a specific deployed contract.
func NewBoekismartFilterer(address common.Address, filterer bind.ContractFilterer) (*BoekismartFilterer, error) {
	contract, err := bindBoekismart(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoekismartFilterer{contract: contract}, nil
}

// bindBoekismart binds a generic wrapper to an already deployed contract.
func bindBoekismart(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoekismartABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boekismart *BoekismartRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Boekismart.Contract.BoekismartCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boekismart *BoekismartRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekismart.Contract.BoekismartTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boekismart *BoekismartRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boekismart.Contract.BoekismartTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boekismart *BoekismartCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Boekismart.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boekismart *BoekismartTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekismart.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boekismart *BoekismartTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boekismart.Contract.contract.Transact(opts, method, params...)
}

// Buy is a paid mutator transaction binding the contract method 0xfcc2b41f.
//
// Solidity: function buy(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartTransactor) Buy(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.contract.Transact(opts, "buy", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xfcc2b41f.
//
// Solidity: function buy(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartSession) Buy(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.Buy(&_Boekismart.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xfcc2b41f.
//
// Solidity: function buy(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartTransactorSession) Buy(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.Buy(&_Boekismart.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// BuyShadow is a paid mutator transaction binding the contract method 0x17c70dd1.
//
// Solidity: function buyShadow(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartTransactor) BuyShadow(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.contract.Transact(opts, "buyShadow", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// BuyShadow is a paid mutator transaction binding the contract method 0x17c70dd1.
//
// Solidity: function buyShadow(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartSession) BuyShadow(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.BuyShadow(&_Boekismart.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// BuyShadow is a paid mutator transaction binding the contract method 0x17c70dd1.
//
// Solidity: function buyShadow(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_Boekismart *BoekismartTransactorSession) BuyShadow(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.BuyShadow(&_Boekismart.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Boekismart *BoekismartTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekismart.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Boekismart *BoekismartSession) Deposit() (*types.Transaction, error) {
	return _Boekismart.Contract.Deposit(&_Boekismart.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Boekismart *BoekismartTransactorSession) Deposit() (*types.Transaction, error) {
	return _Boekismart.Contract.Deposit(&_Boekismart.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Boekismart *BoekismartTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Boekismart *BoekismartSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.Withdraw(&_Boekismart.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_Boekismart *BoekismartTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Boekismart.Contract.Withdraw(&_Boekismart.TransactOpts, amount)
}
