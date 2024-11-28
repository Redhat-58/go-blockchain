package blockchain

import (
	"errors"
	"fmt"
	"strings"
)

const mintTargetBits = 6

func NewToken(name string, symbol string, totalSupply float64) *Token {
	balances := make(map[string]float64)
	balances["creator"] = totalSupply
	newToken := &Token{
		Name:        name,
		Symbol:      symbol,
		TotalSupply: totalSupply,
		Balance:     balances,
	}
	SaveToken(newToken)
	return newToken
}

func (t *Token) GetBalance(account string) float64 {
	if balance, exists := t.Balance[account]; exists {
		return balance
	}
	return 0
}

func (t *Token) Transfer(bc *BlockChain, from string, to string, amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New("Amount must be greater than zero: 0")
	}
	if t.Balance[from] < amount {
		return "", errors.New("Insufficient balance")
	}
	t.Balance[from] -= amount
	t.Balance[to] += amount

	fromBlock, err := bc.GetBlockByHash(from)
	if err != nil {
		return "", errors.New("Error retreving block")
	}
	toBlock, err := bc.GetBlockByHash(from)
	if err != nil {
		return "", errors.New("Error retreving block")
	}

	UpdateBalance(fromBlock, t)
	UpdateBalance(toBlock, t)

	return ("Transfer successfull"), nil
}

func (b *Block) MintToken(t *Token, bc *BlockChain) {
	fmt.Println("Mining Token....")
	targetBits := strings.Repeat("0", mintTargetBits)
	for {
		hash := b.calculateHash()
		if strings.HasPrefix(hash, targetBits) {
			t.Balance[b.Hash]++
			t.TotalSupply++
			bc.FundWallet(b, 1)
			bc.NewTransaction(b.Hash, "Mint", 1)
			break
		}
		b.Nonce++
	}
	targetBits = ""
	return
}
