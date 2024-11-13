package blockchain

import (
	"math/big"
	"time"
)

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
	Token        float64
	Transaction  []*Transaction
}
