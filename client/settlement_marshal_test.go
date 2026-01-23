package client

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestSettlementParamsMarshalJSON(t *testing.T) {
	settlementAmount := big.NewInt(1000000)
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	params := &SettlementParams{
		AcceptedTokens: map[ChainID][]common.Address{
			1: {usdcAddress},
		},
		SettlementChainId: 1,
		SettlementToken:   usdcAddress,
		SettlementAmount:  hexutil.Big(*settlementAmount),
		RecipientAddress:  recipientAddress,
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Failed to marshal SettlementParams: %v", err)
	}

	// Verify externally tagged format
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if _, ok := result["settlement"]; !ok {
		t.Errorf("Expected 'settlement' key in externally tagged format, got keys: %v", result)
	}

	t.Logf("Marshaled SettlementParams: %s", string(data))
}

func TestVaultWithdrawSettlementParamsMarshalJSON(t *testing.T) {
	withdrawAmount := big.NewInt(5000000000)
	vaultAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	params := &VaultWithdrawSettlementParams{
		VaultAddress:         vaultAddress,
		VaultUnderlyingToken: usdcAddress,
		VaultChainId:         1,
		SettlementChainId:    1,
		SettlementToken:      usdcAddress,
		RecipientAddress:     recipientAddress,
		WithdrawAmount:       hexutil.Big(*withdrawAmount),
	}

	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Failed to marshal VaultWithdrawSettlementParams: %v", err)
	}

	// Verify externally tagged format
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if _, ok := result["vaultWithdrawSettlement"]; !ok {
		t.Errorf("Expected 'vaultWithdrawSettlement' key in externally tagged format, got keys: %v", result)
	}

	t.Logf("Marshaled VaultWithdrawSettlementParams: %s", string(data))
}

func TestPaymentRequestMetadataMarshalJSON(t *testing.T) {
	dueDate := time.Now()
	note := "Test payment"
	invoiceId := "inv_12345"

	metadata := &PaymentRequestMetadata{
		Token:              "USDC",
		AmountDue:          "150.00",
		Currency:           "USD",
		DueDate:            &dueDate,
		Note:               &note,
		FromAccountAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Payer: &Payer{
			Name:    "John Doe",
			Address: common.HexToAddress("0x9876543210987654321098765432109876543210"),
		},
		Source:    "QuickBooks",
		InvoiceId: &invoiceId,
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal PaymentRequestMetadata: %v", err)
	}

	// Verify the type field is present
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if result["type"] != "PaymentRequest" {
		t.Errorf("Expected type 'PaymentRequest', got %v", result["type"])
	}

	t.Logf("Marshaled PaymentRequestMetadata: %s", string(data))
}

func TestAutoEarnMetadataMarshalJSON(t *testing.T) {
	metadata := &AutoEarnMetadata{
		TokenSymbol:  "USDC",
		VaultAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		VaultName:    "Aave USDC Vault",
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		t.Fatalf("Failed to marshal AutoEarnMetadata: %v", err)
	}

	// Verify the type field is present
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	if result["type"] != "AutoEarn" {
		t.Errorf("Expected type 'AutoEarn', got %v", result["type"])
	}

	t.Logf("Marshaled AutoEarnMetadata: %s", string(data))
}

func TestBuildSettlementOrchestrationRequestMarshalUnmarshal(t *testing.T) {
	settlementAmount := big.NewInt(1000000)
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")
	payerAddress := common.HexToAddress("0x9876543210987654321098765432109876543210")

	dueDate := time.Now().Add(30 * 24 * time.Hour)
	maxRuns := uint64(5)

	note := "Test payment"
	payerWalletAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")

	paymentMetadata := &PaymentRequestMetadata{
		Token:              "USDC",
		AmountDue:          "1.00",
		Currency:           "USD",
		DueDate:            &dueDate,
		Note:               &note,
		FromAccountAddress: recipientAddress,
		Payer: &Payer{
			Name:    "John Doe",
			Address: payerWalletAddress,
		},
		Source: "QuickBooks",
	}

	var metadataInterface OrchestrationMetadata = paymentMetadata

	req := &BuildSettlementOrchestrationRequest{
		Params: &SettlementParams{
			AcceptedTokens: map[ChainID][]common.Address{
				1: {usdcAddress},
			},
			SettlementChainId: 1,
			SettlementToken:   usdcAddress,
			SettlementAmount:  hexutil.Big(*settlementAmount),
			RecipientAddress:  recipientAddress,
		},
		PayerAddress: &payerAddress,
		Metadata:     &metadataInterface,
		DueDate:      &dueDate,
		MaxRuns:      &maxRuns,
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal BuildSettlementOrchestrationRequest: %v", err)
	}

	t.Logf("Marshaled request: %s", string(data))

	// Unmarshal back
	var decoded BuildSettlementOrchestrationRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal BuildSettlementOrchestrationRequest: %v", err)
	}

	// Verify the params type
	if _, ok := decoded.Params.(*SettlementParams); !ok {
		t.Errorf("Expected params to be *SettlementParams, got %T", decoded.Params)
	}

	// Verify the metadata type
	if decoded.Metadata != nil {
		if _, ok := (*decoded.Metadata).(*PaymentRequestMetadata); !ok {
			t.Errorf("Expected metadata to be *PaymentRequestMetadata, got %T", *decoded.Metadata)
		}
	} else {
		t.Error("Expected metadata to be present")
	}

	// Verify other fields
	if decoded.PayerAddress == nil || *decoded.PayerAddress != payerAddress {
		t.Errorf("PayerAddress mismatch")
	}

	if decoded.MaxRuns == nil || *decoded.MaxRuns != maxRuns {
		t.Errorf("MaxRuns mismatch")
	}
}

func TestBuildSettlementOrchestrationRequestWithVaultParams(t *testing.T) {
	withdrawAmount := big.NewInt(5000000000)
	vaultAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	autoEarnMetadata := &AutoEarnMetadata{
		TokenSymbol:  "USDC",
		VaultAddress: vaultAddress,
		VaultName:    "Aave USDC Vault",
	}

	var metadataInterface OrchestrationMetadata = autoEarnMetadata

	req := &BuildSettlementOrchestrationRequest{
		Params: &VaultWithdrawSettlementParams{
			VaultAddress:         vaultAddress,
			VaultUnderlyingToken: usdcAddress,
			VaultChainId:         1,
			SettlementChainId:    1,
			SettlementToken:      usdcAddress,
			RecipientAddress:     recipientAddress,
			WithdrawAmount:       hexutil.Big(*withdrawAmount),
		},
		Metadata: &metadataInterface,
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal BuildSettlementOrchestrationRequest: %v", err)
	}

	t.Logf("Marshaled vault request: %s", string(data))

	// Unmarshal back
	var decoded BuildSettlementOrchestrationRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal BuildSettlementOrchestrationRequest: %v", err)
	}

	// Verify the params type
	vaultParams, ok := decoded.Params.(*VaultWithdrawSettlementParams)
	if !ok {
		t.Fatalf("Expected params to be *VaultWithdrawSettlementParams, got %T", decoded.Params)
	}

	// Verify vault address
	if vaultParams.VaultAddress != vaultAddress {
		t.Errorf("VaultAddress mismatch: got %s, want %s", vaultParams.VaultAddress.Hex(), vaultAddress.Hex())
	}

	// Verify the metadata type
	if decoded.Metadata != nil {
		if _, ok := (*decoded.Metadata).(*AutoEarnMetadata); !ok {
			t.Errorf("Expected metadata to be *AutoEarnMetadata, got %T", *decoded.Metadata)
		}
	} else {
		t.Error("Expected metadata to be present")
	}
}

func TestBuildSettlementOrchestrationRequestWithoutMetadata(t *testing.T) {
	settlementAmount := big.NewInt(1000000)
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	req := &BuildSettlementOrchestrationRequest{
		Params: &SettlementParams{
			AcceptedTokens: map[ChainID][]common.Address{
				1: {usdcAddress},
			},
			SettlementChainId: 1,
			SettlementToken:   usdcAddress,
			SettlementAmount:  hexutil.Big(*settlementAmount),
			RecipientAddress:  recipientAddress,
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal BuildSettlementOrchestrationRequest: %v", err)
	}

	t.Logf("Marshaled request without metadata: %s", string(data))

	// Unmarshal back
	var decoded BuildSettlementOrchestrationRequest
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal BuildSettlementOrchestrationRequest: %v", err)
	}

	// Verify metadata is nil
	if decoded.Metadata != nil {
		t.Error("Expected metadata to be nil")
	}
}
