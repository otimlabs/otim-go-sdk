package signer

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/tkhq/go-sdk/pkg/api/models"
)

// Signer is an interface for signing data
type Signer interface {
	// Sign signs data using the signer's private key
	Sign(data []byte) ([]byte, error)
	// TKSign signs data using the Turnkey API, signer's private key acts as an api key.
	TKSign(data []byte, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error)
}

// Signature is a struct representing an Ethereum ECDSA signature.
type Signature struct {
	V uint8
	R uint256.Int
	S uint256.Int
}

// SigFromTurnkeyResult converts a Turnkey SignRawPayloadResult to a Signature
// TODO: Needs to be tested.
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

	// R and S are hex strings (0x-prefixed); convert to uint256.Int
	r, err := uint256.FromHex(*res.R)
	if err != nil {
		return nil, fmt.Errorf("parse R: %w", err)
	}
	s, err := uint256.FromHex(*res.S)
	if err != nil {
		return nil, fmt.Errorf("parse S: %w", err)
	}

	return &Signature{
		V: uint8(vParsed),
		R: *r,
		S: *s,
	}, nil
}
