// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sweepdepositerc4626action

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

// ISweepDepositERC4626ActionSweepDepositERC4626 is an auto generated low-level Go binding around an user-defined struct.
type ISweepDepositERC4626ActionSweepDepositERC4626 struct {
	Vault          common.Address
	Recipient      common.Address
	Threshold      *big.Int
	EndBalance     *big.Int
	MinDeposit     *big.Int
	MinTotalShares *big.Int
	Fee            IOtimFeeFee
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

// Sweepdepositerc4626actionMetaData contains all meta data concerning the Sweepdepositerc4626action contract.
var Sweepdepositerc4626actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structISweepDepositERC4626Action.SweepDepositERC4626\",\"components\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTotalShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxDepositReached\",\"inputs\":[{\"name\":\"maxDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newEndBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceUnderThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDepositTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TotalSharesTooLow\",\"inputs\":[]}]",
}

// Sweepdepositerc4626actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Sweepdepositerc4626actionMetaData.ABI instead.
var Sweepdepositerc4626actionABI = Sweepdepositerc4626actionMetaData.ABI

// Sweepdepositerc4626action is an auto generated Go binding around an Ethereum contract.
type Sweepdepositerc4626action struct {
	Sweepdepositerc4626actionCaller     // Read-only binding to the contract
	Sweepdepositerc4626actionTransactor // Write-only binding to the contract
	Sweepdepositerc4626actionFilterer   // Log filterer for contract events
}

// Sweepdepositerc4626actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Sweepdepositerc4626actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepdepositerc4626actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Sweepdepositerc4626actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepdepositerc4626actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sweepdepositerc4626actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepdepositerc4626actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sweepdepositerc4626actionSession struct {
	Contract     *Sweepdepositerc4626action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Sweepdepositerc4626actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sweepdepositerc4626actionCallerSession struct {
	Contract *Sweepdepositerc4626actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Sweepdepositerc4626actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sweepdepositerc4626actionTransactorSession struct {
	Contract     *Sweepdepositerc4626actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Sweepdepositerc4626actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Sweepdepositerc4626actionRaw struct {
	Contract *Sweepdepositerc4626action // Generic contract binding to access the raw methods on
}

// Sweepdepositerc4626actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sweepdepositerc4626actionCallerRaw struct {
	Contract *Sweepdepositerc4626actionCaller // Generic read-only contract binding to access the raw methods on
}

// Sweepdepositerc4626actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sweepdepositerc4626actionTransactorRaw struct {
	Contract *Sweepdepositerc4626actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweepdepositerc4626action creates a new instance of Sweepdepositerc4626action, bound to a specific deployed contract.
func NewSweepdepositerc4626action(address common.Address, backend bind.ContractBackend) (*Sweepdepositerc4626action, error) {
	contract, err := bindSweepdepositerc4626action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sweepdepositerc4626action{Sweepdepositerc4626actionCaller: Sweepdepositerc4626actionCaller{contract: contract}, Sweepdepositerc4626actionTransactor: Sweepdepositerc4626actionTransactor{contract: contract}, Sweepdepositerc4626actionFilterer: Sweepdepositerc4626actionFilterer{contract: contract}}, nil
}

// NewSweepdepositerc4626actionCaller creates a new read-only instance of Sweepdepositerc4626action, bound to a specific deployed contract.
func NewSweepdepositerc4626actionCaller(address common.Address, caller bind.ContractCaller) (*Sweepdepositerc4626actionCaller, error) {
	contract, err := bindSweepdepositerc4626action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sweepdepositerc4626actionCaller{contract: contract}, nil
}

// NewSweepdepositerc4626actionTransactor creates a new write-only instance of Sweepdepositerc4626action, bound to a specific deployed contract.
func NewSweepdepositerc4626actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Sweepdepositerc4626actionTransactor, error) {
	contract, err := bindSweepdepositerc4626action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sweepdepositerc4626actionTransactor{contract: contract}, nil
}

// NewSweepdepositerc4626actionFilterer creates a new log filterer instance of Sweepdepositerc4626action, bound to a specific deployed contract.
func NewSweepdepositerc4626actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Sweepdepositerc4626actionFilterer, error) {
	contract, err := bindSweepdepositerc4626action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sweepdepositerc4626actionFilterer{contract: contract}, nil
}

// bindSweepdepositerc4626action binds a generic wrapper to an already deployed contract.
func bindSweepdepositerc4626action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Sweepdepositerc4626actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepdepositerc4626action.Contract.Sweepdepositerc4626actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.Sweepdepositerc4626actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.Sweepdepositerc4626actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepdepositerc4626action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepdepositerc4626action.Contract.ArgumentsHash(&_Sweepdepositerc4626action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepdepositerc4626action.Contract.ArgumentsHash(&_Sweepdepositerc4626action.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepdepositerc4626action.Contract.FeeTokenRegistry(&_Sweepdepositerc4626action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepdepositerc4626action.Contract.FeeTokenRegistry(&_Sweepdepositerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) GasConstant() (*big.Int, error) {
	return _Sweepdepositerc4626action.Contract.GasConstant(&_Sweepdepositerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) GasConstant() (*big.Int, error) {
	return _Sweepdepositerc4626action.Contract.GasConstant(&_Sweepdepositerc4626action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x6570ecdd.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) Hash(opts *bind.CallOpts, arguments ISweepDepositERC4626ActionSweepDepositERC4626) ([32]byte, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x6570ecdd.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) Hash(arguments ISweepDepositERC4626ActionSweepDepositERC4626) ([32]byte, error) {
	return _Sweepdepositerc4626action.Contract.Hash(&_Sweepdepositerc4626action.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x6570ecdd.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) Hash(arguments ISweepDepositERC4626ActionSweepDepositERC4626) ([32]byte, error) {
	return _Sweepdepositerc4626action.Contract.Hash(&_Sweepdepositerc4626action.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepdepositerc4626action.Contract.Hash0(&_Sweepdepositerc4626action.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepdepositerc4626action.Contract.Hash0(&_Sweepdepositerc4626action.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepdepositerc4626action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) Treasury() (common.Address, error) {
	return _Sweepdepositerc4626action.Contract.Treasury(&_Sweepdepositerc4626action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionCallerSession) Treasury() (common.Address, error) {
	return _Sweepdepositerc4626action.Contract.Treasury(&_Sweepdepositerc4626action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.ChargeFee(&_Sweepdepositerc4626action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.ChargeFee(&_Sweepdepositerc4626action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.Execute(&_Sweepdepositerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepdepositerc4626action.Contract.Execute(&_Sweepdepositerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Sweepdepositerc4626actionMaxDepositReachedIterator is returned from FilterMaxDepositReached and is used to iterate over the raw logs and unpacked data for MaxDepositReached events raised by the Sweepdepositerc4626action contract.
type Sweepdepositerc4626actionMaxDepositReachedIterator struct {
	Event *Sweepdepositerc4626actionMaxDepositReached // Event containing the contract specifics and raw log

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
func (it *Sweepdepositerc4626actionMaxDepositReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Sweepdepositerc4626actionMaxDepositReached)
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
		it.Event = new(Sweepdepositerc4626actionMaxDepositReached)
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
func (it *Sweepdepositerc4626actionMaxDepositReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Sweepdepositerc4626actionMaxDepositReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Sweepdepositerc4626actionMaxDepositReached represents a MaxDepositReached event raised by the Sweepdepositerc4626action contract.
type Sweepdepositerc4626actionMaxDepositReached struct {
	MaxDeposit    *big.Int
	NewEndBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxDepositReached is a free log retrieval operation binding the contract event 0xc9dac0507dbc768d685d40569f8b2dece6391e20a6e5a3af16aa49ca27ae9ade.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit, uint256 newEndBalance)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionFilterer) FilterMaxDepositReached(opts *bind.FilterOpts) (*Sweepdepositerc4626actionMaxDepositReachedIterator, error) {

	logs, sub, err := _Sweepdepositerc4626action.contract.FilterLogs(opts, "MaxDepositReached")
	if err != nil {
		return nil, err
	}
	return &Sweepdepositerc4626actionMaxDepositReachedIterator{contract: _Sweepdepositerc4626action.contract, event: "MaxDepositReached", logs: logs, sub: sub}, nil
}

// WatchMaxDepositReached is a free log subscription operation binding the contract event 0xc9dac0507dbc768d685d40569f8b2dece6391e20a6e5a3af16aa49ca27ae9ade.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit, uint256 newEndBalance)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionFilterer) WatchMaxDepositReached(opts *bind.WatchOpts, sink chan<- *Sweepdepositerc4626actionMaxDepositReached) (event.Subscription, error) {

	logs, sub, err := _Sweepdepositerc4626action.contract.WatchLogs(opts, "MaxDepositReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Sweepdepositerc4626actionMaxDepositReached)
				if err := _Sweepdepositerc4626action.contract.UnpackLog(event, "MaxDepositReached", log); err != nil {
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

// ParseMaxDepositReached is a log parse operation binding the contract event 0xc9dac0507dbc768d685d40569f8b2dece6391e20a6e5a3af16aa49ca27ae9ade.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit, uint256 newEndBalance)
func (_Sweepdepositerc4626action *Sweepdepositerc4626actionFilterer) ParseMaxDepositReached(log types.Log) (*Sweepdepositerc4626actionMaxDepositReached, error) {
	event := new(Sweepdepositerc4626actionMaxDepositReached)
	if err := _Sweepdepositerc4626action.contract.UnpackLog(event, "MaxDepositReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
