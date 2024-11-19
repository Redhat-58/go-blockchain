package main

import (
	"fmt"
	"go-blockchain/blockchain"
	"go-blockchain/controllers"
	"go-blockchain/routes"
	"log"
	"net/http"
)

func main() {

	blockchain.InitDB()
	fmt.Println("....Database initialized....")

	controllers.CreateBlockchain()
	fmt.Println("....Blockchain created....")

	router := routes.InitRoutes()
	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// myBlock :=

	// johnBlock := bc.AddBlock("Block 2: Johns Block")
	// myBlock.MintToken(&bc.Token, bc)
	// myBlock.MintToken(&bc.Token, bc)

	// bc.Token.Transfer(bc, myBlock.Hash, johnBlock.Hash, 0.5)

	// fmt.Println(bc.Token)
	// fmt.Println(bc.Transactions)
	// fmt.Println(bc.Wallets)

}
