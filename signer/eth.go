package signer

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tkhq/go-sdk"
	"github.com/tkhq/go-sdk/pkg/api/client/signing"
	"github.com/tkhq/go-sdk/pkg/api/models"
	"github.com/tkhq/go-sdk/pkg/apikey"
	"github.com/tkhq/go-sdk/pkg/util"
)

// EthSigner implements the Signer interface using an ECDSA private key.
// It supports both local signing and remote signing via the Turnkey API.
type EthSigner struct {
	privateKey *ecdsa.PrivateKey
	tkClient   *sdk.Client
}

// EIP7702AuthorizationPayload represents the EIP-7702 authorization tuple for Turnkey signing
type EIP7702AuthorizationPayload struct {
	ChainID string `json:"chainId"`
	Address string `json:"address"`
	Nonce   string `json:"nonce"`
}

// NewEthSigner creates a new EthSigner from a hex-encoded private key.
// The private key is used as an API key for authenticating with the Turnkey API.
// turnkeyOptions can be provided to customize the underlying Turnkey SDK client.
// These options enable observability features such as custom HTTP clients with tracing,
// request/response logging, and metrics collection. For example:
//   - sdk.WithHTTPClient(client) - use a custom http.Client with OpenTelemetry instrumentation
//   - sdk.WithBaseURL(url) - override the default Turnkey API endpoint
func NewEthSigner(privateKeyHex string, turnkeyOptions ...sdk.OptionFunc) (*EthSigner, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	tkApiKey, err := apikey.FromECDSAPrivateKey(privateKey, apikey.SchemeSECP256K1)
	if err != nil {
		return nil, err
	}

	turnkeyOptions = append(turnkeyOptions, sdk.WithAPIKey(tkApiKey))

	tkClient, err := sdk.New(turnkeyOptions...)

	return &EthSigner{
		privateKey: privateKey,
		tkClient:   tkClient,
	}, nil
}

func (s *EthSigner) pkToHex() string {
	return fmt.Sprintf("%x", crypto.FromECDSA(s.privateKey))
}

func (s *EthSigner) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(data, s.privateKey)
}

func (s *EthSigner) TKSign(ctx context.Context, data []byte, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error) {
	payload := string(data)
	walletAccountAddressString := walletAccountAddress.String()
	params := signing.NewSignRawPayloadParams().WithBody(&models.SignRawPayloadRequest{
		OrganizationID: &subOrganizationId,
		TimestampMs:    util.RequestTimestamp(),
		Parameters: &models.SignRawPayloadIntentV2{
			Encoding:     models.PayloadEncodingHexadecimal.Pointer(),
			HashFunction: models.HashFunctionKeccak256.Pointer(),
			Payload:      &payload,
			SignWith:     &walletAccountAddressString,
		},
		Type: (*string)(models.ActivityTypeSignRawPayloadV2.Pointer()),
	})

	res, err := s.tkClient.V0().Signing.SignRawPayload(params, s.tkClient.Authenticator, withContext(ctx))
	if err != nil {
		return nil, err
	}

	return SigFromTurnkeyResult(res.Payload.Activity.Result.SignRawPayloadResult)
}

func (s *EthSigner) TKSignEIP7702(
	ctx context.Context,
	authorization types.SetCodeAuthorization,
	subOrganizationId string,
	walletAccountAddress common.Address,
) (*Signature, error) {
	// Format the EIP-7702 authorization as a JSON object
	// Turnkey expects the authorization tuple: (chainId, address, nonce)
	authPayload := EIP7702AuthorizationPayload{
		ChainID: hexutil.EncodeBig(authorization.ChainID.ToBig()),
		Address: authorization.Address.Hex(),
		Nonce:   fmt.Sprintf("0x%x", authorization.Nonce),
	}

	// Serialize to JSON
	authJSON, err := json.Marshal(authPayload)
	if err != nil {
		return nil, fmt.Errorf("marshal authorization: %w", err)
	}

	payload := string(authJSON)
	walletAccountAddressString := walletAccountAddress.String()

	params := signing.NewSignRawPayloadParams().WithBody(&models.SignRawPayloadRequest{
		OrganizationID: &subOrganizationId,
		TimestampMs:    util.RequestTimestamp(),
		Parameters: &models.SignRawPayloadIntentV2{
			Encoding:     models.PayloadEncodingEip7702Authorization.Pointer(),
			HashFunction: models.HashFunctionNoOp.Pointer(),
			Payload:      &payload,
			SignWith:     &walletAccountAddressString,
		},
		Type: (*string)(models.ActivityTypeSignRawPayloadV2.Pointer()),
	})

	res, err := s.tkClient.V0().Signing.SignRawPayload(params, s.tkClient.Authenticator, withContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("turnkey sign EIP-7702: %w", err)
	}

	return SigFromTurnkeyResult(res.Payload.Activity.Result.SignRawPayloadResult)
}

func (s *EthSigner) TKSignEIP712(
	ctx context.Context,
	typedData map[string]interface{},
	subOrganizationId string,
	walletAccountAddress common.Address,
) (*Signature, error) {
	// Serialize the EIP-712 typed data to JSON
	// Turnkey will handle the hashing internally when PAYLOAD_ENCODING_EIP712 is set
	typedDataJSON, err := json.Marshal(typedData)
	if err != nil {
		return nil, fmt.Errorf("marshal typed data: %w", err)
	}

	payload := string(typedDataJSON)
	walletAccountAddressString := walletAccountAddress.String()

	params := signing.NewSignRawPayloadParams().WithBody(&models.SignRawPayloadRequest{
		OrganizationID: &subOrganizationId,
		TimestampMs:    util.RequestTimestamp(),
		Parameters: &models.SignRawPayloadIntentV2{
			Encoding:     models.PayloadEncodingEip712.Pointer(),
			HashFunction: models.HashFunctionNoOp.Pointer(),
			Payload:      &payload,
			SignWith:     &walletAccountAddressString,
		},
		Type: (*string)(models.ActivityTypeSignRawPayloadV2.Pointer()),
	})

	res, err := s.tkClient.V0().Signing.SignRawPayload(params, s.tkClient.Authenticator, withContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("turnkey sign EIP-712: %w", err)
	}

	return SigFromTurnkeyResult(res.Payload.Activity.Result.SignRawPayloadResult)
}

func (s *EthSigner) TKSignEIP712Batch(
	ctx context.Context,
	typedDataList []map[string]interface{},
	subOrganizationId string,
	walletAccountAddress common.Address,
) ([]*Signature, error) {
	// Build array of JSON payloads to sign
	payloads := make([]string, len(typedDataList))
	walletAccountAddressString := walletAccountAddress.String()

	for i, typedData := range typedDataList {
		// Serialize each EIP-712 typed data to JSON
		typedDataJSON, err := json.Marshal(typedData)
		if err != nil {
			return nil, fmt.Errorf("marshal typed data %d: %w", i, err)
		}

		payloads[i] = string(typedDataJSON)
	}

	params := signing.NewSignRawPayloadsParams().WithBody(&models.SignRawPayloadsRequest{
		OrganizationID: &subOrganizationId,
		TimestampMs:    util.RequestTimestamp(),
		Parameters: &models.SignRawPayloadsIntent{
			SignWith:     &walletAccountAddressString,
			Encoding:     models.PayloadEncodingEip712.Pointer(),
			HashFunction: models.HashFunctionNoOp.Pointer(),
			Payloads:     payloads,
		},
		Type: (*string)(models.ActivityTypeSignRawPayloads.Pointer()),
	})

	res, err := s.tkClient.V0().Signing.SignRawPayloads(params, s.tkClient.Authenticator, withContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("turnkey batch sign EIP-712: %w", err)
	}

	// Extract signatures from batch result
	if res.Payload == nil || res.Payload.Activity == nil || res.Payload.Activity.Result == nil || res.Payload.Activity.Result.SignRawPayloadsResult == nil {
		return nil, fmt.Errorf("nil turnkey batch result")
	}

	return SigsFromTurnkeyBatchResult(res.Payload.Activity.Result.SignRawPayloadsResult.Signatures)
}
