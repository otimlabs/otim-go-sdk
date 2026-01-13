package client

import (
	"encoding/json"
	"testing"
)

func TestBuildOrchestrationResponseDecoding(t *testing.T) {
	jsonData := `{"requestId":"019ba380-0cfa-7550-92a4-edf1a62b03d4","subOrgId":"73ca31ba-540a-49b5-9848-e6ed41501b79","walletId":"e6055fae-6d22-53e8-a5f9-5f7659b3a386","ephemeralWalletAddress":"0x762858df70e61c808e9345bae3de735db7816cd5","completionInstructions":[{"address":"0x762858df70e61c808e9345bae3de735db7816cd5","chainId":1,"salt":"0x0","maxExecutions":"0x0","action":"0x992267fd20d9f928e533a08cd61f6d57220c42d2","arguments":"0x000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000742d35cc6634c0532925a3b844bc9e7595f0beb000000000000000000000000000000000000000000000000000000000000f3e5800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","actionName":"sweepERC20"}],"instructions":[]}`

	var response BuildOrchestrationResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify the fields
	if response.RequestID != "019ba380-0cfa-7550-92a4-edf1a62b03d4" {
		t.Errorf("RequestID mismatch: got %s", response.RequestID)
	}

	if len(response.CompletionInstructions) != 1 {
		t.Fatalf("Expected 1 completion instruction, got %d", len(response.CompletionInstructions))
	}

	instr := response.CompletionInstructions[0]
	if instr.ActionName != "sweepERC20" {
		t.Errorf("ActionName mismatch: got %s", instr.ActionName)
	}

	// Check that arguments were decoded properly (should be 256 bytes based on the hex string)
	expectedArgsLen := 256
	if len(instr.Arguments) != expectedArgsLen {
		t.Errorf("Arguments length mismatch: expected %d, got %d", expectedArgsLen, len(instr.Arguments))
	}
}
