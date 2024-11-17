package blockchain

func (bc *BlockChain) NewTransaction(to, from string, value float64) {
	tx := &Transaction{
		To:    to,
		From:  from,
		Value: value,
	}
	tx.ID = len(bc.Transactions) + 1
	bc.Transactions = append(bc.Transactions, tx)

}
