// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transfererc20onceaction

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

// ITransferERC20OnceActionTransferERC20Once is an auto generated low-level Go binding around an user-defined struct.
type ITransferERC20OnceActionTransferERC20Once struct {
	Token  common.Address
	Target common.Address
	Value  *big.Int
	Fee    IOtimFeeFee
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

// Transfererc20onceactionMetaData contains all meta data concerning the Transfererc20onceaction contract.
var Transfererc20onceactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structITransferERC20OnceAction.TransferERC20Once\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// Transfererc20onceactionABI is the input ABI used to generate the binding from.
// Deprecated: Use Transfererc20onceactionMetaData.ABI instead.
var Transfererc20onceactionABI = Transfererc20onceactionMetaData.ABI

// Transfererc20onceaction is an auto generated Go binding around an Ethereum contract.
type Transfererc20onceaction struct {
	Transfererc20onceactionCaller     // Read-only binding to the contract
	Transfererc20onceactionTransactor // Write-only binding to the contract
	Transfererc20onceactionFilterer   // Log filterer for contract events
}

// Transfererc20onceactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Transfererc20onceactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20onceactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Transfererc20onceactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20onceactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Transfererc20onceactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20onceactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Transfererc20onceactionSession struct {
	Contract     *Transfererc20onceaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Transfererc20onceactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Transfererc20onceactionCallerSession struct {
	Contract *Transfererc20onceactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// Transfererc20onceactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Transfererc20onceactionTransactorSession struct {
	Contract     *Transfererc20onceactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// Transfererc20onceactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Transfererc20onceactionRaw struct {
	Contract *Transfererc20onceaction // Generic contract binding to access the raw methods on
}

// Transfererc20onceactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Transfererc20onceactionCallerRaw struct {
	Contract *Transfererc20onceactionCaller // Generic read-only contract binding to access the raw methods on
}

// Transfererc20onceactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Transfererc20onceactionTransactorRaw struct {
	Contract *Transfererc20onceactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfererc20onceaction creates a new instance of Transfererc20onceaction, bound to a specific deployed contract.
func NewTransfererc20onceaction(address common.Address, backend bind.ContractBackend) (*Transfererc20onceaction, error) {
	contract, err := bindTransfererc20onceaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfererc20onceaction{Transfererc20onceactionCaller: Transfererc20onceactionCaller{contract: contract}, Transfererc20onceactionTransactor: Transfererc20onceactionTransactor{contract: contract}, Transfererc20onceactionFilterer: Transfererc20onceactionFilterer{contract: contract}}, nil
}

// NewTransfererc20onceactionCaller creates a new read-only instance of Transfererc20onceaction, bound to a specific deployed contract.
func NewTransfererc20onceactionCaller(address common.Address, caller bind.ContractCaller) (*Transfererc20onceactionCaller, error) {
	contract, err := bindTransfererc20onceaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Transfererc20onceactionCaller{contract: contract}, nil
}

// NewTransfererc20onceactionTransactor creates a new write-only instance of Transfererc20onceaction, bound to a specific deployed contract.
func NewTransfererc20onceactionTransactor(address common.Address, transactor bind.ContractTransactor) (*Transfererc20onceactionTransactor, error) {
	contract, err := bindTransfererc20onceaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Transfererc20onceactionTransactor{contract: contract}, nil
}

// NewTransfererc20onceactionFilterer creates a new log filterer instance of Transfererc20onceaction, bound to a specific deployed contract.
func NewTransfererc20onceactionFilterer(address common.Address, filterer bind.ContractFilterer) (*Transfererc20onceactionFilterer, error) {
	contract, err := bindTransfererc20onceaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Transfererc20onceactionFilterer{contract: contract}, nil
}

// bindTransfererc20onceaction binds a generic wrapper to an already deployed contract.
func bindTransfererc20onceaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Transfererc20onceactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfererc20onceaction *Transfererc20onceactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfererc20onceaction.Contract.Transfererc20onceactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfererc20onceaction *Transfererc20onceactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.Transfererc20onceactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfererc20onceaction *Transfererc20onceactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.Transfererc20onceactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfererc20onceaction *Transfererc20onceactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfererc20onceaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfererc20onceaction *Transfererc20onceactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfererc20onceaction *Transfererc20onceactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Transfererc20onceaction *Transfererc20onceactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfererc20onceaction.Contract.ArgumentsHash(&_Transfererc20onceaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfererc20onceaction.Contract.ArgumentsHash(&_Transfererc20onceaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfererc20onceaction.Contract.FeeTokenRegistry(&_Transfererc20onceaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfererc20onceaction.Contract.FeeTokenRegistry(&_Transfererc20onceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20onceaction *Transfererc20onceactionSession) GasConstant() (*big.Int, error) {
	return _Transfererc20onceaction.Contract.GasConstant(&_Transfererc20onceaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) GasConstant() (*big.Int, error) {
	return _Transfererc20onceaction.Contract.GasConstant(&_Transfererc20onceaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0xeeb36a39.
//
// Solidity: function hash((address,address,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) Hash(opts *bind.CallOpts, arguments ITransferERC20OnceActionTransferERC20Once) ([32]byte, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xeeb36a39.
//
// Solidity: function hash((address,address,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionSession) Hash(arguments ITransferERC20OnceActionTransferERC20Once) ([32]byte, error) {
	return _Transfererc20onceaction.Contract.Hash(&_Transfererc20onceaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0xeeb36a39.
//
// Solidity: function hash((address,address,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) Hash(arguments ITransferERC20OnceActionTransferERC20Once) ([32]byte, error) {
	return _Transfererc20onceaction.Contract.Hash(&_Transfererc20onceaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfererc20onceaction.Contract.Hash0(&_Transfererc20onceaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfererc20onceaction.Contract.Hash0(&_Transfererc20onceaction.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfererc20onceaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionSession) Treasury() (common.Address, error) {
	return _Transfererc20onceaction.Contract.Treasury(&_Transfererc20onceaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20onceaction *Transfererc20onceactionCallerSession) Treasury() (common.Address, error) {
	return _Transfererc20onceaction.Contract.Treasury(&_Transfererc20onceaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20onceaction *Transfererc20onceactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20onceaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20onceaction *Transfererc20onceactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.ChargeFee(&_Transfererc20onceaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20onceaction *Transfererc20onceactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.ChargeFee(&_Transfererc20onceaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Transfererc20onceaction *Transfererc20onceactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20onceaction.contract.Transact(opts, "execute", instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Transfererc20onceaction *Transfererc20onceactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.Execute(&_Transfererc20onceaction.TransactOpts, instruction, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) ) returns(bool)
func (_Transfererc20onceaction *Transfererc20onceactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, arg2 InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20onceaction.Contract.Execute(&_Transfererc20onceaction.TransactOpts, instruction, arg1, arg2)
}
