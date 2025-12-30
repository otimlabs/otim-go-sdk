package client

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Test data: Known ABI-encoded structs
const (
	// Sweep
	validSweepEncoded = "0x0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd05000000000000000000000000000000000000000000000000000000012a05f200000000000000000000000000000000000000000000000000000000007735940000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// Transfer Actions
	validTransferEncoded         = "0x0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd05000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000052080000000000000000000000000000000000000000000000000000019b70833b0f0000000000000000000000000000000000000000000000000000019b708362200000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validTransferOnceEncoded     = "0x0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd05000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000052080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validTransferERC20Encoded    = "0x0000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd0500000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000019b708646cf0000000000000000000000000000000000000000000000000000019b70866de00000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validTransferERC20OnceEncoded = "0x0000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd0500000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// CCTP Actions
	validTransferCCTPEncoded = "0x0000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c72380000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validSweepCCTPEncoded    = "0x0000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000002faf0800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// Sweep Actions
	validSweepERC20Encoded           = "0x0000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd0500000000000000000000000000000000000000000000000000000000000001f400000000000000000000000000000000000000000000000000000000000000c80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validSweepUniswapV3Encoded       = "0x0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae900000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c723800000000000000000000000000000000000000000000000000000000000001f40000000000000000000000000000000000000000000000000000000077359400000000000000000000000000000000000000000000000000000000003b9aca000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000038400000000000000000000000000000000000000000000000000000000000003e80000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validSweepDepositERC4626Encoded  = "0x000000000000000000000000beef01735c132ada46aa9aa4c54623caa92a64cb0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae90000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validSweepWithdrawERC4626Encoded = "0x000000000000000000000000beef01735c132ada46aa9aa4c54623caa92a64cb0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae90000000000000000000000000000000000000000000000000000000002faf0800000000000000000000000000000000000000000000000000000000001312d0000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// Refuel Actions
	validRefuelEncoded     = "0x0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd050000000000000000000000000000000000000000000000000000000077359400000000000000000000000000000000000000000000000000000000012a05f20000000000000000000000000000000000000000000000000000000000000052080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validRefuelERC20Encoded = "0x0000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd05000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000001f40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// DeFi Actions
	validUniswapV3ExactInputEncoded = "0x0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae900000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c723800000000000000000000000000000000000000000000000000000000000001f40000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000038400000000000000000000000000000000000000000000000000000000000001f40000000000000000000000000000000000000000000000000000000069541f2b000000000000000000000000000000000000000000000000000000006954463c0000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000008ca00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validDepositERC4626Encoded      = "0x000000000000000000000000beef01735c132ada46aa9aa4c54623caa92a64cb0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae90000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	validWithdrawERC4626Encoded     = "0x000000000000000000000000beef01735c132ada46aa9aa4c54623caa92a64cb0000000000000000000000003af524037e28ce53933405a7753f520a0e4e5ae90000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	
	// Utility Actions
	validCallOnceEncoded              = "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000f62849f9a0b5bf2913b396098f7c7019b51a820a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000008fc00000000000000000000000000000000000000000000000000000000000002006bff20900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"
	validDeactivateInstructionEncoded = "0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
)

// E2E Test constants for TypedData generation
const (
	testSaltHex = "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
	testMaxExecutions = 100
	testActionAddress = "0x1111111111111111111111111111111111111111"
	testChainID = 1
	testOtimDelegateAddress = "0x2222222222222222222222222222222222222222"
)

func TestDecodeArguments_Sweep(t *testing.T) {
	tests := []struct {
		name                     string
		encodedHex               string
		actionType               ActionType
		expectedTarget           string
		expectedThreshold        string
		expectedEndBalance       string
		expectedGasLimit         string
		expectedFeeToken         string
		expectedMaxBaseFee       string
		expectedMaxPriorityFee   string
		expectedExecutionFee     string
		wantErr                  bool
	}{
		{
			name:                   "valid sweep with realistic values",
			encodedHex:             validSweepEncoded,
			actionType:             ActionTypeSweep,
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedThreshold:      "5000000000",
			expectedEndBalance:     "2000000000",
			expectedGasLimit:       "0",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert hex string to bytes
			encodedBytes := common.FromHex(tt.encodedHex)

			// Decode the arguments
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			// Check for expected errors
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			// Use reflection to access fields since DecodeArguments returns an anonymous struct
			v := reflect.ValueOf(result)
			if v.Kind() != reflect.Struct {
				t.Fatalf("expected struct, got %v (type: %T)", v.Kind(), result)
			}

			// Debug: print the type and fields
			t.Logf("Decoded struct type: %T", result)
			t.Logf("Struct has %d fields", v.NumField())
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				t.Logf("  Field %d: %s (type: %v)", i, field.Name, field.Type)
			}

			// Helper to get struct field value
			getField := func(fieldName string) reflect.Value {
				field := v.FieldByName(fieldName)
				if !field.IsValid() {
					t.Fatalf("field %s not found in struct", fieldName)
				}
				return field
			}

			// Extract and verify fields
			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target address mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			gasLimit := getField("GasLimit").Interface().(*big.Int)
			if got := gasLimit.String(); got != tt.expectedGasLimit {
				t.Errorf("GasLimit mismatch: got %s, want %s", got, tt.expectedGasLimit)
			}

			// Verify Fee struct fields
			feeValue := getField("Fee")
			if feeValue.Kind() != reflect.Struct {
				t.Fatalf("Fee field is not a struct, got %v", feeValue.Kind())
			}

			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func getFieldHelper(t *testing.T, v reflect.Value, fieldName string) reflect.Value {
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		t.Fatalf("field %s not found in struct", fieldName)
	}
	return field
}

// E2E Test Helper Functions

// mockBuildInstruction creates a BuildInstructionResponse with test constants
func mockBuildInstruction(arguments []byte, actionName string) BuildInstructionResponse {
	saltBig := new(big.Int)
	saltBig.SetString(testSaltHex[2:], 16) // Remove "0x" prefix
	
	return BuildInstructionResponse{
		Address:       common.HexToAddress(testActionAddress),
		ChainID:       ChainID(testChainID),
		Salt:          hexutil.Big(*saltBig),
		MaxExecutions: hexutil.Big(*big.NewInt(testMaxExecutions)),
		Action:        common.HexToAddress(testActionAddress),
		Arguments:     arguments,
		ActionName:    actionName,
	}
}

// verifyTypedDataStructure validates TypedData has the correct top-level structure
func verifyTypedDataStructure(t *testing.T, typedData map[string]interface{}, expectedPrimaryType string) {
	// Check top-level keys
	if _, ok := typedData["types"]; !ok {
		t.Error("TypedData missing 'types' field")
	}
	if _, ok := typedData["primaryType"]; !ok {
		t.Error("TypedData missing 'primaryType' field")
	}
	if _, ok := typedData["domain"]; !ok {
		t.Error("TypedData missing 'domain' field")
	}
	if _, ok := typedData["message"]; !ok {
		t.Error("TypedData missing 'message' field")
	}
	
	// Check primaryType
	if primaryType, ok := typedData["primaryType"].(string); !ok || primaryType != expectedPrimaryType {
		t.Errorf("primaryType mismatch: got %v, want %s", typedData["primaryType"], expectedPrimaryType)
	}
}

// verifyDomain verifies the EIP712Domain fields match test constants
func verifyDomain(t *testing.T, domain map[string]interface{}) {
	if name, ok := domain["name"].(string); !ok || name != "OtimDelegate" {
		t.Errorf("domain.name mismatch: got %v, want OtimDelegate", domain["name"])
	}
	if version, ok := domain["version"].(string); !ok || version != "1" {
		t.Errorf("domain.version mismatch: got %v, want 1", domain["version"])
	}
	if chainId, ok := domain["chainId"].(string); !ok || chainId != fmt.Sprintf("%d", testChainID) {
		t.Errorf("domain.chainId mismatch: got %v, want %d", domain["chainId"], testChainID)
	}
	if verifyingContract, ok := domain["verifyingContract"].(string); !ok || verifyingContract != testOtimDelegateAddress {
		t.Errorf("domain.verifyingContract mismatch: got %v, want %s", domain["verifyingContract"], testOtimDelegateAddress)
	}
	if _, ok := domain["salt"]; !ok {
		t.Error("domain.salt missing")
	}
}

// verifyInstructionMessage verifies the instruction-level message fields
func verifyInstructionMessage(t *testing.T, message map[string]interface{}) {
	saltBig := new(big.Int)
	saltBig.SetString(testSaltHex[2:], 16)
	
	if salt, ok := message["salt"].(string); !ok || salt != saltBig.String() {
		t.Errorf("message.salt mismatch: got %v, want %s", message["salt"], saltBig.String())
	}
	if maxExecutions, ok := message["maxExecutions"].(string); !ok || maxExecutions != fmt.Sprintf("%d", testMaxExecutions) {
		t.Errorf("message.maxExecutions mismatch: got %v, want %d", message["maxExecutions"], testMaxExecutions)
	}
	if action, ok := message["action"].(string); !ok || action != testActionAddress {
		t.Errorf("message.action mismatch: got %v, want %s", message["action"], testActionAddress)
	}
}

// verifyTypes checks that the types object contains expected type definitions
func verifyTypes(t *testing.T, types map[string]interface{}, expectedTypes []string) {
	for _, typeName := range expectedTypes {
		if _, ok := types[typeName]; !ok {
			t.Errorf("types missing definition for '%s'", typeName)
		}
	}
}

// compareTypedData does a deep comparison of two TypedData structures
func compareTypedData(t *testing.T, expected, actual map[string]interface{}) {
	t.Helper()
	
	// Compare primaryType
	if expected["primaryType"] != actual["primaryType"] {
		t.Errorf("primaryType mismatch: expected %v, got %v", expected["primaryType"], actual["primaryType"])
	}
	
	// Compare domain
	expectedDomain := expected["domain"].(map[string]interface{})
	actualDomain := actual["domain"].(map[string]interface{})
	compareMaps(t, "domain", expectedDomain, actualDomain)
	
	// Compare message
	expectedMessage := expected["message"].(map[string]interface{})
	actualMessage := actual["message"].(map[string]interface{})
	compareMaps(t, "message", expectedMessage, actualMessage)
	
	// Compare types structure (just verify keys exist, not full structure)
	expectedTypes := expected["types"].(map[string]interface{})
	actualTypes := actual["types"].(map[string]interface{})
	for key := range expectedTypes {
		if _, ok := actualTypes[key]; !ok {
			t.Errorf("types.%s missing in actual", key)
		}
	}
}

// compareMaps recursively compares two maps
func compareMaps(t *testing.T, path string, expected, actual map[string]interface{}) {
	t.Helper()
	
	// Check all expected keys exist and have correct values
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			t.Errorf("%s.%s missing in actual", path, key)
			continue
		}
		
		compareValues(t, path+"."+key, expectedValue, actualValue)
	}
	
	// Check no extra keys in actual
	for key := range actual {
		if _, ok := expected[key]; !ok {
			t.Errorf("%s.%s exists in actual but not in expected", path, key)
		}
	}
}

// compareValues compares two values, handling nested maps
func compareValues(t *testing.T, path string, expected, actual interface{}) {
	t.Helper()
	
	switch expectedVal := expected.(type) {
	case map[string]interface{}:
		actualMap, ok := actual.(map[string]interface{})
		if !ok {
			t.Errorf("%s: expected map, got %T", path, actual)
			return
		}
		compareMaps(t, path, expectedVal, actualMap)
	case string:
		actualStr, ok := actual.(string)
		if !ok {
			t.Errorf("%s: expected string, got %T", path, actual)
			return
		}
		if expectedVal != actualStr {
			t.Errorf("%s: expected %v, got %v", path, expectedVal, actualStr)
		}
	default:
		// For other types, use direct comparison
		if expected != actual {
			t.Errorf("%s: expected %v, got %v", path, expected, actual)
		}
	}
}

// expectedDomain creates the expected domain map for E2E tests
func expectedDomain(chainID *big.Int) map[string]interface{} {
	return map[string]interface{}{
		"name":              "OtimDelegate",
		"version":           "1",
		"chainId":           chainID.String(),
		"verifyingContract": testOtimDelegateAddress,
		"salt":              crypto.Keccak256Hash([]byte("ON_TIME_INSTRUCTED_MONEY")).Hex(),
	}
}

// expectedMessageBase creates the base instruction message fields for E2E tests
func expectedMessageBase() map[string]interface{} {
	saltBig := new(big.Int)
	saltBig.SetString(testSaltHex[2:], 16)
	
	return map[string]interface{}{
		"salt":          saltBig.String(),
		"maxExecutions": fmt.Sprintf("%d", testMaxExecutions),
		"action":        testActionAddress,
	}
}

func TestDecodeArguments_Transfer(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedTarget         string
		expectedValue          string
		expectedGasLimit       string
		expectedStartAt        string
		expectedStartBy        string
		expectedInterval       string
		expectedTimeout        string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid transfer",
			encodedHex:             validTransferEncoded,
			actionType:             ActionTypeTransfer,
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedValue:          "100",
			expectedGasLimit:       "21000",
			expectedStartAt:        "1767119207183",
			expectedStartBy:        "1767119217184",
			expectedInterval:       "36000",
			expectedTimeout:        "36000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			if v.Kind() != reflect.Struct {
				t.Fatalf("expected struct, got %v", v.Kind())
			}

			getField := func(fieldName string) reflect.Value {
				field := v.FieldByName(fieldName)
				if !field.IsValid() {
					t.Fatalf("field %s not found in struct", fieldName)
				}
				return field
			}

			// Verify Transfer fields
			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			gasLimit := getField("GasLimit").Interface().(*big.Int)
			if got := gasLimit.String(); got != tt.expectedGasLimit {
				t.Errorf("GasLimit mismatch: got %s, want %s", got, tt.expectedGasLimit)
			}

			// Verify Schedule struct
			scheduleValue := getField("Schedule")
			if scheduleValue.Kind() != reflect.Struct {
				t.Fatalf("Schedule field is not a struct")
			}

			startAt := scheduleValue.FieldByName("StartAt").Interface().(*big.Int)
			if got := startAt.String(); got != tt.expectedStartAt {
				t.Errorf("Schedule.StartAt mismatch: got %s, want %s", got, tt.expectedStartAt)
			}

			startBy := scheduleValue.FieldByName("StartBy").Interface().(*big.Int)
			if got := startBy.String(); got != tt.expectedStartBy {
				t.Errorf("Schedule.StartBy mismatch: got %s, want %s", got, tt.expectedStartBy)
			}

			interval := scheduleValue.FieldByName("Interval").Interface().(*big.Int)
			if got := interval.String(); got != tt.expectedInterval {
				t.Errorf("Schedule.Interval mismatch: got %s, want %s", got, tt.expectedInterval)
			}

			timeout := scheduleValue.FieldByName("Timeout").Interface().(*big.Int)
			if got := timeout.String(); got != tt.expectedTimeout {
				t.Errorf("Schedule.Timeout mismatch: got %s, want %s", got, tt.expectedTimeout)
			}

			// Verify Fee struct
			feeValue := getField("Fee")
			if feeValue.Kind() != reflect.Struct {
				t.Fatalf("Fee field is not a struct")
			}

			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_TransferOnce(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedTarget         string
		expectedValue          string
		expectedGasLimit       string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid transfer once",
			encodedHex:             validTransferOnceEncoded,
			actionType:             ActionTypeTransferOnce,
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedValue:          "100",
			expectedGasLimit:       "21000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			if v.Kind() != reflect.Struct {
				t.Fatalf("expected struct, got %v", v.Kind())
			}

			getField := func(fieldName string) reflect.Value {
				field := v.FieldByName(fieldName)
				if !field.IsValid() {
					t.Fatalf("field %s not found in struct", fieldName)
				}
				return field
			}

			// Verify TransferOnce fields
			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			gasLimit := getField("GasLimit").Interface().(*big.Int)
			if got := gasLimit.String(); got != tt.expectedGasLimit {
				t.Errorf("GasLimit mismatch: got %s, want %s", got, tt.expectedGasLimit)
			}

			// Verify Fee struct
			feeValue := getField("Fee")
			if feeValue.Kind() != reflect.Struct {
				t.Fatalf("Fee field is not a struct")
			}

			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_TransferERC20(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedToken          string
		expectedTarget         string
		expectedValue          string
		expectedStartAt        string
		expectedStartBy        string
		expectedInterval       string
		expectedTimeout        string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid transfer erc20",
			encodedHex:             validTransferERC20Encoded,
			actionType:             ActionTypeTransferERC20,
			expectedToken:          "0x2e234DAe75C793f67A35089C9d99245E1C58470b", // TODO: Update with expected values
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedValue:          "100",
			expectedStartAt:        "1767119406799",
			expectedStartBy:        "1767119416800",
			expectedInterval:       "36000",
			expectedTimeout:        "36000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			// Verify Schedule
			scheduleValue := getField("Schedule")
			startAt := scheduleValue.FieldByName("StartAt").Interface().(*big.Int)
			if got := startAt.String(); got != tt.expectedStartAt {
				t.Errorf("Schedule.StartAt mismatch: got %s, want %s", got, tt.expectedStartAt)
			}

			startBy := scheduleValue.FieldByName("StartBy").Interface().(*big.Int)
			if got := startBy.String(); got != tt.expectedStartBy {
				t.Errorf("Schedule.StartBy mismatch: got %s, want %s", got, tt.expectedStartBy)
			}

			interval := scheduleValue.FieldByName("Interval").Interface().(*big.Int)
			if got := interval.String(); got != tt.expectedInterval {
				t.Errorf("Schedule.Interval mismatch: got %s, want %s", got, tt.expectedInterval)
			}

			timeout := scheduleValue.FieldByName("Timeout").Interface().(*big.Int)
			if got := timeout.String(); got != tt.expectedTimeout {
				t.Errorf("Schedule.Timeout mismatch: got %s, want %s", got, tt.expectedTimeout)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_TransferERC20Once(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedToken          string
		expectedTarget         string
		expectedValue          string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid transfer erc20 once",
			encodedHex:             validTransferERC20OnceEncoded,
			actionType:             ActionTypeTransferERC20Once,
			expectedToken:          "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedValue:          "100",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_TransferCCTP(t *testing.T) {
	tests := []struct {
		name                          string
		encodedHex                    string
		actionType                    ActionType
		expectedToken                 string
		expectedAmount                string
		expectedDestinationDomain     uint32
		expectedDestinationMintRecipient string
		expectedStartAt               string
		expectedStartBy               string
		expectedInterval              string
		expectedTimeout               string
		expectedFeeToken              string
		expectedMaxBaseFee            string
		expectedMaxPriorityFee        string
		expectedExecutionFee          string
		wantErr                       bool
	}{
		{
			name:                          "valid transfer cctp",
			encodedHex:                    validTransferCCTPEncoded,
			actionType:                    ActionTypeTransferCCTP,
			expectedToken:                 "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
			expectedAmount:                "100000000",
			expectedDestinationDomain:     4,
			expectedDestinationMintRecipient: "0x0000000000000000000000000000000000000000000000000000000000000001",
			expectedStartAt:               "0",
			expectedStartBy:               "0",
			expectedInterval:              "0",
			expectedTimeout:               "0",
			expectedFeeToken:              "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:            "0",
			expectedMaxPriorityFee:        "0",
			expectedExecutionFee:          "0",
			wantErr:                       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			amount := getField("Amount").Interface().(*big.Int)
			if got := amount.String(); got != tt.expectedAmount {
				t.Errorf("Amount mismatch: got %s, want %s", got, tt.expectedAmount)
			}

			destDomain := getField("DestinationDomain").Interface().(uint32)
			if got := destDomain; got != tt.expectedDestinationDomain {
				t.Errorf("DestinationDomain mismatch: got %d, want %d", got, tt.expectedDestinationDomain)
			}

			destRecipient := getField("DestinationMintRecipient").Interface().([32]byte)
			if got := hex.EncodeToString(destRecipient[:]); "0x"+got != tt.expectedDestinationMintRecipient {
				t.Errorf("DestinationMintRecipient mismatch: got 0x%s, want %s", got, tt.expectedDestinationMintRecipient)
			}

			// Verify Schedule
			scheduleValue := getField("Schedule")
			startAt := scheduleValue.FieldByName("StartAt").Interface().(*big.Int)
			if got := startAt.String(); got != tt.expectedStartAt {
				t.Errorf("Schedule.StartAt mismatch: got %s, want %s", got, tt.expectedStartAt)
			}

			startBy := scheduleValue.FieldByName("StartBy").Interface().(*big.Int)
			if got := startBy.String(); got != tt.expectedStartBy {
				t.Errorf("Schedule.StartBy mismatch: got %s, want %s", got, tt.expectedStartBy)
			}

			interval := scheduleValue.FieldByName("Interval").Interface().(*big.Int)
			if got := interval.String(); got != tt.expectedInterval {
				t.Errorf("Schedule.Interval mismatch: got %s, want %s", got, tt.expectedInterval)
			}

			timeout := scheduleValue.FieldByName("Timeout").Interface().(*big.Int)
			if got := timeout.String(); got != tt.expectedTimeout {
				t.Errorf("Schedule.Timeout mismatch: got %s, want %s", got, tt.expectedTimeout)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_SweepCCTP(t *testing.T) {
	tests := []struct {
		name                          string
		encodedHex                    string
		actionType                    ActionType
		expectedToken                 string
		expectedDestinationDomain     uint32
		expectedDestinationMintRecipient string
		expectedThreshold             string
		expectedEndBalance            string
		expectedFeeToken              string
		expectedMaxBaseFee            string
		expectedMaxPriorityFee        string
		expectedExecutionFee          string
		wantErr                       bool
	}{
		{
			name:                          "valid sweep cctp",
			encodedHex:                    validSweepCCTPEncoded,
			actionType:                    ActionTypeSweepCCTP,
			expectedToken:                 "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
			expectedDestinationDomain:     4,
			expectedDestinationMintRecipient: "0x0000000000000000000000000000000000000000000000000000000000000001",
			expectedThreshold:             "100000000",
			expectedEndBalance:            "50000000",
			expectedFeeToken:              "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:            "0",
			expectedMaxPriorityFee:        "0",
			expectedExecutionFee:          "0",
			wantErr:                       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			destDomain := getField("DestinationDomain").Interface().(uint32)
			if got := destDomain; got != tt.expectedDestinationDomain {
				t.Errorf("DestinationDomain mismatch: got %d, want %d", got, tt.expectedDestinationDomain)
			}

			destRecipient := getField("DestinationMintRecipient").Interface().([32]byte)
			if got := hex.EncodeToString(destRecipient[:]); "0x"+got != tt.expectedDestinationMintRecipient {
				t.Errorf("DestinationMintRecipient mismatch: got 0x%s, want %s", got, tt.expectedDestinationMintRecipient)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_SweepERC20(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedToken          string
		expectedTarget         string
		expectedThreshold      string
		expectedEndBalance     string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid sweep erc20",
			encodedHex:             validSweepERC20Encoded,
			actionType:             ActionTypeSweepERC20,
			expectedToken:          "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedThreshold:      "500",
			expectedEndBalance:     "200",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_RefuelERC20(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedToken          string
		expectedTarget         string
		expectedThreshold      string
		expectedEndBalance     string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid refuel erc20",
			encodedHex:             validRefuelERC20Encoded,
			actionType:             ActionTypeRefuelERC20,
			expectedToken:          "0x2e234DAe75C793f67A35089C9d99245E1C58470b", // TODO: Update with expected values
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedThreshold:      "100",
			expectedEndBalance:     "500",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			token := getField("Token").Interface().(common.Address)
			if got := token.Hex(); got != tt.expectedToken {
				t.Errorf("Token mismatch: got %s, want %s", got, tt.expectedToken)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_SweepUniswapV3(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedRecipient      string
		expectedTokenIn        string
		expectedTokenOut       string
		expectedFeeTier        string
		expectedThreshold      string
		expectedEndBalance     string
		expectedFloorAmountOut   string
		expectedMeanPriceLookBack string
		expectedMaxPriceDeviationBPS string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid sweep uniswap v3",
			encodedHex:             validSweepUniswapV3Encoded,
			actionType:             ActionTypeSweepUniswapV3,
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
			expectedTokenIn:        "0x0000000000000000000000000000000000000000",
			expectedTokenOut:       "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
			expectedFeeTier:        "500",
			expectedThreshold:      "2000000000",
			expectedEndBalance:     "1000000000",
			expectedFloorAmountOut:   "1",
			expectedMeanPriceLookBack: "900",
			expectedMaxPriceDeviationBPS: "1000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			tokenIn := getField("TokenIn").Interface().(common.Address)
			if got := tokenIn.Hex(); got != tt.expectedTokenIn {
				t.Errorf("TokenIn mismatch: got %s, want %s", got, tt.expectedTokenIn)
			}

			tokenOut := getField("TokenOut").Interface().(common.Address)
			if got := tokenOut.Hex(); got != tt.expectedTokenOut {
				t.Errorf("TokenOut mismatch: got %s, want %s", got, tt.expectedTokenOut)
			}

			feeTier := getField("FeeTier").Interface().(*big.Int)
			if got := feeTier.String(); got != tt.expectedFeeTier {
				t.Errorf("FeeTier mismatch: got %s, want %s", got, tt.expectedFeeTier)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			minAmountOut := getField("FloorAmountOut").Interface().(*big.Int)
			if got := minAmountOut.String(); got != tt.expectedFloorAmountOut {
				t.Errorf("MinAmountOut mismatch: got %s, want %s", got, tt.expectedFloorAmountOut)
			}

			meanPriceLookBack := getField("MeanPriceLookBack").Interface().(uint32)
			if got := fmt.Sprintf("%d", meanPriceLookBack); got != tt.expectedMeanPriceLookBack {
				t.Errorf("MeanPriceLookBack mismatch: got %s, want %s", got, tt.expectedMeanPriceLookBack)
			}

			maxPriceDeviationBPS := getField("MaxPriceDeviationBPS").Interface().(uint32)
			if got := fmt.Sprintf("%d", maxPriceDeviationBPS); got != tt.expectedMaxPriceDeviationBPS {
				t.Errorf("MaxPriceDeviationBPS mismatch: got %s, want %s", got, tt.expectedMaxPriceDeviationBPS)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_SweepDepositERC4626(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedVault          string
		expectedRecipient      string
		expectedThreshold      string
		expectedEndBalance     string
		expectedMinDeposit     string
		expectedMinTotalShares string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid sweep deposit erc4626",
			encodedHex:             validSweepDepositERC4626Encoded,
			actionType:             ActionTypeSweepDepositERC4626,
			expectedVault:          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
			expectedThreshold:      "100000000",
			expectedEndBalance:     "0",
			expectedMinDeposit:     "1",
			expectedMinTotalShares: "100000000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			vault := getField("Vault").Interface().(common.Address)
			if got := vault.Hex(); got != tt.expectedVault {
				t.Errorf("Vault mismatch: got %s, want %s", got, tt.expectedVault)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			minDeposit := getField("MinDeposit").Interface().(*big.Int)
			if got := minDeposit.String(); got != tt.expectedMinDeposit {
				t.Errorf("MinDeposit mismatch: got %s, want %s", got, tt.expectedMinDeposit)
			}

			minTotalShares := getField("MinTotalShares").Interface().(*big.Int)
			if got := minTotalShares.String(); got != tt.expectedMinTotalShares {
				t.Errorf("MinTotalShares mismatch: got %s, want %s", got, tt.expectedMinTotalShares)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_SweepWithdrawERC4626(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedVault          string
		expectedRecipient      string
		expectedThreshold      string
		expectedEndBalance     string
		expectedMinWithdraw    string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid sweep withdraw erc4626",
			encodedHex:             validSweepWithdrawERC4626Encoded,
			actionType:             ActionTypeSweepWithdrawERC4626,
			expectedVault:          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
			expectedThreshold:      "50000000",
			expectedEndBalance:     "20000000",
			expectedMinWithdraw:    "1",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			vault := getField("Vault").Interface().(common.Address)
			if got := vault.Hex(); got != tt.expectedVault {
				t.Errorf("Vault mismatch: got %s, want %s", got, tt.expectedVault)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			minWithdraw := getField("MinWithdraw").Interface().(*big.Int)
			if got := minWithdraw.String(); got != tt.expectedMinWithdraw {
				t.Errorf("MinWithdraw mismatch: got %s, want %s", got, tt.expectedMinWithdraw)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_Refuel(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedTarget         string
		expectedThreshold      string
		expectedEndBalance     string
		expectedGasLimit       string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid refuel",
			encodedHex:             validRefuelEncoded,
			actionType:             ActionTypeRefuel,
			expectedTarget:         "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
			expectedThreshold:      "2000000000",
			expectedEndBalance:     "5000000000",
			expectedGasLimit:       "21000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			threshold := getField("Threshold").Interface().(*big.Int)
			if got := threshold.String(); got != tt.expectedThreshold {
				t.Errorf("Threshold mismatch: got %s, want %s", got, tt.expectedThreshold)
			}

			endBalance := getField("EndBalance").Interface().(*big.Int)
			if got := endBalance.String(); got != tt.expectedEndBalance {
				t.Errorf("EndBalance mismatch: got %s, want %s", got, tt.expectedEndBalance)
			}

			gasLimit := getField("GasLimit").Interface().(*big.Int)
			if got := gasLimit.String(); got != tt.expectedGasLimit {
				t.Errorf("GasLimit mismatch: got %s, want %s", got, tt.expectedGasLimit)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_UniswapV3ExactInput(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedRecipient      string
		expectedTokenIn        string
		expectedTokenOut       string
		expectedFeeTier        string
		expectedAmountIn       string
		expectedFloorAmountOut   string
		expectedMeanPriceLookBack string
		expectedMaxPriceDeviationBPS string
		expectedStartAt        string
		expectedStartBy        string
		expectedInterval       string
		expectedTimeout        string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid uniswap v3 exact input",
			encodedHex:             validUniswapV3ExactInputEncoded,
			actionType:             ActionTypeUniswapV3ExactInput,
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9", // TODO: Update with expected values
			expectedTokenIn:        "0x0000000000000000000000000000000000000000",
			expectedTokenOut:       "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
			expectedFeeTier:        "500",
			expectedAmountIn:       "1000000000000000000",
			expectedFloorAmountOut:   "1",
			expectedMeanPriceLookBack: "900",
			expectedMaxPriceDeviationBPS: "500",
			expectedStartAt:        "1767120683",
			expectedStartBy:        "1767130684",
			expectedInterval:       "36000",
			expectedTimeout:        "36000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			tokenIn := getField("TokenIn").Interface().(common.Address)
			if got := tokenIn.Hex(); got != tt.expectedTokenIn {
				t.Errorf("TokenIn mismatch: got %s, want %s", got, tt.expectedTokenIn)
			}

			tokenOut := getField("TokenOut").Interface().(common.Address)
			if got := tokenOut.Hex(); got != tt.expectedTokenOut {
				t.Errorf("TokenOut mismatch: got %s, want %s", got, tt.expectedTokenOut)
			}

			feeTier := getField("FeeTier").Interface().(*big.Int)
			if got := feeTier.String(); got != tt.expectedFeeTier {
				t.Errorf("FeeTier mismatch: got %s, want %s", got, tt.expectedFeeTier)
			}

			amountIn := getField("AmountIn").Interface().(*big.Int)
			if got := amountIn.String(); got != tt.expectedAmountIn {
				t.Errorf("AmountIn mismatch: got %s, want %s", got, tt.expectedAmountIn)
			}

			minAmountOut := getField("FloorAmountOut").Interface().(*big.Int)
			if got := minAmountOut.String(); got != tt.expectedFloorAmountOut {
				t.Errorf("MinAmountOut mismatch: got %s, want %s", got, tt.expectedFloorAmountOut)
			}

			meanPriceLookBack := getField("MeanPriceLookBack").Interface().(uint32)
			if got := fmt.Sprintf("%d", meanPriceLookBack); got != tt.expectedMeanPriceLookBack {
				t.Errorf("MeanPriceLookBack mismatch: got %s, want %s", got, tt.expectedMeanPriceLookBack)
			}

			maxPriceDeviationBPS := getField("MaxPriceDeviationBPS").Interface().(uint32)
			if got := fmt.Sprintf("%d", maxPriceDeviationBPS); got != tt.expectedMaxPriceDeviationBPS {
				t.Errorf("MaxPriceDeviationBPS mismatch: got %s, want %s", got, tt.expectedMaxPriceDeviationBPS)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_DepositERC4626(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedVault          string
		expectedRecipient      string
		expectedValue          string
		expectedMinDeposit     string
		expectedMinTotalShares string
		expectedStartAt        string
		expectedStartBy        string
		expectedInterval       string
		expectedTimeout        string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid deposit erc4626",
			encodedHex:             validDepositERC4626Encoded,
			actionType:             ActionTypeDepositERC4626,
			expectedVault:          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
			expectedValue:          "100000000",
			expectedMinDeposit:     "1",
			expectedMinTotalShares: "100000000",
			expectedStartAt:        "0",
			expectedStartBy:        "0",
			expectedInterval:       "0",
			expectedTimeout:        "0",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			vault := getField("Vault").Interface().(common.Address)
			if got := vault.Hex(); got != tt.expectedVault {
				t.Errorf("Vault mismatch: got %s, want %s", got, tt.expectedVault)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			minDeposit := getField("MinDeposit").Interface().(*big.Int)
			if got := minDeposit.String(); got != tt.expectedMinDeposit {
				t.Errorf("MinDeposit mismatch: got %s, want %s", got, tt.expectedMinDeposit)
			}

			minTotalShares := getField("MinTotalShares").Interface().(*big.Int)
			if got := minTotalShares.String(); got != tt.expectedMinTotalShares {
				t.Errorf("MinTotalShares mismatch: got %s, want %s", got, tt.expectedMinTotalShares)
			}

			// Verify Schedule
			scheduleValue := getField("Schedule")
			startAt := scheduleValue.FieldByName("StartAt").Interface().(*big.Int)
			if got := startAt.String(); got != tt.expectedStartAt {
				t.Errorf("Schedule.StartAt mismatch: got %s, want %s", got, tt.expectedStartAt)
			}

			startBy := scheduleValue.FieldByName("StartBy").Interface().(*big.Int)
			if got := startBy.String(); got != tt.expectedStartBy {
				t.Errorf("Schedule.StartBy mismatch: got %s, want %s", got, tt.expectedStartBy)
			}

			interval := scheduleValue.FieldByName("Interval").Interface().(*big.Int)
			if got := interval.String(); got != tt.expectedInterval {
				t.Errorf("Schedule.Interval mismatch: got %s, want %s", got, tt.expectedInterval)
			}

			timeout := scheduleValue.FieldByName("Timeout").Interface().(*big.Int)
			if got := timeout.String(); got != tt.expectedTimeout {
				t.Errorf("Schedule.Timeout mismatch: got %s, want %s", got, tt.expectedTimeout)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_WithdrawERC4626(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedVault          string
		expectedRecipient      string
		expectedValue          string
		expectedMinWithdraw    string
		expectedStartAt        string
		expectedStartBy        string
		expectedInterval       string
		expectedTimeout        string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid withdraw erc4626",
			encodedHex:             validWithdrawERC4626Encoded,
			actionType:             ActionTypeWithdrawERC4626,
			expectedVault:          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
			expectedRecipient:      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
			expectedValue:          "100000000",
			expectedMinWithdraw:    "1",
			expectedStartAt:        "0",
			expectedStartBy:        "0",
			expectedInterval:       "0",
			expectedTimeout:        "0",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			vault := getField("Vault").Interface().(common.Address)
			if got := vault.Hex(); got != tt.expectedVault {
				t.Errorf("Vault mismatch: got %s, want %s", got, tt.expectedVault)
			}

			recipient := getField("Recipient").Interface().(common.Address)
			if got := recipient.Hex(); got != tt.expectedRecipient {
				t.Errorf("Recipient mismatch: got %s, want %s", got, tt.expectedRecipient)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			minWithdraw := getField("MinWithdraw").Interface().(*big.Int)
			if got := minWithdraw.String(); got != tt.expectedMinWithdraw {
				t.Errorf("MinWithdraw mismatch: got %s, want %s", got, tt.expectedMinWithdraw)
			}

			// Verify Schedule
			scheduleValue := getField("Schedule")
			startAt := scheduleValue.FieldByName("StartAt").Interface().(*big.Int)
			if got := startAt.String(); got != tt.expectedStartAt {
				t.Errorf("Schedule.StartAt mismatch: got %s, want %s", got, tt.expectedStartAt)
			}

			startBy := scheduleValue.FieldByName("StartBy").Interface().(*big.Int)
			if got := startBy.String(); got != tt.expectedStartBy {
				t.Errorf("Schedule.StartBy mismatch: got %s, want %s", got, tt.expectedStartBy)
			}

			interval := scheduleValue.FieldByName("Interval").Interface().(*big.Int)
			if got := interval.String(); got != tt.expectedInterval {
				t.Errorf("Schedule.Interval mismatch: got %s, want %s", got, tt.expectedInterval)
			}

			timeout := scheduleValue.FieldByName("Timeout").Interface().(*big.Int)
			if got := timeout.String(); got != tt.expectedTimeout {
				t.Errorf("Schedule.Timeout mismatch: got %s, want %s", got, tt.expectedTimeout)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_CallOnce(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedTarget         string
		expectedAllowFailure   bool
		expectedValue          string
		expectedGasLimit       string
		expectedReturnSizeLimit uint16
	    expectedSelector       string
		expectedData           string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid call once",
			encodedHex:             validCallOnceEncoded,
			actionType:             ActionTypeCallOnce,
			expectedTarget:         "0xF62849F9A0B5Bf2913b396098F7c7019b51A820a",
			expectedAllowFailure:   false,
			expectedValue:          "100",
			expectedGasLimit:       "2300",
			expectedReturnSizeLimit: 512,
			expectedSelector:       "0x6bff2090",
			expectedData:           "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			target := getField("Target").Interface().(common.Address)
			if got := target.Hex(); got != tt.expectedTarget {
				t.Errorf("Target mismatch: got %s, want %s", got, tt.expectedTarget)
			}

			allowFailure := getField("AllowFailure").Interface().(bool)
			if got := allowFailure; got != tt.expectedAllowFailure {
				t.Errorf("AllowFailure mismatch: got %v, want %v", got, tt.expectedAllowFailure)
			}

			value := getField("Value").Interface().(*big.Int)
			if got := value.String(); got != tt.expectedValue {
				t.Errorf("Value mismatch: got %s, want %s", got, tt.expectedValue)
			}

			gasLimit := getField("GasLimit").Interface().(*big.Int)
			if got := gasLimit.String(); got != tt.expectedGasLimit {
				t.Errorf("GasLimit mismatch: got %s, want %s", got, tt.expectedGasLimit)
			}

			returnSizeLimit := getField("ReturnSizeLimit").Interface().(uint16)
			if got := returnSizeLimit; got != tt.expectedReturnSizeLimit {
				t.Errorf("ReturnSizeLimit mismatch: got %d, want %d", got, tt.expectedReturnSizeLimit)
			}

			selector := getField("Selector").Interface().([4]byte)
			if got := "0x" + hex.EncodeToString(selector[:]); got != tt.expectedSelector {
				t.Errorf("Selector mismatch: got %s, want %s", got, tt.expectedSelector)
			}

			data := getField("Data").Interface().([]byte)
			if got := "0x" + hex.EncodeToString(data); got != tt.expectedData {
				t.Errorf("Data mismatch: got %s, want %s", got, tt.expectedData)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

func TestDecodeArguments_DeactivateInstruction(t *testing.T) {
	tests := []struct {
		name                   string
		encodedHex             string
		actionType             ActionType
		expectedInstructionId  string
		expectedFeeToken       string
		expectedMaxBaseFee     string
		expectedMaxPriorityFee string
		expectedExecutionFee   string
		wantErr                bool
	}{
		{
			name:                   "valid deactivate instruction",
			encodedHex:             validDeactivateInstructionEncoded,
			actionType:             ActionTypeDeactivateInstruction,
			expectedInstructionId:  "0x0000000000000000000000000000000000000000000000000000000000000001",
			expectedFeeToken:       "0x0000000000000000000000000000000000000000",
			expectedMaxBaseFee:     "0",
			expectedMaxPriorityFee: "0",
			expectedExecutionFee:   "0",
			wantErr:                false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedBytes := common.FromHex(tt.encodedHex)
			result, err := DecodeArguments(tt.actionType, encodedBytes)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result == nil {
				t.Fatal("result is nil")
			}

			v := reflect.ValueOf(result)
			getField := func(fieldName string) reflect.Value {
				return getFieldHelper(t, v, fieldName)
			}

			instructionId := getField("InstructionId").Interface().([32]byte)
			if got := "0x" + hex.EncodeToString(instructionId[:]); got != tt.expectedInstructionId {
				t.Errorf("InstructionId mismatch: got %s, want %s", got, tt.expectedInstructionId)
			}

			// Verify Fee
			feeValue := getField("Fee")
			feeToken := feeValue.FieldByName("Token").Interface().(common.Address)
			if got := feeToken.Hex(); got != tt.expectedFeeToken {
				t.Errorf("Fee.Token mismatch: got %s, want %s", got, tt.expectedFeeToken)
			}

			maxBaseFee := feeValue.FieldByName("MaxBaseFeePerGas").Interface().(*big.Int)
			if got := maxBaseFee.String(); got != tt.expectedMaxBaseFee {
				t.Errorf("Fee.MaxBaseFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxBaseFee)
			}

			maxPriorityFee := feeValue.FieldByName("MaxPriorityFeePerGas").Interface().(*big.Int)
			if got := maxPriorityFee.String(); got != tt.expectedMaxPriorityFee {
				t.Errorf("Fee.MaxPriorityFeePerGas mismatch: got %s, want %s", got, tt.expectedMaxPriorityFee)
			}

			executionFee := feeValue.FieldByName("ExecutionFee").Interface().(*big.Int)
			if got := executionFee.String(); got != tt.expectedExecutionFee {
				t.Errorf("Fee.ExecutionFee mismatch: got %s, want %s", got, tt.expectedExecutionFee)
			}
		})
	}
}

// ============================================================================
// E2E TypedData Generation Tests
// ============================================================================

// TestE2E_TransferTypedData tests the complete flow: decode → generate TypedData
func TestE2E_TransferTypedData(t *testing.T) {
	// Create mock instruction
	encodedBytes := common.FromHex(validTransferEncoded)
	instruction := mockBuildInstruction(encodedBytes, "Transfer")

	// Generate TypedData (which handles action type determination and decoding internally)
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	// Construct expected TypedData
	message := expectedMessageBase()
	message["transfer"] = map[string]interface{}{
		"target":   "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"value":    "100",
		"gasLimit": "21000",
		"schedule": map[string]interface{}{
			"startAt":  "1767119207183",
			"startBy":  "1767119217184",
			"interval": "36000",
			"timeout":  "36000",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"Transfer":     "present",
			"Schedule":     "present",
			"Fee":          "present",
		},
	}

	// Compare expected vs actual
	compareTypedData(t, expected, typedData)
}

// TestE2E_TransferOnceTypedData tests TransferOnce action TypedData generation
func TestE2E_TransferOnceTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validTransferOnceEncoded)
	instruction := mockBuildInstruction(encodedBytes, "TransferOnce")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["transferOnce"] = map[string]interface{}{
		"target":   "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"value":    "100",
		"gasLimit": "21000",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"TransferOnce": "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_TransferERC20TypedData tests TransferERC20 action TypedData generation
func TestE2E_TransferERC20TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validTransferERC20Encoded)
	instruction := mockBuildInstruction(encodedBytes, "TransferERC20")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["transferERC20"] = map[string]interface{}{
		"token":  "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
		"target": "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"value":  "100",
		"schedule": map[string]interface{}{
			"startAt":  "1767119406799",
			"startBy":  "1767119416800",
			"interval": "36000",
			"timeout":  "36000",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":  "present",
			"Instruction":   "present",
			"TransferERC20": "present",
			"Schedule":      "present",
			"Fee":           "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_TransferERC20OnceTypedData tests TransferERC20Once action TypedData generation
func TestE2E_TransferERC20OnceTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validTransferERC20OnceEncoded)
	instruction := mockBuildInstruction(encodedBytes, "TransferERC20Once")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["transferERC20Once"] = map[string]interface{}{
		"token":  "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
		"target": "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"value":  "100",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":       "present",
			"Instruction":        "present",
			"TransferERC20Once": "present",
			"Fee":                "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_TransferCCTPTypedData tests TransferCCTP action TypedData generation
func TestE2E_TransferCCTPTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validTransferCCTPEncoded)
	instruction := mockBuildInstruction(encodedBytes, "TransferCCTP")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["transferCCTP"] = map[string]interface{}{
		"token":                      "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
		"amount":                     "100000000",
		"destinationDomain":          "4",
		"destinationMintRecipient":   "0x0000000000000000000000000000000000000000000000000000000000000001",
		"schedule": map[string]interface{}{
			"startAt":  "0",
			"startBy":  "0",
			"interval": "0",
			"timeout":  "0",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"TransferCCTP": "present",
			"Schedule":     "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepCCTPTypedData tests SweepCCTP action TypedData generation
func TestE2E_SweepCCTPTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepCCTPEncoded)
	instruction := mockBuildInstruction(encodedBytes, "SweepCCTP")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweepCCTP"] = map[string]interface{}{
		"token":                     "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
		"threshold":                 "100000000",
		"endBalance":                "50000000",
		"destinationDomain":         "4",
		"destinationMintRecipient":  "0x0000000000000000000000000000000000000000000000000000000000000001",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"SweepCCTP":    "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepTypedData tests Sweep action TypedData generation
func TestE2E_SweepTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepEncoded)
	instruction := mockBuildInstruction(encodedBytes, "Sweep")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweep"] = map[string]interface{}{
		"target":     "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"threshold":  "5000000000",
		"endBalance": "2000000000",
		"gasLimit":   "0",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"Sweep":        "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepERC20TypedData tests SweepERC20 action TypedData generation
func TestE2E_SweepERC20TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepERC20Encoded)
	instruction := mockBuildInstruction(encodedBytes, "SweepERC20")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweepERC20"] = map[string]interface{}{
		"token":      "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
		"target":     "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"threshold":  "500",
		"endBalance": "200",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"SweepERC20":   "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepUniswapV3TypedData tests SweepUniswapV3 action TypedData generation
func TestE2E_SweepUniswapV3TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepUniswapV3Encoded)
	instruction := mockBuildInstruction(encodedBytes, "SweepUniswapV3")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweepUniswapV3"] = map[string]interface{}{
		"recipient":              "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"tokenIn":                "0x0000000000000000000000000000000000000000",
		"tokenOut":               "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
		"feeTier":                "500",
		"threshold":              "2000000000",
		"endBalance":             "1000000000",
		"floorAmountOut":         "1",
		"meanPriceLookBack":      "900",
		"maxPriceDeviationBPS":   "1000",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":   "present",
			"Instruction":    "present",
			"SweepUniswapV3": "present",
			"Fee":            "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepDepositERC4626TypedData tests SweepDepositERC4626 action TypedData generation
func TestE2E_SweepDepositERC4626TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepDepositERC4626Encoded)
	instruction := mockBuildInstruction(encodedBytes, "SweepDepositERC4626")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweepDepositERC4626"] = map[string]interface{}{
		"vault":          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
		"recipient":      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"threshold":      "100000000",
		"endBalance":     "0",
		"minDeposit":     "1",
		"minTotalShares": "100000000",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":         "present",
			"Instruction":          "present",
			"SweepDepositERC4626": "present",
			"Fee":                  "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_SweepWithdrawERC4626TypedData tests SweepWithdrawERC4626 action TypedData generation
func TestE2E_SweepWithdrawERC4626TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validSweepWithdrawERC4626Encoded)
	instruction := mockBuildInstruction(encodedBytes, "SweepWithdrawERC4626")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["sweepWithdrawERC4626"] = map[string]interface{}{
		"vault":       "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
		"recipient":   "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"threshold":   "50000000",
		"endBalance":  "20000000",
		"minWithdraw": "1",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":          "present",
			"Instruction":           "present",
			"SweepWithdrawERC4626": "present",
			"Fee":                   "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_RefuelTypedData tests Refuel action TypedData generation
func TestE2E_RefuelTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validRefuelEncoded)
	instruction := mockBuildInstruction(encodedBytes, "Refuel")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["refuel"] = map[string]interface{}{
		"target":     "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"threshold":  "2000000000",
		"endBalance": "5000000000",
		"gasLimit":   "21000",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"Refuel":       "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_RefuelERC20TypedData tests RefuelERC20 action TypedData generation
func TestE2E_RefuelERC20TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validRefuelERC20Encoded)
	instruction := mockBuildInstruction(encodedBytes, "RefuelERC20")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["refuelERC20"] = map[string]interface{}{
		"token":      "0x2e234DAe75C793f67A35089C9d99245E1C58470b",
		"target":     "0x2d1d989af240B673C84cEeb3E6279Ea98a2CFd05",
		"threshold":  "100",
		"endBalance": "500",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"RefuelERC20":  "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_UniswapV3ExactInputTypedData tests UniswapV3ExactInput action TypedData generation
func TestE2E_UniswapV3ExactInputTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validUniswapV3ExactInputEncoded)
	instruction := mockBuildInstruction(encodedBytes, "UniswapV3ExactInput")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["uniswapV3ExactInput"] = map[string]interface{}{
		"recipient":              "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"tokenIn":                "0x0000000000000000000000000000000000000000",
		"tokenOut":               "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
		"feeTier":                "500",
		"amountIn":               "1000000000000000000",
		"floorAmountOut":         "1",
		"meanPriceLookBack":      "900",
		"maxPriceDeviationBPS":   "500",
		"schedule": map[string]interface{}{
			"startAt":  "1767120683",
			"startBy":  "1767130684",
			"interval": "36000",
			"timeout":  "36000",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":        "present",
			"Instruction":         "present",
			"UniswapV3ExactInput": "present",
			"Schedule":            "present",
			"Fee":                 "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_DepositERC4626TypedData tests DepositERC4626 action TypedData generation
func TestE2E_DepositERC4626TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validDepositERC4626Encoded)
	instruction := mockBuildInstruction(encodedBytes, "DepositERC4626")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["depositERC4626"] = map[string]interface{}{
		"vault":          "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
		"recipient":      "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"value":          "100000000",
		"minDeposit":     "1",
		"minTotalShares": "100000000",
		"schedule": map[string]interface{}{
			"startAt":  "0",
			"startBy":  "0",
			"interval": "0",
			"timeout":  "0",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":   "present",
			"Instruction":    "present",
			"DepositERC4626": "present",
			"Schedule":       "present",
			"Fee":            "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_WithdrawERC4626TypedData tests WithdrawERC4626 action TypedData generation
func TestE2E_WithdrawERC4626TypedData(t *testing.T) {
	encodedBytes := common.FromHex(validWithdrawERC4626Encoded)
	instruction := mockBuildInstruction(encodedBytes, "WithdrawERC4626")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["withdrawERC4626"] = map[string]interface{}{
		"vault":       "0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB",
		"recipient":   "0x3aF524037e28ce53933405A7753f520A0e4E5ae9",
		"value":       "100000000",
		"minWithdraw": "1",
		"schedule": map[string]interface{}{
			"startAt":  "0",
			"startBy":  "0",
			"interval": "0",
			"timeout":  "0",
		},
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":    "present",
			"Instruction":     "present",
			"WithdrawERC4626": "present",
			"Schedule":        "present",
			"Fee":             "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_CallOnceTypedData tests CallOnce action TypedData generation
func TestE2E_CallOnceTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validCallOnceEncoded)
	instruction := mockBuildInstruction(encodedBytes, "CallOnce")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["callOnce"] = map[string]interface{}{
		"target":          "0xF62849F9A0B5Bf2913b396098F7c7019b51A820a",
		"allowFailure":    false,
		"value":           "100",
		"gasLimit":        "2300",
		"returnSizeLimit": "512",
		"selector":        "0x6bff2090",
		"data":            "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain": "present",
			"Instruction":  "present",
			"CallOnce":     "present",
			"Fee":          "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

// TestE2E_DeactivateInstructionTypedData tests DeactivateInstruction action TypedData generation
func TestE2E_DeactivateInstructionTypedData(t *testing.T) {
	encodedBytes := common.FromHex(validDeactivateInstructionEncoded)
	instruction := mockBuildInstruction(encodedBytes, "DeactivateInstruction")
	chainID := big.NewInt(testChainID)
	otimDelegate := common.HexToAddress(testOtimDelegateAddress)
	typedData, err := BuildTypedDataForAction(instruction, chainID, otimDelegate)
	if err != nil {
		t.Fatalf("BuildTypedDataForAction failed: %v", err)
	}

	message := expectedMessageBase()
	message["deactivateInstruction"] = map[string]interface{}{
		"instructionId": "0x0000000000000000000000000000000000000000000000000000000000000001",
		"fee": map[string]interface{}{
			"token":                "0x0000000000000000000000000000000000000000",
			"maxBaseFeePerGas":     "0",
			"maxPriorityFeePerGas": "0",
			"executionFee":         "0",
		},
	}
	
	expected := map[string]interface{}{
		"primaryType": "Instruction",
		"domain":      expectedDomain(chainID),
		"message":     message,
		"types": map[string]interface{}{
			"EIP712Domain":          "present",
			"Instruction":           "present",
			"DeactivateInstruction": "present",
			"Fee":                   "present",
		},
	}

	compareTypedData(t, expected, typedData)
}

