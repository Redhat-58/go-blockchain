package blockchain

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{createGenesisBlock()},
	}
}

func createGenesisBlock() *Block {
	return NewBlock(0, "GenesisBlock", "")
}

func (bc *BlockChain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.GetLastBlock()
	newBlock := NewBlock(lastBlock.Index+1, data, lastBlock.Hash)
	newBlock.ProofOfWork()
	bc.Blocks = append(bc.Blocks, newBlock)

}
