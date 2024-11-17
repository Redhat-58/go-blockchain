package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {

	bc := blockchain.NewBlockChain()

	myBlock := bc.AddBlock("Block 1: Jalals Block")

	johnBlock := bc.AddBlock("Block 2: Johns Block")
	myBlock.MintToken(&bc.Token, bc)
	myBlock.MintToken(&bc.Token, bc)

	bc.Token.Transfer(myBlock.Hash, johnBlock.Hash, 0.5)

	// for _, block := range bc.Blocks {

	// 	fmt.Println("Index: ", block.Index)
	// 	fmt.Println("Timestamp: ", block.Timestamp)
	// 	fmt.Println("Data: ", block.Data)
	// 	fmt.Println("PrvHash: ", block.PreviousHash)
	// 	fmt.Println("Hash: ", block.Hash)
	// 	fmt.Println()
	// }
	// fmt.Println()

	fmt.Println(bc.Token)
	fmt.Println(bc.Transactions)
	fmt.Println(bc.Wallets)

}
