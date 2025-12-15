package signer

type Signer interface {
	Sign(data []byte) ([]byte, error)
}
