package block_chain

import (
	"fmt"
	"testing"
	"time"
	"web3/transaction"
)

var chain BlockChain

func init() {
	transactions := make([]transaction.Transaction, 0)
	blocks := make([]Block, 0)
	m := make(map[string]struct{})
	chain := NewBlockChain(transactions, blocks, m)
	block := chain.NewBlock(time.Now().UnixMilli(), 1, "")
	fmt.Println(block)

}

func TestBlockChain_ConfirmOfWork(t *testing.T) {
	confirmed := chain.ConfirmOfWork()
	fmt.Println("confirmed:", confirmed)
}
