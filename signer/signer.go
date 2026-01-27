package signer

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-openapi/runtime"
	"github.com/holiman/uint256"
	"github.com/tkhq/go-sdk/pkg/api/models"
)

// withContext allows for passing the current context to Turnkey SDK client.
func withContext(ctx context.Context) func(*runtime.ClientOperation) {
	return func(op *runtime.ClientOperation) {
		op.Context = ctx
	}
}

// Signer is an interface for signing data
type Signer interface {
	// Sign signs data using the signer's private key
	Sign(data []byte) ([]byte, error)
	// TKSign signs data using the Turnkey API, signer's private key acts as an api key.
	TKSign(ctx context.Context, data []byte, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error)
	// TKSignEIP7702 signs EIP-7702 authorization using the Turnkey API
	TKSignEIP7702(ctx context.Context, authorization types.SetCodeAuthorization, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error)
	// TKSignEIP712 signs EIP-712 typed data using the Turnkey API
	TKSignEIP712(ctx context.Context, typedData map[string]interface{}, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error)
	// TKSignEIP712Batch signs multiple EIP-712 typed data payloads using the Turnkey batch API
	TKSignEIP712Batch(ctx context.Context, typedDataList []map[string]interface{}, subOrganizationId string, walletAccountAddress common.Address) ([]*Signature, error)
}

// Signature is a struct representing an Ethereum ECDSA signature.
type Signature struct {
	V uint8
	R uint256.Int
	S uint256.Int
}

// SigFromTurnkeyResult converts a Turnkey SignRawPayloadResult to a Signature
func SigFromTurnkeyResult(res *models.SignRawPayloadResult) (*Signature, error) {
	if res == nil || res.V == nil || res.S == nil || res.R == nil {
		return nil, fmt.Errorf("nil turnkey result or signature")
	}

	// V is provided as a string, parse as base-10 uint8
	v := *res.V
	vParsed, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("parse V: %w", err)
	}

	// R and S are hex strings; decode to bytes and use SetBytes
	// (uint256.FromHex rejects hex with leading zeros, which Turnkey may return)
	rHex := *res.R
	sHex := *res.S

	rBytes, err := hex.DecodeString(rHex)
	if err != nil {
		return nil, fmt.Errorf("decode R: %w", err)
	}
	sBytes, err := hex.DecodeString(sHex)
	if err != nil {
		return nil, fmt.Errorf("decode S: %w", err)
	}

	r := new(uint256.Int).SetBytes(rBytes)
	s := new(uint256.Int).SetBytes(sBytes)

	return &Signature{
		V: uint8(vParsed),
		R: *r,
		S: *s,
	}, nil
}

// SigsFromTurnkeyBatchResult converts a Turnkey SignRawPayloadsResult to a slice of Signatures
func SigsFromTurnkeyBatchResult(results []*models.SignRawPayloadResult) ([]*Signature, error) {
	if results == nil {
		return nil, fmt.Errorf("nil turnkey batch results")
	}

	signatures := make([]*Signature, len(results))
	for i, res := range results {
		sig, err := SigFromTurnkeyResult(res)
		if err != nil {
			return nil, fmt.Errorf("parse signature %d: %w", i, err)
		}
		signatures[i] = sig
	}

	return signatures, nil
}
