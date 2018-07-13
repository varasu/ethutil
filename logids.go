// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// LogidABI is the input ABI used to generate the binding from.
const LogidABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"deleteUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minEther\",\"type\":\"uint256\"}],\"name\":\"setMinEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"m\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ban\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minEther\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"NewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"DeleteUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"BanUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// LogidBin is the compiled bytecode used for deploying new contracts.
const LogidBin = `608060405234801561001057600080fd5b5060405160408061070f833981016040528051602090910151600160a060020a038116151561004c5760008054600160a060020a031916331790555b60008054600160a060020a03909216600160a060020a03199092169190911790556002556106908061007f6000396000f3006080604052600436106100a35763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663026ff05e81146100b5578063051af15d146100cc5780630e50cee5146100f35780631aa3a008146101145780633ccfd60b1461011c5780637aef1d4d146101315780637bdaa43d146101495780638da5cb5b146101c557806397c3ccd8146101f6578063a6f9dae114610217575b3480156100af57600080fd5b50600080fd5b3480156100c157600080fd5b506100ca610238565b005b3480156100d857600080fd5b506100e16102c4565b60408051918252519081900360200190f35b3480156100ff57600080fd5b506100e1600160a060020a03600435166102ca565b6100ca6102dc565b34801561012857600080fd5b506100ca610351565b34801561013d57600080fd5b506100ca6004356103ef565b34801561015557600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101b19436949293602493928401919081908401838280828437509497505050833560ff169450505060208201359160400135905061040b565b604080519115158252519081900360200190f35b3480156101d157600080fd5b506101da610573565b60408051600160a060020a039092168252519081900360200190f35b34801561020257600080fd5b506100ca600160a060020a0360043516610582565b34801561022357600080fd5b506100ca600160a060020a036004351661061e565b33600090815260016020526040812054151561025357600080fd5b5033600081815260016020526040808220805490839055905190929183156108fc02918491818181858888f19350505050158015610295573d6000803e3d6000fd5b5060405133907fa590e9c1b5854cc60d3ea7ecf49814ae8688f092b38adda95a86ad98c424401790600090a250565b60025481565b60016020526000908152604090205481565b6002543410156102eb57600080fd5b33600090815260016020526040812054111561030657600080fd5b336000818152600160209081526040918290203490819055825190815291517f82ff713a8585d7b6ad9cfe09df409e3b9a7fe6f712d4a85d5cdc57af76d02b499281900390910190a2565b336000908152600160205260408120548190151561036e57600080fd5b505033600081815260016020819052604080832080549083905590519093919284156108fc02918591818181858888f193505050501580156103b4573d6000803e3d6000fd5b5060408051838152905133917f8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918919081900360200190a25050565b600054600160a060020a0316331461040657600080fd5b600255565b60008060006002876040518082805190602001908083835b602083106104425780518252601f199092019160209182019101610423565b51815160209384036101000a600019018019909216911617905260405191909301945091925050808303816000865af1158015610483573d6000803e3d6000fd5b5050506040513d602081101561049857600080fd5b5051604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052815190819003603c018120600080835260208381018086529290925260ff8b1683850152606083018a905260808301899052925193955060019360a0808401949293601f19830193908390039091019190865af115801561052d573d6000803e3d6000fd5b505060408051601f190151600160a060020a0381166000908152600160205291822054909350111590506105645760019250610569565b600092505b5050949350505050565b600054600160a060020a031681565b60008054600160a060020a0316331461059a57600080fd5b50600160a060020a038116600081815260016020526040808220805490839055905190929183156108fc02918491818181858888f193505050501580156105e5573d6000803e3d6000fd5b50604051600160a060020a038316907fa0f5d8014fa8679b7c163b85bf4316c52253365ca6fed5740ebd5e3fdd7b767e90600090a25050565b600054600160a060020a0316331461063557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820d89e02cfe7b0dc0b77e995a212d844e0fea28e8513e3a72a7d6eb60474b1fc910029`

// DeployLogid deploys a new Ethereum contract, binding an instance of Logid to it.
func DeployLogid(auth *bind.TransactOpts, backend bind.ContractBackend, _minEther *big.Int, _owner common.Address) (common.Address, *types.Transaction, *Logid, error) {
	parsed, err := abi.JSON(strings.NewReader(LogidABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LogidBin), backend, _minEther, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logid{LogidCaller: LogidCaller{contract: contract}, LogidTransactor: LogidTransactor{contract: contract}, LogidFilterer: LogidFilterer{contract: contract}}, nil
}

// Logid is an auto generated Go binding around an Ethereum contract.
type Logid struct {
	LogidCaller     // Read-only binding to the contract
	LogidTransactor // Write-only binding to the contract
	LogidFilterer   // Log filterer for contract events
}

// LogidCaller is an auto generated read-only Go binding around an Ethereum contract.
type LogidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LogidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LogidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LogidSession struct {
	Contract     *Logid            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LogidCallerSession struct {
	Contract *LogidCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LogidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LogidTransactorSession struct {
	Contract     *LogidTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogidRaw is an auto generated low-level Go binding around an Ethereum contract.
type LogidRaw struct {
	Contract *Logid // Generic contract binding to access the raw methods on
}

// LogidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LogidCallerRaw struct {
	Contract *LogidCaller // Generic read-only contract binding to access the raw methods on
}

// LogidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LogidTransactorRaw struct {
	Contract *LogidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogid creates a new instance of Logid, bound to a specific deployed contract.
func NewLogid(address common.Address, backend bind.ContractBackend) (*Logid, error) {
	contract, err := bindLogid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logid{LogidCaller: LogidCaller{contract: contract}, LogidTransactor: LogidTransactor{contract: contract}, LogidFilterer: LogidFilterer{contract: contract}}, nil
}

// NewLogidCaller creates a new read-only instance of Logid, bound to a specific deployed contract.
func NewLogidCaller(address common.Address, caller bind.ContractCaller) (*LogidCaller, error) {
	contract, err := bindLogid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LogidCaller{contract: contract}, nil
}

// NewLogidTransactor creates a new write-only instance of Logid, bound to a specific deployed contract.
func NewLogidTransactor(address common.Address, transactor bind.ContractTransactor) (*LogidTransactor, error) {
	contract, err := bindLogid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LogidTransactor{contract: contract}, nil
}

// NewLogidFilterer creates a new log filterer instance of Logid, bound to a specific deployed contract.
func NewLogidFilterer(address common.Address, filterer bind.ContractFilterer) (*LogidFilterer, error) {
	contract, err := bindLogid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LogidFilterer{contract: contract}, nil
}

// bindLogid binds a generic wrapper to an already deployed contract.
func bindLogid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LogidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logid *LogidRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logid.Contract.LogidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logid *LogidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logid.Contract.LogidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logid *LogidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logid.Contract.LogidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logid *LogidCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Logid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logid *LogidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logid *LogidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logid.Contract.contract.Transact(opts, method, params...)
}

// IsAllowed is a free data retrieval call binding the contract method 0x7bdaa43d.
//
// Solidity: function isAllowed(m bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Logid *LogidCaller) IsAllowed(opts *bind.CallOpts, m []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Logid.contract.Call(opts, out, "isAllowed", m, v, r, s)
	return *ret0, err
}

// IsAllowed is a free data retrieval call binding the contract method 0x7bdaa43d.
//
// Solidity: function isAllowed(m bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Logid *LogidSession) IsAllowed(m []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Logid.Contract.IsAllowed(&_Logid.CallOpts, m, v, r, s)
}

// IsAllowed is a free data retrieval call binding the contract method 0x7bdaa43d.
//
// Solidity: function isAllowed(m bytes, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Logid *LogidCallerSession) IsAllowed(m []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Logid.Contract.IsAllowed(&_Logid.CallOpts, m, v, r, s)
}

// MinEther is a free data retrieval call binding the contract method 0x051af15d.
//
// Solidity: function minEther() constant returns(uint256)
func (_Logid *LogidCaller) MinEther(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Logid.contract.Call(opts, out, "minEther")
	return *ret0, err
}

// MinEther is a free data retrieval call binding the contract method 0x051af15d.
//
// Solidity: function minEther() constant returns(uint256)
func (_Logid *LogidSession) MinEther() (*big.Int, error) {
	return _Logid.Contract.MinEther(&_Logid.CallOpts)
}

// MinEther is a free data retrieval call binding the contract method 0x051af15d.
//
// Solidity: function minEther() constant returns(uint256)
func (_Logid *LogidCallerSession) MinEther() (*big.Int, error) {
	return _Logid.Contract.MinEther(&_Logid.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Logid *LogidCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Logid.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Logid *LogidSession) Owner() (common.Address, error) {
	return _Logid.Contract.Owner(&_Logid.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Logid *LogidCallerSession) Owner() (common.Address, error) {
	return _Logid.Contract.Owner(&_Logid.CallOpts)
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x0e50cee5.
//
// Solidity: function registeredUsers( address) constant returns(uint256)
func (_Logid *LogidCaller) RegisteredUsers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Logid.contract.Call(opts, out, "registeredUsers", arg0)
	return *ret0, err
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x0e50cee5.
//
// Solidity: function registeredUsers( address) constant returns(uint256)
func (_Logid *LogidSession) RegisteredUsers(arg0 common.Address) (*big.Int, error) {
	return _Logid.Contract.RegisteredUsers(&_Logid.CallOpts, arg0)
}

// RegisteredUsers is a free data retrieval call binding the contract method 0x0e50cee5.
//
// Solidity: function registeredUsers( address) constant returns(uint256)
func (_Logid *LogidCallerSession) RegisteredUsers(arg0 common.Address) (*big.Int, error) {
	return _Logid.Contract.RegisteredUsers(&_Logid.CallOpts, arg0)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(user address) returns()
func (_Logid *LogidTransactor) Ban(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "ban", user)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(user address) returns()
func (_Logid *LogidSession) Ban(user common.Address) (*types.Transaction, error) {
	return _Logid.Contract.Ban(&_Logid.TransactOpts, user)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(user address) returns()
func (_Logid *LogidTransactorSession) Ban(user common.Address) (*types.Transaction, error) {
	return _Logid.Contract.Ban(&_Logid.TransactOpts, user)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Logid *LogidTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Logid *LogidSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Logid.Contract.ChangeOwner(&_Logid.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(_newOwner address) returns()
func (_Logid *LogidTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Logid.Contract.ChangeOwner(&_Logid.TransactOpts, _newOwner)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x026ff05e.
//
// Solidity: function deleteUser() returns()
func (_Logid *LogidTransactor) DeleteUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "deleteUser")
}

// DeleteUser is a paid mutator transaction binding the contract method 0x026ff05e.
//
// Solidity: function deleteUser() returns()
func (_Logid *LogidSession) DeleteUser() (*types.Transaction, error) {
	return _Logid.Contract.DeleteUser(&_Logid.TransactOpts)
}

// DeleteUser is a paid mutator transaction binding the contract method 0x026ff05e.
//
// Solidity: function deleteUser() returns()
func (_Logid *LogidTransactorSession) DeleteUser() (*types.Transaction, error) {
	return _Logid.Contract.DeleteUser(&_Logid.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Logid *LogidTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Logid *LogidSession) Register() (*types.Transaction, error) {
	return _Logid.Contract.Register(&_Logid.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Logid *LogidTransactorSession) Register() (*types.Transaction, error) {
	return _Logid.Contract.Register(&_Logid.TransactOpts)
}

// SetMinEther is a paid mutator transaction binding the contract method 0x7aef1d4d.
//
// Solidity: function setMinEther(_minEther uint256) returns()
func (_Logid *LogidTransactor) SetMinEther(opts *bind.TransactOpts, _minEther *big.Int) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "setMinEther", _minEther)
}

// SetMinEther is a paid mutator transaction binding the contract method 0x7aef1d4d.
//
// Solidity: function setMinEther(_minEther uint256) returns()
func (_Logid *LogidSession) SetMinEther(_minEther *big.Int) (*types.Transaction, error) {
	return _Logid.Contract.SetMinEther(&_Logid.TransactOpts, _minEther)
}

// SetMinEther is a paid mutator transaction binding the contract method 0x7aef1d4d.
//
// Solidity: function setMinEther(_minEther uint256) returns()
func (_Logid *LogidTransactorSession) SetMinEther(_minEther *big.Int) (*types.Transaction, error) {
	return _Logid.Contract.SetMinEther(&_Logid.TransactOpts, _minEther)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Logid *LogidTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logid.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Logid *LogidSession) Withdraw() (*types.Transaction, error) {
	return _Logid.Contract.Withdraw(&_Logid.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Logid *LogidTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Logid.Contract.Withdraw(&_Logid.TransactOpts)
}

// LogidBanUserIterator is returned from FilterBanUser and is used to iterate over the raw logs and unpacked data for BanUser events raised by the Logid contract.
type LogidBanUserIterator struct {
	Event *LogidBanUser // Event containing the contract specifics and raw log

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
func (it *LogidBanUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogidBanUser)
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
		it.Event = new(LogidBanUser)
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
func (it *LogidBanUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogidBanUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogidBanUser represents a BanUser event raised by the Logid contract.
type LogidBanUser struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBanUser is a free log retrieval operation binding the contract event 0xa0f5d8014fa8679b7c163b85bf4316c52253365ca6fed5740ebd5e3fdd7b767e.
//
// Solidity: event BanUser(_user indexed address)
func (_Logid *LogidFilterer) FilterBanUser(opts *bind.FilterOpts, _user []common.Address) (*LogidBanUserIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.FilterLogs(opts, "BanUser", _userRule)
	if err != nil {
		return nil, err
	}
	return &LogidBanUserIterator{contract: _Logid.contract, event: "BanUser", logs: logs, sub: sub}, nil
}

// WatchBanUser is a free log subscription operation binding the contract event 0xa0f5d8014fa8679b7c163b85bf4316c52253365ca6fed5740ebd5e3fdd7b767e.
//
// Solidity: event BanUser(_user indexed address)
func (_Logid *LogidFilterer) WatchBanUser(opts *bind.WatchOpts, sink chan<- *LogidBanUser, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.WatchLogs(opts, "BanUser", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogidBanUser)
				if err := _Logid.contract.UnpackLog(event, "BanUser", log); err != nil {
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

// LogidDeleteUserIterator is returned from FilterDeleteUser and is used to iterate over the raw logs and unpacked data for DeleteUser events raised by the Logid contract.
type LogidDeleteUserIterator struct {
	Event *LogidDeleteUser // Event containing the contract specifics and raw log

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
func (it *LogidDeleteUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogidDeleteUser)
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
		it.Event = new(LogidDeleteUser)
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
func (it *LogidDeleteUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogidDeleteUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogidDeleteUser represents a DeleteUser event raised by the Logid contract.
type LogidDeleteUser struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteUser is a free log retrieval operation binding the contract event 0xa590e9c1b5854cc60d3ea7ecf49814ae8688f092b38adda95a86ad98c4244017.
//
// Solidity: event DeleteUser(_user indexed address)
func (_Logid *LogidFilterer) FilterDeleteUser(opts *bind.FilterOpts, _user []common.Address) (*LogidDeleteUserIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.FilterLogs(opts, "DeleteUser", _userRule)
	if err != nil {
		return nil, err
	}
	return &LogidDeleteUserIterator{contract: _Logid.contract, event: "DeleteUser", logs: logs, sub: sub}, nil
}

// WatchDeleteUser is a free log subscription operation binding the contract event 0xa590e9c1b5854cc60d3ea7ecf49814ae8688f092b38adda95a86ad98c4244017.
//
// Solidity: event DeleteUser(_user indexed address)
func (_Logid *LogidFilterer) WatchDeleteUser(opts *bind.WatchOpts, sink chan<- *LogidDeleteUser, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.WatchLogs(opts, "DeleteUser", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogidDeleteUser)
				if err := _Logid.contract.UnpackLog(event, "DeleteUser", log); err != nil {
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

// LogidNewUserIterator is returned from FilterNewUser and is used to iterate over the raw logs and unpacked data for NewUser events raised by the Logid contract.
type LogidNewUserIterator struct {
	Event *LogidNewUser // Event containing the contract specifics and raw log

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
func (it *LogidNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogidNewUser)
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
		it.Event = new(LogidNewUser)
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
func (it *LogidNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogidNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogidNewUser represents a NewUser event raised by the Logid contract.
type LogidNewUser struct {
	Value *big.Int
	User  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewUser is a free log retrieval operation binding the contract event 0x82ff713a8585d7b6ad9cfe09df409e3b9a7fe6f712d4a85d5cdc57af76d02b49.
//
// Solidity: event NewUser(_value uint256, _user indexed address)
func (_Logid *LogidFilterer) FilterNewUser(opts *bind.FilterOpts, _user []common.Address) (*LogidNewUserIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.FilterLogs(opts, "NewUser", _userRule)
	if err != nil {
		return nil, err
	}
	return &LogidNewUserIterator{contract: _Logid.contract, event: "NewUser", logs: logs, sub: sub}, nil
}

// WatchNewUser is a free log subscription operation binding the contract event 0x82ff713a8585d7b6ad9cfe09df409e3b9a7fe6f712d4a85d5cdc57af76d02b49.
//
// Solidity: event NewUser(_value uint256, _user indexed address)
func (_Logid *LogidFilterer) WatchNewUser(opts *bind.WatchOpts, sink chan<- *LogidNewUser, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.WatchLogs(opts, "NewUser", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogidNewUser)
				if err := _Logid.contract.UnpackLog(event, "NewUser", log); err != nil {
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

// LogidWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Logid contract.
type LogidWithdrawIterator struct {
	Event *LogidWithdraw // Event containing the contract specifics and raw log

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
func (it *LogidWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LogidWithdraw)
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
		it.Event = new(LogidWithdraw)
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
func (it *LogidWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LogidWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LogidWithdraw represents a Withdraw event raised by the Logid contract.
type LogidWithdraw struct {
	Value *big.Int
	User  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(_value uint256, _user indexed address)
func (_Logid *LogidFilterer) FilterWithdraw(opts *bind.FilterOpts, _user []common.Address) (*LogidWithdrawIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.FilterLogs(opts, "Withdraw", _userRule)
	if err != nil {
		return nil, err
	}
	return &LogidWithdrawIterator{contract: _Logid.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8353ffcac0876ad14e226d9783c04540bfebf13871e868157d2a391cad98e918.
//
// Solidity: event Withdraw(_value uint256, _user indexed address)
func (_Logid *LogidFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LogidWithdraw, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Logid.contract.WatchLogs(opts, "Withdraw", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LogidWithdraw)
				if err := _Logid.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
