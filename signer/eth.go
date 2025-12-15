package signer

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type EthSigner struct {
	privateKey *ecdsa.PrivateKey
}

func NewEthSigner(privateKey *ecdsa.PrivateKey) *EthSigner {
	return &EthSigner{
		privateKey: privateKey,
	}
}

func (s *EthSigner) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(data, s.privateKey)
}
