package blockchain

import (
	"math/big"
	"time"
)

type BlockChain struct {
	Blocks []*Block
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PreviousHash []byte
	Hash         []byte
}
