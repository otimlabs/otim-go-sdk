// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package refuelaction

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

// IRefuelActionRefuel is an auto generated low-level Go binding around an user-defined struct.
type IRefuelActionRefuel struct {
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

// RefuelactionMetaData contains all meta data concerning the Refuelaction contract.
var RefuelactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"deactivate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"refuel\",\"type\":\"tuple\",\"internalType\":\"structIRefuelAction.Refuel\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RefuelActionFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceOverThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// RefuelactionABI is the input ABI used to generate the binding from.
// Deprecated: Use RefuelactionMetaData.ABI instead.
var RefuelactionABI = RefuelactionMetaData.ABI

// Refuelaction is an auto generated Go binding around an Ethereum contract.
type Refuelaction struct {
	RefuelactionCaller     // Read-only binding to the contract
	RefuelactionTransactor // Write-only binding to the contract
	RefuelactionFilterer   // Log filterer for contract events
}

// RefuelactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type RefuelactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RefuelactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RefuelactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RefuelactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RefuelactionSession struct {
	Contract     *Refuelaction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RefuelactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RefuelactionCallerSession struct {
	Contract *RefuelactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RefuelactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RefuelactionTransactorSession struct {
	Contract     *RefuelactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RefuelactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type RefuelactionRaw struct {
	Contract *Refuelaction // Generic contract binding to access the raw methods on
}

// RefuelactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RefuelactionCallerRaw struct {
	Contract *RefuelactionCaller // Generic read-only contract binding to access the raw methods on
}

// RefuelactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RefuelactionTransactorRaw struct {
	Contract *RefuelactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRefuelaction creates a new instance of Refuelaction, bound to a specific deployed contract.
func NewRefuelaction(address common.Address, backend bind.ContractBackend) (*Refuelaction, error) {
	contract, err := bindRefuelaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Refuelaction{RefuelactionCaller: RefuelactionCaller{contract: contract}, RefuelactionTransactor: RefuelactionTransactor{contract: contract}, RefuelactionFilterer: RefuelactionFilterer{contract: contract}}, nil
}

// NewRefuelactionCaller creates a new read-only instance of Refuelaction, bound to a specific deployed contract.
func NewRefuelactionCaller(address common.Address, caller bind.ContractCaller) (*RefuelactionCaller, error) {
	contract, err := bindRefuelaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RefuelactionCaller{contract: contract}, nil
}

// NewRefuelactionTransactor creates a new write-only instance of Refuelaction, bound to a specific deployed contract.
func NewRefuelactionTransactor(address common.Address, transactor bind.ContractTransactor) (*RefuelactionTransactor, error) {
	contract, err := bindRefuelaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RefuelactionTransactor{contract: contract}, nil
}

// NewRefuelactionFilterer creates a new log filterer instance of Refuelaction, bound to a specific deployed contract.
func NewRefuelactionFilterer(address common.Address, filterer bind.ContractFilterer) (*RefuelactionFilterer, error) {
	contract, err := bindRefuelaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RefuelactionFilterer{contract: contract}, nil
}

// bindRefuelaction binds a generic wrapper to an already deployed contract.
func bindRefuelaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RefuelactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Refuelaction *RefuelactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Refuelaction.Contract.RefuelactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Refuelaction *RefuelactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Refuelaction.Contract.RefuelactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Refuelaction *RefuelactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Refuelaction.Contract.RefuelactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Refuelaction *RefuelactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Refuelaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Refuelaction *RefuelactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Refuelaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Refuelaction *RefuelactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Refuelaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Refuelaction *RefuelactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Refuelaction *RefuelactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Refuelaction.Contract.ArgumentsHash(&_Refuelaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Refuelaction *RefuelactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Refuelaction.Contract.ArgumentsHash(&_Refuelaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Refuelaction *RefuelactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Refuelaction *RefuelactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Refuelaction.Contract.FeeTokenRegistry(&_Refuelaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Refuelaction *RefuelactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Refuelaction.Contract.FeeTokenRegistry(&_Refuelaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Refuelaction *RefuelactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Refuelaction *RefuelactionSession) GasConstant() (*big.Int, error) {
	return _Refuelaction.Contract.GasConstant(&_Refuelaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Refuelaction *RefuelactionCallerSession) GasConstant() (*big.Int, error) {
	return _Refuelaction.Contract.GasConstant(&_Refuelaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) refuel) pure returns(bytes32)
func (_Refuelaction *RefuelactionCaller) Hash(opts *bind.CallOpts, refuel IRefuelActionRefuel) ([32]byte, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "hash", refuel)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) refuel) pure returns(bytes32)
func (_Refuelaction *RefuelactionSession) Hash(refuel IRefuelActionRefuel) ([32]byte, error) {
	return _Refuelaction.Contract.Hash(&_Refuelaction.CallOpts, refuel)
}

// Hash is a free data retrieval call binding the contract method 0x745835fb.
//
// Solidity: function hash((address,uint256,uint256,uint256,(address,uint256,uint256,uint256)) refuel) pure returns(bytes32)
func (_Refuelaction *RefuelactionCallerSession) Hash(refuel IRefuelActionRefuel) ([32]byte, error) {
	return _Refuelaction.Contract.Hash(&_Refuelaction.CallOpts, refuel)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Refuelaction *RefuelactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Refuelaction *RefuelactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Refuelaction.Contract.Hash0(&_Refuelaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Refuelaction *RefuelactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Refuelaction.Contract.Hash0(&_Refuelaction.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Refuelaction *RefuelactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Refuelaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Refuelaction *RefuelactionSession) Treasury() (common.Address, error) {
	return _Refuelaction.Contract.Treasury(&_Refuelaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Refuelaction *RefuelactionCallerSession) Treasury() (common.Address, error) {
	return _Refuelaction.Contract.Treasury(&_Refuelaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Refuelaction *RefuelactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Refuelaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Refuelaction *RefuelactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Refuelaction.Contract.ChargeFee(&_Refuelaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Refuelaction *RefuelactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Refuelaction.Contract.ChargeFee(&_Refuelaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Refuelaction *RefuelactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Refuelaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Refuelaction *RefuelactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Refuelaction.Contract.Execute(&_Refuelaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Refuelaction *RefuelactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Refuelaction.Contract.Execute(&_Refuelaction.TransactOpts, instruction, arg1, executionState)
}

// RefuelactionRefuelActionFailedIterator is returned from FilterRefuelActionFailed and is used to iterate over the raw logs and unpacked data for RefuelActionFailed events raised by the Refuelaction contract.
type RefuelactionRefuelActionFailedIterator struct {
	Event *RefuelactionRefuelActionFailed // Event containing the contract specifics and raw log

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
func (it *RefuelactionRefuelActionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RefuelactionRefuelActionFailed)
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
		it.Event = new(RefuelactionRefuelActionFailed)
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
func (it *RefuelactionRefuelActionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RefuelactionRefuelActionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RefuelactionRefuelActionFailed represents a RefuelActionFailed event raised by the Refuelaction contract.
type RefuelactionRefuelActionFailed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefuelActionFailed is a free log retrieval operation binding the contract event 0x562e5484ba6302d3b527d81d8e0785d4219c6fc724fefeaace98fb00b7b5296d.
//
// Solidity: event RefuelActionFailed(address indexed target)
func (_Refuelaction *RefuelactionFilterer) FilterRefuelActionFailed(opts *bind.FilterOpts, target []common.Address) (*RefuelactionRefuelActionFailedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Refuelaction.contract.FilterLogs(opts, "RefuelActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return &RefuelactionRefuelActionFailedIterator{contract: _Refuelaction.contract, event: "RefuelActionFailed", logs: logs, sub: sub}, nil
}

// WatchRefuelActionFailed is a free log subscription operation binding the contract event 0x562e5484ba6302d3b527d81d8e0785d4219c6fc724fefeaace98fb00b7b5296d.
//
// Solidity: event RefuelActionFailed(address indexed target)
func (_Refuelaction *RefuelactionFilterer) WatchRefuelActionFailed(opts *bind.WatchOpts, sink chan<- *RefuelactionRefuelActionFailed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Refuelaction.contract.WatchLogs(opts, "RefuelActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RefuelactionRefuelActionFailed)
				if err := _Refuelaction.contract.UnpackLog(event, "RefuelActionFailed", log); err != nil {
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

// ParseRefuelActionFailed is a log parse operation binding the contract event 0x562e5484ba6302d3b527d81d8e0785d4219c6fc724fefeaace98fb00b7b5296d.
//
// Solidity: event RefuelActionFailed(address indexed target)
func (_Refuelaction *RefuelactionFilterer) ParseRefuelActionFailed(log types.Log) (*RefuelactionRefuelActionFailed, error) {
	event := new(RefuelactionRefuelActionFailed)
	if err := _Refuelaction.contract.UnpackLog(event, "RefuelActionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
