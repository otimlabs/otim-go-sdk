// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sweeperc20action

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

// ISweepERC20ActionSweepERC20 is an auto generated low-level Go binding around an user-defined struct.
type ISweepERC20ActionSweepERC20 struct {
	Token      common.Address
	Target     common.Address
	Threshold  *big.Int
	EndBalance *big.Int
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

// Sweeperc20actionMetaData contains all meta data concerning the Sweeperc20action contract.
var Sweeperc20actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"deactivate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structISweepERC20Action.SweepERC20\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BalanceUnderThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// Sweeperc20actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Sweeperc20actionMetaData.ABI instead.
var Sweeperc20actionABI = Sweeperc20actionMetaData.ABI

// Sweeperc20action is an auto generated Go binding around an Ethereum contract.
type Sweeperc20action struct {
	Sweeperc20actionCaller     // Read-only binding to the contract
	Sweeperc20actionTransactor // Write-only binding to the contract
	Sweeperc20actionFilterer   // Log filterer for contract events
}

// Sweeperc20actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Sweeperc20actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweeperc20actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Sweeperc20actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweeperc20actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sweeperc20actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweeperc20actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sweeperc20actionSession struct {
	Contract     *Sweeperc20action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Sweeperc20actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sweeperc20actionCallerSession struct {
	Contract *Sweeperc20actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Sweeperc20actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sweeperc20actionTransactorSession struct {
	Contract     *Sweeperc20actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Sweeperc20actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Sweeperc20actionRaw struct {
	Contract *Sweeperc20action // Generic contract binding to access the raw methods on
}

// Sweeperc20actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sweeperc20actionCallerRaw struct {
	Contract *Sweeperc20actionCaller // Generic read-only contract binding to access the raw methods on
}

// Sweeperc20actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sweeperc20actionTransactorRaw struct {
	Contract *Sweeperc20actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweeperc20action creates a new instance of Sweeperc20action, bound to a specific deployed contract.
func NewSweeperc20action(address common.Address, backend bind.ContractBackend) (*Sweeperc20action, error) {
	contract, err := bindSweeperc20action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sweeperc20action{Sweeperc20actionCaller: Sweeperc20actionCaller{contract: contract}, Sweeperc20actionTransactor: Sweeperc20actionTransactor{contract: contract}, Sweeperc20actionFilterer: Sweeperc20actionFilterer{contract: contract}}, nil
}

// NewSweeperc20actionCaller creates a new read-only instance of Sweeperc20action, bound to a specific deployed contract.
func NewSweeperc20actionCaller(address common.Address, caller bind.ContractCaller) (*Sweeperc20actionCaller, error) {
	contract, err := bindSweeperc20action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sweeperc20actionCaller{contract: contract}, nil
}

// NewSweeperc20actionTransactor creates a new write-only instance of Sweeperc20action, bound to a specific deployed contract.
func NewSweeperc20actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Sweeperc20actionTransactor, error) {
	contract, err := bindSweeperc20action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sweeperc20actionTransactor{contract: contract}, nil
}

// NewSweeperc20actionFilterer creates a new log filterer instance of Sweeperc20action, bound to a specific deployed contract.
func NewSweeperc20actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Sweeperc20actionFilterer, error) {
	contract, err := bindSweeperc20action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sweeperc20actionFilterer{contract: contract}, nil
}

// bindSweeperc20action binds a generic wrapper to an already deployed contract.
func bindSweeperc20action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Sweeperc20actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweeperc20action *Sweeperc20actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweeperc20action.Contract.Sweeperc20actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweeperc20action *Sweeperc20actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.Sweeperc20actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweeperc20action *Sweeperc20actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.Sweeperc20actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweeperc20action *Sweeperc20actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweeperc20action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweeperc20action *Sweeperc20actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweeperc20action *Sweeperc20actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweeperc20action *Sweeperc20actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Sweeperc20action *Sweeperc20actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweeperc20action.Contract.ArgumentsHash(&_Sweeperc20action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweeperc20action *Sweeperc20actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweeperc20action.Contract.ArgumentsHash(&_Sweeperc20action.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweeperc20action *Sweeperc20actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweeperc20action *Sweeperc20actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweeperc20action.Contract.FeeTokenRegistry(&_Sweeperc20action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweeperc20action *Sweeperc20actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweeperc20action.Contract.FeeTokenRegistry(&_Sweeperc20action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweeperc20action *Sweeperc20actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweeperc20action *Sweeperc20actionSession) GasConstant() (*big.Int, error) {
	return _Sweeperc20action.Contract.GasConstant(&_Sweeperc20action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweeperc20action *Sweeperc20actionCallerSession) GasConstant() (*big.Int, error) {
	return _Sweeperc20action.Contract.GasConstant(&_Sweeperc20action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x343c2f1c.
//
// Solidity: function hash((address,address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionCaller) Hash(opts *bind.CallOpts, arguments ISweepERC20ActionSweepERC20) ([32]byte, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x343c2f1c.
//
// Solidity: function hash((address,address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionSession) Hash(arguments ISweepERC20ActionSweepERC20) ([32]byte, error) {
	return _Sweeperc20action.Contract.Hash(&_Sweeperc20action.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x343c2f1c.
//
// Solidity: function hash((address,address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionCallerSession) Hash(arguments ISweepERC20ActionSweepERC20) ([32]byte, error) {
	return _Sweeperc20action.Contract.Hash(&_Sweeperc20action.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweeperc20action.Contract.Hash0(&_Sweeperc20action.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweeperc20action *Sweeperc20actionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweeperc20action.Contract.Hash0(&_Sweeperc20action.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweeperc20action *Sweeperc20actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweeperc20action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweeperc20action *Sweeperc20actionSession) Treasury() (common.Address, error) {
	return _Sweeperc20action.Contract.Treasury(&_Sweeperc20action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweeperc20action *Sweeperc20actionCallerSession) Treasury() (common.Address, error) {
	return _Sweeperc20action.Contract.Treasury(&_Sweeperc20action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweeperc20action *Sweeperc20actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweeperc20action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweeperc20action *Sweeperc20actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.ChargeFee(&_Sweeperc20action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweeperc20action *Sweeperc20actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.ChargeFee(&_Sweeperc20action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweeperc20action *Sweeperc20actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweeperc20action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweeperc20action *Sweeperc20actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.Execute(&_Sweeperc20action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Sweeperc20action *Sweeperc20actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweeperc20action.Contract.Execute(&_Sweeperc20action.TransactOpts, instruction, arg1, executionState)
}
