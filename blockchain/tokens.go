package blockchain

import (
	"errors"
	"fmt"
	"strings"
)

const mintTargetBits = 3

func NewToken(name string, symbol string, totalSupply float64) *Token {
	balances := make(map[string]float64)
	balances["creator"] = totalSupply
	return &Token{
		Name:        name,
		Symbol:      "",
		TotalSupply: totalSupply,
		Balance:     balances,
	}
}

func (t *Token) GetBalance(account string) float64 {
	if balance, exists := t.Balance[account]; exists {
		return balance
	}
	return 0
}

func (t *Token) Transfer(from string, to string, amount float64) error {
	if amount <= 0 {
		return errors.New("Amount must be greater than zero: 0")
	}
	if t.Balance[from] < amount {
		return errors.New("Insufficient balance")
	}
	t.Balance[from] -= amount
	t.Balance[to] += amount
	fmt.Println("Transferred %.2f %s from %s to %s\n", amount, t.Symbol, from, to)

	return nil
}

func (b *Block) MintToken(t *Token) {
	targetBits := strings.Repeat("0", mintTargetBits)
	for {
		hash := b.calculateHash()
		if strings.HasPrefix(hash, targetBits) {
			t.Balance[b.Hash]++
			fmt.Printf("Token mined! Nonce: %d, Hash: %s\n", b.Nonce, b.Hash)
			break
		}
		b.Nonce++
	}
}
