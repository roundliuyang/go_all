package block_chain

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"web3/config"
	"web3/transaction"
)

type BlockChain struct {
	CurrentTransactions []transaction.Transaction
	BlockList           []Block             `json:"blockList"`
	LastBlock           Block               `json:"lastBlock"`
	Nodes               map[string]struct{} `json:"nodes"`
}

func NewBlockChain(transactions []transaction.Transaction, blocks []Block, nodes map[string]struct{}) *BlockChain {
	return &BlockChain{
		CurrentTransactions: transactions,
		BlockList:           blocks,
		Nodes:               nodes,
	}
}

func (bc *BlockChain) Hash(block Block) (string, error) {
	key, err := json.Marshal(config.SecretKey)
	if err != nil {
		return "", err
	}
	m := hmac.New(sha256.New, key)
	b, err := json.Marshal(block)
	if err != nil {
		return "", err
	}
	m.Write(b)
	signature := hex.EncodeToString(m.Sum(nil))
	return signature, nil
}

func (bc *BlockChain) GetLastBlock() Block {
	return bc.BlockList[len(bc.BlockList)-1]
}

func (bc *BlockChain) ConfirmOfWork() int {
	lastBlock := bc.LastBlock
	confirm := 0
	for !bc.ValidateConfirm(lastBlock.PreHash, confirm) {
		confirm++
	}
	return confirm
}

func (bc *BlockChain) ValidateConfirm(preHash string, confirm int) bool {
	s := fmt.Sprintf("%s%d", preHash, confirm)
	key, err := json.Marshal(config.SecretKey)
	if err != nil {
		return false
	}
	m := hmac.New(sha256.New, key)
	m.Write([]byte(s))
	signature := hex.EncodeToString(m.Sum(nil))
	fmt.Println(signature)
	// 比如说前面6位、8位、10位是0的一个数字哈希值是要非常多的时间的
	if signature[:1] == "0" {
		return true
	}
	return false
}

func (bc *BlockChain) NewBlock(timeStamp int64, confirm int, preHash string) Block {
	if preHash == "" && len(bc.BlockList) >= 1 {
		preHash, _ = bc.Hash(bc.LastBlock)
	}
	b := NewBlock(len(bc.BlockList)+1, timeStamp, bc.CurrentTransactions, confirm, preHash)
	bc.BlockList = append(bc.BlockList, b)
	bc.CurrentTransactions = make([]transaction.Transaction, 0)
	return b
}

func (bc *BlockChain) NewTransaction(from, to string, amount float64) int {
	t := transaction.NewTransaction(from, to, amount)
	bc.CurrentTransactions = append(bc.CurrentTransactions, t)
	lastBlock := bc.BlockList[len(bc.BlockList)-1]
	return lastBlock.Index + 1
}

func (bc *BlockChain) AddNode(urlAddr string) error {
	l, err := url.Parse(urlAddr)
	if err != nil {
		return err
	}

	// http://192.168.1.1:8080
	s := fmt.Sprintf("%s://%s:%s", l.Scheme, l.Hostname(), l.Port())
	bc.Nodes[s] = struct{}{}
	return nil
}
