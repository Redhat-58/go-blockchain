package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {

	blockchain.InitDB()

	bc := blockchain.NewBlockChain()

	myBlock := bc.AddBlock("Block 1: Jalals Block")

	johnBlock := bc.AddBlock("Block 2: Johns Block")
	myBlock.MintToken(&bc.Token, bc)
	myBlock.MintToken(&bc.Token, bc)

	bc.Token.Transfer(myBlock.Hash, johnBlock.Hash, 0.5)

	fmt.Println(bc.Token)
	fmt.Println(bc.Transactions)
	fmt.Println(bc.Wallets)

}
