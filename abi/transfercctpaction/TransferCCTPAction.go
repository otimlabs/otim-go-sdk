// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transfercctpaction

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

// ITransferCCTPActionTransferCCTP is an auto generated low-level Go binding around an user-defined struct.
type ITransferCCTPActionTransferCCTP struct {
	Token                    common.Address
	Amount                   *big.Int
	DestinationDomain        uint32
	DestinationMintRecipient [32]byte
	Schedule                 IIntervalSchedule
	Fee                      IOtimFeeFee
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

// TransfercctpactionMetaData contains all meta data concerning the Transfercctpaction contract.
var TransfercctpactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessengerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenMinterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structITransferCCTPAction.TransferCCTP\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destinationMintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMinter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCTPBurnLimitReached\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"burnLimitPerMessage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCTPTokenNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// TransfercctpactionABI is the input ABI used to generate the binding from.
// Deprecated: Use TransfercctpactionMetaData.ABI instead.
var TransfercctpactionABI = TransfercctpactionMetaData.ABI

// Transfercctpaction is an auto generated Go binding around an Ethereum contract.
type Transfercctpaction struct {
	TransfercctpactionCaller     // Read-only binding to the contract
	TransfercctpactionTransactor // Write-only binding to the contract
	TransfercctpactionFilterer   // Log filterer for contract events
}

// TransfercctpactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransfercctpactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransfercctpactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransfercctpactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransfercctpactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransfercctpactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransfercctpactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransfercctpactionSession struct {
	Contract     *Transfercctpaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TransfercctpactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransfercctpactionCallerSession struct {
	Contract *TransfercctpactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TransfercctpactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransfercctpactionTransactorSession struct {
	Contract     *TransfercctpactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TransfercctpactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransfercctpactionRaw struct {
	Contract *Transfercctpaction // Generic contract binding to access the raw methods on
}

// TransfercctpactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransfercctpactionCallerRaw struct {
	Contract *TransfercctpactionCaller // Generic read-only contract binding to access the raw methods on
}

// TransfercctpactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransfercctpactionTransactorRaw struct {
	Contract *TransfercctpactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransfercctpaction creates a new instance of Transfercctpaction, bound to a specific deployed contract.
func NewTransfercctpaction(address common.Address, backend bind.ContractBackend) (*Transfercctpaction, error) {
	contract, err := bindTransfercctpaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transfercctpaction{TransfercctpactionCaller: TransfercctpactionCaller{contract: contract}, TransfercctpactionTransactor: TransfercctpactionTransactor{contract: contract}, TransfercctpactionFilterer: TransfercctpactionFilterer{contract: contract}}, nil
}

// NewTransfercctpactionCaller creates a new read-only instance of Transfercctpaction, bound to a specific deployed contract.
func NewTransfercctpactionCaller(address common.Address, caller bind.ContractCaller) (*TransfercctpactionCaller, error) {
	contract, err := bindTransfercctpaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransfercctpactionCaller{contract: contract}, nil
}

// NewTransfercctpactionTransactor creates a new write-only instance of Transfercctpaction, bound to a specific deployed contract.
func NewTransfercctpactionTransactor(address common.Address, transactor bind.ContractTransactor) (*TransfercctpactionTransactor, error) {
	contract, err := bindTransfercctpaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransfercctpactionTransactor{contract: contract}, nil
}

// NewTransfercctpactionFilterer creates a new log filterer instance of Transfercctpaction, bound to a specific deployed contract.
func NewTransfercctpactionFilterer(address common.Address, filterer bind.ContractFilterer) (*TransfercctpactionFilterer, error) {
	contract, err := bindTransfercctpaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransfercctpactionFilterer{contract: contract}, nil
}

// bindTransfercctpaction binds a generic wrapper to an already deployed contract.
func bindTransfercctpaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransfercctpactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfercctpaction *TransfercctpactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfercctpaction.Contract.TransfercctpactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfercctpaction *TransfercctpactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.TransfercctpactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfercctpaction *TransfercctpactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.TransfercctpactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transfercctpaction *TransfercctpactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Transfercctpaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transfercctpaction *TransfercctpactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transfercctpaction *TransfercctpactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfercctpaction *TransfercctpactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Transfercctpaction *TransfercctpactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfercctpaction.Contract.ArgumentsHash(&_Transfercctpaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Transfercctpaction *TransfercctpactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Transfercctpaction.Contract.ArgumentsHash(&_Transfercctpaction.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfercctpaction *TransfercctpactionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfercctpaction *TransfercctpactionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transfercctpaction.Contract.CheckInterval(&_Transfercctpaction.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Transfercctpaction *TransfercctpactionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Transfercctpaction.Contract.CheckInterval(&_Transfercctpaction.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfercctpaction *TransfercctpactionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfercctpaction *TransfercctpactionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transfercctpaction.Contract.CheckStart(&_Transfercctpaction.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Transfercctpaction *TransfercctpactionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Transfercctpaction.Contract.CheckStart(&_Transfercctpaction.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfercctpaction *TransfercctpactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfercctpaction *TransfercctpactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfercctpaction.Contract.FeeTokenRegistry(&_Transfercctpaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Transfercctpaction *TransfercctpactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Transfercctpaction.Contract.FeeTokenRegistry(&_Transfercctpaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfercctpaction *TransfercctpactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfercctpaction *TransfercctpactionSession) GasConstant() (*big.Int, error) {
	return _Transfercctpaction.Contract.GasConstant(&_Transfercctpaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Transfercctpaction *TransfercctpactionCallerSession) GasConstant() (*big.Int, error) {
	return _Transfercctpaction.Contract.GasConstant(&_Transfercctpaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCaller) Hash(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "hash", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionSession) Hash(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash(&_Transfercctpaction.CallOpts, schedule)
}

// Hash is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCallerSession) Hash(schedule IIntervalSchedule) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash(&_Transfercctpaction.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xdd5f3318.
//
// Solidity: function hash((address,uint256,uint32,bytes32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCaller) Hash0(opts *bind.CallOpts, arguments ITransferCCTPActionTransferCCTP) ([32]byte, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "hash0", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xdd5f3318.
//
// Solidity: function hash((address,uint256,uint32,bytes32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionSession) Hash0(arguments ITransferCCTPActionTransferCCTP) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash0(&_Transfercctpaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xdd5f3318.
//
// Solidity: function hash((address,uint256,uint32,bytes32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCallerSession) Hash0(arguments ITransferCCTPActionTransferCCTP) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash0(&_Transfercctpaction.CallOpts, arguments)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash1(&_Transfercctpaction.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Transfercctpaction *TransfercctpactionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Transfercctpaction.Contract.Hash1(&_Transfercctpaction.CallOpts, fee)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Transfercctpaction *TransfercctpactionCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Transfercctpaction *TransfercctpactionSession) TokenMessenger() (common.Address, error) {
	return _Transfercctpaction.Contract.TokenMessenger(&_Transfercctpaction.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Transfercctpaction *TransfercctpactionCallerSession) TokenMessenger() (common.Address, error) {
	return _Transfercctpaction.Contract.TokenMessenger(&_Transfercctpaction.CallOpts)
}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Transfercctpaction *TransfercctpactionCaller) TokenMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "tokenMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Transfercctpaction *TransfercctpactionSession) TokenMinter() (common.Address, error) {
	return _Transfercctpaction.Contract.TokenMinter(&_Transfercctpaction.CallOpts)
}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Transfercctpaction *TransfercctpactionCallerSession) TokenMinter() (common.Address, error) {
	return _Transfercctpaction.Contract.TokenMinter(&_Transfercctpaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfercctpaction *TransfercctpactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Transfercctpaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfercctpaction *TransfercctpactionSession) Treasury() (common.Address, error) {
	return _Transfercctpaction.Contract.Treasury(&_Transfercctpaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Transfercctpaction *TransfercctpactionCallerSession) Treasury() (common.Address, error) {
	return _Transfercctpaction.Contract.Treasury(&_Transfercctpaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfercctpaction *TransfercctpactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfercctpaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfercctpaction *TransfercctpactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.ChargeFee(&_Transfercctpaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Transfercctpaction *TransfercctpactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.ChargeFee(&_Transfercctpaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfercctpaction *TransfercctpactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfercctpaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfercctpaction *TransfercctpactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.Execute(&_Transfercctpaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Transfercctpaction *TransfercctpactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Transfercctpaction.Contract.Execute(&_Transfercctpaction.TransactOpts, instruction, arg1, executionState)
}

// TransfercctpactionCCTPBurnLimitReachedIterator is returned from FilterCCTPBurnLimitReached and is used to iterate over the raw logs and unpacked data for CCTPBurnLimitReached events raised by the Transfercctpaction contract.
type TransfercctpactionCCTPBurnLimitReachedIterator struct {
	Event *TransfercctpactionCCTPBurnLimitReached // Event containing the contract specifics and raw log

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
func (it *TransfercctpactionCCTPBurnLimitReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransfercctpactionCCTPBurnLimitReached)
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
		it.Event = new(TransfercctpactionCCTPBurnLimitReached)
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
func (it *TransfercctpactionCCTPBurnLimitReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransfercctpactionCCTPBurnLimitReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransfercctpactionCCTPBurnLimitReached represents a CCTPBurnLimitReached event raised by the Transfercctpaction contract.
type TransfercctpactionCCTPBurnLimitReached struct {
	Token               common.Address
	BurnLimitPerMessage *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCCTPBurnLimitReached is a free log retrieval operation binding the contract event 0x760af14bc606afee102f609faf989edd3f0ab0cd6e7739387b3138df52bf7d3a.
//
// Solidity: event CCTPBurnLimitReached(address indexed token, uint256 burnLimitPerMessage)
func (_Transfercctpaction *TransfercctpactionFilterer) FilterCCTPBurnLimitReached(opts *bind.FilterOpts, token []common.Address) (*TransfercctpactionCCTPBurnLimitReachedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Transfercctpaction.contract.FilterLogs(opts, "CCTPBurnLimitReached", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TransfercctpactionCCTPBurnLimitReachedIterator{contract: _Transfercctpaction.contract, event: "CCTPBurnLimitReached", logs: logs, sub: sub}, nil
}

// WatchCCTPBurnLimitReached is a free log subscription operation binding the contract event 0x760af14bc606afee102f609faf989edd3f0ab0cd6e7739387b3138df52bf7d3a.
//
// Solidity: event CCTPBurnLimitReached(address indexed token, uint256 burnLimitPerMessage)
func (_Transfercctpaction *TransfercctpactionFilterer) WatchCCTPBurnLimitReached(opts *bind.WatchOpts, sink chan<- *TransfercctpactionCCTPBurnLimitReached, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Transfercctpaction.contract.WatchLogs(opts, "CCTPBurnLimitReached", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransfercctpactionCCTPBurnLimitReached)
				if err := _Transfercctpaction.contract.UnpackLog(event, "CCTPBurnLimitReached", log); err != nil {
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

// ParseCCTPBurnLimitReached is a log parse operation binding the contract event 0x760af14bc606afee102f609faf989edd3f0ab0cd6e7739387b3138df52bf7d3a.
//
// Solidity: event CCTPBurnLimitReached(address indexed token, uint256 burnLimitPerMessage)
func (_Transfercctpaction *TransfercctpactionFilterer) ParseCCTPBurnLimitReached(log types.Log) (*TransfercctpactionCCTPBurnLimitReached, error) {
	event := new(TransfercctpactionCCTPBurnLimitReached)
	if err := _Transfercctpaction.contract.UnpackLog(event, "CCTPBurnLimitReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
