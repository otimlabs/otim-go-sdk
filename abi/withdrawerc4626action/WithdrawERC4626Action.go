// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawerc4626action

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

// IWithdrawERC4626ActionWithdrawERC4626 is an auto generated low-level Go binding around an user-defined struct.
type IWithdrawERC4626ActionWithdrawERC4626 struct {
	Vault       common.Address
	Recipient   common.Address
	Value       *big.Int
	MinWithdraw *big.Int
	Schedule    IIntervalSchedule
	Fee         IOtimFeeFee
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

// Withdrawerc4626actionMetaData contains all meta data concerning the Withdrawerc4626action contract.
var Withdrawerc4626actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structIWithdrawERC4626Action.WithdrawERC4626\",\"components\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minWithdraw\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MaxWithdrawReached\",\"inputs\":[{\"name\":\"maxWithdraw\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxWithdrawTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// Withdrawerc4626actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Withdrawerc4626actionMetaData.ABI instead.
var Withdrawerc4626actionABI = Withdrawerc4626actionMetaData.ABI

// Withdrawerc4626action is an auto generated Go binding around an Ethereum contract.
type Withdrawerc4626action struct {
	Withdrawerc4626actionCaller     // Read-only binding to the contract
	Withdrawerc4626actionTransactor // Write-only binding to the contract
	Withdrawerc4626actionFilterer   // Log filterer for contract events
}

// Withdrawerc4626actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Withdrawerc4626actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Withdrawerc4626actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Withdrawerc4626actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Withdrawerc4626actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Withdrawerc4626actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Withdrawerc4626actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Withdrawerc4626actionSession struct {
	Contract     *Withdrawerc4626action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Withdrawerc4626actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Withdrawerc4626actionCallerSession struct {
	Contract *Withdrawerc4626actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Withdrawerc4626actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Withdrawerc4626actionTransactorSession struct {
	Contract     *Withdrawerc4626actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Withdrawerc4626actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Withdrawerc4626actionRaw struct {
	Contract *Withdrawerc4626action // Generic contract binding to access the raw methods on
}

// Withdrawerc4626actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Withdrawerc4626actionCallerRaw struct {
	Contract *Withdrawerc4626actionCaller // Generic read-only contract binding to access the raw methods on
}

// Withdrawerc4626actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Withdrawerc4626actionTransactorRaw struct {
	Contract *Withdrawerc4626actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawerc4626action creates a new instance of Withdrawerc4626action, bound to a specific deployed contract.
func NewWithdrawerc4626action(address common.Address, backend bind.ContractBackend) (*Withdrawerc4626action, error) {
	contract, err := bindWithdrawerc4626action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawerc4626action{Withdrawerc4626actionCaller: Withdrawerc4626actionCaller{contract: contract}, Withdrawerc4626actionTransactor: Withdrawerc4626actionTransactor{contract: contract}, Withdrawerc4626actionFilterer: Withdrawerc4626actionFilterer{contract: contract}}, nil
}

// NewWithdrawerc4626actionCaller creates a new read-only instance of Withdrawerc4626action, bound to a specific deployed contract.
func NewWithdrawerc4626actionCaller(address common.Address, caller bind.ContractCaller) (*Withdrawerc4626actionCaller, error) {
	contract, err := bindWithdrawerc4626action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Withdrawerc4626actionCaller{contract: contract}, nil
}

// NewWithdrawerc4626actionTransactor creates a new write-only instance of Withdrawerc4626action, bound to a specific deployed contract.
func NewWithdrawerc4626actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Withdrawerc4626actionTransactor, error) {
	contract, err := bindWithdrawerc4626action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Withdrawerc4626actionTransactor{contract: contract}, nil
}

// NewWithdrawerc4626actionFilterer creates a new log filterer instance of Withdrawerc4626action, bound to a specific deployed contract.
func NewWithdrawerc4626actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Withdrawerc4626actionFilterer, error) {
	contract, err := bindWithdrawerc4626action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Withdrawerc4626actionFilterer{contract: contract}, nil
}

// bindWithdrawerc4626action binds a generic wrapper to an already deployed contract.
func bindWithdrawerc4626action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Withdrawerc4626actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawerc4626action *Withdrawerc4626actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawerc4626action.Contract.Withdrawerc4626actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawerc4626action *Withdrawerc4626actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.Withdrawerc4626actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawerc4626action *Withdrawerc4626actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.Withdrawerc4626actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawerc4626action *Withdrawerc4626actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawerc4626action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawerc4626action *Withdrawerc4626actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawerc4626action *Withdrawerc4626actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Withdrawerc4626action *Withdrawerc4626actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Withdrawerc4626action.Contract.ArgumentsHash(&_Withdrawerc4626action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Withdrawerc4626action.Contract.ArgumentsHash(&_Withdrawerc4626action.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Withdrawerc4626action.Contract.CheckInterval(&_Withdrawerc4626action.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Withdrawerc4626action.Contract.CheckInterval(&_Withdrawerc4626action.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Withdrawerc4626action.Contract.CheckStart(&_Withdrawerc4626action.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Withdrawerc4626action.Contract.CheckStart(&_Withdrawerc4626action.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Withdrawerc4626action.Contract.FeeTokenRegistry(&_Withdrawerc4626action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Withdrawerc4626action.Contract.FeeTokenRegistry(&_Withdrawerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) GasConstant() (*big.Int, error) {
	return _Withdrawerc4626action.Contract.GasConstant(&_Withdrawerc4626action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) GasConstant() (*big.Int, error) {
	return _Withdrawerc4626action.Contract.GasConstant(&_Withdrawerc4626action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x129044a3.
//
// Solidity: function hash((address,address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) Hash(opts *bind.CallOpts, arguments IWithdrawERC4626ActionWithdrawERC4626) ([32]byte, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x129044a3.
//
// Solidity: function hash((address,address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) Hash(arguments IWithdrawERC4626ActionWithdrawERC4626) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash(&_Withdrawerc4626action.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x129044a3.
//
// Solidity: function hash((address,address,uint256,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) Hash(arguments IWithdrawERC4626ActionWithdrawERC4626) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash(&_Withdrawerc4626action.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) Hash0(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "hash0", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash0(&_Withdrawerc4626action.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash0(&_Withdrawerc4626action.CallOpts, schedule)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash1(&_Withdrawerc4626action.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Withdrawerc4626action.Contract.Hash1(&_Withdrawerc4626action.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawerc4626action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) Treasury() (common.Address, error) {
	return _Withdrawerc4626action.Contract.Treasury(&_Withdrawerc4626action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Withdrawerc4626action *Withdrawerc4626actionCallerSession) Treasury() (common.Address, error) {
	return _Withdrawerc4626action.Contract.Treasury(&_Withdrawerc4626action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Withdrawerc4626action *Withdrawerc4626actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Withdrawerc4626action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Withdrawerc4626action *Withdrawerc4626actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.ChargeFee(&_Withdrawerc4626action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Withdrawerc4626action *Withdrawerc4626actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.ChargeFee(&_Withdrawerc4626action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Withdrawerc4626action *Withdrawerc4626actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Withdrawerc4626action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Withdrawerc4626action *Withdrawerc4626actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.Execute(&_Withdrawerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Withdrawerc4626action *Withdrawerc4626actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Withdrawerc4626action.Contract.Execute(&_Withdrawerc4626action.TransactOpts, instruction, arg1, executionState)
}

// Withdrawerc4626actionMaxWithdrawReachedIterator is returned from FilterMaxWithdrawReached and is used to iterate over the raw logs and unpacked data for MaxWithdrawReached events raised by the Withdrawerc4626action contract.
type Withdrawerc4626actionMaxWithdrawReachedIterator struct {
	Event *Withdrawerc4626actionMaxWithdrawReached // Event containing the contract specifics and raw log

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
func (it *Withdrawerc4626actionMaxWithdrawReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Withdrawerc4626actionMaxWithdrawReached)
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
		it.Event = new(Withdrawerc4626actionMaxWithdrawReached)
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
func (it *Withdrawerc4626actionMaxWithdrawReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Withdrawerc4626actionMaxWithdrawReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Withdrawerc4626actionMaxWithdrawReached represents a MaxWithdrawReached event raised by the Withdrawerc4626action contract.
type Withdrawerc4626actionMaxWithdrawReached struct {
	MaxWithdraw *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxWithdrawReached is a free log retrieval operation binding the contract event 0x8f4f8d0ea98e816c0e7964aafe1a600f28ca53d3bc05007daf746eae761a46cc.
//
// Solidity: event MaxWithdrawReached(uint256 maxWithdraw)
func (_Withdrawerc4626action *Withdrawerc4626actionFilterer) FilterMaxWithdrawReached(opts *bind.FilterOpts) (*Withdrawerc4626actionMaxWithdrawReachedIterator, error) {

	logs, sub, err := _Withdrawerc4626action.contract.FilterLogs(opts, "MaxWithdrawReached")
	if err != nil {
		return nil, err
	}
	return &Withdrawerc4626actionMaxWithdrawReachedIterator{contract: _Withdrawerc4626action.contract, event: "MaxWithdrawReached", logs: logs, sub: sub}, nil
}

// WatchMaxWithdrawReached is a free log subscription operation binding the contract event 0x8f4f8d0ea98e816c0e7964aafe1a600f28ca53d3bc05007daf746eae761a46cc.
//
// Solidity: event MaxWithdrawReached(uint256 maxWithdraw)
func (_Withdrawerc4626action *Withdrawerc4626actionFilterer) WatchMaxWithdrawReached(opts *bind.WatchOpts, sink chan<- *Withdrawerc4626actionMaxWithdrawReached) (event.Subscription, error) {

	logs, sub, err := _Withdrawerc4626action.contract.WatchLogs(opts, "MaxWithdrawReached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Withdrawerc4626actionMaxWithdrawReached)
				if err := _Withdrawerc4626action.contract.UnpackLog(event, "MaxWithdrawReached", log); err != nil {
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

// ParseMaxWithdrawReached is a log parse operation binding the contract event 0x8f4f8d0ea98e816c0e7964aafe1a600f28ca53d3bc05007daf746eae761a46cc.
//
// Solidity: event MaxWithdrawReached(uint256 maxWithdraw)
func (_Withdrawerc4626action *Withdrawerc4626actionFilterer) ParseMaxWithdrawReached(log types.Log) (*Withdrawerc4626actionMaxWithdrawReached, error) {
	event := new(Withdrawerc4626actionMaxWithdrawReached)
	if err := _Withdrawerc4626action.contract.UnpackLog(event, "MaxWithdrawReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
