// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sweepaction

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

// IOtimFeeFee is an auto generated low-level Go binding around an user-defined struct.
type IOtimFeeFee struct {
	Token                common.Address
	MaxBaseFeePerGas     *big.Int
	MaxPriorityFeePerGas *big.Int
	ExecutionFee         *big.Int
}

// ISweepActionSweep is an auto generated low-level Go binding around an user-defined struct.
type ISweepActionSweep struct {
	Target     common.Address
	Threshold  *big.Int
	EndBalance *big.Int
	GasLimit   *big.Int
	Fee        IOtimFeeFee
}

// InstructionLibExecutionState is an auto generated low-level Go binding around an user-defined struct.
type InstructionLibExecutionState struct {
	Deactivated    bool
	ExecutionCount *big.Int
	LastExecuted   *big.Int
}

// InstructionLibInstruction is an auto generated low-level Go binding around an user-defined struct.
type InstructionLibInstruction struct {
	Salt          *big.Int
	MaxExecutions *big.Int
	Action        common.Address
	Arguments     []byte
}

// InstructionLibSignature is an auto generated low-level Go binding around an user-defined struct.
type InstructionLibSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SweepactionMetaData contains all meta data concerning the Sweepaction contract.
var SweepactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"deactivate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structISweepAction.Sweep\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SweepActionFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceUnderThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// SweepactionABI is the input ABI used to generate the binding from.
// Deprecated: Use SweepactionMetaData.ABI instead.
var SweepactionABI = SweepactionMetaData.ABI

// Sweepaction is an auto generated Go binding around an Ethereum contract.
type Sweepaction struct {
	SweepactionCaller     // Read-only binding to the contract
	SweepactionTransactor // Write-only binding to the contract
	SweepactionFilterer   // Log filterer for contract events
}

// SweepactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweepactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweepactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SweepactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweepactionSession struct {
	Contract     *Sweepaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweepactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweepactionCallerSession struct {
	Contract *SweepactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SweepactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweepactionTransactorSession struct {
	Contract     *SweepactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SweepactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweepactionRaw struct {
	Contract *Sweepaction // Generic contract binding to access the raw methods on
}

// SweepactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweepactionCallerRaw struct {
	Contract *SweepactionCaller // Generic read-only contract binding to access the raw methods on
}

// SweepactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweepactionTransactorRaw struct {
	Contract *SweepactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweepaction creates a new instance of Sweepaction, bound to a specific deployed contract.
func NewSweepaction(address common.Address, backend bind.ContractBackend) (*Sweepaction, error) {
	contract, err := bindSweepaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sweepaction{SweepactionCaller: SweepactionCaller{contract: contract}, SweepactionTransactor: SweepactionTransactor{contract: contract}, SweepactionFilterer: SweepactionFilterer{contract: contract}}, nil
}

// NewSweepactionCaller creates a new read-only instance of Sweepaction, bound to a specific deployed contract.
func NewSweepactionCaller(address common.Address, caller bind.ContractCaller) (*SweepactionCaller, error) {
	contract, err := bindSweepaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SweepactionCaller{contract: contract}, nil
}

// NewSweepactionTransactor creates a new write-only instance of Sweepaction, bound to a specific deployed contract.
func NewSweepactionTransactor(address common.Address, transactor bind.ContractTransactor) (*SweepactionTransactor, error) {
	contract, err := bindSweepaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SweepactionTransactor{contract: contract}, nil
}

// NewSweepactionFilterer creates a new log filterer instance of Sweepaction, bound to a specific deployed contract.
func NewSweepactionFilterer(address common.Address, filterer bind.ContractFilterer) (*SweepactionFilterer, error) {
	contract, err := bindSweepaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SweepactionFilterer{contract: contract}, nil
}

// bindSweepaction binds a generic wrapper to an already deployed contract.
func bindSweepaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SweepactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepaction *SweepactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepaction.Contract.SweepactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepaction *SweepactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepaction.Contract.SweepactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepaction *SweepactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepaction.Contract.SweepactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepaction *SweepactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepaction *SweepactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepaction *SweepactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepaction *SweepactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "argumentsHash", arguments)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepaction *SweepactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepaction.Contract.ArgumentsHash(&_Sweepaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepaction *SweepactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepaction.Contract.ArgumentsHash(&_Sweepaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepaction *SweepactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepaction *SweepactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepaction.Contract.FeeTokenRegistry(&_Sweepaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepaction *SweepactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepaction.Contract.FeeTokenRegistry(&_Sweepaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepaction *SweepactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepaction *SweepactionSession) GasConstant() (*big.Int, error) {
	return _Sweepaction.Contract.GasConstant(&_Sweepaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepaction *SweepactionCallerSession) GasConstant() (*big.Int, error) {
	return _Sweepaction.Contract.GasConstant(&_Sweepaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepaction *SweepactionCaller) Hash(opts *bind.CallOpts, arguments ISweepActionSweep) ([32]byte, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepaction *SweepactionSession) Hash(arguments ISweepActionSweep) ([32]byte, error) {
	return _Sweepaction.Contract.Hash(&_Sweepaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepaction *SweepactionCallerSession) Hash(arguments ISweepActionSweep) ([32]byte, error) {
	return _Sweepaction.Contract.Hash(&_Sweepaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepaction *SweepactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepaction *SweepactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepaction.Contract.Hash0(&_Sweepaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepaction *SweepactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepaction.Contract.Hash0(&_Sweepaction.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepaction *SweepactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepaction *SweepactionSession) Treasury() (common.Address, error) {
	return _Sweepaction.Contract.Treasury(&_Sweepaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepaction *SweepactionCallerSession) Treasury() (common.Address, error) {
	return _Sweepaction.Contract.Treasury(&_Sweepaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepaction *SweepactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepaction *SweepactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepaction.Contract.ChargeFee(&_Sweepaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepaction *SweepactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepaction.Contract.ChargeFee(&_Sweepaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweepaction *SweepactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweepaction *SweepactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepaction.Contract.Execute(&_Sweepaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweepaction *SweepactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepaction.Contract.Execute(&_Sweepaction.TransactOpts, instruction, arg1, executionState)
}

// SweepactionSweepActionFailedIterator is returned from FilterSweepActionFailed and is used to iterate over the raw logs and unpacked data for SweepActionFailed events raised by the Sweepaction contract.
type SweepactionSweepActionFailedIterator struct {
	Event *SweepactionSweepActionFailed // Event containing the contract specifics and raw log

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
func (it *SweepactionSweepActionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SweepactionSweepActionFailed)
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
		it.Event = new(SweepactionSweepActionFailed)
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
func (it *SweepactionSweepActionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SweepactionSweepActionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SweepactionSweepActionFailed represents a SweepActionFailed event raised by the Sweepaction contract.
type SweepactionSweepActionFailed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSweepActionFailed is a free log retrieval operation binding the contract event 0x2cdc3d1b2b051d61bfda6d41e23436aca4690fef7d88be13f3475fef5cc6b01b.
//
// Solidity: event SweepActionFailed(address indexed target)
func (_Sweepaction *SweepactionFilterer) FilterSweepActionFailed(opts *bind.FilterOpts, target []common.Address) (*SweepactionSweepActionFailedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Sweepaction.contract.FilterLogs(opts, "SweepActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return &SweepactionSweepActionFailedIterator{contract: _Sweepaction.contract, event: "SweepActionFailed", logs: logs, sub: sub}, nil
}

// WatchSweepActionFailed is a free log subscription operation binding the contract event 0x2cdc3d1b2b051d61bfda6d41e23436aca4690fef7d88be13f3475fef5cc6b01b.
//
// Solidity: event SweepActionFailed(address indexed target)
func (_Sweepaction *SweepactionFilterer) WatchSweepActionFailed(opts *bind.WatchOpts, sink chan<- *SweepactionSweepActionFailed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Sweepaction.contract.WatchLogs(opts, "SweepActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SweepactionSweepActionFailed)
				if err := _Sweepaction.contract.UnpackLog(event, "SweepActionFailed", log); err != nil {
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

// ParseSweepActionFailed is a log parse operation binding the contract event 0x2cdc3d1b2b051d61bfda6d41e23436aca4690fef7d88be13f3475fef5cc6b01b.
//
// Solidity: event SweepActionFailed(address indexed target)
func (_Sweepaction *SweepactionFilterer) ParseSweepActionFailed(log types.Log) (*SweepactionSweepActionFailed, error) {
	event := new(SweepactionSweepActionFailed)
	if err := _Sweepaction.contract.UnpackLog(event, "SweepActionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
