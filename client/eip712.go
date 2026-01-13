package client

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"otim-go-sdk/abi/callonceaction"
	"otim-go-sdk/abi/deactivateinstructionaction"
	"otim-go-sdk/abi/depositerc4626action"
	"otim-go-sdk/abi/refuelaction"
	"otim-go-sdk/abi/refuelerc20action"
	"otim-go-sdk/abi/sweepaction"
	"otim-go-sdk/abi/sweepcctpaction"
	"otim-go-sdk/abi/sweepdepositerc4626action"
	"otim-go-sdk/abi/sweeperc20action"
	"otim-go-sdk/abi/sweepuniswapv3action"
	"otim-go-sdk/abi/sweepwithdrawerc4626action"
	"otim-go-sdk/abi/transferaction"
	"otim-go-sdk/abi/transfercctpaction"
	"otim-go-sdk/abi/transfererc20action"
	"otim-go-sdk/abi/transfererc20onceaction"
	"otim-go-sdk/abi/transferonceaction"
	"otim-go-sdk/abi/uniswapv3exactinputaction"
	"otim-go-sdk/abi/withdrawerc4626action"
)

// ActionType represents the type of action
type ActionType string

const (
	ActionTypeTransfer              ActionType = "transfer"
	ActionTypeTransferOnce          ActionType = "transferOnce"
	ActionTypeTransferERC20         ActionType = "transferERC20"
	ActionTypeTransferERC20Once     ActionType = "transferERC20Once"
	ActionTypeTransferCCTP          ActionType = "transferCCTP"
	ActionTypeSweep                 ActionType = "sweep"
	ActionTypeSweepERC20            ActionType = "sweepERC20"
	ActionTypeSweepCCTP             ActionType = "sweepCCTP"
	ActionTypeSweepUniswapV3        ActionType = "sweepUniswapV3"
	ActionTypeRefuel                ActionType = "refuel"
	ActionTypeRefuelERC20           ActionType = "refuelERC20"
	ActionTypeUniswapV3ExactInput   ActionType = "uniswapV3ExactInput"
	ActionTypeDepositERC4626        ActionType = "depositERC4626"
	ActionTypeWithdrawERC4626       ActionType = "withdrawERC4626"
	ActionTypeSweepDepositERC4626   ActionType = "sweepDepositERC4626"
	ActionTypeSweepWithdrawERC4626  ActionType = "sweepWithdrawERC4626"
	ActionTypeCallOnce              ActionType = "callOnce"
	ActionTypeDeactivateInstruction ActionType = "deactivateInstruction"
)

// ActionTypeFromName converts an action name string to ActionType
func ActionTypeFromName(name string) (ActionType, error) {
	// Normalize the name to handle both PascalCase and camelCase inputs
	// e.g., "Transfer" -> "transfer", "TransferOnce" -> "transferOnce"
	normalizedName := name
	if len(name) > 0 {
		normalizedName = strings.ToLower(name[:1]) + name[1:]
	}
	actionType := ActionType(normalizedName)

	// Validate that it's a known action type
	switch actionType {
	case ActionTypeTransfer,
		ActionTypeTransferOnce,
		ActionTypeTransferERC20,
		ActionTypeTransferERC20Once,
		ActionTypeTransferCCTP,
		ActionTypeSweep,
		ActionTypeSweepERC20,
		ActionTypeSweepCCTP,
		ActionTypeSweepUniswapV3,
		ActionTypeRefuel,
		ActionTypeRefuelERC20,
		ActionTypeUniswapV3ExactInput,
		ActionTypeDepositERC4626,
		ActionTypeWithdrawERC4626,
		ActionTypeSweepDepositERC4626,
		ActionTypeSweepWithdrawERC4626,
		ActionTypeCallOnce,
		ActionTypeDeactivateInstruction:
		return actionType, nil
	default:
		return "", fmt.Errorf("unknown action name: %s", name)
	}
}

// convertToTypedStruct converts an anonymous struct from ABI decoding to a typed struct
func convertToTypedStruct(from interface{}, to interface{}) error {
	// Use JSON marshaling/unmarshaling as a simple conversion method
	data, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err := json.Unmarshal(data, to); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}
	return nil
}

// BuildTypedDataForAction builds EIP-712 TypedData for a given action type
func BuildTypedDataForAction(
	instruction BuildInstructionResponse,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	// Determine the action type from the instruction's action name
	actionType, err := ActionTypeFromName(instruction.ActionName)
	if err != nil {
		return nil, fmt.Errorf("invalid action name: %w", err)
	}

	// Decode the arguments from the instruction
	actionArgs, err := DecodeArguments(actionType, instruction.Arguments)
	if err != nil {
		return nil, fmt.Errorf("decode %s arguments: %w", actionType, err)
	}

	switch actionType {
	case ActionTypeTransfer:
		var typedArgs transferaction.ITransferActionTransfer
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert Transfer args: %w", err)
		}
		return buildTransferInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeTransferOnce:
		var typedArgs transferonceaction.ITransferOnceActionTransferOnce
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert TransferOnce args: %w", err)
		}
		return buildTransferOnceInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeTransferERC20:
		var typedArgs transfererc20action.ITransferERC20ActionTransferERC20
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert TransferERC20 args: %w", err)
		}
		return buildTransferERC20InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeTransferERC20Once:
		var typedArgs transfererc20onceaction.ITransferERC20OnceActionTransferERC20Once
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert TransferERC20Once args: %w", err)
		}
		return buildTransferERC20OnceInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeTransferCCTP:
		var typedArgs transfercctpaction.ITransferCCTPActionTransferCCTP
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert TransferCCTP args: %w", err)
		}
		return buildTransferCCTPInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweep:
		var typedArgs sweepaction.ISweepActionSweep
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert Sweep args: %w", err)
		}
		return buildSweepInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweepERC20:
		var typedArgs sweeperc20action.ISweepERC20ActionSweepERC20
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert SweepERC20 args: %w", err)
		}
		return buildSweepERC20InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweepCCTP:
		var typedArgs sweepcctpaction.ISweepCCTPActionSweepCCTP
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert SweepCCTP args: %w", err)
		}
		return buildSweepCCTPInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweepUniswapV3:
		var typedArgs sweepuniswapv3action.ISweepUniswapV3ActionSweepUniswapV3
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert SweepUniswapV3 args: %w", err)
		}
		return buildSweepUniswapV3InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeRefuel:
		var typedArgs refuelaction.IRefuelActionRefuel
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert Refuel args: %w", err)
		}
		return buildRefuelInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeRefuelERC20:
		var typedArgs refuelerc20action.IRefuelERC20ActionRefuelERC20
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert RefuelERC20 args: %w", err)
		}
		return buildRefuelERC20InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeUniswapV3ExactInput:
		var typedArgs uniswapv3exactinputaction.IUniswapV3ExactInputActionUniswapV3ExactInput
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert UniswapV3ExactInput args: %w", err)
		}
		return buildUniswapV3ExactInputInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeDepositERC4626:
		var typedArgs depositerc4626action.IDepositERC4626ActionDepositERC4626
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert DepositERC4626 args: %w", err)
		}
		return buildDepositERC4626InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeWithdrawERC4626:
		var typedArgs withdrawerc4626action.IWithdrawERC4626ActionWithdrawERC4626
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert WithdrawERC4626 args: %w", err)
		}
		return buildWithdrawERC4626InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweepDepositERC4626:
		var typedArgs sweepdepositerc4626action.ISweepDepositERC4626ActionSweepDepositERC4626
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert SweepDepositERC4626 args: %w", err)
		}
		return buildSweepDepositERC4626InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeSweepWithdrawERC4626:
		var typedArgs sweepwithdrawerc4626action.ISweepWithdrawERC4626ActionSweepWithdrawERC4626
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert SweepWithdrawERC4626 args: %w", err)
		}
		return buildSweepWithdrawERC4626InstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeCallOnce:
		var typedArgs callonceaction.ICallOnceActionCallOnce
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert CallOnce args: %w", err)
		}
		return buildCallOnceInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	case ActionTypeDeactivateInstruction:
		var typedArgs deactivateinstructionaction.IDeactivateInstructionActionDeactivateInstruction
		if err := convertToTypedStruct(actionArgs, &typedArgs); err != nil {
			return nil, fmt.Errorf("convert DeactivateInstruction args: %w", err)
		}
		return buildDeactivateInstructionInstructionTypedData(instruction, typedArgs, otimDelegateAddr)
	default:
		return nil, fmt.Errorf("unsupported action type: %s", actionType)
	}
}

// Helper to create domain separator
func createDomain(chainID *big.Int, otimDelegateAddr common.Address) map[string]interface{} {
	return map[string]interface{}{
		"name":              "OtimDelegate",
		"version":           "1",
		"chainId":           chainID.String(),
		"verifyingContract": otimDelegateAddr.Hex(),
		"salt":              crypto.Keccak256Hash([]byte("ON_TIME_INSTRUCTED_MONEY")).Hex(),
	}
}

// Helper to convert Fee struct to map
func feeToMap(fee interface{}) map[string]interface{} {
	// Handle different fee types - they all have the same structure
	var token common.Address
	var maxBaseFeePerGas, maxPriorityFeePerGas, executionFee *big.Int

	switch f := fee.(type) {
	case transferaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case transferonceaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case transfererc20action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case transfererc20onceaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case transfercctpaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweepaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweepcctpaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweeperc20action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case refuelaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case refuelerc20action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case uniswapv3exactinputaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweepuniswapv3action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case depositerc4626action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case withdrawerc4626action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweepdepositerc4626action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case sweepwithdrawerc4626action.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case callonceaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	case deactivateinstructionaction.IOtimFeeFee:
		token = f.Token
		maxBaseFeePerGas = f.MaxBaseFeePerGas
		maxPriorityFeePerGas = f.MaxPriorityFeePerGas
		executionFee = f.ExecutionFee
	// Add other fee types as needed - they all have the same structure
	default:
		// Use reflection as fallback if needed
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"token":                token.Hex(),
		"maxBaseFeePerGas":     maxBaseFeePerGas.String(),
		"maxPriorityFeePerGas": maxPriorityFeePerGas.String(),
		"executionFee":         executionFee.String(),
	}
}

// Helper to convert Schedule struct to map
func scheduleToMap(schedule interface{}) map[string]interface{} {
	// Handle different schedule types - they all have the same structure
	var startAt, startBy, interval, timeout *big.Int

	switch s := schedule.(type) {
	case transferaction.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	case transfererc20action.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	case transfercctpaction.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	case uniswapv3exactinputaction.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	case depositerc4626action.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	case withdrawerc4626action.IIntervalSchedule:
		startAt = s.StartAt
		startBy = s.StartBy
		interval = s.Interval
		timeout = s.Timeout
	// Add other schedule types as needed - they all have the same structure
	default:
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"startAt":  startAt.String(),
		"startBy":  startBy.String(),
		"interval": interval.String(),
		"timeout":  timeout.String(),
	}
}

// Fee type definition for EIP-712
// EIP712Domain type definition for EIP-712
func eip712DomainTypeDefinition() []map[string]string {
	return []map[string]string{
		{"name": "name", "type": "string"},
		{"name": "version", "type": "string"},
		{"name": "chainId", "type": "uint256"},
		{"name": "verifyingContract", "type": "address"},
		{"name": "salt", "type": "bytes32"},
	}
}

// instructionTypeDefinition creates an Instruction type definition with a specific action field
// actionFieldName is the name of the action-specific field (e.g., "transfer", "sweep", etc.)
// actionTypeName is the type of the action field (e.g., "Transfer", "Sweep", etc.)
func instructionTypeDefinition(actionFieldName, actionTypeName string) []map[string]string {
	return []map[string]string{
		{"name": "salt", "type": "uint256"},
		{"name": "maxExecutions", "type": "uint256"},
		{"name": "action", "type": "address"},
		{"name": actionFieldName, "type": actionTypeName},
	}
}

func feeTypeDefinition() []map[string]string {
	return []map[string]string{
		{"name": "token", "type": "address"},
		{"name": "maxBaseFeePerGas", "type": "uint256"},
		{"name": "maxPriorityFeePerGas", "type": "uint256"},
		{"name": "executionFee", "type": "uint256"},
	}
}

// Schedule type definition for EIP-712
func scheduleTypeDefinition() []map[string]string {
	return []map[string]string{
		{"name": "startAt", "type": "uint256"},
		{"name": "startBy", "type": "uint256"},
		{"name": "interval", "type": "uint256"},
		{"name": "timeout", "type": "uint256"},
	}
}

// buildTransferInstructionTypedData builds EIP-712 TypedData for Transfer action
func buildTransferInstructionTypedData(
	instruction BuildInstructionResponse,
	transfer transferaction.ITransferActionTransfer,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("transfer", "Transfer"),
			"Transfer": []map[string]string{
				{"name": "target", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "gasLimit", "type": "uint256"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"transfer": map[string]interface{}{
				"target":   transfer.Target.Hex(),
				"value":    transfer.Value.String(),
				"gasLimit": transfer.GasLimit.String(),
				"schedule": scheduleToMap(transfer.Schedule),
				"fee":      feeToMap(transfer.Fee),
			},
		},
	}, nil
}

// buildTransferOnceInstructionTypedData builds EIP-712 TypedData for TransferOnce action
func buildTransferOnceInstructionTypedData(
	instruction BuildInstructionResponse,
	transferOnce transferonceaction.ITransferOnceActionTransferOnce,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("transferOnce", "TransferOnce"),
			"TransferOnce": []map[string]string{
				{"name": "target", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "gasLimit", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"transferOnce": map[string]interface{}{
				"target":   transferOnce.Target.Hex(),
				"value":    transferOnce.Value.String(),
				"gasLimit": transferOnce.GasLimit.String(),
				"fee":      feeToMap(transferOnce.Fee),
			},
		},
	}, nil
}

// buildSweepInstructionTypedData builds EIP-712 TypedData for Sweep action
func buildSweepInstructionTypedData(
	instruction BuildInstructionResponse,
	sweep sweepaction.ISweepActionSweep,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweep", "Sweep"),
			"Sweep": []map[string]string{
				{"name": "target", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "gasLimit", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweep": map[string]interface{}{
				"target":     sweep.Target.Hex(),
				"threshold":  sweep.Threshold.String(),
				"endBalance": sweep.EndBalance.String(),
				"gasLimit":   sweep.GasLimit.String(),
				"fee":        feeToMap(sweep.Fee),
			},
		},
	}, nil
}

// buildRefuelInstructionTypedData builds EIP-712 TypedData for Refuel action
func buildRefuelInstructionTypedData(
	instruction BuildInstructionResponse,
	refuel refuelaction.IRefuelActionRefuel,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("refuel", "Refuel"),
			"Refuel": []map[string]string{
				{"name": "target", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "gasLimit", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"refuel": map[string]interface{}{
				"target":     refuel.Target.Hex(),
				"threshold":  refuel.Threshold.String(),
				"endBalance": refuel.EndBalance.String(),
				"gasLimit":   refuel.GasLimit.String(),
				"fee":        feeToMap(refuel.Fee),
			},
		},
	}, nil
}

// buildTransferERC20InstructionTypedData builds EIP-712 TypedData for TransferERC20 action
func buildTransferERC20InstructionTypedData(
	instruction BuildInstructionResponse,
	transferERC20 transfererc20action.ITransferERC20ActionTransferERC20,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("transferERC20", "TransferERC20"),
			"TransferERC20": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "target", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"transferERC20": map[string]interface{}{
				"token":    transferERC20.Token.Hex(),
				"target":   transferERC20.Target.Hex(),
				"value":    transferERC20.Value.String(),
				"schedule": scheduleToMap(transferERC20.Schedule),
				"fee":      feeToMap(transferERC20.Fee),
			},
		},
	}, nil
}

// buildTransferERC20OnceInstructionTypedData builds EIP-712 TypedData for TransferERC20Once action
func buildTransferERC20OnceInstructionTypedData(
	instruction BuildInstructionResponse,
	transferERC20Once transfererc20onceaction.ITransferERC20OnceActionTransferERC20Once,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("transferERC20Once", "TransferERC20Once"),
			"TransferERC20Once": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "target", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"transferERC20Once": map[string]interface{}{
				"token":  transferERC20Once.Token.Hex(),
				"target": transferERC20Once.Target.Hex(),
				"value":  transferERC20Once.Value.String(),
				"fee":    feeToMap(transferERC20Once.Fee),
			},
		},
	}, nil
}

// buildSweepERC20InstructionTypedData builds EIP-712 TypedData for SweepERC20 action
func buildSweepERC20InstructionTypedData(
	instruction BuildInstructionResponse,
	sweepERC20 sweeperc20action.ISweepERC20ActionSweepERC20,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweepERC20", "SweepERC20"),
			"SweepERC20": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "target", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweepERC20": map[string]interface{}{
				"token":      sweepERC20.Token.Hex(),
				"target":     sweepERC20.Target.Hex(),
				"threshold":  sweepERC20.Threshold.String(),
				"endBalance": sweepERC20.EndBalance.String(),
				"fee":        feeToMap(sweepERC20.Fee),
			},
		},
	}, nil
}

// buildRefuelERC20InstructionTypedData builds EIP-712 TypedData for RefuelERC20 action
func buildRefuelERC20InstructionTypedData(
	instruction BuildInstructionResponse,
	refuelERC20 refuelerc20action.IRefuelERC20ActionRefuelERC20,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("refuelERC20", "RefuelERC20"),
			"RefuelERC20": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "target", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"refuelERC20": map[string]interface{}{
				"token":      refuelERC20.Token.Hex(),
				"target":     refuelERC20.Target.Hex(),
				"threshold":  refuelERC20.Threshold.String(),
				"endBalance": refuelERC20.EndBalance.String(),
				"fee":        feeToMap(refuelERC20.Fee),
			},
		},
	}, nil
}

// buildTransferCCTPInstructionTypedData builds EIP-712 TypedData for TransferCCTP action
func buildTransferCCTPInstructionTypedData(
	instruction BuildInstructionResponse,
	transferCCTP transfercctpaction.ITransferCCTPActionTransferCCTP,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("transferCCTP", "TransferCCTP"),
			"TransferCCTP": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "amount", "type": "uint256"},
				{"name": "destinationDomain", "type": "uint32"},
				{"name": "destinationMintRecipient", "type": "bytes32"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"transferCCTP": map[string]interface{}{
				"token":                    transferCCTP.Token.Hex(),
				"amount":                   transferCCTP.Amount.String(),
				"destinationDomain":        fmt.Sprintf("%d", transferCCTP.DestinationDomain),
				"destinationMintRecipient": hexutil.Encode(transferCCTP.DestinationMintRecipient[:]),
				"schedule":                 scheduleToMap(transferCCTP.Schedule),
				"fee":                      feeToMap(transferCCTP.Fee),
			},
		},
	}, nil
}

// buildSweepCCTPInstructionTypedData builds EIP-712 TypedData for SweepCCTP action
func buildSweepCCTPInstructionTypedData(
	instruction BuildInstructionResponse,
	sweepCCTP sweepcctpaction.ISweepCCTPActionSweepCCTP,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweepCCTP", "SweepCCTP"),
			"SweepCCTP": []map[string]string{
				{"name": "token", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "destinationDomain", "type": "uint32"},
				{"name": "destinationMintRecipient", "type": "bytes32"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweepCCTP": map[string]interface{}{
				"token":                    sweepCCTP.Token.Hex(),
				"threshold":                sweepCCTP.Threshold.String(),
				"endBalance":               sweepCCTP.EndBalance.String(),
				"destinationDomain":        fmt.Sprintf("%d", sweepCCTP.DestinationDomain),
				"destinationMintRecipient": hexutil.Encode(sweepCCTP.DestinationMintRecipient[:]),
				"fee":                      feeToMap(sweepCCTP.Fee),
			},
		},
	}, nil
}

// buildUniswapV3ExactInputInstructionTypedData builds EIP-712 TypedData for UniswapV3ExactInput action
func buildUniswapV3ExactInputInstructionTypedData(
	instruction BuildInstructionResponse,
	uniswapV3 uniswapv3exactinputaction.IUniswapV3ExactInputActionUniswapV3ExactInput,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("uniswapV3ExactInput", "UniswapV3ExactInput"),
			"UniswapV3ExactInput": []map[string]string{
				{"name": "recipient", "type": "address"},
				{"name": "tokenIn", "type": "address"},
				{"name": "tokenOut", "type": "address"},
				{"name": "feeTier", "type": "uint256"},
				{"name": "amountIn", "type": "uint256"},
				{"name": "floorAmountOut", "type": "uint256"},
				{"name": "meanPriceLookBack", "type": "uint32"},
				{"name": "maxPriceDeviationBPS", "type": "uint32"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"uniswapV3ExactInput": map[string]interface{}{
				"recipient":            uniswapV3.Recipient.Hex(),
				"tokenIn":              uniswapV3.TokenIn.Hex(),
				"tokenOut":             uniswapV3.TokenOut.Hex(),
				"feeTier":              uniswapV3.FeeTier.String(),
				"amountIn":             uniswapV3.AmountIn.String(),
				"floorAmountOut":       uniswapV3.FloorAmountOut.String(),
				"meanPriceLookBack":    fmt.Sprintf("%d", uniswapV3.MeanPriceLookBack),
				"maxPriceDeviationBPS": fmt.Sprintf("%d", uniswapV3.MaxPriceDeviationBPS),
				"schedule":             scheduleToMap(uniswapV3.Schedule),
				"fee":                  feeToMap(uniswapV3.Fee),
			},
		},
	}, nil
}

// buildSweepUniswapV3InstructionTypedData builds EIP-712 TypedData for SweepUniswapV3 action
func buildSweepUniswapV3InstructionTypedData(
	instruction BuildInstructionResponse,
	sweepUniswapV3 sweepuniswapv3action.ISweepUniswapV3ActionSweepUniswapV3,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweepUniswapV3", "SweepUniswapV3"),
			"SweepUniswapV3": []map[string]string{
				{"name": "recipient", "type": "address"},
				{"name": "tokenIn", "type": "address"},
				{"name": "tokenOut", "type": "address"},
				{"name": "feeTier", "type": "uint256"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "floorAmountOut", "type": "uint256"},
				{"name": "meanPriceLookBack", "type": "uint32"},
				{"name": "maxPriceDeviationBPS", "type": "uint32"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweepUniswapV3": map[string]interface{}{
				"recipient":            sweepUniswapV3.Recipient.Hex(),
				"tokenIn":              sweepUniswapV3.TokenIn.Hex(),
				"tokenOut":             sweepUniswapV3.TokenOut.Hex(),
				"feeTier":              sweepUniswapV3.FeeTier.String(),
				"threshold":            sweepUniswapV3.Threshold.String(),
				"endBalance":           sweepUniswapV3.EndBalance.String(),
				"floorAmountOut":       sweepUniswapV3.FloorAmountOut.String(),
				"meanPriceLookBack":    fmt.Sprintf("%d", sweepUniswapV3.MeanPriceLookBack),
				"maxPriceDeviationBPS": fmt.Sprintf("%d", sweepUniswapV3.MaxPriceDeviationBPS),
				"fee":                  feeToMap(sweepUniswapV3.Fee),
			},
		},
	}, nil
}

// buildDepositERC4626InstructionTypedData builds EIP-712 TypedData for DepositERC4626 action
func buildDepositERC4626InstructionTypedData(
	instruction BuildInstructionResponse,
	depositERC4626 depositerc4626action.IDepositERC4626ActionDepositERC4626,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("depositERC4626", "DepositERC4626"),
			"DepositERC4626": []map[string]string{
				{"name": "vault", "type": "address"},
				{"name": "recipient", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "minDeposit", "type": "uint256"},
				{"name": "minTotalShares", "type": "uint256"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"depositERC4626": map[string]interface{}{
				"vault":          depositERC4626.Vault.Hex(),
				"recipient":      depositERC4626.Recipient.Hex(),
				"value":          depositERC4626.Value.String(),
				"minDeposit":     depositERC4626.MinDeposit.String(),
				"minTotalShares": depositERC4626.MinTotalShares.String(),
				"schedule":       scheduleToMap(depositERC4626.Schedule),
				"fee":            feeToMap(depositERC4626.Fee),
			},
		},
	}, nil
}

// buildWithdrawERC4626InstructionTypedData builds EIP-712 TypedData for WithdrawERC4626 action
func buildWithdrawERC4626InstructionTypedData(
	instruction BuildInstructionResponse,
	withdrawERC4626 withdrawerc4626action.IWithdrawERC4626ActionWithdrawERC4626,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("withdrawERC4626", "WithdrawERC4626"),
			"WithdrawERC4626": []map[string]string{
				{"name": "vault", "type": "address"},
				{"name": "recipient", "type": "address"},
				{"name": "value", "type": "uint256"},
				{"name": "minWithdraw", "type": "uint256"},
				{"name": "schedule", "type": "Schedule"},
				{"name": "fee", "type": "Fee"},
			},
			"Schedule": scheduleTypeDefinition(),
			"Fee":      feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"withdrawERC4626": map[string]interface{}{
				"vault":       withdrawERC4626.Vault.Hex(),
				"recipient":   withdrawERC4626.Recipient.Hex(),
				"value":       withdrawERC4626.Value.String(),
				"minWithdraw": withdrawERC4626.MinWithdraw.String(),
				"schedule":    scheduleToMap(withdrawERC4626.Schedule),
				"fee":         feeToMap(withdrawERC4626.Fee),
			},
		},
	}, nil
}

// buildSweepDepositERC4626InstructionTypedData builds EIP-712 TypedData for SweepDepositERC4626 action
func buildSweepDepositERC4626InstructionTypedData(
	instruction BuildInstructionResponse,
	sweepDepositERC4626 sweepdepositerc4626action.ISweepDepositERC4626ActionSweepDepositERC4626,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweepDepositERC4626", "SweepDepositERC4626"),
			"SweepDepositERC4626": []map[string]string{
				{"name": "vault", "type": "address"},
				{"name": "recipient", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "minDeposit", "type": "uint256"},
				{"name": "minTotalShares", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweepDepositERC4626": map[string]interface{}{
				"vault":          sweepDepositERC4626.Vault.Hex(),
				"recipient":      sweepDepositERC4626.Recipient.Hex(),
				"threshold":      sweepDepositERC4626.Threshold.String(),
				"endBalance":     sweepDepositERC4626.EndBalance.String(),
				"minDeposit":     sweepDepositERC4626.MinDeposit.String(),
				"minTotalShares": sweepDepositERC4626.MinTotalShares.String(),
				"fee":            feeToMap(sweepDepositERC4626.Fee),
			},
		},
	}, nil
}

// buildSweepWithdrawERC4626InstructionTypedData builds EIP-712 TypedData for SweepWithdrawERC4626 action
func buildSweepWithdrawERC4626InstructionTypedData(
	instruction BuildInstructionResponse,
	sweepWithdrawERC4626 sweepwithdrawerc4626action.ISweepWithdrawERC4626ActionSweepWithdrawERC4626,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("sweepWithdrawERC4626", "SweepWithdrawERC4626"),
			"SweepWithdrawERC4626": []map[string]string{
				{"name": "vault", "type": "address"},
				{"name": "recipient", "type": "address"},
				{"name": "threshold", "type": "uint256"},
				{"name": "endBalance", "type": "uint256"},
				{"name": "minWithdraw", "type": "uint256"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"sweepWithdrawERC4626": map[string]interface{}{
				"vault":       sweepWithdrawERC4626.Vault.Hex(),
				"recipient":   sweepWithdrawERC4626.Recipient.Hex(),
				"threshold":   sweepWithdrawERC4626.Threshold.String(),
				"endBalance":  sweepWithdrawERC4626.EndBalance.String(),
				"minWithdraw": sweepWithdrawERC4626.MinWithdraw.String(),
				"fee":         feeToMap(sweepWithdrawERC4626.Fee),
			},
		},
	}, nil
}

// buildCallOnceInstructionTypedData builds EIP-712 TypedData for CallOnce action
func buildCallOnceInstructionTypedData(
	instruction BuildInstructionResponse,
	callOnce callonceaction.ICallOnceActionCallOnce,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("callOnce", "CallOnce"),
			"CallOnce": []map[string]string{
				{"name": "target", "type": "address"},
				{"name": "allowFailure", "type": "bool"},
				{"name": "value", "type": "uint256"},
				{"name": "gasLimit", "type": "uint256"},
				{"name": "returnSizeLimit", "type": "uint16"},
				{"name": "selector", "type": "bytes4"},
				{"name": "data", "type": "bytes"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"callOnce": map[string]interface{}{
				"target":          callOnce.Target.Hex(),
				"allowFailure":    callOnce.AllowFailure,
				"value":           callOnce.Value.String(),
				"gasLimit":        callOnce.GasLimit.String(),
				"returnSizeLimit": fmt.Sprintf("%d", callOnce.ReturnSizeLimit),
				"selector":        hexutil.Encode(callOnce.Selector[:]),
				"data":            hexutil.Encode(callOnce.Data),
				"fee":             feeToMap(callOnce.Fee),
			},
		},
	}, nil
}

// buildDeactivateInstructionInstructionTypedData builds EIP-712 TypedData for DeactivateInstruction action
func buildDeactivateInstructionInstructionTypedData(
	instruction BuildInstructionResponse,
	deactivateInstruction deactivateinstructionaction.IDeactivateInstructionActionDeactivateInstruction,
	otimDelegateAddr common.Address,
) (map[string]interface{}, error) {
	chainID := big.NewInt(int64(instruction.ChainID))
	return map[string]interface{}{
		"types": map[string]interface{}{
			"EIP712Domain": eip712DomainTypeDefinition(),
			"Instruction":  instructionTypeDefinition("deactivateInstruction", "DeactivateInstruction"),
			"DeactivateInstruction": []map[string]string{
				{"name": "instructionId", "type": "bytes32"},
				{"name": "fee", "type": "Fee"},
			},
			"Fee": feeTypeDefinition(),
		},
		"primaryType": "Instruction",
		"domain":      createDomain(chainID, otimDelegateAddr),
		"message": map[string]interface{}{
			"salt":          instruction.Salt.ToInt().String(),
			"maxExecutions": instruction.MaxExecutions.ToInt().String(),
			"action":        instruction.Action.Hex(),
			"deactivateInstruction": map[string]interface{}{
				"instructionId": hexutil.Encode(deactivateInstruction.InstructionId[:]),
				"fee":           feeToMap(deactivateInstruction.Fee),
			},
		},
	}, nil
}

// DecodeArguments decodes the instruction arguments based on action type
// The arguments bytes should be ABI-encoded representation of the action struct
func DecodeArguments(actionType ActionType, arguments []byte) (interface{}, error) {
	// Get the appropriate ABI for the action
	var abiJSON string

	switch actionType {
	case ActionTypeTransfer:
		abiJSON = transferaction.TransferactionMetaData.ABI
	case ActionTypeTransferOnce:
		abiJSON = transferonceaction.TransferonceactionMetaData.ABI
	case ActionTypeTransferERC20:
		abiJSON = transfererc20action.Transfererc20actionMetaData.ABI
	case ActionTypeTransferERC20Once:
		abiJSON = transfererc20onceaction.Transfererc20onceactionMetaData.ABI
	case ActionTypeTransferCCTP:
		abiJSON = transfercctpaction.TransfercctpactionMetaData.ABI
	case ActionTypeSweep:
		abiJSON = sweepaction.SweepactionMetaData.ABI
	case ActionTypeSweepERC20:
		abiJSON = sweeperc20action.Sweeperc20actionMetaData.ABI
	case ActionTypeSweepCCTP:
		abiJSON = sweepcctpaction.SweepcctpactionMetaData.ABI
	case ActionTypeSweepUniswapV3:
		abiJSON = sweepuniswapv3action.Sweepuniswapv3actionMetaData.ABI
	case ActionTypeRefuel:
		abiJSON = refuelaction.RefuelactionMetaData.ABI
	case ActionTypeRefuelERC20:
		abiJSON = refuelerc20action.Refuelerc20actionMetaData.ABI
	case ActionTypeUniswapV3ExactInput:
		abiJSON = uniswapv3exactinputaction.Uniswapv3exactinputactionMetaData.ABI
	case ActionTypeDepositERC4626:
		abiJSON = depositerc4626action.Depositerc4626actionMetaData.ABI
	case ActionTypeWithdrawERC4626:
		abiJSON = withdrawerc4626action.Withdrawerc4626actionMetaData.ABI
	case ActionTypeSweepDepositERC4626:
		abiJSON = sweepdepositerc4626action.Sweepdepositerc4626actionMetaData.ABI
	case ActionTypeSweepWithdrawERC4626:
		abiJSON = sweepwithdrawerc4626action.Sweepwithdrawerc4626actionMetaData.ABI
	case ActionTypeCallOnce:
		abiJSON = callonceaction.CallonceactionMetaData.ABI
	case ActionTypeDeactivateInstruction:
		abiJSON = deactivateinstructionaction.DeactivateinstructionactionMetaData.ABI
	default:
		return nil, fmt.Errorf("unsupported action type: %s", actionType)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("parse ABI: %w", err)
	}

	// Find the `hash` method that takes the action struct as input
	// This method has the complete struct type definition with all components
	// Note: When Solidity functions are overloaded, go-ethereum names them hash, hash0, hash1, etc.
	var structArgument *abi.Argument

	// Look through all methods starting with "hash" to find the action struct
	// Filter out nested structs like Schedule and Fee by checking the parameter name
	for _, method := range parsedABI.Methods {
		if strings.HasPrefix(method.Name, "hash") {
			arg := &method.Inputs[0]

			// Skip nested structs - we want the main action struct
			// Nested structs typically have names like "schedule", "fee", etc.
			paramName := strings.ToLower(arg.Name)
			if paramName == "schedule" || paramName == "fee" {
				continue
			}

			structArgument = arg
			break
		}
	}

	if structArgument == nil {
		return nil, fmt.Errorf("could not find hash method with action struct in ABI")
	}

	// Create ABI arguments wrapper for unpacking
	argsDef := abi.Arguments{*structArgument}

	// Unpack the arguments into the target struct
	values, err := argsDef.Unpack(arguments)
	if err != nil {
		return nil, fmt.Errorf("unpack arguments: %w", err)
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("no values unpacked from arguments")
	}

	// Return the first (and only) value
	return values[0], nil
}
