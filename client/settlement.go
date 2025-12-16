package client

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
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

// TODO: Likely needs to be moved out from a client to a workflow helper.
func (c *Client) NewOrchestrationFromBuild(
	ctx context.Context,
	buildOrchestrationResponse *BuildOrchestrationResponse,
) (*NewOrchestrationRequest, error) {
	authorization := types.SetCodeAuthorization{
		ChainID: *uint256.NewInt(0),
		Nonce:   0,
		Address: c.otimDelegateAddr,
		// Those should just be set to 0
		V: 0,
		R: *uint256.NewInt(0),
		S: *uint256.NewInt(0),
	}

	hash := authorization.SigHash()

	authorizationSignature, err := c.TKSign(hash[:], buildOrchestrationResponse.SubOrgID, buildOrchestrationResponse.EphemeralWalletAddress)
	if err != nil {
		return nil, err
	}

	// TODO: Figure out how to encode this so it matches alloy format.
	authorization.V = authorizationSignature.V
	authorization.R = authorizationSignature.R
	authorization.S = authorizationSignature.S

	return nil, nil
}
