// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transferonceaction

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

// ITransferOnceActionTransferOnce is an auto generated low-level Go binding around an user-defined struct.
type ITransferOnceActionTransferOnce struct {
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Fee      IOtimFeeFee
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

// TransferonceactionMetaData contains all meta data concerning the Transferonceaction contract.
var TransferonceactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"deactivate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structITransferOnceAction.TransferOnce\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TransferOnceActionFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// TransferonceactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferonceactionMetaData.ABI instead.
var TransferonceactionABI = TransferonceactionMetaData.ABI

// Transferonceaction is an auto generated Go binding around an Ethereum contract.
type Transferonceaction struct {
	TransferonceactionCaller     // Read-only binding to the contract
	TransferonceactionTransactor // Write-only binding to the contract
	TransferonceactionFilterer   // Log filterer for contract events
}

// TransferonceactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferonceactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferonceactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferonceactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferonceactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferonceactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferonceactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferonceactionSession struct {
	Contract     *Transferonceaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransferonceactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferonceactionCallerSession struct {
	Contract *TransferonceactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TransferonceactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferonceactionTransactorSession struct {
	Contract     *TransferonceactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TransferonceactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferonceactionRaw struct {
	Contract *Transferonceaction // Generic contract binding to access the raw methods on
}

// TransferonceactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferonceactionCallerRaw struct {
	Contract *TransferonceactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransferonceactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferonceactionTransactorRaw struct {
	Contract *TransferonceactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferonceaction creates a new instance of Transferonceaction, bound to a specific deployed contract.
func NewTransferonceaction(address common.Address, backend bind.ContractBackend) (*Transferonceaction, error) {
	contract, err := bindTransferonceaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transferonceaction{TransferonceactionCaller: TransferonceactionCaller{contract: contract}, TransferonceactionTransactor: TransferonceactionTransactor{contract: contract}, TransferonceactionFilterer: TransferonceactionFilterer{contract: contract}}, nil
}

// NewTransferonceactionCaller creates a new read-only instance of Transferonceaction, bound to a specific deployed contract.
func NewTransferonceactionCaller(address common.Address, caller bind.ContractCaller) (*TransferonceactionCaller, error) {
	contract, err := bindTransferonceaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferonceactionCaller{contract: contract}, nil
}

// NewTransferonceactionTransactor creates a new write-only instance of Transferonceaction, bound to a specific deployed contract.
func NewTransferonceactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferonceactionTransactor, error) {
	contract, err := bindTransferonceaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferonceactionTransactor{contract: contract}, nil
}

// NewTransferonceactionFilterer creates a new log filterer instance of Transferonceaction, bound to a specific deployed contract.
func NewTransferonceactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferonceactionFilterer, error) {
	contract, err := bindTransferonceaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferonceactionFilterer{contract: contract}, nil
}

// bindTransferonceaction binds a generic wrapper to an already deployed contract.
func bindTransferonceaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferonceactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferonceaction *TransferonceactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferonceaction.Contract.TransferonceactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferonceaction *TransferonceactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferonceaction.Contract.TransferonceactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferonceaction *TransferonceactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferonceaction.Contract.TransferonceactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferonceaction *TransferonceactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferonceaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferonceaction *TransferonceactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferonceaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferonceaction *TransferonceactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferonceaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transferonceaction *TransferonceactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Transferonceaction *TransferonceactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transferonceaction.Contract.ArgumentsHash(&_Transferonceaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transferonceaction *TransferonceactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transferonceaction.Contract.ArgumentsHash(&_Transferonceaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferonceaction *TransferonceactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferonceaction *TransferonceactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Transferonceaction.Contract.FeeTokenRegistry(&_Transferonceaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferonceaction *TransferonceactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Transferonceaction.Contract.FeeTokenRegistry(&_Transferonceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferonceaction *TransferonceactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferonceaction *TransferonceactionSession) GasConstant() (*big.Int, error) {
	return _Transferonceaction.Contract.GasConstant(&_Transferonceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferonceaction *TransferonceactionCallerSession) GasConstant() (*big.Int, error) {
	return _Transferonceaction.Contract.GasConstant(&_Transferonceaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x220f28a3.
//
// Solidity: function hash((address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionCaller) Hash(opts *bind.CallOpts, arguments ITransferOnceActionTransferOnce) ([32]byte, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x220f28a3.
//
// Solidity: function hash((address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionSession) Hash(arguments ITransferOnceActionTransferOnce) ([32]byte, error) {
	return _Transferonceaction.Contract.Hash(&_Transferonceaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x220f28a3.
//
// Solidity: function hash((address,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionCallerSession) Hash(arguments ITransferOnceActionTransferOnce) ([32]byte, error) {
	return _Transferonceaction.Contract.Hash(&_Transferonceaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Transferonceaction.Contract.Hash0(&_Transferonceaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferonceaction *TransferonceactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Transferonceaction.Contract.Hash0(&_Transferonceaction.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferonceaction *TransferonceactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transferonceaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferonceaction *TransferonceactionSession) Treasury() (common.Address, error) {
	return _Transferonceaction.Contract.Treasury(&_Transferonceaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferonceaction *TransferonceactionCallerSession) Treasury() (common.Address, error) {
	return _Transferonceaction.Contract.Treasury(&_Transferonceaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferonceaction *TransferonceactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferonceaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferonceaction *TransferonceactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferonceaction.Contract.ChargeFee(&_Transferonceaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferonceaction *TransferonceactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferonceaction.Contract.ChargeFee(&_Transferonceaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool deactivate)
func (_Transferonceaction *TransferonceactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferonceaction.contract.Transact(opts, "execute", instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool deactivate)
func (_Transferonceaction *TransferonceactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferonceaction.Contract.Execute(&_Transferonceaction.TransactOpts, instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool deactivate)
func (_Transferonceaction *TransferonceactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferonceaction.Contract.Execute(&_Transferonceaction.TransactOpts, instruction, arg1, arg2)
}

// TransferonceactionTransferOnceActionFailedIterator is returned from FilterTransferOnceActionFailed and is used to iterate over the raw logs and unpacked data for TransferOnceActionFailed events raised by the Transferonceaction contract.
type TransferonceactionTransferOnceActionFailedIterator struct {
	Event *TransferonceactionTransferOnceActionFailed // Event containing the contract specifics and raw log

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
func (it *TransferonceactionTransferOnceActionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferonceactionTransferOnceActionFailed)
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
		it.Event = new(TransferonceactionTransferOnceActionFailed)
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
func (it *TransferonceactionTransferOnceActionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferonceactionTransferOnceActionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferonceactionTransferOnceActionFailed represents a TransferOnceActionFailed event raised by the Transferonceaction contract.
type TransferonceactionTransferOnceActionFailed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferOnceActionFailed is a free log retrieval operation binding the contract event 0x1b7c8ea9837921f2032ffb7e1482e6b636f4f97f92952da6d6eecbe9651f874d.
//
// Solidity: event TransferOnceActionFailed(address indexed target)
func (_Transferonceaction *TransferonceactionFilterer) FilterTransferOnceActionFailed(opts *bind.FilterOpts, target []common.Address) (*TransferonceactionTransferOnceActionFailedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Transferonceaction.contract.FilterLogs(opts, "TransferOnceActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return &TransferonceactionTransferOnceActionFailedIterator{contract: _Transferonceaction.contract, event: "TransferOnceActionFailed", logs: logs, sub: sub}, nil
}

// WatchTransferOnceActionFailed is a free log subscription operation binding the contract event 0x1b7c8ea9837921f2032ffb7e1482e6b636f4f97f92952da6d6eecbe9651f874d.
//
// Solidity: event TransferOnceActionFailed(address indexed target)
func (_Transferonceaction *TransferonceactionFilterer) WatchTransferOnceActionFailed(opts *bind.WatchOpts, sink chan<- *TransferonceactionTransferOnceActionFailed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Transferonceaction.contract.WatchLogs(opts, "TransferOnceActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferonceactionTransferOnceActionFailed)
				if err := _Transferonceaction.contract.UnpackLog(event, "TransferOnceActionFailed", log); err != nil {
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

// ParseTransferOnceActionFailed is a log parse operation binding the contract event 0x1b7c8ea9837921f2032ffb7e1482e6b636f4f97f92952da6d6eecbe9651f874d.
//
// Solidity: event TransferOnceActionFailed(address indexed target)
func (_Transferonceaction *TransferonceactionFilterer) ParseTransferOnceActionFailed(log types.Log) (*TransferonceactionTransferOnceActionFailed, error) {
	event := new(TransferonceactionTransferOnceActionFailed)
	if err := _Transferonceaction.contract.UnpackLog(event, "TransferOnceActionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
