package client

import (
	"context"
	"encoding/json"
	"fmt"
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
	Metadata         json.RawMessage              `json:"metadata,omitempty"`
	DueDate          *time.Time                   `json:"dueDate,omitempty"`
	MaxRuns          *uint64                      `json:"maxRuns,omitempty"`
}

type BuildInstructionResponse struct {
	Address       common.Address `json:"address"`
	ChainID       ChainID        `json:"chainId"`
	Salt          U256           `json:"salt"`
	MaxExecutions U256           `json:"maxExecutions"`
	Action        common.Address `json:"action"`
	Arguments     hexutil.Bytes  `json:"arguments"`
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

type Signature struct {
	V uint8  `json:"v"`
	R string `json:"r"`
	S string `json:"s"`
}

type NewInstructionRequest struct {
	Address             common.Address `json:"address"`
	ChainID             ChainID        `json:"chainId"`
	Salt                U256           `json:"salt"`
	MaxExecutions       U256           `json:"maxExecutions"`
	Action              common.Address `json:"action"`
	Arguments           hexutil.Bytes  `json:"arguments"`
	ActivationSignature Signature      `json:"activationSignature"`
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
) error {
	return c.postJSON(ctx, "/orchestration/new", req, nil)
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
		ctx,
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, err
	}

	// Step 2: Sign all instructions with EIP-712
	instructions, err := c.signInstructions(
		ctx,
		buildOrchestrationResponse.Instructions,
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("sign instructions: %w", err)
	}

	// Step 3: Sign all completion instructions with EIP-712
	completionInstructions, err := c.signInstructions(
		ctx,
		buildOrchestrationResponse.CompletionInstructions,
		buildOrchestrationResponse.SubOrgID,
		buildOrchestrationResponse.EphemeralWalletAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("sign completion instructions: %w", err)
	}

	result := &NewOrchestrationRequest{
		RequestID:              buildOrchestrationResponse.RequestID,
		SignedAuthorization:    hexutil.Encode(signedAuthorization),
		CompletionInstructions: completionInstructions,
		Instructions:           instructions,
	}

	return result, nil
}

// signAuthorization creates and signs an EIP-7702 authorization, returning the RLP-encoded result
func (c *Client) signAuthorization(
	ctx context.Context,
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
	authSig, err := c.TKSignEIP7702(ctx, authorization, subOrgID, walletAddress)
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
	ctx context.Context,
	buildInstructions []BuildInstructionResponse,
	subOrgID string,
	walletAddress common.Address,
) ([]NewInstructionRequest, error) {
	if len(buildInstructions) == 0 {
		return []NewInstructionRequest{}, nil
	}

	// Build all TypedData payloads
	typedDataList := make([]map[string]interface{}, len(buildInstructions))
	for i, instr := range buildInstructions {
		typedData, err := BuildTypedDataForAction(instr, c.otimDelegateAddr)
		if err != nil {
			return nil, fmt.Errorf("build typed data for instruction %d: %w", i, err)
		}
		typedDataList[i] = typedData
	}

	// Sign all instructions in a single batch call to Turnkey
	signatures, err := c.TKSignEIP712Batch(ctx, typedDataList, subOrgID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("batch sign instructions: %w", err)
	}

	// Pack signatures and build instruction requests
	instructions := make([]NewInstructionRequest, len(buildInstructions))
	for i, instr := range buildInstructions {
		sig := signatures[i]

		instructions[i] = NewInstructionRequest{
			Address:       instr.Address,
			ChainID:       instr.ChainID,
			Salt:          instr.Salt,
			MaxExecutions: instr.MaxExecutions,
			Action:        instr.Action,
			Arguments:     instr.Arguments,
			ActivationSignature: Signature{
				V: sig.V,
				R: sig.R.Hex(),
				S: sig.S.Hex(),
			},
		}
	}

	return instructions, nil
}
