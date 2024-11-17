package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func CreateGenesisBlock() *Block {
	block := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         "Genesis Block",
		PreviousHash: "000000",
		Hash:         "",
		Nonce:        0,
		Transaction:  []*Transaction{},
	}
	block.Hash = block.calculateHash()
	return block
}

func (bc *BlockChain) NewBlock(index int, data string, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Data:         data,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
		Nonce:        0,
		Transaction:  []*Transaction{},
	}

	fmt.Println("Mining new block...")
	block.ProofOfWork()
	bc.Blocks = append(bc.Blocks, block)
	block.NewWallet(bc)
	return block
}

func (b *Block) calculateHash() string {
	data := fmt.Sprintf("%d%s%d%s%d", b.Index, b.PreviousHash, b.Nonce, b.Timestamp)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (bc *BlockChain) GetBlockByHash(hash string) (*Block, error) {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return block, nil
		}
	}
	return nil, fmt.Errorf("block not found for hash: %s", hash)
}
