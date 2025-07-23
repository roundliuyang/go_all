package block_chain

import "web3/transaction"

type Block struct {
	Index        int                       `json:"index"`
	TimeStamp    int64                     `json:"timeStamp"`
	Transactions []transaction.Transaction `json:"transactions"`
	Confirm      int                       `json:"confirm"`
	PreHash      string                    `json:"preHash"`
}

func NewBlock(index int, timeStamp int64, transactions []transaction.Transaction, confirm int, preHash string) Block {
	return Block{
		Index:        index,
		TimeStamp:    timeStamp,
		Transactions: transactions,
		Confirm:      confirm,
		PreHash:      preHash,
	}
}
