package blockchain

import (
	"crypto/ecdsa"
	"math/big"
	"time"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Value      float64
}

type Transaction struct {
	ID    int
	To    string
	From  string
	Value float64
}

type Token struct {
	Name        string
	Symbol      string
	TotalSupply float64
	Balance     map[string]float64
}

type BlockChain struct {
	Blocks       []*Block
	Token        Token
	Transactions []*Transaction
	Wallets      map[string]*Wallet
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
	Transaction  []*Transaction
}
