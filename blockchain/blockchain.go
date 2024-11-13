package blockchain

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{createGenesisBlock()},
	}
}

func createGenesisBlock() *Block {
	return NewBlock(0, "Initial Block", []byte{})
}

func (bc *BlockChain) GetLastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.GetLastBlock()
	newBlock := NewBlock(lastBlock.Index+1, data, lastBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

}
