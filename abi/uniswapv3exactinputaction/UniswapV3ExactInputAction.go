// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3exactinputaction

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

// IUniswapV3ExactInputActionUniswapV3ExactInput is an auto generated low-level Go binding around an user-defined struct.
type IUniswapV3ExactInputActionUniswapV3ExactInput struct {
	Recipient            common.Address
	TokenIn              common.Address
	TokenOut             common.Address
	FeeTier              *big.Int
	AmountIn             *big.Int
	FloorAmountOut       *big.Int
	MeanPriceLookBack    uint32
	MaxPriceDeviationBPS uint32
	Schedule             IIntervalSchedule
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

// Uniswapv3exactinputactionMetaData contains all meta data concerning the Uniswapv3exactinputaction contract.
var Uniswapv3exactinputactionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"routerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factoryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wethAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"treasuryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasConstant_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ERC20_TO_ERC20_COMMAND\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_TO_ETH_COMMANDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ETH_TO_ERC20_COMMANDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"argumentsHash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"chargeFee\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkInterval\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"lastExecuted\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkStart\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"instruction\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Instruction\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxExecutions\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.Signature\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"executionState\",\"type\":\"tuple\",\"internalType\":\"structInstructionLib.ExecutionState\",\"components\":[{\"name\":\"deactivated\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionCount\",\"type\":\"uint120\",\"internalType\":\"uint120\"},{\"name\":\"lastExecuted\",\"type\":\"uint120\",\"internalType\":\"uint120\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeTokenRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeTokenRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasConstant\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"arguments\",\"type\":\"tuple\",\"internalType\":\"structIUniswapV3ExactInputAction.UniswapV3ExactInput\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTier\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"floorAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"meanPriceLookBack\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPriceDeviationBPS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"tuple\",\"internalType\":\"structIInterval.Schedule\",\"components\":[{\"name\":\"startAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBy\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structIOtimFee.Fee\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxBaseFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniversalRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITreasury\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uniswapV3Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIUniswapV3Factory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BaseFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooEarly\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionTooLate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArguments\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriorityFeePerGasTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"T\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UniswapV3PoolDoesNotExist\",\"inputs\":[]}]",
}

// Uniswapv3exactinputactionABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv3exactinputactionMetaData.ABI instead.
var Uniswapv3exactinputactionABI = Uniswapv3exactinputactionMetaData.ABI

// Uniswapv3exactinputaction is an auto generated Go binding around an Ethereum contract.
type Uniswapv3exactinputaction struct {
	Uniswapv3exactinputactionCaller     // Read-only binding to the contract
	Uniswapv3exactinputactionTransactor // Write-only binding to the contract
	Uniswapv3exactinputactionFilterer   // Log filterer for contract events
}

// Uniswapv3exactinputactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv3exactinputactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3exactinputactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv3exactinputactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3exactinputactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv3exactinputactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv3exactinputactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv3exactinputactionSession struct {
	Contract     *Uniswapv3exactinputaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Uniswapv3exactinputactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv3exactinputactionCallerSession struct {
	Contract *Uniswapv3exactinputactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// Uniswapv3exactinputactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv3exactinputactionTransactorSession struct {
	Contract     *Uniswapv3exactinputactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// Uniswapv3exactinputactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv3exactinputactionRaw struct {
	Contract *Uniswapv3exactinputaction // Generic contract binding to access the raw methods on
}

// Uniswapv3exactinputactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv3exactinputactionCallerRaw struct {
	Contract *Uniswapv3exactinputactionCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv3exactinputactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv3exactinputactionTransactorRaw struct {
	Contract *Uniswapv3exactinputactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv3exactinputaction creates a new instance of Uniswapv3exactinputaction, bound to a specific deployed contract.
func NewUniswapv3exactinputaction(address common.Address, backend bind.ContractBackend) (*Uniswapv3exactinputaction, error) {
	contract, err := bindUniswapv3exactinputaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3exactinputaction{Uniswapv3exactinputactionCaller: Uniswapv3exactinputactionCaller{contract: contract}, Uniswapv3exactinputactionTransactor: Uniswapv3exactinputactionTransactor{contract: contract}, Uniswapv3exactinputactionFilterer: Uniswapv3exactinputactionFilterer{contract: contract}}, nil
}

// NewUniswapv3exactinputactionCaller creates a new read-only instance of Uniswapv3exactinputaction, bound to a specific deployed contract.
func NewUniswapv3exactinputactionCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv3exactinputactionCaller, error) {
	contract, err := bindUniswapv3exactinputaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3exactinputactionCaller{contract: contract}, nil
}

// NewUniswapv3exactinputactionTransactor creates a new write-only instance of Uniswapv3exactinputaction, bound to a specific deployed contract.
func NewUniswapv3exactinputactionTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv3exactinputactionTransactor, error) {
	contract, err := bindUniswapv3exactinputaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3exactinputactionTransactor{contract: contract}, nil
}

// NewUniswapv3exactinputactionFilterer creates a new log filterer instance of Uniswapv3exactinputaction, bound to a specific deployed contract.
func NewUniswapv3exactinputactionFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv3exactinputactionFilterer, error) {
	contract, err := bindUniswapv3exactinputaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv3exactinputactionFilterer{contract: contract}, nil
}

// bindUniswapv3exactinputaction binds a generic wrapper to an already deployed contract.
func bindUniswapv3exactinputaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv3exactinputactionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3exactinputaction.Contract.Uniswapv3exactinputactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.Uniswapv3exactinputactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.Uniswapv3exactinputactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv3exactinputaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.contract.Transact(opts, method, params...)
}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) ERC20TOERC20COMMAND(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "ERC20_TO_ERC20_COMMAND")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) ERC20TOERC20COMMAND() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ERC20TOERC20COMMAND(&_Uniswapv3exactinputaction.CallOpts)
}

// ERC20TOERC20COMMAND is a free data retrieval call binding the contract method 0xadf98436.
//
// Solidity: function ERC20_TO_ERC20_COMMAND() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) ERC20TOERC20COMMAND() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ERC20TOERC20COMMAND(&_Uniswapv3exactinputaction.CallOpts)
}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) ERC20TOETHCOMMANDS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "ERC20_TO_ETH_COMMANDS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) ERC20TOETHCOMMANDS() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ERC20TOETHCOMMANDS(&_Uniswapv3exactinputaction.CallOpts)
}

// ERC20TOETHCOMMANDS is a free data retrieval call binding the contract method 0xb7f9bc1d.
//
// Solidity: function ERC20_TO_ETH_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) ERC20TOETHCOMMANDS() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ERC20TOETHCOMMANDS(&_Uniswapv3exactinputaction.CallOpts)
}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) ETHTOERC20COMMANDS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "ETH_TO_ERC20_COMMANDS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) ETHTOERC20COMMANDS() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ETHTOERC20COMMANDS(&_Uniswapv3exactinputaction.CallOpts)
}

// ETHTOERC20COMMANDS is a free data retrieval call binding the contract method 0xf4543ecd.
//
// Solidity: function ETH_TO_ERC20_COMMANDS() view returns(bytes)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) ETHTOERC20COMMANDS() ([]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ETHTOERC20COMMANDS(&_Uniswapv3exactinputaction.CallOpts)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) ArgumentsHash(opts *bind.CallOpts, arguments []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "argumentsHash", arguments)

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
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ArgumentsHash(&_Uniswapv3exactinputaction.CallOpts, arguments)
}

// ArgumentsHash is a free data retrieval call binding the contract method 0x9fcd91b3.
//
// Solidity: function argumentsHash(bytes arguments) pure returns(bytes32, bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) ArgumentsHash(arguments []byte) ([32]byte, [32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.ArgumentsHash(&_Uniswapv3exactinputaction.CallOpts, arguments)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) CheckInterval(opts *bind.CallOpts, schedule IIntervalSchedule, lastExecuted *big.Int) error {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "checkInterval", schedule, lastExecuted)

	if err != nil {
		return err
	}

	return err

}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Uniswapv3exactinputaction.Contract.CheckInterval(&_Uniswapv3exactinputaction.CallOpts, schedule, lastExecuted)
}

// CheckInterval is a free data retrieval call binding the contract method 0xd808980b.
//
// Solidity: function checkInterval((uint256,uint256,uint256,uint256) schedule, uint256 lastExecuted) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) CheckInterval(schedule IIntervalSchedule, lastExecuted *big.Int) error {
	return _Uniswapv3exactinputaction.Contract.CheckInterval(&_Uniswapv3exactinputaction.CallOpts, schedule, lastExecuted)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) CheckStart(opts *bind.CallOpts, schedule IIntervalSchedule) error {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "checkStart", schedule)

	if err != nil {
		return err
	}

	return err

}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) CheckStart(schedule IIntervalSchedule) error {
	return _Uniswapv3exactinputaction.Contract.CheckStart(&_Uniswapv3exactinputaction.CallOpts, schedule)
}

// CheckStart is a free data retrieval call binding the contract method 0x8dccf7b2.
//
// Solidity: function checkStart((uint256,uint256,uint256,uint256) schedule) view returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) CheckStart(schedule IIntervalSchedule) error {
	return _Uniswapv3exactinputaction.Contract.CheckStart(&_Uniswapv3exactinputaction.CallOpts, schedule)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) FeeTokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "feeTokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) FeeTokenRegistry() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.FeeTokenRegistry(&_Uniswapv3exactinputaction.CallOpts)
}

// FeeTokenRegistry is a free data retrieval call binding the contract method 0xaebfab1b.
//
// Solidity: function feeTokenRegistry() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) FeeTokenRegistry() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.FeeTokenRegistry(&_Uniswapv3exactinputaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) GasConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "gasConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) GasConstant() (*big.Int, error) {
	return _Uniswapv3exactinputaction.Contract.GasConstant(&_Uniswapv3exactinputaction.CallOpts)
}

// GasConstant is a free data retrieval call binding the contract method 0x2cb0f8a2.
//
// Solidity: function gasConstant() view returns(uint256)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) GasConstant() (*big.Int, error) {
	return _Uniswapv3exactinputaction.Contract.GasConstant(&_Uniswapv3exactinputaction.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x30137460.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint32,uint32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) Hash(opts *bind.CallOpts, arguments IUniswapV3ExactInputActionUniswapV3ExactInput) ([32]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "hash", arguments)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x30137460.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint32,uint32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Hash(arguments IUniswapV3ExactInputActionUniswapV3ExactInput) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash(&_Uniswapv3exactinputaction.CallOpts, arguments)
}

// Hash is a free data retrieval call binding the contract method 0x30137460.
//
// Solidity: function hash((address,address,address,uint24,uint256,uint256,uint32,uint32,(uint256,uint256,uint256,uint256),(address,uint256,uint256,uint256)) arguments) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) Hash(arguments IUniswapV3ExactInputActionUniswapV3ExactInput) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash(&_Uniswapv3exactinputaction.CallOpts, arguments)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) Hash0(opts *bind.CallOpts, schedule IIntervalSchedule) ([32]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "hash0", schedule)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash0(&_Uniswapv3exactinputaction.CallOpts, schedule)
}

// Hash0 is a free data retrieval call binding the contract method 0xd607eb95.
//
// Solidity: function hash((uint256,uint256,uint256,uint256) schedule) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) Hash0(schedule IIntervalSchedule) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash0(&_Uniswapv3exactinputaction.CallOpts, schedule)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) Hash1(opts *bind.CallOpts, fee IOtimFeeFee) ([32]byte, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "hash1", fee)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash1(&_Uniswapv3exactinputaction.CallOpts, fee)
}

// Hash1 is a free data retrieval call binding the contract method 0xfd89ea0c.
//
// Solidity: function hash((address,uint256,uint256,uint256) fee) pure returns(bytes32)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) Hash1(fee IOtimFeeFee) ([32]byte, error) {
	return _Uniswapv3exactinputaction.Contract.Hash1(&_Uniswapv3exactinputaction.CallOpts, fee)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Router() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.Router(&_Uniswapv3exactinputaction.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) Router() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.Router(&_Uniswapv3exactinputaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Treasury() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.Treasury(&_Uniswapv3exactinputaction.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) Treasury() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.Treasury(&_Uniswapv3exactinputaction.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) UniswapV3Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "uniswapV3Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) UniswapV3Factory() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.UniswapV3Factory(&_Uniswapv3exactinputaction.CallOpts)
}

// UniswapV3Factory is a free data retrieval call binding the contract method 0x5b549182.
//
// Solidity: function uniswapV3Factory() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) UniswapV3Factory() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.UniswapV3Factory(&_Uniswapv3exactinputaction.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv3exactinputaction.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) WethAddress() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.WethAddress(&_Uniswapv3exactinputaction.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionCallerSession) WethAddress() (common.Address, error) {
	return _Uniswapv3exactinputaction.Contract.WethAddress(&_Uniswapv3exactinputaction.CallOpts)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactor) ChargeFee(opts *bind.TransactOpts, gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.contract.Transact(opts, "chargeFee", gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.ChargeFee(&_Uniswapv3exactinputaction.TransactOpts, gasUsed, fee)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x99d0dd37.
//
// Solidity: function chargeFee(uint256 gasUsed, (address,uint256,uint256,uint256) fee) returns()
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactorSession) ChargeFee(gasUsed *big.Int, fee IOtimFeeFee) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.ChargeFee(&_Uniswapv3exactinputaction.TransactOpts, gasUsed, fee)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactor) Execute(opts *bind.TransactOpts, instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.contract.Transact(opts, "execute", instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.Execute(&_Uniswapv3exactinputaction.TransactOpts, instruction, arg1, executionState)
}

// Execute is a paid mutator transaction binding the contract method 0x69179de3.
//
// Solidity: function execute((uint256,uint256,address,bytes) instruction, (uint8,bytes32,bytes32) , (bool,uint120,uint120) executionState) returns(bool)
func (_Uniswapv3exactinputaction *Uniswapv3exactinputactionTransactorSession) Execute(instruction InstructionLibInstruction, arg1 InstructionLibSignature, executionState InstructionLibExecutionState) (*types.Transaction, error) {
	return _Uniswapv3exactinputaction.Contract.Execute(&_Uniswapv3exactinputaction.TransactOpts, instruction, arg1, executionState)
}
