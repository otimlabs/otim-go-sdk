package client

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Test data: Known ABI-encoded Sweep structs
// These can be replaced with actual encoded data from your other repo
const (
	// Example: Valid Sweep struct with realistic values
	validSweepEncoded = "0x0000000000000000000000002d1d989af240b673c84ceeb3e6279ea98a2cfd05000000000000000000000000000000000000000000000000000000012a05f200000000000000000000000000000000000000000000000000000000007735940000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
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
