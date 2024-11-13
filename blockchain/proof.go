package blockchain

import (
	"fmt"
	"strings"
)

const targetBits = 6

func (b *Block) ProofOfWork() {
	targetBits := strings.Repeat("0", targetBits)
	for {
		b.Hash = b.calculateHash()
		if strings.HasPrefix(b.Hash, targetBits) {
			fmt.Printf("Block mined! Nonce: %d, Hash: %s\n", b.Nonce, b.Hash)
			break
		}
		b.Nonce++
	}
}
