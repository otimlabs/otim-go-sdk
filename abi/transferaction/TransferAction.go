// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transferaction

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

// IIntervalSchedule is an auto generated low-level Go binding around an user-defined struct.
type IIntervalSchedule struct {
	StartAt  *big.Int
	StartBy  *big.Int
	Interval *big.Int
	Timeout  *big.Int
}

// IOtimFeeFee is an auto generated low-level Go binding around an user-defined struct.
type IOtimFeeFee struct {
	Token                common.Address
	MaxBaseFeePerGas     *big.Int
	MaxPriorityFeePerGas *big.Int
	ExecutionFee         *big.Int
}

// ITransferActionTransfer is an auto generated low-level Go binding around an user-defined struct.
type ITransferActionTransfer struct {
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Schedule IIntervalSchedule
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

// TransferactionMetaData contains all meta data concerning the Transferaction contract.
var TransferactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"deactivate\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"transfer\",\"type\":\"tuple\",\"internalType\":\"structITransferAction.Transfer\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TransferActionFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// TransferactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferactionMetaData.ABI instead.
var TransferactionABI = TransferactionMetaData.ABI

// Transferaction is an auto generated Go binding around an Ethereum contract.
type Transferaction struct {
	TransferactionCaller     // Read-only binding to the contract
	TransferactionTransactor // Write-only binding to the contract
	TransferactionFilterer   // Log filterer for contract events
}

// TransferactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferactionSession struct {
	Contract     *Transferaction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferactionCallerSession struct {
	Contract *TransferactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferactionTransactorSession struct {
	Contract     *TransferactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferactionRaw struct {
	Contract *Transferaction // Generic contract binding to access the raw methods on
}

// TransferactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferactionCallerRaw struct {
	Contract *TransferactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransferactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferactionTransactorRaw struct {
	Contract *TransferactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferaction creates a new instance of Transferaction, bound to a specific deployed contract.
func NewTransferaction(address common.Address, backend bind.ContractBackend) (*Transferaction, error) {
	contract, err := bindTransferaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transferaction{TransferactionCaller: TransferactionCaller{contract: contract}, TransferactionTransactor: TransferactionTransactor{contract: contract}, TransferactionFilterer: TransferactionFilterer{contract: contract}}, nil
}

// NewTransferactionCaller creates a new read-only instance of Transferaction, bound to a specific deployed contract.
func NewTransferactionCaller(address common.Address, caller bind.ContractCaller) (*TransferactionCaller, error) {
	contract, err := bindTransferaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferactionCaller{contract: contract}, nil
}

// NewTransferactionTransactor creates a new write-only instance of Transferaction, bound to a specific deployed contract.
func NewTransferactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferactionTransactor, error) {
	contract, err := bindTransferaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferactionTransactor{contract: contract}, nil
}

// NewTransferactionFilterer creates a new log filterer instance of Transferaction, bound to a specific deployed contract.
func NewTransferactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferactionFilterer, error) {
	contract, err := bindTransferaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferactionFilterer{contract: contract}, nil
}

// bindTransferaction binds a generic wrapper to an already deployed contract.
func bindTransferaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferaction *TransferactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferaction.Contract.TransferactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferaction *TransferactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferaction.Contract.TransferactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferaction *TransferactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferaction.Contract.TransferactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transferaction *TransferactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transferaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transferaction *TransferactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transferaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transferaction *TransferactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transferaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transferaction *TransferactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Transferaction *TransferactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transferaction.Contract.ArgumentsHash(&_Transferaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transferaction *TransferactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transferaction.Contract.ArgumentsHash(&_Transferaction.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transferaction *TransferactionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transferaction *TransferactionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transferaction.Contract.CheckInterval(&_Transferaction.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transferaction *TransferactionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transferaction.Contract.CheckInterval(&_Transferaction.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transferaction *TransferactionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transferaction *TransferactionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transferaction.Contract.CheckStart(&_Transferaction.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transferaction *TransferactionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transferaction.Contract.CheckStart(&_Transferaction.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferaction *TransferactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferaction *TransferactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Transferaction.Contract.FeeTokenRegistry(&_Transferaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transferaction *TransferactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Transferaction.Contract.FeeTokenRegistry(&_Transferaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferaction *TransferactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferaction *TransferactionSession) GasConstant() (*big.Int, error) {
	return _Transferaction.Contract.GasConstant(&_Transferaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transferaction *TransferactionCallerSession) GasConstant() (*big.Int, error) {
	return _Transferaction.Contract.GasConstant(&_Transferaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x13220799.
//
// Solidity: function hash((address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transfer) pure returns(bytes32)
func (_Transferaction *TransferactionCaller) Hash(opts *bind.CallOpts, transfer ITransferActionTransfer) ([32]byte, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "hash", transfer)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x13220799.
//
// Solidity: function hash((address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transfer) pure returns(bytes32)
func (_Transferaction *TransferactionSession) Hash(transfer ITransferActionTransfer) ([32]byte, error) {
	return _Transferaction.Contract.Hash(&_Transferaction.CallOpts, transfer)
}

// Hash is a free data retrieval call binding the contract method 0x13220799.
//
// Solidity: function hash((address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transfer) pure returns(bytes32)
func (_Transferaction *TransferactionCallerSession) Hash(transfer ITransferActionTransfer) ([32]byte, error) {
	return _Transferaction.Contract.Hash(&_Transferaction.CallOpts, transfer)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transferaction *TransferactionCaller) Hash0(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "hash0", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transferaction *TransferactionSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transferaction.Contract.Hash0(&_Transferaction.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transferaction *TransferactionCallerSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transferaction.Contract.Hash0(&_Transferaction.CallOpts, schedule)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferaction *TransferactionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferaction *TransferactionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transferaction.Contract.Hash1(&_Transferaction.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transferaction *TransferactionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transferaction.Contract.Hash1(&_Transferaction.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferaction *TransferactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transferaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferaction *TransferactionSession) Treasury() (common.Address, error) {
	return _Transferaction.Contract.Treasury(&_Transferaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transferaction *TransferactionCallerSession) Treasury() (common.Address, error) {
	return _Transferaction.Contract.Treasury(&_Transferaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferaction *TransferactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferaction *TransferactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferaction.Contract.ChargeFee(&_Transferaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transferaction *TransferactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transferaction.Contract.ChargeFee(&_Transferaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Transferaction *TransferactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Transferaction *TransferactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferaction.Contract.Execute(&_Transferaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool deactivate)
func (_Transferaction *TransferactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transferaction.Contract.Execute(&_Transferaction.TransactOpts, instruction, arg1, executionState)
}

// TransferactionTransferActionFailedIterator is returned from FilterTransferActionFailed and is used to iterate over the raw logs and unpacked data for TransferActionFailed events raised by the Transferaction contract.
type TransferactionTransferActionFailedIterator struct {
	Event *TransferactionTransferActionFailed // Event containing the contract specifics and raw log

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
func (it *TransferactionTransferActionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferactionTransferActionFailed)
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
		it.Event = new(TransferactionTransferActionFailed)
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
func (it *TransferactionTransferActionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferactionTransferActionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferactionTransferActionFailed represents a TransferActionFailed event raised by the Transferaction contract.
type TransferactionTransferActionFailed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferActionFailed is a free log retrieval operation binding the contract event 0xc541ead5f363cbc1aca6b3b45932c1d02bb72eb471e96fc4597eed922f6bbb02.
//
// Solidity: event TransferActionFailed(address indexed target)
func (_Transferaction *TransferactionFilterer) FilterTransferActionFailed(opts *bind.FilterOpts, target []common.Address) (*TransferactionTransferActionFailedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Transferaction.contract.FilterLogs(opts, "TransferActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return &TransferactionTransferActionFailedIterator{contract: _Transferaction.contract, event: "TransferActionFailed", logs: logs, sub: sub}, nil
}

// WatchTransferActionFailed is a free log subscription operation binding the contract event 0xc541ead5f363cbc1aca6b3b45932c1d02bb72eb471e96fc4597eed922f6bbb02.
//
// Solidity: event TransferActionFailed(address indexed target)
func (_Transferaction *TransferactionFilterer) WatchTransferActionFailed(opts *bind.WatchOpts, sink chan<- *TransferactionTransferActionFailed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Transferaction.contract.WatchLogs(opts, "TransferActionFailed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferactionTransferActionFailed)
				if err := _Transferaction.contract.UnpackLog(event, "TransferActionFailed", log); err != nil {
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

// ParseTransferActionFailed is a log parse operation binding the contract event 0xc541ead5f363cbc1aca6b3b45932c1d02bb72eb471e96fc4597eed922f6bbb02.
//
// Solidity: event TransferActionFailed(address indexed target)
func (_Transferaction *TransferactionFilterer) ParseTransferActionFailed(log types.Log) (*TransferactionTransferActionFailed, error) {
	event := new(TransferactionTransferActionFailed)
	if err := _Transferaction.contract.UnpackLog(event, "TransferActionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
