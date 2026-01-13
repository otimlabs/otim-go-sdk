// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sweepuniswapv3action

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

// ISweepUniswapV3ActionSweepUniswapV3 is an auto generated low-level Go binding around an user-defined struct.
type ISweepUniswapV3ActionSweepUniswapV3 struct {
	Recipient            common.Address
	TokenIn              common.Address
	TokenOut             common.Address
	FeeTier              *big.Int
	Threshold            *big.Int
	EndBalance           *big.Int
	FloorAmountOut       *big.Int
	MeanPriceLookBack    uint32
	MaxPriceDeviationBPS uint32
	Fee                  IOtimFeeFee
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

// Sweepuniswapv3actionMetaData contains all meta data concerning the Sweepuniswapv3action contract.
var Sweepuniswapv3actionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"routerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factoryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wethAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ERC20_TO_ERC20_COMMAND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_TO_ETH_COMMANDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETH_TO_ERC20_COMMANDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structISweepUniswapV3Action.SweepUniswapV3\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTier\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"floorAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"meanPriceLookBack\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPriceDeviationBPS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniversalRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uniswapV3Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniswapV3Factory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BalanceUnderThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"T\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UniswapV3PoolDoesNotExist\",\"inputs\":[]}]",
}

// Sweepuniswapv3actionABI is the input ABI used to generate the binding from.
// Deprecated: Use Sweepuniswapv3actionMetaData.ABI instead.
var Sweepuniswapv3actionABI = Sweepuniswapv3actionMetaData.ABI

// Sweepuniswapv3action is an auto generated Go binding around an Ethereum contract.
type Sweepuniswapv3action struct {
	Sweepuniswapv3actionCaller     // Read-only binding to the contract
	Sweepuniswapv3actionTransactor // Write-only binding to the contract
	Sweepuniswapv3actionFilterer   // Log filterer for contract events
}

// Sweepuniswapv3actionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Sweepuniswapv3actionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepuniswapv3actionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Sweepuniswapv3actionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepuniswapv3actionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Sweepuniswapv3actionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Sweepuniswapv3actionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Sweepuniswapv3actionSession struct {
	Contract     *Sweepuniswapv3action // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Sweepuniswapv3actionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Sweepuniswapv3actionCallerSession struct {
	Contract *Sweepuniswapv3actionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// Sweepuniswapv3actionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Sweepuniswapv3actionTransactorSession struct {
	Contract     *Sweepuniswapv3actionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Sweepuniswapv3actionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Sweepuniswapv3actionRaw struct {
	Contract *Sweepuniswapv3action // Generic contract binding to access the raw methods on
}

// Sweepuniswapv3actionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Sweepuniswapv3actionCallerRaw struct {
	Contract *Sweepuniswapv3actionCaller // Generic read-only contract binding to access the raw methods on
}

// Sweepuniswapv3actionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Sweepuniswapv3actionTransactorRaw struct {
	Contract *Sweepuniswapv3actionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweepuniswapv3action creates a new instance of Sweepuniswapv3action, bound to a specific deployed contract.
func NewSweepuniswapv3action(address common.Address, backend bind.ContractBackend) (*Sweepuniswapv3action, error) {
	contract, err := bindSweepuniswapv3action(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sweepuniswapv3action{Sweepuniswapv3actionCaller: Sweepuniswapv3actionCaller{contract: contract}, Sweepuniswapv3actionTransactor: Sweepuniswapv3actionTransactor{contract: contract}, Sweepuniswapv3actionFilterer: Sweepuniswapv3actionFilterer{contract: contract}}, nil
}

// NewSweepuniswapv3actionCaller creates a new read-only instance of Sweepuniswapv3action, bound to a specific deployed contract.
func NewSweepuniswapv3actionCaller(address common.Address, caller bind.ContractCaller) (*Sweepuniswapv3actionCaller, error) {
	contract, err := bindSweepuniswapv3action(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Sweepuniswapv3actionCaller{contract: contract}, nil
}

// NewSweepuniswapv3actionTransactor creates a new write-only instance of Sweepuniswapv3action, bound to a specific deployed contract.
func NewSweepuniswapv3actionTransactor(address common.Address, transactor bind.ContractTransactor) (*Sweepuniswapv3actionTransactor, error) {
	contract, err := bindSweepuniswapv3action(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Sweepuniswapv3actionTransactor{contract: contract}, nil
}

// NewSweepuniswapv3actionFilterer creates a new log filterer instance of Sweepuniswapv3action, bound to a specific deployed contract.
func NewSweepuniswapv3actionFilterer(address common.Address, filterer bind.ContractFilterer) (*Sweepuniswapv3actionFilterer, error) {
	contract, err := bindSweepuniswapv3action(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Sweepuniswapv3actionFilterer{contract: contract}, nil
}

// bindSweepuniswapv3action binds a generic wrapper to an already deployed contract.
func bindSweepuniswapv3action(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Sweepuniswapv3actionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepuniswapv3action *Sweepuniswapv3actionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepuniswapv3action.Contract.Sweepuniswapv3actionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepuniswapv3action *Sweepuniswapv3actionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.Sweepuniswapv3actionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepuniswapv3action *Sweepuniswapv3actionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.Sweepuniswapv3actionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sweepuniswapv3action.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) ERC20TOERC20COMMAND(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "ERC20_TO_ERC20_COMMAND")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) ERC20TOERC20COMMAND() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ERC20TOERC20COMMAND(&_Sweepuniswapv3action.CallOpts)
}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) ERC20TOERC20COMMAND() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ERC20TOERC20COMMAND(&_Sweepuniswapv3action.CallOpts)
}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) ERC20TOETHCOMMANDS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "ERC20_TO_ETH_COMMANDS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) ERC20TOETHCOMMANDS() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ERC20TOETHCOMMANDS(&_Sweepuniswapv3action.CallOpts)
}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) ERC20TOETHCOMMANDS() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ERC20TOETHCOMMANDS(&_Sweepuniswapv3action.CallOpts)
}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) ETHTOERC20COMMANDS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "ETH_TO_ERC20_COMMANDS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) ETHTOERC20COMMANDS() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ETHTOERC20COMMANDS(&_Sweepuniswapv3action.CallOpts)
}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) ETHTOERC20COMMANDS() ([]byte, error) {
	return _Sweepuniswapv3action.Contract.ETHTOERC20COMMANDS(&_Sweepuniswapv3action.CallOpts)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepuniswapv3action.Contract.ArgumentsHash(&_Sweepuniswapv3action.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Sweepuniswapv3action.Contract.ArgumentsHash(&_Sweepuniswapv3action.CallOpts, arguments)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.FeeTokenRegistry(&_Sweepuniswapv3action.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.FeeTokenRegistry(&_Sweepuniswapv3action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) GasConstant() (*big.Int, error) {
	return _Sweepuniswapv3action.Contract.GasConstant(&_Sweepuniswapv3action.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) GasConstant() (*big.Int, error) {
	return _Sweepuniswapv3action.Contract.GasConstant(&_Sweepuniswapv3action.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0xc0e32606.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint256,uint32,uint32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) Hash(opts *bind.CallOpts, arguments ISweepUniswapV3ActionSweepUniswapV3) ([32]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xc0e32606.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint256,uint32,uint32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) Hash(arguments ISweepUniswapV3ActionSweepUniswapV3) ([32]byte, error) {
	return _Sweepuniswapv3action.Contract.Hash(&_Sweepuniswapv3action.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0xc0e32606.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint256,uint32,uint32,(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) Hash(arguments ISweepUniswapV3ActionSweepUniswapV3) ([32]byte, error) {
	return _Sweepuniswapv3action.Contract.Hash(&_Sweepuniswapv3action.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) Hash0(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "hash0", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepuniswapv3action.Contract.Hash0(&_Sweepuniswapv3action.CallOpts, fee)
}

// Hash0 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) Hash0(fee IOtimFeeFee) ([32]byte, error) {
	return _Sweepuniswapv3action.Contract.Hash0(&_Sweepuniswapv3action.CallOpts, fee)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) Router() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.Router(&_Sweepuniswapv3action.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) Router() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.Router(&_Sweepuniswapv3action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) Treasury() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.Treasury(&_Sweepuniswapv3action.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) Treasury() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.Treasury(&_Sweepuniswapv3action.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) UniswapV3Factory() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.UniswapV3Factory(&_Sweepuniswapv3action.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) UniswapV3Factory() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.UniswapV3Factory(&_Sweepuniswapv3action.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sweepuniswapv3action.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) WethAddress() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.WethAddress(&_Sweepuniswapv3action.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Sweepuniswapv3action *Sweepuniswapv3actionCallerSession) WethAddress() (common.Address, error) {
	return _Sweepuniswapv3action.Contract.WethAddress(&_Sweepuniswapv3action.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepuniswapv3action.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.ChargeFee(&_Sweepuniswapv3action.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.ChargeFee(&_Sweepuniswapv3action.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepuniswapv3action.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepuniswapv3action *Sweepuniswapv3actionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.Execute(&_Sweepuniswapv3action.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Sweepuniswapv3action *Sweepuniswapv3actionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Sweepuniswapv3action.Contract.Execute(&_Sweepuniswapv3action.TransactOpts, instruction, arg1, executionState)
}
