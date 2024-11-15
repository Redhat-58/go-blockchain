package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func (b *Block) NewWallet(bc *BlockChain) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error creating private key", err)
	}
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	if bc.Wallets == nil {
		bc.Wallets = make(map[string]*Wallet)
	}
	bc.Wallets[b.Hash] = &Wallet{privateKey, publicKey, 0}
	fmt.Println("created user wallet")

}

func (bc BlockChain) FundWallet(b *Block, value float64) {
	bc.Wallets[b.Hash].Value += value
}
