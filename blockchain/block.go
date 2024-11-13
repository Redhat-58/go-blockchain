package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func NewBlock(index int, data string, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Data:         data,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Nonce:        0,
		Token:        0,
	}

	fmt.Println("Mining new block...")
	block.ProofOfWork()
	return block
}

func (b *Block) calculateHash() string {
	data := fmt.Sprintf("%d%s%d%s%d", b.Index, b.PreviousHash, b.Nonce, b.Timestamp)
	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}
