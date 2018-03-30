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

// GreetABI is the input ABI used to generate the binding from.
const GreetABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"digit\",\"type\":\"uint256\"},{\"name\":\"other\",\"type\":\"address\"}],\"name\":\"complexHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"digit\",\"type\":\"uint256\"}],\"name\":\"simpleHashDigit\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"deltaContract\",\"type\":\"address\"},{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"checkSignatureKeccak\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"deltaContract\",\"type\":\"address\"},{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"checkSignatureSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"other\",\"type\":\"address\"}],\"name\":\"simpleHashAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"deltaContract\",\"type\":\"address\"},{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"hashTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes32\"}],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// GreetBin is the compiled bytecode used for deploying new contracts.
const GreetBin = `6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556105288061003b6000396000f3006060604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166341bb2177811461009d57806341c0e1b5146100d15780635d220de0146100e657806370d39ef1146100fc5780638244f83a146100fc578063a025502b1461015e578063e8be747b1461017d578063e9ebeafe146101b5578063f851a440146101cb575b600080fd5b34156100a857600080fd5b6100bf600435600160a060020a03602435166101de565b60405190815260200160405180910390f35b34156100dc57600080fd5b6100e461025b565b005b34156100f157600080fd5b6100bf600435610282565b341561010757600080fd5b610142600160a060020a03600435811690602435811690604435906064351660843560a43560c43560ff60e4351661010435610124356102ba565b604051600160a060020a03909116815260200160405180910390f35b341561016957600080fd5b6100bf600160a060020a03600435166103ea565b341561018857600080fd5b6100bf600160a060020a03600435811690602435811690604435906064351660843560a43560c43561042b565b34156101c057600080fd5b6100bf6004356104b9565b34156101d657600080fd5b6101426104ed565b6000805481906002903090600160a060020a031686866040516c01000000000000000000000000600160a060020a0395861681028252938516840260148201526028810192909252909216026048820152605c016020604051808303816000865af1151561024b57600080fd5b5050604051805195945050505050565b60005433600160a060020a039081169116141561028057600054600160a060020a0316ff5b565b6000600282604051908152602090810190604051808303816000865af115156102aa57600080fd5b505060405180519150505b919050565b60008060028c8c8c8c8c8c8c6040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc016020604051808303816000865af1151561033557600080fd5b50506040518051905090506001816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0160405180910390208686866040516000815260200160405260405193845260ff9092166020808501919091526040808501929092526060840192909252608090920191516020810390808403906000865af115156103d157600080fd5b5050602060405103519c9b505050505050505050505050565b6000600282604051600160a060020a03919091166c010000000000000000000000000281526014016020604051808303816000865af115156102aa57600080fd5b60006002888888888888886040516c01000000000000000000000000600160a060020a0398891681028252968816870260148201526028810195909552929095169093026048830152605c820192909252607c810192909252609c82015260bc016020604051808303816000865af115156104a557600080fd5b505060405180519998505050505050505050565b60008115156104e957507f48656c6c6f2c20576f726c6400000000000000000000000000000000000000006102b5565b5090565b600054600160a060020a0316815600a165627a7a7230582067ea92ec322ab4a590701f4a46c6956f2b5697c845fdd1beb4ac28c876bdc6840029`

// DeployGreet deploys a new Ethereum contract, binding an instance of Greet to it.
func DeployGreet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Greet, error) {
	parsed, err := abi.JSON(strings.NewReader(GreetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GreetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Greet{GreetCaller: GreetCaller{contract: contract}, GreetTransactor: GreetTransactor{contract: contract}, GreetFilterer: GreetFilterer{contract: contract}}, nil
}

// Greet is an auto generated Go binding around an Ethereum contract.
type Greet struct {
	GreetCaller     // Read-only binding to the contract
	GreetTransactor // Write-only binding to the contract
	GreetFilterer   // Log filterer for contract events
}

// GreetCaller is an auto generated read-only Go binding around an Ethereum contract.
type GreetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GreetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GreetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GreetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GreetSession struct {
	Contract     *Greet            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GreetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GreetCallerSession struct {
	Contract *GreetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GreetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GreetTransactorSession struct {
	Contract     *GreetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GreetRaw is an auto generated low-level Go binding around an Ethereum contract.
type GreetRaw struct {
	Contract *Greet // Generic contract binding to access the raw methods on
}

// GreetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GreetCallerRaw struct {
	Contract *GreetCaller // Generic read-only contract binding to access the raw methods on
}

// GreetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GreetTransactorRaw struct {
	Contract *GreetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGreet creates a new instance of Greet, bound to a specific deployed contract.
func NewGreet(address common.Address, backend bind.ContractBackend) (*Greet, error) {
	contract, err := bindGreet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Greet{GreetCaller: GreetCaller{contract: contract}, GreetTransactor: GreetTransactor{contract: contract}, GreetFilterer: GreetFilterer{contract: contract}}, nil
}

// NewGreetCaller creates a new read-only instance of Greet, bound to a specific deployed contract.
func NewGreetCaller(address common.Address, caller bind.ContractCaller) (*GreetCaller, error) {
	contract, err := bindGreet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GreetCaller{contract: contract}, nil
}

// NewGreetTransactor creates a new write-only instance of Greet, bound to a specific deployed contract.
func NewGreetTransactor(address common.Address, transactor bind.ContractTransactor) (*GreetTransactor, error) {
	contract, err := bindGreet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GreetTransactor{contract: contract}, nil
}

// NewGreetFilterer creates a new log filterer instance of Greet, bound to a specific deployed contract.
func NewGreetFilterer(address common.Address, filterer bind.ContractFilterer) (*GreetFilterer, error) {
	contract, err := bindGreet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GreetFilterer{contract: contract}, nil
}

// bindGreet binds a generic wrapper to an already deployed contract.
func bindGreet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GreetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greet *GreetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Greet.Contract.GreetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greet *GreetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greet.Contract.GreetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greet *GreetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greet.Contract.GreetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Greet *GreetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Greet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Greet *GreetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Greet *GreetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Greet.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Greet *GreetCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Greet *GreetSession) Admin() (common.Address, error) {
	return _Greet.Contract.Admin(&_Greet.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Greet *GreetCallerSession) Admin() (common.Address, error) {
	return _Greet.Contract.Admin(&_Greet.CallOpts)
}

// CheckSignatureKeccak is a free data retrieval call binding the contract method 0x70d39ef1.
//
// Solidity: function checkSignatureKeccak(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetCaller) CheckSignatureKeccak(opts *bind.CallOpts, deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "checkSignatureKeccak", deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
	return *ret0, err
}

// CheckSignatureKeccak is a free data retrieval call binding the contract method 0x70d39ef1.
//
// Solidity: function checkSignatureKeccak(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetSession) CheckSignatureKeccak(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Greet.Contract.CheckSignatureKeccak(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CheckSignatureKeccak is a free data retrieval call binding the contract method 0x70d39ef1.
//
// Solidity: function checkSignatureKeccak(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetCallerSession) CheckSignatureKeccak(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Greet.Contract.CheckSignatureKeccak(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CheckSignatureSha3 is a free data retrieval call binding the contract method 0x8244f83a.
//
// Solidity: function checkSignatureSha3(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetCaller) CheckSignatureSha3(opts *bind.CallOpts, deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "checkSignatureSha3", deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
	return *ret0, err
}

// CheckSignatureSha3 is a free data retrieval call binding the contract method 0x8244f83a.
//
// Solidity: function checkSignatureSha3(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetSession) CheckSignatureSha3(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Greet.Contract.CheckSignatureSha3(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CheckSignatureSha3 is a free data retrieval call binding the contract method 0x8244f83a.
//
// Solidity: function checkSignatureSha3(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) constant returns(address)
func (_Greet *GreetCallerSession) CheckSignatureSha3(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Greet.Contract.CheckSignatureSha3(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// ComplexHash is a free data retrieval call binding the contract method 0x41bb2177.
//
// Solidity: function complexHash(digit uint256, other address) constant returns(bytes32)
func (_Greet *GreetCaller) ComplexHash(opts *bind.CallOpts, digit *big.Int, other common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "complexHash", digit, other)
	return *ret0, err
}

// ComplexHash is a free data retrieval call binding the contract method 0x41bb2177.
//
// Solidity: function complexHash(digit uint256, other address) constant returns(bytes32)
func (_Greet *GreetSession) ComplexHash(digit *big.Int, other common.Address) ([32]byte, error) {
	return _Greet.Contract.ComplexHash(&_Greet.CallOpts, digit, other)
}

// ComplexHash is a free data retrieval call binding the contract method 0x41bb2177.
//
// Solidity: function complexHash(digit uint256, other address) constant returns(bytes32)
func (_Greet *GreetCallerSession) ComplexHash(digit *big.Int, other common.Address) ([32]byte, error) {
	return _Greet.Contract.ComplexHash(&_Greet.CallOpts, digit, other)
}

// Greet is a free data retrieval call binding the contract method 0xe9ebeafe.
//
// Solidity: function greet(input bytes32) constant returns(bytes32)
func (_Greet *GreetCaller) Greet(opts *bind.CallOpts, input [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "greet", input)
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xe9ebeafe.
//
// Solidity: function greet(input bytes32) constant returns(bytes32)
func (_Greet *GreetSession) Greet(input [32]byte) ([32]byte, error) {
	return _Greet.Contract.Greet(&_Greet.CallOpts, input)
}

// Greet is a free data retrieval call binding the contract method 0xe9ebeafe.
//
// Solidity: function greet(input bytes32) constant returns(bytes32)
func (_Greet *GreetCallerSession) Greet(input [32]byte) ([32]byte, error) {
	return _Greet.Contract.Greet(&_Greet.CallOpts, input)
}

// HashTrade is a free data retrieval call binding the contract method 0xe8be747b.
//
// Solidity: function hashTrade(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) constant returns(bytes32)
func (_Greet *GreetCaller) HashTrade(opts *bind.CallOpts, deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "hashTrade", deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
	return *ret0, err
}

// HashTrade is a free data retrieval call binding the contract method 0xe8be747b.
//
// Solidity: function hashTrade(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) constant returns(bytes32)
func (_Greet *GreetSession) HashTrade(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Greet.Contract.HashTrade(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// HashTrade is a free data retrieval call binding the contract method 0xe8be747b.
//
// Solidity: function hashTrade(deltaContract address, tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) constant returns(bytes32)
func (_Greet *GreetCallerSession) HashTrade(deltaContract common.Address, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) ([32]byte, error) {
	return _Greet.Contract.HashTrade(&_Greet.CallOpts, deltaContract, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// SimpleHashAddress is a free data retrieval call binding the contract method 0xa025502b.
//
// Solidity: function simpleHashAddress(other address) constant returns(bytes32)
func (_Greet *GreetCaller) SimpleHashAddress(opts *bind.CallOpts, other common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "simpleHashAddress", other)
	return *ret0, err
}

// SimpleHashAddress is a free data retrieval call binding the contract method 0xa025502b.
//
// Solidity: function simpleHashAddress(other address) constant returns(bytes32)
func (_Greet *GreetSession) SimpleHashAddress(other common.Address) ([32]byte, error) {
	return _Greet.Contract.SimpleHashAddress(&_Greet.CallOpts, other)
}

// SimpleHashAddress is a free data retrieval call binding the contract method 0xa025502b.
//
// Solidity: function simpleHashAddress(other address) constant returns(bytes32)
func (_Greet *GreetCallerSession) SimpleHashAddress(other common.Address) ([32]byte, error) {
	return _Greet.Contract.SimpleHashAddress(&_Greet.CallOpts, other)
}

// SimpleHashDigit is a free data retrieval call binding the contract method 0x5d220de0.
//
// Solidity: function simpleHashDigit(digit uint256) constant returns(bytes32)
func (_Greet *GreetCaller) SimpleHashDigit(opts *bind.CallOpts, digit *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Greet.contract.Call(opts, out, "simpleHashDigit", digit)
	return *ret0, err
}

// SimpleHashDigit is a free data retrieval call binding the contract method 0x5d220de0.
//
// Solidity: function simpleHashDigit(digit uint256) constant returns(bytes32)
func (_Greet *GreetSession) SimpleHashDigit(digit *big.Int) ([32]byte, error) {
	return _Greet.Contract.SimpleHashDigit(&_Greet.CallOpts, digit)
}

// SimpleHashDigit is a free data retrieval call binding the contract method 0x5d220de0.
//
// Solidity: function simpleHashDigit(digit uint256) constant returns(bytes32)
func (_Greet *GreetCallerSession) SimpleHashDigit(digit *big.Int) ([32]byte, error) {
	return _Greet.Contract.SimpleHashDigit(&_Greet.CallOpts, digit)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greet *GreetTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Greet.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greet *GreetSession) Kill() (*types.Transaction, error) {
	return _Greet.Contract.Kill(&_Greet.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Greet *GreetTransactorSession) Kill() (*types.Transaction, error) {
	return _Greet.Contract.Kill(&_Greet.TransactOpts)
}
