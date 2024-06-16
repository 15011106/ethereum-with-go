// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chapter2

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

// Chapter2MetaData contains all meta data concerning the Chapter2 contract.
var Chapter2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea264697066735822122093ac87ddeabe43fb56b5755bc230e2cf05db2461367907c65f49c4405fc2b17464736f6c63430008130033",
}

// Chapter2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Chapter2MetaData.ABI instead.
var Chapter2ABI = Chapter2MetaData.ABI

// Chapter2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Chapter2MetaData.Bin instead.
var Chapter2Bin = Chapter2MetaData.Bin

// DeployChapter2 deploys a new Ethereum contract, binding an instance of Chapter2 to it.
func DeployChapter2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Chapter2, error) {
	parsed, err := Chapter2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Chapter2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Chapter2{Chapter2Caller: Chapter2Caller{contract: contract}, Chapter2Transactor: Chapter2Transactor{contract: contract}, Chapter2Filterer: Chapter2Filterer{contract: contract}}, nil
}

// Chapter2 is an auto generated Go binding around an Ethereum contract.
type Chapter2 struct {
	Chapter2Caller     // Read-only binding to the contract
	Chapter2Transactor // Write-only binding to the contract
	Chapter2Filterer   // Log filterer for contract events
}

// Chapter2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Chapter2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chapter2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Chapter2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chapter2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Chapter2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Chapter2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Chapter2Session struct {
	Contract     *Chapter2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Chapter2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Chapter2CallerSession struct {
	Contract *Chapter2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Chapter2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Chapter2TransactorSession struct {
	Contract     *Chapter2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Chapter2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Chapter2Raw struct {
	Contract *Chapter2 // Generic contract binding to access the raw methods on
}

// Chapter2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Chapter2CallerRaw struct {
	Contract *Chapter2Caller // Generic read-only contract binding to access the raw methods on
}

// Chapter2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Chapter2TransactorRaw struct {
	Contract *Chapter2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewChapter2 creates a new instance of Chapter2, bound to a specific deployed contract.
func NewChapter2(address common.Address, backend bind.ContractBackend) (*Chapter2, error) {
	contract, err := bindChapter2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chapter2{Chapter2Caller: Chapter2Caller{contract: contract}, Chapter2Transactor: Chapter2Transactor{contract: contract}, Chapter2Filterer: Chapter2Filterer{contract: contract}}, nil
}

// NewChapter2Caller creates a new read-only instance of Chapter2, bound to a specific deployed contract.
func NewChapter2Caller(address common.Address, caller bind.ContractCaller) (*Chapter2Caller, error) {
	contract, err := bindChapter2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Chapter2Caller{contract: contract}, nil
}

// NewChapter2Transactor creates a new write-only instance of Chapter2, bound to a specific deployed contract.
func NewChapter2Transactor(address common.Address, transactor bind.ContractTransactor) (*Chapter2Transactor, error) {
	contract, err := bindChapter2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Chapter2Transactor{contract: contract}, nil
}

// NewChapter2Filterer creates a new log filterer instance of Chapter2, bound to a specific deployed contract.
func NewChapter2Filterer(address common.Address, filterer bind.ContractFilterer) (*Chapter2Filterer, error) {
	contract, err := bindChapter2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Chapter2Filterer{contract: contract}, nil
}

// bindChapter2 binds a generic wrapper to an already deployed contract.
func bindChapter2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Chapter2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chapter2 *Chapter2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chapter2.Contract.Chapter2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chapter2 *Chapter2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chapter2.Contract.Chapter2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chapter2 *Chapter2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chapter2.Contract.Chapter2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chapter2 *Chapter2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chapter2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chapter2 *Chapter2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chapter2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chapter2 *Chapter2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chapter2.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Chapter2 *Chapter2Caller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chapter2.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Chapter2 *Chapter2Session) Retrieve() (*big.Int, error) {
	return _Chapter2.Contract.Retrieve(&_Chapter2.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Chapter2 *Chapter2CallerSession) Retrieve() (*big.Int, error) {
	return _Chapter2.Contract.Retrieve(&_Chapter2.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Chapter2 *Chapter2Transactor) Store(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Chapter2.contract.Transact(opts, "store", number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Chapter2 *Chapter2Session) Store(number *big.Int) (*types.Transaction, error) {
	return _Chapter2.Contract.Store(&_Chapter2.TransactOpts, number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Chapter2 *Chapter2TransactorSession) Store(number *big.Int) (*types.Transaction, error) {
	return _Chapter2.Contract.Store(&_Chapter2.TransactOpts, number)
}
