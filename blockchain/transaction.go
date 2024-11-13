package blockchain

import (
	"bytes"
	"crypto/sha256"
)

func (bc *Block) NewTransaction(input, output string) *Transaction {
	tx := &Transaction{
		ID:     []byte{},
		Input:  []byte(input),
		Output: []byte(output),
	}
	tx.ID = tx.calculateHash()
	bc.Transaction = append(bc.Transaction, tx)
	return tx

}

func (tx *Transaction) calculateHash() []byte {
	data := bytes.Join([][]byte{
		tx.Input,
		tx.Output,
	}, []byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}
