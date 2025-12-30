package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

type ChainID uint64

type U256 = hexutil.Big

type BuildSettlementOrchestrationRequest struct {
	AcceptedTokens   map[ChainID][]common.Address `json:"acceptedTokens"`
	SettlementChain  ChainID                      `json:"settlementChain"`
	SettlementToken  common.Address               `json:"settlementToken"`
	SettlementAmount U256                         `json:"settlementAmount"`
	PayerAddress     *common.Address              `json:"payerAddress,omitempty"`
	RecipientAddress common.Address               `json:"recipientAddress"`
	Metadata         json.RawMessage              `json:"metadata"`
	DueDate          *time.Time                   `json:"dueDate,omitempty"`
	MaxRuns          *uint64                      `json:"maxRuns,omitempty"`
}

type BuildInstructionResponse struct {
	Address       common.Address `json:"address"`
	ChainID       ChainID        `json:"chainId"`
	Salt          U256           `json:"salt"`
	MaxExecutions U256           `json:"maxExecutions"`
	Action        common.Address `json:"action"`
	Arguments     []byte         `json:"arguments"`
	ActionName    string         `json:"actionName"`
}

type BuildOrchestrationResponse struct {
	RequestID              string                     `json:"requestId"`
	SubOrgID               string                     `json:"subOrgId"`
	WalletID               string                     `json:"walletId"`
	EphemeralWalletAddress common.Address             `json:"ephemeralWalletAddress"`
	CompletionInstructions []BuildInstructionResponse `json:"completionInstructions"`
	Instructions           []BuildInstructionResponse `json:"instructions"`
}

type NewInstructionRequest struct {
	Address             common.Address `json:"address"`
	ChainID             ChainID        `json:"chainId"`
	Salt                U256           `json:"salt"`
	MaxExecutions       U256           `json:"maxExecutions"`
	Action              common.Address `json:"action"`
	Arguments           []byte         `json:"arguments"`
	ActivationSignature []byte         `json:"activationSignature"`
	Nickname            *string        `json:"nickname,omitempty"`
}

type NewOrchestrationRequest struct {
	RequestID              string                  `json:"requestId"`
	SignedAuthorization    string                  `json:"signedAuthorization"` // RLP-encoded SignedAuthorization as hex string
	CompletionInstructions []NewInstructionRequest `json:"completionInstructions"`
	Instructions           []NewInstructionRequest `json:"instructions"`
}

func (c *Client) BuildSettlementOrchestration(
	ctx context.Context,
	req *BuildSettlementOrchestrationRequest,
) (*BuildOrchestrationResponse, error) {
	var result BuildOrchestrationResponse
	if err := c.postJSON(ctx, "/orchestration/build/settlement", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) NewOrchestration(
	ctx context.Context,
	req *NewOrchestrationRequest,
) (*BuildOrchestrationResponse, error) {

	var result BuildOrchestrationResponse
	if err := c.postJSON(ctx, "/orchestration/new", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}



// NewOrchestrationFromBuild creates a NewOrchestrationRequest from a BuildOrchestrationResponse
// by signing the authorization and all instructions with EIP-712 signatures via Turnkey.
// Action types are determined from the ActionName field in BuildInstructionResponse.
// If ActionName is not provided, the function will attempt to detect action types from the
// instruction arguments as a fallback.
func (c *Client) NewOrchestrationFromBuild(
	ctx context.Context,
	buildOrchestrationResponse *BuildOrchestrationResponse,
) (*NewOrchestrationRequest, error) {
	// Step 1: Sign the EIP-7702 authorization
	signedAuthorization, err := c.signAuthorization(
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, err
	}

	// Step 2: Sign all instructions with EIP-712
	instructions, err := c.signInstructions(
		buildOrchestrationResponse.Instructions,
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("sign instructions: %w", err)
	}

	// Step 3: Sign all completion instructions with EIP-712
	completionInstructions, err := c.signInstructions(
		buildOrchestrationResponse.CompletionInstructions,
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("sign completion instructions: %w", err)
	}

	return &NewOrchestrationRequest{
		RequestID:              buildOrchestrationResponse.RequestID,
		SignedAuthorization:    hexutil.Encode(signedAuthorization),
		CompletionInstructions: completionInstructions,
		Instructions:           instructions,
	}, nil
}

// signAuthorization creates and signs an EIP-7702 authorization, returning the RLP-encoded result
func (c *Client) signAuthorization(
	subOrgID string,
	walletAddress common.Address,
) ([]byte, error) {
	// Create the EIP-7702 authorization
	authorization := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(0),
		Nonce:   0,
		Address: c.otimDelegateAddr,
		V:       0,
		R:       *uint256.NewInt(0),
		S:       *uint256.NewInt(0),
	}

	// Sign the authorization via Turnkey
	authSig, err := c.TKSignEIP7702(authorization, subOrgID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("sign authorization: %w", err)
	}

	// Apply the signature to the authorization
	authorization.V = authSig.V
	authorization.R = authSig.R
	authorization.S = authSig.S

	// RLP encode the signed authorization
	signedAuthorization, err := rlp.EncodeToBytes(authorization)
	if err != nil {
		return nil, fmt.Errorf("encode authorization: %w", err)
	}

	return signedAuthorization, nil
}

// signInstructions signs a list of instructions with EIP-712 signatures.
// Action types are determined in the following priority order:
// 1. ActionName field from BuildInstructionResponse (preferred)
// 2. Automatic detection from instruction arguments (fallback)
func (c *Client) signInstructions(
	buildInstructions []BuildInstructionResponse,
	subOrgID string,
	walletAddress common.Address,
) ([]NewInstructionRequest, error) {
	instructions := make([]NewInstructionRequest, len(buildInstructions))

	for i, instr := range buildInstructions {
		// Determine the action type
		var actionType ActionType
		var err error

		// Use ActionName if available
		actionType, err = ActionTypeFromName(instr.ActionName)
		if err != nil {
			return nil, fmt.Errorf("invalid action name for instruction %d: %w", i, err)
		}
		

		// Decode the arguments to get the typed struct
		actionArgs, err := DecodeArguments(actionType, instr.Arguments)
		if err != nil {
			return nil, fmt.Errorf("decode arguments for instruction %d: %w", i, err)
		}

		// Build EIP-712 TypedData
		chainID := big.NewInt(int64(instr.ChainID))
		typedData, err := BuildTypedDataForAction(actionType, instr, actionArgs, chainID, c.otimDelegateAddr)
		if err != nil {
			return nil, fmt.Errorf("build typed data for instruction %d: %w", i, err)
		}

		// Sign via Turnkey - Turnkey will handle the EIP-712 hashing internally
		sig, err := c.TKSignEIP712(typedData, subOrgID, walletAddress)
		if err != nil {
			return nil, fmt.Errorf("sign instruction %d: %w", i, err)
		}

		// Pack signature into 65-byte format: R (32 bytes) || S (32 bytes) || V (1 byte)
		sigBytes := make([]byte, 65)
		rBytes := sig.R.Bytes32()
		sBytes := sig.S.Bytes32()
		copy(sigBytes[0:32], rBytes[:])
		copy(sigBytes[32:64], sBytes[:])
		sigBytes[64] = sig.V

		instructions[i] = NewInstructionRequest{
			Address:             instr.Address,
			ChainID:             instr.ChainID,
			Salt:                instr.Salt,
			MaxExecutions:       instr.MaxExecutions,
			Action:              instr.Action,
			Arguments:           instr.Arguments,
			ActivationSignature: sigBytes,
		}
	}

	return instructions, nil
}
