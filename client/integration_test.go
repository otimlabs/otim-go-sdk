package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/otimlabs/otim-go-sdk/signer"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-openapi/runtime"
	"github.com/tkhq/go-sdk/pkg/api/client/wallets"
	"github.com/tkhq/go-sdk/pkg/api/models"
	"github.com/tkhq/go-sdk/pkg/util"
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

// deleteWalletsInternal is a test-only helper that deletes wallets from a Turnkey sub-organization.
// This function is intentionally not part of the public SDK API to prevent accidental data loss.
// It sets deleteWithoutExport=true to allow deletion of non-exported wallets.
func deleteWalletsInternal(ctx context.Context, ethSigner *signer.EthSigner, subOrgID string, walletIds []string) error {
	if len(walletIds) == 0 {
		return nil
	}

	tkClient := ethSigner.TKClientForTesting()

	deleteWithoutExport := true
	params := wallets.NewDeleteWalletsParams().WithBody(&models.DeleteWalletsRequest{
		OrganizationID: &subOrgID,
		TimestampMs:    util.RequestTimestamp(),
		Parameters: &models.DeleteWalletsIntent{
			WalletIds:           walletIds,
			DeleteWithoutExport: &deleteWithoutExport,
		},
		Type: (*string)(models.ActivityTypeDeleteWallets.Pointer()),
	})

	withContext := func(ctx context.Context) func(*runtime.ClientOperation) {
		return func(op *runtime.ClientOperation) {
			op.Context = ctx
		}
	}

	_, err := tkClient.V0().Wallets.DeleteWallets(params, tkClient.Authenticator, withContext(ctx))
	if err != nil {
		return fmt.Errorf("turnkey delete wallets: %w", err)
	}

	return nil
}

// cleanupWallets lists and deletes all wallets in the specified Turnkey sub-organization.
// This function logs cleanup actions but does not fail the test if cleanup encounters errors.
func cleanupWallets(t *testing.T, ctx context.Context, client *Client, subOrgID string, ethSigner *signer.EthSigner) {
	t.Helper()

	t.Logf("Starting wallet cleanup for sub-org: %s", subOrgID)

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

	err = deleteWalletsInternal(ctx, ethSigner, subOrgID, walletIds)
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
		Params: &SettlementParams{
			AcceptedTokens:    acceptedTokens,
			SettlementChainId: 1,           // Ethereum mainnet
			SettlementToken:   usdcAddress, // USDC
			SettlementAmount:  hexutil.Big(*settlementAmount),
			RecipientAddress:  recipientAddress,
		},
	}
}

// createTestVaultWithdrawBuildRequest creates a VaultWithdrawSettlement request for testing
func createTestVaultWithdrawSettlementBuildRequest() *BuildSettlementOrchestrationRequest {
	withdrawAmount := new(big.Int)
	withdrawAmount.SetString("5000000000", 10) // 5000 USDC (6 decimals)

	// Steakhouse USDC
	vaultAddress := common.HexToAddress("0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB")

	// USDC contract address on Ethereum mainnet
	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// Hardcoded recipient address for testing
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	return &BuildSettlementOrchestrationRequest{
		Params: &VaultWithdrawSettlementRequest{
			VaultAddress:      vaultAddress,
			VaultChainId:      1, // Ethereum mainnet
			SettlementChainId: 1,
			SettlementToken:   usdcAddress,
			RecipientAddress:  recipientAddress,
			WithdrawAmount:    hexutil.Big(*withdrawAmount),
		},
	}
}

// createTestSettlementVaultDepositBuildRequest creates a SettlementVaultDeposit request for testing
func createTestSettlementVaultDepositBuildRequest() *BuildSettlementOrchestrationRequest {
	depositAmount := new(big.Int)
	depositAmount.SetString("1000000000", 10) // 1000 USDC (6 decimals)

	// Steakhouse USDC
	vaultAddress := common.HexToAddress("0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB")

	// USDC contract addresses
	usdcEthereum := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	usdcOptimism := common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607")

	// Hardcoded recipient address for testing
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	// Accepted tokens: allow USDC on Ethereum and Optimism
	acceptedTokens := map[ChainID][]common.Address{
		1:  {usdcEthereum}, // Ethereum mainnet
		10: {usdcOptimism}, // Optimism
	}

	return &BuildSettlementOrchestrationRequest{
		Params: &SettlementVaultDepositRequest{
			AcceptedTokens:   acceptedTokens,
			VaultChainId:     1, // Ethereum mainnet
			VaultAddress:     vaultAddress,
			DepositAmount:    hexutil.Big(*depositAmount),
			RecipientAddress: recipientAddress,
		},
	}
}

// createTestVaultMigrateBuildRequest creates a VaultMigrate request for testing
func createTestVaultMigrateBuildRequest() *BuildSettlementOrchestrationRequest {
	withdrawAmount := new(big.Int)
	withdrawAmount.SetString("1000000000", 10) // 1000 USDC (6 decimals)

	// Steakhouse USDC (ETH mainnet)
	sourceVaultAddress := common.HexToAddress("0xBEEF01735c132Ada46AA9aA4c54623cAA92A64CB")

	// Steakhouse USDC (Base mainnet)
	destVaultAddress := common.HexToAddress("0xbeeF010f9cb27031ad51e3333f9aF9C6B1228183")

	// Recipient address
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	return &BuildSettlementOrchestrationRequest{
		Params: &VaultMigrateRequest{
			SourceVaultAddress: sourceVaultAddress,
			SourceVaultChainId: 1, // Ethereum mainnet
			WithdrawAmount:     hexutil.Big(*withdrawAmount),
			DestVaultAddress:   destVaultAddress,
			DestVaultChainId:   8453, // Base mainnet
			RecipientAddress:   recipientAddress,
		},
	}
}

// createTestCoveredVaultDepositBuildRequest creates a CoveredVaultDeposit request for testing
func createTestCoveredVaultDepositBuildRequest() *BuildSettlementOrchestrationRequest {
	depositAmount := new(big.Int)
	depositAmount.SetString("1000000000", 10) // 1000 tokens (assuming 6 decimals)

	// ERC7540 covered vault address (example address)
	coveredVaultAddress := common.HexToAddress("0xC0FFEE01735c132Ada46AA9aA4c54623cAA92A64CB")

	// Recipient address
	recipientAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb0")

	return &BuildSettlementOrchestrationRequest{
		Params: &CoveredVaultDepositRequest{
			CoveredVaultAddress: coveredVaultAddress,
			CoveredVaultChainId: 1, // Ethereum mainnet
			DepositAmount:       hexutil.Big(*depositAmount),
			RecipientAddress:    recipientAddress,
		},
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

	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)

	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	t.Log("Creating client")
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID, ethSigner)

	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID, ethSigner)
	})

	t.Log("Phase 2: Building settlement orchestration request")
	buildRequest := createTestBuildRequest()

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

	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	t.Log("✓ Integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}

// TestVaultWithdrawSettlementIntegration tests the full vault withdrawal settlement orchestration flow:
// 1. Call BuildSettlementOrchestration API with VaultWithdrawSettlementRequest
// 2. Sign the returned delegation and instructions with Turnkey
// 3. Submit to NewOrchestration API
func TestVaultWithdrawSettlementIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)

	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	t.Log("Creating client")
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID, ethSigner)

	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID, ethSigner)
	})

	t.Log("Phase 2: Building vault withdrawal settlement orchestration request")
	buildRequest := createTestVaultWithdrawSettlementBuildRequest()

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

	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	t.Log("✓ Vault withdrawal settlement integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}

// TestSettlementVaultDepositIntegration tests the full settlement vault deposit orchestration flow:
// 1. Call BuildSettlementOrchestration API with SettlementVaultDepositRequest
// 2. Sign the returned delegation and instructions with Turnkey
// 3. Submit to NewOrchestration API
func TestSettlementVaultDepositIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)

	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	t.Log("Creating client")
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID, ethSigner)

	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID, ethSigner)
	})

	t.Log("Phase 2: Building settlement vault deposit orchestration request")
	buildRequest := createTestSettlementVaultDepositBuildRequest()

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

	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	t.Log("✓ Settlement vault deposit integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}

// TestVaultMigrateIntegration tests the full vault migration orchestration flow:
// 1. Call BuildSettlementOrchestration API with VaultMigrateRequest
// 2. Sign the returned delegation and instructions with Turnkey
// 3. Submit to NewOrchestration API
func TestVaultMigrateIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)

	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	t.Log("Creating client")
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID, ethSigner)

	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID, ethSigner)
	})

	t.Log("Phase 2: Building vault migration orchestration request")
	buildRequest := createTestVaultMigrateBuildRequest()

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

	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	t.Log("✓ Vault migration integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}

// TestCoveredVaultDepositIntegration tests the full covered vault deposit orchestration flow:
// 1. Call BuildSettlementOrchestration API with CoveredVaultDepositRequest
// 2. Sign the returned delegation and instructions with Turnkey
// 3. Submit to NewOrchestration API
func TestCoveredVaultDepositIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	ctx := context.Background()

	t.Log("Phase 1: Loading configuration from environment variables")
	config, err := loadTestConfig(t)
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	t.Logf("API URL: %s", config.apiURL)

	t.Log("Initializing EthSigner with Turnkey credentials")
	ethSigner, err := signer.NewEthSigner(config.developerPrivateKey)
	if err != nil {
		t.Fatalf("Failed to create EthSigner: %v", err)
	}

	t.Log("Creating client")
	client, err := NewClient(ethSigner, config.apiURL, config.apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	t.Log("Fetching sub-organization ID from Otim API")
	subOrgID, err := client.GetSubOrganization(ctx)
	if err != nil {
		t.Fatalf("Failed to fetch sub-organization ID: %v", err)
	}
	t.Logf("Sub-organization ID: %s", subOrgID)

	t.Log("Cleaning up wallets before test")
	cleanupWallets(t, ctx, client, subOrgID, ethSigner)

	t.Cleanup(func() {
		t.Log("Cleaning up wallets after test")
		cleanupWallets(t, context.Background(), client, subOrgID, ethSigner)
	})

	t.Log("Phase 2: Building covered vault deposit orchestration request")
	buildRequest := createTestCoveredVaultDepositBuildRequest()

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

	t.Log("Phase 5: Submitting to NewOrchestration API")
	err = client.NewOrchestration(ctx, newRequest)
	if err != nil {
		t.Fatalf("NewOrchestration failed: %v", err)
	}

	t.Log("✓ Covered vault deposit integration test completed successfully")
	t.Logf("✓ Full orchestration flow validated for RequestID: %s", buildResponse.RequestID)
}
