package controllers

import "go-blockchain/blockchain"

var BC *blockchain.BlockChain

func CreateBlockchain() {
	bc := blockchain.NewBlockChain()
	BC = bc
}
