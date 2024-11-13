package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {

	// bc := NewBlockChain()
	bc := blockchain.NewBlockChain()

	bc.AddBlock("Block 1: Jalal just created the blockchain")

	for _, block := range bc.Blocks {

		fmt.Println("Index", block.Index)
		fmt.Println("Timestamp", block.Timestamp)
		fmt.Println("Data", block.Data)
		fmt.Println("PrvHash", block.PreviousHash)
		fmt.Println("Hash", block.Hash)
		fmt.Println()
	}

}
