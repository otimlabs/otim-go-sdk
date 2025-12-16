// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositerc4626action

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

// IDepositERC4626ActionDepositERC4626 is an auto generated low-level Go binding around an user-defined struct.
type IDepositERC4626ActionDepositERC4626 struct {
	Vault          common.Address
	Recipient      common.Address
	Value          *big.Int
	MinDeposit     *big.Int
	MinTotalShares *big.Int
	Schedule       IIntervalSchedule
	Fee            IOtimFeeFee
}

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

// Depositerc4626actionMetaData contains all meta data concerning the Depositerc4626action contract.
var Depositerc4626actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structIDepositERC4626Action.DepositERC4626\",\"components\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minTotalShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxDepositReached\",\"inputs\":[{\"name\":\"maxDeposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxDepositTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TotalSharesTooLow\",\"inputs\":[]}]",
}

// Depositerc4626actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Depositerc4626actionMetaData.ABI instead.
var Depositerc4626actionABI = Depositerc4626actionMetaData.ABI

// Depositerc4626action is an auto generated Go binding around an Ethereum contract.
type Depositerc4626action struct {
	Depositerc4626actionCaller     // Read-only binding to the contract
	Depositerc4626actionTransactor // Write-only binding to the contract
	Depositerc4626actionFilterer   // Log filterer for contract events
}

// Depositerc4626actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Depositerc4626actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositerc4626actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Depositerc4626actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositerc4626actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Depositerc4626actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Depositerc4626actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Depositerc4626actionSession struct {
	Contract     *Depositerc4626action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Depositerc4626actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Depositerc4626actionCallerSession struct {
	Contract *Depositerc4626actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// Depositerc4626actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Depositerc4626actionTransactorSession struct {
	Contract     *Depositerc4626actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Depositerc4626actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Depositerc4626actionRaw struct {
	Contract *Depositerc4626action // Generic contract binding to access the raw methods on
}

// Depositerc4626actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Depositerc4626actionCallerRaw struct {
	Contract *Depositerc4626actionCaller // Generic read-only contract binding to access the raw methods on
}

// Depositerc4626actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Depositerc4626actionTransactorRaw struct {
	Contract *Depositerc4626actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositerc4626action creates a new instance of Depositerc4626action, bound to a specific deployed contract.
func NewDepositerc4626action(address common.Address, backend bind.ContractBackend) (*Depositerc4626action, error) {
	contract, err := bindDepositerc4626action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositerc4626action{Depositerc4626actionCaller: Depositerc4626actionCaller{contract: contract}, Depositerc4626actionTransactor: Depositerc4626actionTransactor{contract: contract}, Depositerc4626actionFilterer: Depositerc4626actionFilterer{contract: contract}}, nil
}

// NewDepositerc4626actionCaller creates a new read-only instance of Depositerc4626action, bound to a specific deployed contract.
func NewDepositerc4626actionCaller(address common.Address, caller bind.ContractCaller) (*Depositerc4626actionCaller, error) {
	contract, err := bindDepositerc4626action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Depositerc4626actionCaller{contract: contract}, nil
}

// NewDepositerc4626actionTransactor creates a new write-only instance of Depositerc4626action, bound to a specific deployed contract.
func NewDepositerc4626actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Depositerc4626actionTransactor, error) {
	contract, err := bindDepositerc4626action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Depositerc4626actionTransactor{contract: contract}, nil
}

// NewDepositerc4626actionFilterer creates a new log filterer instance of Depositerc4626action, bound to a specific deployed contract.
func NewDepositerc4626actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Depositerc4626actionFilterer, error) {
	contract, err := bindDepositerc4626action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Depositerc4626actionFilterer{contract: contract}, nil
}

// bindDepositerc4626action binds a generic wrapper to an already deployed contract.
func bindDepositerc4626action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Depositerc4626actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositerc4626action *Depositerc4626actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositerc4626action.Contract.Depositerc4626actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositerc4626action *Depositerc4626actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.Depositerc4626actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositerc4626action *Depositerc4626actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.Depositerc4626actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositerc4626action *Depositerc4626actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositerc4626action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositerc4626action *Depositerc4626actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositerc4626action *Depositerc4626actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Depositerc4626action *Depositerc4626actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Depositerc4626action *Depositerc4626actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Depositerc4626action.Contract.ArgumentsHash(&_Depositerc4626action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Depositerc4626action *Depositerc4626actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Depositerc4626action.Contract.ArgumentsHash(&_Depositerc4626action.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Depositerc4626action *Depositerc4626actionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Depositerc4626action *Depositerc4626actionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Depositerc4626action.Contract.CheckInterval(&_Depositerc4626action.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Depositerc4626action *Depositerc4626actionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Depositerc4626action.Contract.CheckInterval(&_Depositerc4626action.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Depositerc4626action *Depositerc4626actionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Depositerc4626action *Depositerc4626actionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Depositerc4626action.Contract.CheckStart(&_Depositerc4626action.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Depositerc4626action *Depositerc4626actionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Depositerc4626action.Contract.CheckStart(&_Depositerc4626action.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Depositerc4626action *Depositerc4626actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Depositerc4626action *Depositerc4626actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Depositerc4626action.Contract.FeeTokenRegistry(&_Depositerc4626action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Depositerc4626action *Depositerc4626actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Depositerc4626action.Contract.FeeTokenRegistry(&_Depositerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Depositerc4626action *Depositerc4626actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Depositerc4626action *Depositerc4626actionSession) GasConstant() (*big.Int, error) {
	return _Depositerc4626action.Contract.GasConstant(&_Depositerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Depositerc4626action *Depositerc4626actionCallerSession) GasConstant() (*big.Int, error) {
	return _Depositerc4626action.Contract.GasConstant(&_Depositerc4626action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x9488e717.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCaller) Hash(opts *bind.CallOpts, arguments IDepositERC4626ActionDepositERC4626) ([32]byte, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x9488e717.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionSession) Hash(arguments IDepositERC4626ActionDepositERC4626) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash(&_Depositerc4626action.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x9488e717.
//
// Solidity: function hash((address,address,uint256,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCallerSession) Hash(arguments IDepositERC4626ActionDepositERC4626) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash(&_Depositerc4626action.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCaller) Hash0(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "hash0", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash0(&_Depositerc4626action.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCallerSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash0(&_Depositerc4626action.CallOpts, schedule)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash1(&_Depositerc4626action.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Depositerc4626action *Depositerc4626actionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Depositerc4626action.Contract.Hash1(&_Depositerc4626action.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Depositerc4626action *Depositerc4626actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositerc4626action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Depositerc4626action *Depositerc4626actionSession) Treasury() (common.Address, error) {
	return _Depositerc4626action.Contract.Treasury(&_Depositerc4626action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Depositerc4626action *Depositerc4626actionCallerSession) Treasury() (common.Address, error) {
	return _Depositerc4626action.Contract.Treasury(&_Depositerc4626action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Depositerc4626action *Depositerc4626actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Depositerc4626action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Depositerc4626action *Depositerc4626actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.ChargeFee(&_Depositerc4626action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Depositerc4626action *Depositerc4626actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.ChargeFee(&_Depositerc4626action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Depositerc4626action *Depositerc4626actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Depositerc4626action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Depositerc4626action *Depositerc4626actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.Execute(&_Depositerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Depositerc4626action *Depositerc4626actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Depositerc4626action.Contract.Execute(&_Depositerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Depositerc4626actionMaxDepositReachedIterator is returned from FilterMaxDepositReached and is used to iterate over the raw logs and unpacked data for MaxDepositReached events raised by the Depositerc4626action contract.
type Depositerc4626actionMaxDepositReachedIterator struct {
	Event *Depositerc4626actionMaxDepositReached // Event containing the contract specifics and raw log

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
func (it *Depositerc4626actionMaxDepositReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Depositerc4626actionMaxDepositReached)
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
		it.Event = new(Depositerc4626actionMaxDepositReached)
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
func (it *Depositerc4626actionMaxDepositReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Depositerc4626actionMaxDepositReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Depositerc4626actionMaxDepositReached represents a MaxDepositReached event raised by the Depositerc4626action contract.
type Depositerc4626actionMaxDepositReached struct {
	MaxDeposit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxDepositReached is a free log retrieval operation binding the contract event 0x1753f9cc55807c3f0e09975a71e3ca0b813c7f55ddf6996049c130fdce40f4ab.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit)
func (_Depositerc4626action *Depositerc4626actionFilterer) FilterMaxDepositReached(opts *bind.FilterOpts) (*Depositerc4626actionMaxDepositReachedIterator, error) {

	logs, sub, err := _Depositerc4626action.contract.FilterLogs(opts, "MaxDepositReached")
	if err != nil {
		return nil, err
	}
	return &Depositerc4626actionMaxDepositReachedIterator{contract: _Depositerc4626action.contract, event: "MaxDepositReached", logs: logs, sub: sub}, nil
}

// WatchMaxDepositReached is a free log subscription operation binding the contract event 0x1753f9cc55807c3f0e09975a71e3ca0b813c7f55ddf6996049c130fdce40f4ab.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit)
func (_Depositerc4626action *Depositerc4626actionFilterer) WatchMaxDepositReached(opts *bind.WatchOpts, sink chan<- *Depositerc4626actionMaxDepositReached) (event.Subscription, error) {

	logs, sub, err := _Depositerc4626action.contract.WatchLogs(opts, "MaxDepositReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Depositerc4626actionMaxDepositReached)
				if err := _Depositerc4626action.contract.UnpackLog(event, "MaxDepositReached", log); err != nil {
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

// ParseMaxDepositReached is a log parse operation binding the contract event 0x1753f9cc55807c3f0e09975a71e3ca0b813c7f55ddf6996049c130fdce40f4ab.
//
// Solidity: event MaxDepositReached(uint256 maxDeposit)
func (_Depositerc4626action *Depositerc4626actionFilterer) ParseMaxDepositReached(log types.Log) (*Depositerc4626actionMaxDepositReached, error) {
	event := new(Depositerc4626actionMaxDepositReached)
	if err := _Depositerc4626action.contract.UnpackLog(event, "MaxDepositReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
