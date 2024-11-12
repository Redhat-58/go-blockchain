package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         string
	PreviousHash []byte
	Hash         []byte
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
