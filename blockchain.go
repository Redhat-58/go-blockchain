package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{createGenesisBlock()},
	}
}

func createGenesisBlock() *Block {
	return NewBlock(0, "Initial Block", []byte{})
}

func (bc *BlockChain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.GetLastBlock()
	newBlock := NewBlock(lastBlock.Index+1, data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)

}
