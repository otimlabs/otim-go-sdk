// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package callonceaction

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

// ICallOnceActionCallOnce is an auto generated low-level Go binding around an user-defined struct.
type ICallOnceActionCallOnce struct {
	Target          common.Address
	AllowFailure    bool
	Value           *big.Int
	GasLimit        *big.Int
	ReturnSizeLimit uint16
	Selector        [4]byte
	Data            []byte
	Fee             IOtimFeeFee
}

// IOtimFeeFee is an auto generated low-level Go binding around an user-defined struct.
type IOtimFeeFee struct {
	Token                common.Address
	MaxBaseFeePerGas     *big.Int
	MaxPriorityFeePerGas *big.Int
	ExecutionFee         *big.Int
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

// CallonceactionMetaData contains all meta data concerning the Callonceaction contract.
var CallonceactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"instructionStorageAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structICallOnceAction.CallOnce\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowFailure\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnSizeLimit\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"instructionStorageAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CallOnceAttempted\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallOnceSucceeded\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"result\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallOnceFailed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// CallonceactionABI is the input ABI used to generate the binding from.
// Deprecated: Use CallonceactionMetaData.ABI instead.
var CallonceactionABI = CallonceactionMetaData.ABI

// Callonceaction is an auto generated Go binding around an Ethereum contract.
type Callonceaction struct {
	CallonceactionCaller     // Read-only binding to the contract
	CallonceactionTransactor // Write-only binding to the contract
	CallonceactionFilterer   // Log filterer for contract events
}

// CallonceactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallonceactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallonceactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallonceactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallonceactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallonceactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallonceactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallonceactionSession struct {
	Contract     *Callonceaction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallonceactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallonceactionCallerSession struct {
	Contract *CallonceactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CallonceactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallonceactionTransactorSession struct {
	Contract     *CallonceactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CallonceactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallonceactionRaw struct {
	Contract *Callonceaction // Generic contract binding to access the raw methods on
}

// CallonceactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallonceactionCallerRaw struct {
	Contract *CallonceactionCaller // Generic read-only contract binding to access the raw methods on
}

// CallonceactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallonceactionTransactorRaw struct {
	Contract *CallonceactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallonceaction creates a new instance of Callonceaction, bound to a specific deployed contract.
func NewCallonceaction(address common.Address, backend bind.ContractBackend) (*Callonceaction, error) {
	contract, err := bindCallonceaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callonceaction{CallonceactionCaller: CallonceactionCaller{contract: contract}, CallonceactionTransactor: CallonceactionTransactor{contract: contract}, CallonceactionFilterer: CallonceactionFilterer{contract: contract}}, nil
}

// NewCallonceactionCaller creates a new read-only instance of Callonceaction, bound to a specific deployed contract.
func NewCallonceactionCaller(address common.Address, caller bind.ContractCaller) (*CallonceactionCaller, error) {
	contract, err := bindCallonceaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallonceactionCaller{contract: contract}, nil
}

// NewCallonceactionTransactor creates a new write-only instance of Callonceaction, bound to a specific deployed contract.
func NewCallonceactionTransactor(address common.Address, transactor bind.ContractTransactor) (*CallonceactionTransactor, error) {
	contract, err := bindCallonceaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallonceactionTransactor{contract: contract}, nil
}

// NewCallonceactionFilterer creates a new log filterer instance of Callonceaction, bound to a specific deployed contract.
func NewCallonceactionFilterer(address common.Address, filterer bind.ContractFilterer) (*CallonceactionFilterer, error) {
	contract, err := bindCallonceaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallonceactionFilterer{contract: contract}, nil
}

// bindCallonceaction binds a generic wrapper to an already deployed contract.
func bindCallonceaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CallonceactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callonceaction *CallonceactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callonceaction.Contract.CallonceactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callonceaction *CallonceactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callonceaction.Contract.CallonceactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callonceaction *CallonceactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callonceaction.Contract.CallonceactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callonceaction *CallonceactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Callonceaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callonceaction *CallonceactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callonceaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callonceaction *CallonceactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callonceaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Callonceaction *CallonceactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Callonceaction *CallonceactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Callonceaction.Contract.ArgumentsHash(&_Callonceaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Callonceaction *CallonceactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Callonceaction.Contract.ArgumentsHash(&_Callonceaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Callonceaction *CallonceactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Callonceaction *CallonceactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Callonceaction.Contract.FeeTokenRegistry(&_Callonceaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Callonceaction *CallonceactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Callonceaction.Contract.FeeTokenRegistry(&_Callonceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Callonceaction *CallonceactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Callonceaction *CallonceactionSession) GasConstant() (*big.Int, error) {
	return _Callonceaction.Contract.GasConstant(&_Callonceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Callonceaction *CallonceactionCallerSession) GasConstant() (*big.Int, error) {
	return _Callonceaction.Contract.GasConstant(&_Callonceaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x8544ba0d.
//
// Solidity: function hash((address,bool,uint256,uint256,uint16,bytes4,bytes,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Callonceaction *CallonceactionCaller) Hash(opts *bind.CallOpts, arguments ICallOnceActionCallOnce) ([32]byte, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x8544ba0d.
//
// Solidity: function hash((address,bool,uint256,uint256,uint16,bytes4,bytes,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Callonceaction *CallonceactionSession) Hash(arguments ICallOnceActionCallOnce) ([32]byte, error) {
	return _Callonceaction.Contract.Hash(&_Callonceaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x8544ba0d.
//
// Solidity: function hash((address,bool,uint256,uint256,uint16,bytes4,bytes,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Callonceaction *CallonceactionCallerSession) Hash(arguments ICallOnceActionCallOnce) ([32]byte, error) {
	return _Callonceaction.Contract.Hash(&_Callonceaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Callonceaction *CallonceactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Callonceaction *CallonceactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Callonceaction.Contract.Hash0(&_Callonceaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Callonceaction *CallonceactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Callonceaction.Contract.Hash0(&_Callonceaction.CallOpts, fee)
}

// InstructionStorageAddress is a free data retrieval call binding the contract method 0x2b17b6d8.
//
// Solidity: function instructionStorageAddress() view returns(address)
func (_Callonceaction *CallonceactionCaller) InstructionStorageAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "instructionStorageAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstructionStorageAddress is a free data retrieval call binding the contract method 0x2b17b6d8.
//
// Solidity: function instructionStorageAddress() view returns(address)
func (_Callonceaction *CallonceactionSession) InstructionStorageAddress() (common.Address, error) {
	return _Callonceaction.Contract.InstructionStorageAddress(&_Callonceaction.CallOpts)
}

// InstructionStorageAddress is a free data retrieval call binding the contract method 0x2b17b6d8.
//
// Solidity: function instructionStorageAddress() view returns(address)
func (_Callonceaction *CallonceactionCallerSession) InstructionStorageAddress() (common.Address, error) {
	return _Callonceaction.Contract.InstructionStorageAddress(&_Callonceaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Callonceaction *CallonceactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Callonceaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Callonceaction *CallonceactionSession) Treasury() (common.Address, error) {
	return _Callonceaction.Contract.Treasury(&_Callonceaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Callonceaction *CallonceactionCallerSession) Treasury() (common.Address, error) {
	return _Callonceaction.Contract.Treasury(&_Callonceaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Callonceaction *CallonceactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Callonceaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Callonceaction *CallonceactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Callonceaction.Contract.ChargeFee(&_Callonceaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Callonceaction *CallonceactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Callonceaction.Contract.ChargeFee(&_Callonceaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Callonceaction *CallonceactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Callonceaction.contract.Transact(opts, "execute", instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Callonceaction *CallonceactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Callonceaction.Contract.Execute(&_Callonceaction.TransactOpts, instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Callonceaction *CallonceactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Callonceaction.Contract.Execute(&_Callonceaction.TransactOpts, instruction, arg1, arg2)
}

// CallonceactionCallOnceAttemptedIterator is returned from FilterCallOnceAttempted and is used to iterate over the raw logs and unpacked data for CallOnceAttempted events raised by the Callonceaction contract.
type CallonceactionCallOnceAttemptedIterator struct {
	Event *CallonceactionCallOnceAttempted // Event containing the contract specifics and raw log

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
func (it *CallonceactionCallOnceAttemptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallonceactionCallOnceAttempted)
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
		it.Event = new(CallonceactionCallOnceAttempted)
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
func (it *CallonceactionCallOnceAttemptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallonceactionCallOnceAttemptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallonceactionCallOnceAttempted represents a CallOnceAttempted event raised by the Callonceaction contract.
type CallonceactionCallOnceAttempted struct {
	Target   common.Address
	Selector [4]byte
	Result   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCallOnceAttempted is a free log retrieval operation binding the contract event 0x7157046e3c71a2c7075735a12b07605696371bd49562367d61423c5875b8f521.
//
// Solidity: event CallOnceAttempted(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) FilterCallOnceAttempted(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*CallonceactionCallOnceAttemptedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _Callonceaction.contract.FilterLogs(opts, "CallOnceAttempted", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &CallonceactionCallOnceAttemptedIterator{contract: _Callonceaction.contract, event: "CallOnceAttempted", logs: logs, sub: sub}, nil
}

// WatchCallOnceAttempted is a free log subscription operation binding the contract event 0x7157046e3c71a2c7075735a12b07605696371bd49562367d61423c5875b8f521.
//
// Solidity: event CallOnceAttempted(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) WatchCallOnceAttempted(opts *bind.WatchOpts, sink chan<- *CallonceactionCallOnceAttempted, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _Callonceaction.contract.WatchLogs(opts, "CallOnceAttempted", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallonceactionCallOnceAttempted)
				if err := _Callonceaction.contract.UnpackLog(event, "CallOnceAttempted", log); err != nil {
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

// ParseCallOnceAttempted is a log parse operation binding the contract event 0x7157046e3c71a2c7075735a12b07605696371bd49562367d61423c5875b8f521.
//
// Solidity: event CallOnceAttempted(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) ParseCallOnceAttempted(log types.Log) (*CallonceactionCallOnceAttempted, error) {
	event := new(CallonceactionCallOnceAttempted)
	if err := _Callonceaction.contract.UnpackLog(event, "CallOnceAttempted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CallonceactionCallOnceSucceededIterator is returned from FilterCallOnceSucceeded and is used to iterate over the raw logs and unpacked data for CallOnceSucceeded events raised by the Callonceaction contract.
type CallonceactionCallOnceSucceededIterator struct {
	Event *CallonceactionCallOnceSucceeded // Event containing the contract specifics and raw log

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
func (it *CallonceactionCallOnceSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallonceactionCallOnceSucceeded)
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
		it.Event = new(CallonceactionCallOnceSucceeded)
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
func (it *CallonceactionCallOnceSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallonceactionCallOnceSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallonceactionCallOnceSucceeded represents a CallOnceSucceeded event raised by the Callonceaction contract.
type CallonceactionCallOnceSucceeded struct {
	Target   common.Address
	Selector [4]byte
	Result   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCallOnceSucceeded is a free log retrieval operation binding the contract event 0x008b13222edaaf9e6d71ef79cd807832c1d5cccc2b0bf3cae96856b931872f75.
//
// Solidity: event CallOnceSucceeded(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) FilterCallOnceSucceeded(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*CallonceactionCallOnceSucceededIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _Callonceaction.contract.FilterLogs(opts, "CallOnceSucceeded", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &CallonceactionCallOnceSucceededIterator{contract: _Callonceaction.contract, event: "CallOnceSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallOnceSucceeded is a free log subscription operation binding the contract event 0x008b13222edaaf9e6d71ef79cd807832c1d5cccc2b0bf3cae96856b931872f75.
//
// Solidity: event CallOnceSucceeded(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) WatchCallOnceSucceeded(opts *bind.WatchOpts, sink chan<- *CallonceactionCallOnceSucceeded, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _Callonceaction.contract.WatchLogs(opts, "CallOnceSucceeded", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallonceactionCallOnceSucceeded)
				if err := _Callonceaction.contract.UnpackLog(event, "CallOnceSucceeded", log); err != nil {
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

// ParseCallOnceSucceeded is a log parse operation binding the contract event 0x008b13222edaaf9e6d71ef79cd807832c1d5cccc2b0bf3cae96856b931872f75.
//
// Solidity: event CallOnceSucceeded(address indexed target, bytes4 indexed selector, bytes result)
func (_Callonceaction *CallonceactionFilterer) ParseCallOnceSucceeded(log types.Log) (*CallonceactionCallOnceSucceeded, error) {
	event := new(CallonceactionCallOnceSucceeded)
	if err := _Callonceaction.contract.UnpackLog(event, "CallOnceSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
