package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func (b *Block) NewWallet(bc *BlockChain) *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error creating private key", err)
	}
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	if bc.Wallets == nil {
		bc.Wallets = make(map[string]*Wallet)
	}
	newWallet := &Wallet{privateKey, publicKey, 0}
	bc.Wallets[b.Hash] = newWallet
	fmt.Println("created user wallet")
	return newWallet

}

func (bc BlockChain) FundWallet(b *Block, value float64) {
	bc.Wallets[b.Hash].Value += value
}
