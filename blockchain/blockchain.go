package blockchain

import (
	"fmt"
)

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks:     []*Block{CreateGenesisBlock()},
		Token:      *createGenesisToken(),
		Wallets:    createWallets(),
		BlockIndex: initializeBlockIndex(),
	}
}

func createGenesisToken() *Token {
	return NewToken("Token", "GTK", 0)
}

func createWallets() map[string]*Wallet {
	fmt.Println("Creating wallets")
	return make(map[string]*Wallet)
}

func initializeBlockIndex() map[string]*Block {
	fmt.Println("Creating Block index")
	return make(map[string]*Block)
}

func (bc *BlockChain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *BlockChain) AddBlock(data string) *Block {
	lastBlock := bc.GetLastBlock()
	newBlock := bc.NewBlock(lastBlock.Index+1, data, lastBlock.Hash)
	bc.BlockIndex[newBlock.Hash] = newBlock
	SaveBlock(newBlock)
	return newBlock

}
