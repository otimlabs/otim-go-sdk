package client

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/otimlabs/otim-go-sdk/signer"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Environment variable names
const (
	envAPIURL     = "OTIM_API_URL"
	envAPIKey     = "OTIM_API_KEY"
	envPrivateKey = "OTIM_PRIVATE_KEY"
)

// testConfig holds the configuration loaded from environment variables
type testConfig struct {
	apiURL              string
	apiKey              string
	developerPrivateKey string
}

// loadTestConfig loads configuration from environment variables
// Returns an error if any required variable is missing
func loadTestConfig(t *testing.T) (*testConfig, error) {
	t.Helper()

	config := &testConfig{
		apiURL:              os.Getenv(envAPIURL),
		apiKey:              os.Getenv(envAPIKey),
		developerPrivateKey: os.Getenv(envPrivateKey),
	}

	// Validate required environment variables
	missingVars := []string{}
	if config.apiURL == "" {
		missingVars = append(missingVars, envAPIURL)
	}
	if config.apiKey == "" {
		missingVars = append(missingVars, envAPIKey)
	}
	if config.developerPrivateKey == "" {
		missingVars = append(missingVars, envPrivateKey)
	}

	if len(missingVars) > 0 {
		t.Fatalf("Missing required environment variables: %v", missingVars)
	}

	return config, nil
}

// cleanupWallets lists and deletes all wallets in the specified Turnkey sub-organization.
// This function logs cleanup actions but does not fail the test if cleanup encounters errors.
func cleanupWallets(t *testing.T, ctx context.Context, client *Client, subOrgID string) {
	t.Helper()

	t.Logf("Starting wallet cleanup for sub-org: %s", subOrgID)

	// List all wallets in the sub-organization
	walletIds, err := client.TKListWallets(ctx, subOrgID)
	if err != nil {
		t.Logf("Warning: Failed to list wallets during cleanup: %v", err)
		return
	}

	if len(walletIds) == 0 {
		t.Log("No wallets to clean up")
		return
	}

	t.Logf("Found %d wallet(s) to delete: %v", len(walletIds), walletIds)

	// Delete all wallets
	err = client.TKDeleteWallets(ctx, subOrgID, walletIds)
	if err != nil {
		t.Logf("Warning: Failed to delete wallets during cleanup: %v", err)
		return
	}

	t.Logf("Successfully deleted %d wallet(s)", len(walletIds))
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
		SettlementChain:  1,           // Ethereum mainnet
		SettlementToken:  usdcAddress, // USDC
		SettlementAmount: hexutil.Big(*settlementAmount),
		RecipientAddress: recipientAddress,
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

	// Initialize EthSigner with Turnkey
	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	// Create Client
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Fetch sub-organization ID from Otim API
	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	// Cleanup wallets before test
	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID)

	// Setup cleanup after test (runs even if test fails)
	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID)
	})

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
