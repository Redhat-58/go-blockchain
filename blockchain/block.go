package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

const targetBits = 16

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PreviousHash []byte
	Hash         []byte
}

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewBlock(index int, data string, previousHash []byte) *Block {
	block := &Block{
		Index:        index,
		Data:         data,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Hash:         []byte{},
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() []byte {
	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", b.Index)),
		b.PreviousHash,
		[]byte(b.Timestamp.String()),
		[]byte(b.Data),
	}, []byte{})

	hash := sha256.Sum256(data)
	return hash[:]
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		[]byte(fmt.Sprintf("%d", pow.Block.Index)),
		pow.Block.PreviousHash,
		[]byte(pow.Block.Timestamp.String()),
		[]byte(pow.Block.Data),
		[]byte(fmt.Sprintf("%d", nonce)),
	}, []byte{})
	return data
}
