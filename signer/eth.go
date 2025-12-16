package signer

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tkhq/go-sdk"
	"github.com/tkhq/go-sdk/pkg/api/client/signing"
	"github.com/tkhq/go-sdk/pkg/api/models"
	"github.com/tkhq/go-sdk/pkg/apikey"
	"github.com/tkhq/go-sdk/pkg/util"
)

type EthSigner struct {
	privateKey *ecdsa.PrivateKey
	tkClient   *sdk.Client
}

func NewEthSigner(privateKeyHex string) (*EthSigner, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	tkApiKey, err := apikey.FromECDSAPrivateKey(privateKey, apikey.SchemeSECP256K1)
	if err != nil {
		return nil, err
	}

	tkClient, err := sdk.New(sdk.WithAPIKey(tkApiKey))

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

func (s *EthSigner) TKSign(data []byte, subOrganizationId string, walletAccountAddress common.Address) (*Signature, error) {
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

	res, err := s.tkClient.V0().Signing.SignRawPayload(params, s.tkClient.Authenticator)
	if err != nil {
		return nil, err
	}

	return SigFromTurnkeyResult(res.Payload.Activity.Result.SignRawPayloadResult)
}
