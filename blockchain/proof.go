package blockchain

import (
	"bytes"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

const targetBits = 16

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", pow.Block.Index)),
		pow.Block.PreviousHash,
		[]byte(pow.Block.Timestamp.String()),
		[]byte(pow.Block.Data),
		[]byte(fmt.Sprintf("%d", nonce)),
	}, []byte{})
	return data
}
