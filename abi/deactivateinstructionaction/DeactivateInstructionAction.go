// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deactivateinstructionaction

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

// IDeactivateInstructionActionDeactivateInstruction is an auto generated low-level Go binding around an user-defined struct.
type IDeactivateInstructionActionDeactivateInstruction struct {
	InstructionId [32]byte
	Fee           IOtimFeeFee
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

// DeactivateinstructionactionMetaData contains all meta data concerning the Deactivateinstructionaction contract.
var DeactivateinstructionactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"instructionStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structIDeactivateInstructionAction.DeactivateInstruction\",\"components\":[{\"name\":\"instructionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"instructionStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIInstructionStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"InstructionDeactivated\",\"inputs\":[{\"name\":\"instructionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InstructionAlreadyDeactivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// DeactivateinstructionactionABI is the input ABI used to generate the binding from.
// Deprecated: Use DeactivateinstructionactionMetaData.ABI instead.
var DeactivateinstructionactionABI = DeactivateinstructionactionMetaData.ABI

// Deactivateinstructionaction is an auto generated Go binding around an Ethereum contract.
type Deactivateinstructionaction struct {
	DeactivateinstructionactionCaller     // Read-only binding to the contract
	DeactivateinstructionactionTransactor // Write-only binding to the contract
	DeactivateinstructionactionFilterer   // Log filterer for contract events
}

// DeactivateinstructionactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeactivateinstructionactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeactivateinstructionactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeactivateinstructionactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeactivateinstructionactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeactivateinstructionactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeactivateinstructionactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeactivateinstructionactionSession struct {
	Contract     *Deactivateinstructionaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DeactivateinstructionactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeactivateinstructionactionCallerSession struct {
	Contract *DeactivateinstructionactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// DeactivateinstructionactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeactivateinstructionactionTransactorSession struct {
	Contract     *DeactivateinstructionactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// DeactivateinstructionactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeactivateinstructionactionRaw struct {
	Contract *Deactivateinstructionaction // Generic contract binding to access the raw methods on
}

// DeactivateinstructionactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeactivateinstructionactionCallerRaw struct {
	Contract *DeactivateinstructionactionCaller // Generic read-only contract binding to access the raw methods on
}

// DeactivateinstructionactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeactivateinstructionactionTransactorRaw struct {
	Contract *DeactivateinstructionactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeactivateinstructionaction creates a new instance of Deactivateinstructionaction, bound to a specific deployed contract.
func NewDeactivateinstructionaction(address common.Address, backend bind.ContractBackend) (*Deactivateinstructionaction, error) {
	contract, err := bindDeactivateinstructionaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deactivateinstructionaction{DeactivateinstructionactionCaller: DeactivateinstructionactionCaller{contract: contract}, DeactivateinstructionactionTransactor: DeactivateinstructionactionTransactor{contract: contract}, DeactivateinstructionactionFilterer: DeactivateinstructionactionFilterer{contract: contract}}, nil
}

// NewDeactivateinstructionactionCaller creates a new read-only instance of Deactivateinstructionaction, bound to a specific deployed contract.
func NewDeactivateinstructionactionCaller(address common.Address, caller bind.ContractCaller) (*DeactivateinstructionactionCaller, error) {
	contract, err := bindDeactivateinstructionaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeactivateinstructionactionCaller{contract: contract}, nil
}

// NewDeactivateinstructionactionTransactor creates a new write-only instance of Deactivateinstructionaction, bound to a specific deployed contract.
func NewDeactivateinstructionactionTransactor(address common.Address, transactor bind.ContractTransactor) (*DeactivateinstructionactionTransactor, error) {
	contract, err := bindDeactivateinstructionaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeactivateinstructionactionTransactor{contract: contract}, nil
}

// NewDeactivateinstructionactionFilterer creates a new log filterer instance of Deactivateinstructionaction, bound to a specific deployed contract.
func NewDeactivateinstructionactionFilterer(address common.Address, filterer bind.ContractFilterer) (*DeactivateinstructionactionFilterer, error) {
	contract, err := bindDeactivateinstructionaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeactivateinstructionactionFilterer{contract: contract}, nil
}

// bindDeactivateinstructionaction binds a generic wrapper to an already deployed contract.
func bindDeactivateinstructionaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeactivateinstructionactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deactivateinstructionaction *DeactivateinstructionactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deactivateinstructionaction.Contract.DeactivateinstructionactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deactivateinstructionaction *DeactivateinstructionactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.DeactivateinstructionactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deactivateinstructionaction *DeactivateinstructionactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.DeactivateinstructionactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deactivateinstructionaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Deactivateinstructionaction.Contract.ArgumentsHash(&_Deactivateinstructionaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Deactivateinstructionaction.Contract.ArgumentsHash(&_Deactivateinstructionaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.FeeTokenRegistry(&_Deactivateinstructionaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.FeeTokenRegistry(&_Deactivateinstructionaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) GasConstant() (*big.Int, error) {
	return _Deactivateinstructionaction.Contract.GasConstant(&_Deactivateinstructionaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) GasConstant() (*big.Int, error) {
	return _Deactivateinstructionaction.Contract.GasConstant(&_Deactivateinstructionaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0xe85ea1a1.
//
// Solidity: function hash((bytes32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) Hash(opts *bind.CallOpts, arguments IDeactivateInstructionActionDeactivateInstruction) ([32]byte, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xe85ea1a1.
//
// Solidity: function hash((bytes32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) Hash(arguments IDeactivateInstructionActionDeactivateInstruction) ([32]byte, error) {
	return _Deactivateinstructionaction.Contract.Hash(&_Deactivateinstructionaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0xe85ea1a1.
//
// Solidity: function hash((bytes32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) Hash(arguments IDeactivateInstructionActionDeactivateInstruction) ([32]byte, error) {
	return _Deactivateinstructionaction.Contract.Hash(&_Deactivateinstructionaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Deactivateinstructionaction.Contract.Hash0(&_Deactivateinstructionaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Deactivateinstructionaction.Contract.Hash0(&_Deactivateinstructionaction.CallOpts, fee)
}

// InstructionStorage is a free data retrieval call binding the contract method 0xce0dafae.
//
// Solidity: function instructionStorage() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) InstructionStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "instructionStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InstructionStorage is a free data retrieval call binding the contract method 0xce0dafae.
//
// Solidity: function instructionStorage() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) InstructionStorage() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.InstructionStorage(&_Deactivateinstructionaction.CallOpts)
}

// InstructionStorage is a free data retrieval call binding the contract method 0xce0dafae.
//
// Solidity: function instructionStorage() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) InstructionStorage() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.InstructionStorage(&_Deactivateinstructionaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deactivateinstructionaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) Treasury() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.Treasury(&_Deactivateinstructionaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Deactivateinstructionaction *DeactivateinstructionactionCallerSession) Treasury() (common.Address, error) {
	return _Deactivateinstructionaction.Contract.Treasury(&_Deactivateinstructionaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Deactivateinstructionaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.ChargeFee(&_Deactivateinstructionaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.ChargeFee(&_Deactivateinstructionaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Deactivateinstructionaction.contract.Transact(opts, "execute", instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Deactivateinstructionaction *DeactivateinstructionactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.Execute(&_Deactivateinstructionaction.TransactOpts, instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Deactivateinstructionaction *DeactivateinstructionactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Deactivateinstructionaction.Contract.Execute(&_Deactivateinstructionaction.TransactOpts, instruction, arg1, arg2)
}

// DeactivateinstructionactionInstructionDeactivatedIterator is returned from FilterInstructionDeactivated and is used to iterate over the raw logs and unpacked data for InstructionDeactivated events raised by the Deactivateinstructionaction contract.
type DeactivateinstructionactionInstructionDeactivatedIterator struct {
	Event *DeactivateinstructionactionInstructionDeactivated // Event containing the contract specifics and raw log

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
func (it *DeactivateinstructionactionInstructionDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeactivateinstructionactionInstructionDeactivated)
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
		it.Event = new(DeactivateinstructionactionInstructionDeactivated)
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
func (it *DeactivateinstructionactionInstructionDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeactivateinstructionactionInstructionDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeactivateinstructionactionInstructionDeactivated represents a InstructionDeactivated event raised by the Deactivateinstructionaction contract.
type DeactivateinstructionactionInstructionDeactivated struct {
	InstructionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInstructionDeactivated is a free log retrieval operation binding the contract event 0x1082dc6fb1ba8d32f074a5e83864cf6cc4b447170a99e122100a5ed6107bd574.
//
// Solidity: event InstructionDeactivated(bytes32 indexed instructionId)
func (_Deactivateinstructionaction *DeactivateinstructionactionFilterer) FilterInstructionDeactivated(opts *bind.FilterOpts, instructionId [][32]byte) (*DeactivateinstructionactionInstructionDeactivatedIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}

	logs, sub, err := _Deactivateinstructionaction.contract.FilterLogs(opts, "InstructionDeactivated", instructionIdRule)
	if err != nil {
		return nil, err
	}
	return &DeactivateinstructionactionInstructionDeactivatedIterator{contract: _Deactivateinstructionaction.contract, event: "InstructionDeactivated", logs: logs, sub: sub}, nil
}

// WatchInstructionDeactivated is a free log subscription operation binding the contract event 0x1082dc6fb1ba8d32f074a5e83864cf6cc4b447170a99e122100a5ed6107bd574.
//
// Solidity: event InstructionDeactivated(bytes32 indexed instructionId)
func (_Deactivateinstructionaction *DeactivateinstructionactionFilterer) WatchInstructionDeactivated(opts *bind.WatchOpts, sink chan<- *DeactivateinstructionactionInstructionDeactivated, instructionId [][32]byte) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}

	logs, sub, err := _Deactivateinstructionaction.contract.WatchLogs(opts, "InstructionDeactivated", instructionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeactivateinstructionactionInstructionDeactivated)
				if err := _Deactivateinstructionaction.contract.UnpackLog(event, "InstructionDeactivated", log); err != nil {
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

// ParseInstructionDeactivated is a log parse operation binding the contract event 0x1082dc6fb1ba8d32f074a5e83864cf6cc4b447170a99e122100a5ed6107bd574.
//
// Solidity: event InstructionDeactivated(bytes32 indexed instructionId)
func (_Deactivateinstructionaction *DeactivateinstructionactionFilterer) ParseInstructionDeactivated(log types.Log) (*DeactivateinstructionactionInstructionDeactivated, error) {
	event := new(DeactivateinstructionactionInstructionDeactivated)
	if err := _Deactivateinstructionaction.contract.UnpackLog(event, "InstructionDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
