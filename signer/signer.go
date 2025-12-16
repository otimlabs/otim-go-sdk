package signer

import "github.com/ethereum/go-ethereum/common"

type Signer interface {
	Sign(data []byte) ([]byte, error)
	TKSign(data []byte, subOrganizationId string, walletAccountAddress common.Address) ([]byte, error)
}
