package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {

	bc := blockchain.NewBlockChain()

	bc.AddBlock("Block 1: Jalal just enterd a tansaction")

	for _, block := range bc.Blocks {

		fmt.Println("Index: ", block.Index)
		fmt.Println("Timestamp: ", block.Timestamp)
		fmt.Println("Data: ", block.Data)
		fmt.Println("PrvHash: ", block.PreviousHash)
		fmt.Println("Hash: ", block.Hash)
		fmt.Println("Token: ", block.Token)
		fmt.Println()
	}

}
