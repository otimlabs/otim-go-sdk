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

// OrchestrationParams is the interface for settlement orchestration parameters
type OrchestrationParams interface {
	orchestrationParams()
}

// SettlementParams represents standard multi-chain settlement orchestration parameters
type SettlementParams struct {
	AcceptedTokens    map[ChainID][]common.Address `json:"acceptedTokens"`
	SettlementChainId ChainID                      `json:"settlementChainId"`
	SettlementToken   common.Address               `json:"settlementToken"`
	SettlementAmount  U256                         `json:"settlementAmount"`
	RecipientAddress  common.Address               `json:"recipientAddress"`
}

// VaultWithdrawSettlementParams represents vault withdrawal settlement orchestration parameters
type VaultWithdrawSettlementParams struct {
	VaultAddress         common.Address `json:"vaultAddress"`
	VaultUnderlyingToken common.Address `json:"vaultUnderlyingToken"`
	VaultChainId         ChainID        `json:"vaultChainId"`
	SettlementChainId    ChainID        `json:"settlementChainId"`
	SettlementToken      common.Address `json:"settlementToken"`
	RecipientAddress     common.Address `json:"recipientAddress"`
	WithdrawAmount       U256           `json:"withdrawAmount"`
}

// Implement OrchestrationParams interface
func (s *SettlementParams) orchestrationParams()              {}
func (v *VaultWithdrawSettlementParams) orchestrationParams() {}

// OrchestrationMetadata is the interface for orchestration metadata
type OrchestrationMetadata interface {
	orchestrationMetadata()
}

// Payer represents payer information for payment requests
type Payer struct {
	Name    string         `json:"name"`
	Address common.Address `json:"address"`
}

// PaymentRequestMetadata represents payment request metadata for an orchestration
type PaymentRequestMetadata struct {
	Token              string         `json:"token"`
	AmountDue          string         `json:"amountDue"`
	Currency           string         `json:"currency"`
	DueDate            *time.Time     `json:"dueDate,omitempty"`
	Note               *string        `json:"note,omitempty"`
	FromAccountAddress common.Address `json:"fromAccountAddress"`
	Payer              *Payer         `json:"payer,omitempty"`
	Source             string         `json:"source"`
	InvoiceId          *string        `json:"invoiceId,omitempty"`
	InvoiceNumber      *string        `json:"invoiceNumber,omitempty"`
	AttachmentUrl      *string        `json:"attachmentUrl,omitempty"`
	AttachmentName     *string        `json:"attachmentName,omitempty"`
}

// AutoEarnMetadata represents auto-earn vault deposit metadata for an orchestration
type AutoEarnMetadata struct {
	TokenSymbol  string         `json:"tokenSymbol"`
	VaultAddress common.Address `json:"vaultAddress"`
	VaultName    string         `json:"vaultName"`
}

// Implement OrchestrationMetadata interface
func (p *PaymentRequestMetadata) orchestrationMetadata() {}
func (a *AutoEarnMetadata) orchestrationMetadata()       {}

// BuildSettlementOrchestrationRequest is the request body for the /orchestration/build/settlement endpoint
type BuildSettlementOrchestrationRequest struct {
	Params       OrchestrationParams    `json:"params"`
	PayerAddress *common.Address        `json:"payerAddress,omitempty"`
	Metadata     *OrchestrationMetadata `json:"metadata,omitempty"`
	DueDate      *time.Time             `json:"dueDate,omitempty"`
	MaxRuns      *uint64                `json:"maxRuns,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for OrchestrationParams (externally tagged)
func (p SettlementParams) MarshalJSON() ([]byte, error) {
	type Alias SettlementParams
	return json.Marshal(map[string]interface{}{
		"settlement": (Alias)(p),
	})
}

// MarshalJSON implements custom JSON marshaling for VaultWithdrawSettlementParams (externally tagged)
func (v VaultWithdrawSettlementParams) MarshalJSON() ([]byte, error) {
	type Alias VaultWithdrawSettlementParams
	return json.Marshal(map[string]interface{}{
		"vaultWithdrawSettlement": (Alias)(v),
	})
}

// MarshalJSON implements custom JSON marshaling for PaymentRequestMetadata
func (p PaymentRequestMetadata) MarshalJSON() ([]byte, error) {
	type Alias PaymentRequestMetadata
	return json.Marshal(&struct {
		Type string `json:"type"`
		Alias
	}{
		Type:  "PaymentRequest",
		Alias: (Alias)(p),
	})
}

// MarshalJSON implements custom JSON marshaling for AutoEarnMetadata
func (a AutoEarnMetadata) MarshalJSON() ([]byte, error) {
	type Alias AutoEarnMetadata
	return json.Marshal(&struct {
		Type string `json:"type"`
		Alias
	}{
		Type:  "AutoEarn",
		Alias: (Alias)(a),
	})
}

// MarshalJSON implements custom JSON marshaling for BuildSettlementOrchestrationRequest
func (r BuildSettlementOrchestrationRequest) MarshalJSON() ([]byte, error) {
	type Alias BuildSettlementOrchestrationRequest
	aux := &struct {
		Params json.RawMessage `json:"params"`
		*Alias
	}{
		Alias: (*Alias)(&r),
	}

	// Marshal the params field
	paramsJSON, err := json.Marshal(r.Params)
	if err != nil {
		return nil, fmt.Errorf("marshal params: %w", err)
	}
	aux.Params = paramsJSON

	// Handle optional metadata
	if r.Metadata != nil {
		metadataJSON, err := json.Marshal(*r.Metadata)
		if err != nil {
			return nil, fmt.Errorf("marshal metadata: %w", err)
		}
		// We need to use a different approach to include metadata
		return json.Marshal(&struct {
			Params       json.RawMessage `json:"params"`
			PayerAddress *common.Address `json:"payerAddress,omitempty"`
			Metadata     json.RawMessage `json:"metadata,omitempty"`
			DueDate      *time.Time      `json:"dueDate,omitempty"`
			MaxRuns      *uint64         `json:"maxRuns,omitempty"`
		}{
			Params:       paramsJSON,
			PayerAddress: r.PayerAddress,
			Metadata:     metadataJSON,
			DueDate:      r.DueDate,
			MaxRuns:      r.MaxRuns,
		})
	}

	return json.Marshal(&struct {
		Params       json.RawMessage `json:"params"`
		PayerAddress *common.Address `json:"payerAddress,omitempty"`
		DueDate      *time.Time      `json:"dueDate,omitempty"`
		MaxRuns      *uint64         `json:"maxRuns,omitempty"`
	}{
		Params:       paramsJSON,
		PayerAddress: r.PayerAddress,
		DueDate:      r.DueDate,
		MaxRuns:      r.MaxRuns,
	})
}

// UnmarshalJSON implements custom JSON unmarshaling for BuildSettlementOrchestrationRequest
func (r *BuildSettlementOrchestrationRequest) UnmarshalJSON(data []byte) error {
	// First unmarshal into a temporary struct to get the raw params and metadata
	var raw struct {
		Params       json.RawMessage `json:"params"`
		PayerAddress *common.Address `json:"payerAddress,omitempty"`
		Metadata     json.RawMessage `json:"metadata,omitempty"`
		DueDate      *time.Time      `json:"dueDate,omitempty"`
		MaxRuns      *uint64         `json:"maxRuns,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Unmarshal params based on variant key (externally tagged)
	var paramsMap map[string]json.RawMessage
	if err := json.Unmarshal(raw.Params, &paramsMap); err != nil {
		return fmt.Errorf("unmarshal params map: %w", err)
	}

	if settlementData, ok := paramsMap["settlement"]; ok {
		var params SettlementParams
		if err := json.Unmarshal(settlementData, &params); err != nil {
			return fmt.Errorf("unmarshal settlement params: %w", err)
		}
		r.Params = &params
	} else if vaultData, ok := paramsMap["vaultWithdrawSettlement"]; ok {
		var params VaultWithdrawSettlementParams
		if err := json.Unmarshal(vaultData, &params); err != nil {
			return fmt.Errorf("unmarshal vault withdraw settlement params: %w", err)
		}
		r.Params = &params
	} else {
		return fmt.Errorf("unknown orchestration params type, expected 'settlement' or 'vaultWithdrawSettlement'")
	}

	// Unmarshal metadata if present
	if len(raw.Metadata) > 0 {
		var metadataType struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(raw.Metadata, &metadataType); err != nil {
			return fmt.Errorf("unmarshal metadata type: %w", err)
		}

		switch metadataType.Type {
		case "PaymentRequest":
			var metadata PaymentRequestMetadata
			if err := json.Unmarshal(raw.Metadata, &metadata); err != nil {
				return fmt.Errorf("unmarshal payment request metadata: %w", err)
			}
			var metadataInterface OrchestrationMetadata = &metadata
			r.Metadata = &metadataInterface
		case "AutoEarn":
			var metadata AutoEarnMetadata
			if err := json.Unmarshal(raw.Metadata, &metadata); err != nil {
				return fmt.Errorf("unmarshal auto earn metadata: %w", err)
			}
			var metadataInterface OrchestrationMetadata = &metadata
			r.Metadata = &metadataInterface
		default:
			return fmt.Errorf("unknown orchestration metadata type: %s", metadataType.Type)
		}
	}

	r.PayerAddress = raw.PayerAddress
	r.DueDate = raw.DueDate
	r.MaxRuns = raw.MaxRuns

	return nil
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
