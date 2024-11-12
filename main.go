package main

import (
	"fmt"
)

func main() {

	bc := NewBlockChain()

	bc.AddBlock("Block 1: Jalal just created the blockchain")

	for _, block := range bc.blocks {

		fmt.Println("Index", block.Index)
		fmt.Println("Timestamp", block.Timestamp)
		fmt.Println("Data", block.Data)
		fmt.Println("PrvHash", block.PreviousHash)
		fmt.Println("Hash", block.Hash)
		fmt.Println()
	}

}
