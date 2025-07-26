package block_chain

import "web3/transaction"

type Block struct {
	Index     int   `json:"index"`
	TimeStamp int64 `json:"timeStamp"`
	// Transactions 是一个包含交易对象的切片。每个交易（transaction.Transaction）通常包含交易的详细信息，如发送方、接收方、金额等。
	// 这个字段表示区块中包含的所有交易。区块的核心任务之一就是记录这些交易。
	Transactions []transaction.Transaction `json:"transactions"`
	// Confirm 表示区块被网络确认的次数。这个字段通常用于记录区块在区块链中的“确认状态”。例如，如果一个区块被矿工确认并添加到区块链，它的 Confirm 值可能会被设置为1。
	// 随着更多区块被添加到链上，这个区块的 Confirm 数量会增加，表示它在区块链中被确认的次数更多，安全性也更高。
	Confirm int `json:"confirm"`
	// PreHash 是前一个区块的哈希值。每个区块都包含前一个区块的哈希值，形成了区块链。这个字段确保了区块链的完整性和不可篡改性。
	// 如果某个区块被修改，它的哈希值会发生变化，导致后续所有区块的哈希值发生变化，从而保证区块链的安全性。
	PreHash string `json:"preHash"`
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
