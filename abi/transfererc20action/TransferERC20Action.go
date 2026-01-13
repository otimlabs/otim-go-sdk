// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transfererc20action

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

// ITransferERC20ActionTransferERC20 is an auto generated low-level Go binding around an user-defined struct.
type ITransferERC20ActionTransferERC20 struct {
	Token    common.Address
	Target   common.Address
	Value    *big.Int
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

// Transfererc20actionMetaData contains all meta data concerning the Transfererc20action contract.
var Transfererc20actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"transferERC20\",\"type\":\"tuple\",\"internalType\":\"structITransferERC20Action.TransferERC20\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// Transfererc20actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Transfererc20actionMetaData.ABI instead.
var Transfererc20actionABI = Transfererc20actionMetaData.ABI

// Transfererc20action is an auto generated Go binding around an Ethereum contract.
type Transfererc20action struct {
	Transfererc20actionCaller     // Read-only binding to the contract
	Transfererc20actionTransactor // Write-only binding to the contract
	Transfererc20actionFilterer   // Log filterer for contract events
}

// Transfererc20actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Transfererc20actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Transfererc20actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Transfererc20actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Transfererc20actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Transfererc20actionSession struct {
	Contract     *Transfererc20action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Transfererc20actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Transfererc20actionCallerSession struct {
	Contract *Transfererc20actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Transfererc20actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Transfererc20actionTransactorSession struct {
	Contract     *Transfererc20actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Transfererc20actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Transfererc20actionRaw struct {
	Contract *Transfererc20action // Generic contract binding to access the raw methods on
}

// Transfererc20actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Transfererc20actionCallerRaw struct {
	Contract *Transfererc20actionCaller // Generic read-only contract binding to access the raw methods on
}

// Transfererc20actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Transfererc20actionTransactorRaw struct {
	Contract *Transfererc20actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfererc20action creates a new instance of Transfererc20action, bound to a specific deployed contract.
func NewTransfererc20action(address common.Address, backend bind.ContractBackend) (*Transfererc20action, error) {
	contract, err := bindTransfererc20action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfererc20action{Transfererc20actionCaller: Transfererc20actionCaller{contract: contract}, Transfererc20actionTransactor: Transfererc20actionTransactor{contract: contract}, Transfererc20actionFilterer: Transfererc20actionFilterer{contract: contract}}, nil
}

// NewTransfererc20actionCaller creates a new read-only instance of Transfererc20action, bound to a specific deployed contract.
func NewTransfererc20actionCaller(address common.Address, caller bind.ContractCaller) (*Transfererc20actionCaller, error) {
	contract, err := bindTransfererc20action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Transfererc20actionCaller{contract: contract}, nil
}

// NewTransfererc20actionTransactor creates a new write-only instance of Transfererc20action, bound to a specific deployed contract.
func NewTransfererc20actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Transfererc20actionTransactor, error) {
	contract, err := bindTransfererc20action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Transfererc20actionTransactor{contract: contract}, nil
}

// NewTransfererc20actionFilterer creates a new log filterer instance of Transfererc20action, bound to a specific deployed contract.
func NewTransfererc20actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Transfererc20actionFilterer, error) {
	contract, err := bindTransfererc20action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Transfererc20actionFilterer{contract: contract}, nil
}

// bindTransfererc20action binds a generic wrapper to an already deployed contract.
func bindTransfererc20action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Transfererc20actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfererc20action *Transfererc20actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfererc20action.Contract.Transfererc20actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfererc20action *Transfererc20actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfererc20action.Contract.Transfererc20actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfererc20action *Transfererc20actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfererc20action.Contract.Transfererc20actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfererc20action *Transfererc20actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfererc20action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfererc20action *Transfererc20actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfererc20action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfererc20action *Transfererc20actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfererc20action.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfererc20action *Transfererc20actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Transfererc20action *Transfererc20actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfererc20action.Contract.ArgumentsHash(&_Transfererc20action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfererc20action *Transfererc20actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfererc20action.Contract.ArgumentsHash(&_Transfererc20action.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfererc20action *Transfererc20actionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfererc20action *Transfererc20actionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transfererc20action.Contract.CheckInterval(&_Transfererc20action.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfererc20action *Transfererc20actionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transfererc20action.Contract.CheckInterval(&_Transfererc20action.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfererc20action *Transfererc20actionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfererc20action *Transfererc20actionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transfererc20action.Contract.CheckStart(&_Transfererc20action.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfererc20action *Transfererc20actionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transfererc20action.Contract.CheckStart(&_Transfererc20action.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20action *Transfererc20actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20action *Transfererc20actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfererc20action.Contract.FeeTokenRegistry(&_Transfererc20action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfererc20action *Transfererc20actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfererc20action.Contract.FeeTokenRegistry(&_Transfererc20action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20action *Transfererc20actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20action *Transfererc20actionSession) GasConstant() (*big.Int, error) {
	return _Transfererc20action.Contract.GasConstant(&_Transfererc20action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfererc20action *Transfererc20actionCallerSession) GasConstant() (*big.Int, error) {
	return _Transfererc20action.Contract.GasConstant(&_Transfererc20action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x19d42734.
//
// Solidity: function hash((address,address,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transferERC20) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCaller) Hash(opts *bind.CallOpts, transferERC20 ITransferERC20ActionTransferERC20) ([32]byte, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "hash", transferERC20)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x19d42734.
//
// Solidity: function hash((address,address,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transferERC20) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionSession) Hash(transferERC20 ITransferERC20ActionTransferERC20) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash(&_Transfererc20action.CallOpts, transferERC20)
}

// Hash is a free data retrieval call binding the contract method 0x19d42734.
//
// Solidity: function hash((address,address,uint256,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) transferERC20) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCallerSession) Hash(transferERC20 ITransferERC20ActionTransferERC20) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash(&_Transfererc20action.CallOpts, transferERC20)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCaller) Hash0(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "hash0", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash0(&_Transfererc20action.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCallerSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash0(&_Transfererc20action.CallOpts, schedule)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash1(&_Transfererc20action.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfererc20action *Transfererc20actionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfererc20action.Contract.Hash1(&_Transfererc20action.CallOpts, fee)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20action *Transfererc20actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfererc20action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20action *Transfererc20actionSession) Treasury() (common.Address, error) {
	return _Transfererc20action.Contract.Treasury(&_Transfererc20action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfererc20action *Transfererc20actionCallerSession) Treasury() (common.Address, error) {
	return _Transfererc20action.Contract.Treasury(&_Transfererc20action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20action *Transfererc20actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20action *Transfererc20actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20action.Contract.ChargeFee(&_Transfererc20action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfererc20action *Transfererc20actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfererc20action.Contract.ChargeFee(&_Transfererc20action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfererc20action *Transfererc20actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfererc20action *Transfererc20actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20action.Contract.Execute(&_Transfererc20action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfererc20action *Transfererc20actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfererc20action.Contract.Execute(&_Transfererc20action.TransactOpts, instruction, arg1, executionState)
}
