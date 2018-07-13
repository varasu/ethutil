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

// BoekiproxyABI is the input ABI used to generate the binding from.
const BoekiproxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"nop\",\"type\":\"address\"}],\"name\":\"changeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken_BOEKI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit_BOEKI\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw_BOEKI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_dca\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// BoekiproxyBin is the compiled bytecode used for deploying new contracts.
const BoekiproxyBin = `6060604052341561000f57600080fd5b6040516040806104a1833981016040528080519190602001805160008054600160a060020a03928316600160a060020a03199182161790915560028054958316958216959095179094556001805433909216919094161790925550506104278061007a6000396000f30060606040526004361061005e5763ffffffff60e060020a60003504166306394c9b81146100a45780633c036b0c146100c5578063570ca735146100e75780636902750814610116578063a6f9dae11461011e578063ef30a6bc1461013d575b60025460009033600160a060020a0390811691161461007c57600080fd5b50600054600160a060020a03166040513660008237602081368334865af18015600257602082f35b34156100af57600080fd5b6100c3600160a060020a0360043516610153565b005b34156100d057600080fd5b6100c3600160a060020a036004351660243561019d565b34156100f257600080fd5b6100fa61028d565b604051600160a060020a03909116815260200160405180910390f35b6100c361029c565b341561012957600080fd5b6100c3600160a060020a0360043516610309565b341561014857600080fd5b6100c3600435610353565b60015433600160a060020a0390811691161461016e57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015433600160a060020a039081169116146101b857600080fd5b600054600160a060020a0316639e281a98838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561020e57600080fd5b5af1151561021b57600080fd5b50505081600160a060020a031663a9059cbb338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561027257600080fd5b5af1151561027f57600080fd5b505050604051805150505050565b600254600160a060020a031681565b60015433600160a060020a039081169116146102b757600080fd5b600054600160a060020a031663d0e30db0346040518263ffffffff1660e060020a0281526004016000604051808303818588803b15156102f657600080fd5b5af1151561030357600080fd5b50505050565b60015433600160a060020a0390811691161461032457600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015433600160a060020a0390811691161461036e57600080fd5b600054600160a060020a0316632e1a7d4d8260405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b15156103b657600080fd5b5af115156103c357600080fd5b5050600160a060020a033316905081156108fc0282604051600060405180830381858888f1935050505015156103f857600080fd5b505600a165627a7a72305820ac60e3022b4fbe241a8a0887d7c10cdf0ec9e56ca050e4d0c9bb4b2aa219bb3e0029`

// DeployBoekiproxy deploys a new Ethereum contract, binding an instance of Boekiproxy to it.
func DeployBoekiproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _dca common.Address) (common.Address, *types.Transaction, *Boekiproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(BoekiproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BoekiproxyBin), backend, _operator, _dca)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Boekiproxy{BoekiproxyCaller: BoekiproxyCaller{contract: contract}, BoekiproxyTransactor: BoekiproxyTransactor{contract: contract}, BoekiproxyFilterer: BoekiproxyFilterer{contract: contract}}, nil
}

// Boekiproxy is an auto generated Go binding around an Ethereum contract.
type Boekiproxy struct {
	BoekiproxyCaller     // Read-only binding to the contract
	BoekiproxyTransactor // Write-only binding to the contract
	BoekiproxyFilterer   // Log filterer for contract events
}

// BoekiproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoekiproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekiproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoekiproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekiproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoekiproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoekiproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoekiproxySession struct {
	Contract     *Boekiproxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoekiproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoekiproxyCallerSession struct {
	Contract *BoekiproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BoekiproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoekiproxyTransactorSession struct {
	Contract     *BoekiproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BoekiproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoekiproxyRaw struct {
	Contract *Boekiproxy // Generic contract binding to access the raw methods on
}

// BoekiproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoekiproxyCallerRaw struct {
	Contract *BoekiproxyCaller // Generic read-only contract binding to access the raw methods on
}

// BoekiproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoekiproxyTransactorRaw struct {
	Contract *BoekiproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoekiproxy creates a new instance of Boekiproxy, bound to a specific deployed contract.
func NewBoekiproxy(address common.Address, backend bind.ContractBackend) (*Boekiproxy, error) {
	contract, err := bindBoekiproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Boekiproxy{BoekiproxyCaller: BoekiproxyCaller{contract: contract}, BoekiproxyTransactor: BoekiproxyTransactor{contract: contract}, BoekiproxyFilterer: BoekiproxyFilterer{contract: contract}}, nil
}

// NewBoekiproxyCaller creates a new read-only instance of Boekiproxy, bound to a specific deployed contract.
func NewBoekiproxyCaller(address common.Address, caller bind.ContractCaller) (*BoekiproxyCaller, error) {
	contract, err := bindBoekiproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoekiproxyCaller{contract: contract}, nil
}

// NewBoekiproxyTransactor creates a new write-only instance of Boekiproxy, bound to a specific deployed contract.
func NewBoekiproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BoekiproxyTransactor, error) {
	contract, err := bindBoekiproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoekiproxyTransactor{contract: contract}, nil
}

// NewBoekiproxyFilterer creates a new log filterer instance of Boekiproxy, bound to a specific deployed contract.
func NewBoekiproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BoekiproxyFilterer, error) {
	contract, err := bindBoekiproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoekiproxyFilterer{contract: contract}, nil
}

// bindBoekiproxy binds a generic wrapper to an already deployed contract.
func bindBoekiproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoekiproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boekiproxy *BoekiproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Boekiproxy.Contract.BoekiproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boekiproxy *BoekiproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekiproxy.Contract.BoekiproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boekiproxy *BoekiproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boekiproxy.Contract.BoekiproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boekiproxy *BoekiproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Boekiproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boekiproxy *BoekiproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekiproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boekiproxy *BoekiproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boekiproxy.Contract.contract.Transact(opts, method, params...)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Boekiproxy *BoekiproxyCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Boekiproxy.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Boekiproxy *BoekiproxySession) Operator() (common.Address, error) {
	return _Boekiproxy.Contract.Operator(&_Boekiproxy.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_Boekiproxy *BoekiproxyCallerSession) Operator() (common.Address, error) {
	return _Boekiproxy.Contract.Operator(&_Boekiproxy.CallOpts)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(nop address) returns()
func (_Boekiproxy *BoekiproxyTransactor) ChangeOperator(opts *bind.TransactOpts, nop common.Address) (*types.Transaction, error) {
	return _Boekiproxy.contract.Transact(opts, "changeOperator", nop)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(nop address) returns()
func (_Boekiproxy *BoekiproxySession) ChangeOperator(nop common.Address) (*types.Transaction, error) {
	return _Boekiproxy.Contract.ChangeOperator(&_Boekiproxy.TransactOpts, nop)
}

// ChangeOperator is a paid mutator transaction binding the contract method 0x06394c9b.
//
// Solidity: function changeOperator(nop address) returns()
func (_Boekiproxy *BoekiproxyTransactorSession) ChangeOperator(nop common.Address) (*types.Transaction, error) {
	return _Boekiproxy.Contract.ChangeOperator(&_Boekiproxy.TransactOpts, nop)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Boekiproxy *BoekiproxyTransactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Boekiproxy.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Boekiproxy *BoekiproxySession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Boekiproxy.Contract.ChangeOwner(&_Boekiproxy.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_owner address) returns()
func (_Boekiproxy *BoekiproxyTransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _Boekiproxy.Contract.ChangeOwner(&_Boekiproxy.TransactOpts, _owner)
}

// Deposit_BOEKI is a paid mutator transaction binding the contract method 0x69027508.
//
// Solidity: function deposit_BOEKI() returns()
func (_Boekiproxy *BoekiproxyTransactor) Deposit_BOEKI(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boekiproxy.contract.Transact(opts, "deposit_BOEKI")
}

// Deposit_BOEKI is a paid mutator transaction binding the contract method 0x69027508.
//
// Solidity: function deposit_BOEKI() returns()
func (_Boekiproxy *BoekiproxySession) Deposit_BOEKI() (*types.Transaction, error) {
	return _Boekiproxy.Contract.Deposit_BOEKI(&_Boekiproxy.TransactOpts)
}

// Deposit_BOEKI is a paid mutator transaction binding the contract method 0x69027508.
//
// Solidity: function deposit_BOEKI() returns()
func (_Boekiproxy *BoekiproxyTransactorSession) Deposit_BOEKI() (*types.Transaction, error) {
	return _Boekiproxy.Contract.Deposit_BOEKI(&_Boekiproxy.TransactOpts)
}

// WithdrawToken_BOEKI is a paid mutator transaction binding the contract method 0x3c036b0c.
//
// Solidity: function withdrawToken_BOEKI(token address, amount uint256) returns()
func (_Boekiproxy *BoekiproxyTransactor) WithdrawToken_BOEKI(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.contract.Transact(opts, "withdrawToken_BOEKI", token, amount)
}

// WithdrawToken_BOEKI is a paid mutator transaction binding the contract method 0x3c036b0c.
//
// Solidity: function withdrawToken_BOEKI(token address, amount uint256) returns()
func (_Boekiproxy *BoekiproxySession) WithdrawToken_BOEKI(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.Contract.WithdrawToken_BOEKI(&_Boekiproxy.TransactOpts, token, amount)
}

// WithdrawToken_BOEKI is a paid mutator transaction binding the contract method 0x3c036b0c.
//
// Solidity: function withdrawToken_BOEKI(token address, amount uint256) returns()
func (_Boekiproxy *BoekiproxyTransactorSession) WithdrawToken_BOEKI(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.Contract.WithdrawToken_BOEKI(&_Boekiproxy.TransactOpts, token, amount)
}

// Withdraw_BOEKI is a paid mutator transaction binding the contract method 0xef30a6bc.
//
// Solidity: function withdraw_BOEKI(amount uint256) returns()
func (_Boekiproxy *BoekiproxyTransactor) Withdraw_BOEKI(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.contract.Transact(opts, "withdraw_BOEKI", amount)
}

// Withdraw_BOEKI is a paid mutator transaction binding the contract method 0xef30a6bc.
//
// Solidity: function withdraw_BOEKI(amount uint256) returns()
func (_Boekiproxy *BoekiproxySession) Withdraw_BOEKI(amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.Contract.Withdraw_BOEKI(&_Boekiproxy.TransactOpts, amount)
}

// Withdraw_BOEKI is a paid mutator transaction binding the contract method 0xef30a6bc.
//
// Solidity: function withdraw_BOEKI(amount uint256) returns()
func (_Boekiproxy *BoekiproxyTransactorSession) Withdraw_BOEKI(amount *big.Int) (*types.Transaction, error) {
	return _Boekiproxy.Contract.Withdraw_BOEKI(&_Boekiproxy.TransactOpts, amount)
}
