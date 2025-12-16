// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sweepcctpaction

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

// ISweepCCTPActionSweepCCTP is an auto generated low-level Go binding around an user-defined struct.
type ISweepCCTPActionSweepCCTP struct {
	Token                    common.Address
	DestinationDomain        uint32
	DestinationMintRecipient [32]byte
	Threshold                *big.Int
	EndBalance               *big.Int
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

// SweepcctpactionMetaData contains all meta data concerning the Sweepcctpaction contract.
var SweepcctpactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessengerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenMinterAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structISweepCCTPAction.SweepCCTP\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destinationMintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokenMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenMinter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITokenController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCTPBurnLimitReached\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"burnLimitPerMessage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceUnderThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCTPTokenNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// SweepcctpactionABI is the input ABI used to generate the binding from.
// Deprecated: Use SweepcctpactionMetaData.ABI instead.
var SweepcctpactionABI = SweepcctpactionMetaData.ABI

// Sweepcctpaction is an auto generated Go binding around an Ethereum contract.
type Sweepcctpaction struct {
	SweepcctpactionCaller     // Read-only binding to the contract
	SweepcctpactionTransactor // Write-only binding to the contract
	SweepcctpactionFilterer   // Log filterer for contract events
}

// SweepcctpactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweepcctpactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepcctpactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweepcctpactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepcctpactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SweepcctpactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweepcctpactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweepcctpactionSession struct {
	Contract     *Sweepcctpaction  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweepcctpactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweepcctpactionCallerSession struct {
	Contract *SweepcctpactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SweepcctpactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweepcctpactionTransactorSession struct {
	Contract     *SweepcctpactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SweepcctpactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweepcctpactionRaw struct {
	Contract *Sweepcctpaction // Generic contract binding to access the raw methods on
}

// SweepcctpactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweepcctpactionCallerRaw struct {
	Contract *SweepcctpactionCaller // Generic read-only contract binding to access the raw methods on
}

// SweepcctpactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweepcctpactionTransactorRaw struct {
	Contract *SweepcctpactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweepcctpaction creates a new instance of Sweepcctpaction, bound to a specific deployed contract.
func NewSweepcctpaction(address common.Address, backend bind.ContractBackend) (*Sweepcctpaction, error) {
	contract, err := bindSweepcctpaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sweepcctpaction{SweepcctpactionCaller: SweepcctpactionCaller{contract: contract}, SweepcctpactionTransactor: SweepcctpactionTransactor{contract: contract}, SweepcctpactionFilterer: SweepcctpactionFilterer{contract: contract}}, nil
}

// NewSweepcctpactionCaller creates a new read-only instance of Sweepcctpaction, bound to a specific deployed contract.
func NewSweepcctpactionCaller(address common.Address, caller bind.ContractCaller) (*SweepcctpactionCaller, error) {
	contract, err := bindSweepcctpaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SweepcctpactionCaller{contract: contract}, nil
}

// NewSweepcctpactionTransactor creates a new write-only instance of Sweepcctpaction, bound to a specific deployed contract.
func NewSweepcctpactionTransactor(address common.Address, transactor bind.ContractTransactor) (*SweepcctpactionTransactor, error) {
	contract, err := bindSweepcctpaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SweepcctpactionTransactor{contract: contract}, nil
}

// NewSweepcctpactionFilterer creates a new log filterer instance of Sweepcctpaction, bound to a specific deployed contract.
func NewSweepcctpactionFilterer(address common.Address, filterer bind.ContractFilterer) (*SweepcctpactionFilterer, error) {
	contract, err := bindSweepcctpaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SweepcctpactionFilterer{contract: contract}, nil
}

// bindSweepcctpaction binds a generic wrapper to an already deployed contract.
func bindSweepcctpaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SweepcctpactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepcctpaction *SweepcctpactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepcctpaction.Contract.SweepcctpactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepcctpaction *SweepcctpactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.SweepcctpactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepcctpaction *SweepcctpactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.SweepcctpactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepcctpaction *SweepcctpactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepcctpaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepcctpaction *SweepcctpactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepcctpaction *SweepcctpactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.contract.Transact(opts, method, params...)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepcctpaction *SweepcctpactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Sweepcctpaction *SweepcctpactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepcctpaction.Contract.ArgumentsHash(&_Sweepcctpaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepcctpaction *SweepcctpactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepcctpaction.Contract.ArgumentsHash(&_Sweepcctpaction.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepcctpaction *SweepcctpactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepcctpaction.Contract.FeeTokenRegistry(&_Sweepcctpaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepcctpaction.Contract.FeeTokenRegistry(&_Sweepcctpaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepcctpaction *SweepcctpactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepcctpaction *SweepcctpactionSession) GasConstant() (*big.Int, error) {
	return _Sweepcctpaction.Contract.GasConstant(&_Sweepcctpaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepcctpaction *SweepcctpactionCallerSession) GasConstant() (*big.Int, error) {
	return _Sweepcctpaction.Contract.GasConstant(&_Sweepcctpaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x9ce4f4df.
//
// Solidity: function hash((address,uint32,bytes32,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionCaller) Hash(opts *bind.CallOpts, arguments ISweepCCTPActionSweepCCTP) ([32]byte, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x9ce4f4df.
//
// Solidity: function hash((address,uint32,bytes32,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionSession) Hash(arguments ISweepCCTPActionSweepCCTP) ([32]byte, error) {
	return _Sweepcctpaction.Contract.Hash(&_Sweepcctpaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x9ce4f4df.
//
// Solidity: function hash((address,uint32,bytes32,uint256,uint256,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionCallerSession) Hash(arguments ISweepCCTPActionSweepCCTP) ([32]byte, error) {
	return _Sweepcctpaction.Contract.Hash(&_Sweepcctpaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepcctpaction.Contract.Hash0(&_Sweepcctpaction.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepcctpaction *SweepcctpactionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepcctpaction.Contract.Hash0(&_Sweepcctpaction.CallOpts, fee)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Sweepcctpaction *SweepcctpactionSession) TokenMessenger() (common.Address, error) {
	return _Sweepcctpaction.Contract.TokenMessenger(&_Sweepcctpaction.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCallerSession) TokenMessenger() (common.Address, error) {
	return _Sweepcctpaction.Contract.TokenMessenger(&_Sweepcctpaction.CallOpts)
}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCaller) TokenMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "tokenMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Sweepcctpaction *SweepcctpactionSession) TokenMinter() (common.Address, error) {
	return _Sweepcctpaction.Contract.TokenMinter(&_Sweepcctpaction.CallOpts)
}

// TokenMinter is a free data retrieval call binding the contract method 0xcfb3647b.
//
// Solidity: function tokenMinter() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCallerSession) TokenMinter() (common.Address, error) {
	return _Sweepcctpaction.Contract.TokenMinter(&_Sweepcctpaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepcctpaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepcctpaction *SweepcctpactionSession) Treasury() (common.Address, error) {
	return _Sweepcctpaction.Contract.Treasury(&_Sweepcctpaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepcctpaction *SweepcctpactionCallerSession) Treasury() (common.Address, error) {
	return _Sweepcctpaction.Contract.Treasury(&_Sweepcctpaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepcctpaction *SweepcctpactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepcctpaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepcctpaction *SweepcctpactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.ChargeFee(&_Sweepcctpaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepcctpaction *SweepcctpactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.ChargeFee(&_Sweepcctpaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepcctpaction *SweepcctpactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepcctpaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepcctpaction *SweepcctpactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.Execute(&_Sweepcctpaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepcctpaction *SweepcctpactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepcctpaction.Contract.Execute(&_Sweepcctpaction.TransactOpts, instruction, arg1, executionState)
}

// SweepcctpactionCCTPBurnLimitReachedIterator is returned from FilterCCTPBurnLimitReached and is used to iterate over the raw logs and unpacked data for CCTPBurnLimitReached events raised by the Sweepcctpaction contract.
type SweepcctpactionCCTPBurnLimitReachedIterator struct {
	Event *SweepcctpactionCCTPBurnLimitReached // Event containing the contract specifics and raw log

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
func (it *SweepcctpactionCCTPBurnLimitReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SweepcctpactionCCTPBurnLimitReached)
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
		it.Event = new(SweepcctpactionCCTPBurnLimitReached)
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
func (it *SweepcctpactionCCTPBurnLimitReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SweepcctpactionCCTPBurnLimitReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SweepcctpactionCCTPBurnLimitReached represents a CCTPBurnLimitReached event raised by the Sweepcctpaction contract.
type SweepcctpactionCCTPBurnLimitReached struct {
	Token               common.Address
	BurnLimitPerMessage *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCCTPBurnLimitReached is a free log retrieval operation binding the contract event 0x760af14bc606afee102f609faf989edd3f0ab0cd6e7739387b3138df52bf7d3a.
//
// Solidity: event CCTPBurnLimitReached(address indexed token, uint256 burnLimitPerMessage)
func (_Sweepcctpaction *SweepcctpactionFilterer) FilterCCTPBurnLimitReached(opts *bind.FilterOpts, token []common.Address) (*SweepcctpactionCCTPBurnLimitReachedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sweepcctpaction.contract.FilterLogs(opts, "CCTPBurnLimitReached", tokenRule)
	if err != nil {
		return nil, err
	}
	return &SweepcctpactionCCTPBurnLimitReachedIterator{contract: _Sweepcctpaction.contract, event: "CCTPBurnLimitReached", logs: logs, sub: sub}, nil
}

// WatchCCTPBurnLimitReached is a free log subscription operation binding the contract event 0x760af14bc606afee102f609faf989edd3f0ab0cd6e7739387b3138df52bf7d3a.
//
// Solidity: event CCTPBurnLimitReached(address indexed token, uint256 burnLimitPerMessage)
func (_Sweepcctpaction *SweepcctpactionFilterer) WatchCCTPBurnLimitReached(opts *bind.WatchOpts, sink chan<- *SweepcctpactionCCTPBurnLimitReached, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Sweepcctpaction.contract.WatchLogs(opts, "CCTPBurnLimitReached", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SweepcctpactionCCTPBurnLimitReached)
				if err := _Sweepcctpaction.contract.UnpackLog(event, "CCTPBurnLimitReached", log); err != nil {
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
func (_Sweepcctpaction *SweepcctpactionFilterer) ParseCCTPBurnLimitReached(log types.Log) (*SweepcctpactionCCTPBurnLimitReached, error) {
	event := new(SweepcctpactionCCTPBurnLimitReached)
	if err := _Sweepcctpaction.contract.UnpackLog(event, "CCTPBurnLimitReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
