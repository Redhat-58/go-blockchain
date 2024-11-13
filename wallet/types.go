package wallet

import "crypto/ecdsa"

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}
