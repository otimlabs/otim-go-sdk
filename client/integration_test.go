package client

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"testing"

	"otim-go-sdk/signer"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Environment variable names
const (
	envAPIURL            = "OTIM_API_URL"
	envAPIKey            = "OTIM_API_KEY"
	envTurnkeyPrivateKey = "TURNKEY_PRIVATE_KEY"
	envOtimDelegateAddr  = "OTIM_DELEGATE_ADDRESS"
)

// testConfig holds the configuration loaded from environment variables
type testConfig struct {
	apiURL            string
	apiKey            string
	turnkeyPrivateKey string
	otimDelegateAddr  common.Address
}

// loadTestConfig loads configuration from environment variables
// Returns an error if any required variable is missing
func loadTestConfig(t *testing.T) (*testConfig, error) {
	t.Helper()

	config := &testConfig{
		apiURL:            os.Getenv(envAPIURL),
		apiKey:            os.Getenv(envAPIKey),
		turnkeyPrivateKey: os.Getenv(envTurnkeyPrivateKey),
	}

	// Validate required environment variables
	missingVars := []string{}
	if config.apiURL == "" {
		missingVars = append(missingVars, envAPIURL)
	}
	if config.apiKey == "" {
		missingVars = append(missingVars, envAPIKey)
	}
	if config.turnkeyPrivateKey == "" {
		missingVars = append(missingVars, envTurnkeyPrivateKey)
	}

	delegateAddr := os.Getenv(envOtimDelegateAddr)
	if delegateAddr == "" {
		missingVars = append(missingVars, envOtimDelegateAddr)
	} else {
		config.otimDelegateAddr = common.HexToAddress(delegateAddr)
	}

	if len(missingVars) > 0 {
		t.Fatalf("Missing required environment variables: %v", missingVars)
	}

	return config, nil
}

// createTestBuildRequest creates a sample BuildSettlementOrchestrationRequest for testing
func createTestBuildRequest() *BuildSettlementOrchestrationRequest {
	// Test data: simple settlement request
	settlementAmount := new(big.Int)
	settlementAmount.SetString("1000000", 10) // 1 USDC (6 decimals)

	// USDC contract address on Ethereum mainnet
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// Hardcoded recipient address for testing
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	// Accepted tokens: allow USDC on Ethereum mainnet (chain 1)
	acceptedTokens := map[ChainID][]common.Address{
		1: {usdcAddress},
	}

	return &BuildSettlementOrchestrationRequest{
		AcceptedTokens:   acceptedTokens,
		SettlementChain:  1,               // Ethereum mainnet
		SettlementToken:  usdcAddress,     // USDC
		SettlementAmount: hexutil.Big(*settlementAmount),
		RecipientAddress: recipientAddress,
		Metadata:         json.RawMessage(`{"test": "integration-test"}`),
	}
}

// TestSettlementOrchestrationIntegration tests the full settlement orchestration flow:
// 1. Call BuildSettlementOrchestration API
// 2. Sign the returned delegation and instructions with Turnkey
// 3. Submit to NewOrchestration API
func TestSettlementOrchestrationIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	// Phase 1: Setup - Load configuration and initialize client
	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)
	t.Logf("Otim Delegate Address: %s", config.otimDelegateAddr.Hex())

	// Initialize EthSigner with Turnkey
	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.turnkeyPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	// Create Client
	client := NewClient(ethSigner, config.apiURL, config.apiKey, config.otimDelegateAddr)
	if client == nil {
		t.Fatal("Failed to create client")
	}

	// Phase 2: Build settlement orchestration request
	t.Log("Phase 2: Building settlement orchestration request")
	buildRequest := createTestBuildRequest()

	// Phase 3: Call BuildSettlementOrchestration API
	t.Log("Phase 3: Calling BuildSettlementOrchestration API")
	buildResponse, err := client.BuildSettlementOrchestration(ctx, buildRequest)
	if err != nil {
		t.Fatalf("BuildSettlementOrchestration failed: %v", err)
	}

	// Validate build response
	if buildResponse == nil {
		t.Fatal("BuildSettlementOrchestration returned nil response")
	}
	if buildResponse.RequestID == "" {
		t.Error("BuildResponse missing RequestID")
	}
	if buildResponse.EphemeralWalletAddress == (common.Address{}) {
		t.Error("BuildResponse missing EphemeralWalletAddress")
	}

	t.Logf("Build Response - RequestID: %s", buildResponse.RequestID)
	t.Logf("Build Response - Ephemeral Wallet: %s", buildResponse.EphemeralWalletAddress.Hex())
	t.Logf("Build Response - Instructions: %d", len(buildResponse.Instructions))
	t.Logf("Build Response - Completion Instructions: %d", len(buildResponse.CompletionInstructions))

	// Phase 4: Sign delegation and instructions via Turnkey
	t.Log("Phase 4: Signing delegation and instructions with Turnkey")
	newRequest, err := client.NewOrchestrationFromBuild(ctx, buildResponse)
	if err != nil {
		t.Fatalf("NewOrchestrationFromBuild failed: %v", err)
	}

	// Validate signed request
	if newRequest == nil {
		t.Fatal("NewOrchestrationFromBuild returned nil")
	}
	if newRequest.RequestID != buildResponse.RequestID {
		t.Errorf("RequestID mismatch: got %s, want %s", newRequest.RequestID, buildResponse.RequestID)
	}
	if newRequest.SignedAuthorization == "" {
		t.Error("SignedAuthorization is empty")
	}

	t.Logf("Signed Authorization length: %d", len(newRequest.SignedAuthorization))

	// Validate all instructions have signatures
	for i, instr := range newRequest.Instructions {
		if len(instr.ActivationSignature) != 65 {
			t.Errorf("Instruction %d: ActivationSignature has wrong length: got %d, want 65", i, len(instr.ActivationSignature))
		}
	}
	for i, instr := range newRequest.CompletionInstructions {
		if len(instr.ActivationSignature) != 65 {
			t.Errorf("CompletionInstruction %d: ActivationSignature has wrong length: got %d, want 65", i, len(instr.ActivationSignature))
		}
	}

	t.Logf("All %d instructions signed successfully", len(newRequest.Instructions)+len(newRequest.CompletionInstructions))

	// Phase 5: Submit to NewOrchestration API
	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	// Success!
	t.Log("✓ Integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}