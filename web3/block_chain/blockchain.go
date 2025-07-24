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

/*
实现区块链中的哈希计算和签名过程。具体来说，它使用了 HMAC（Hash-based Message Authentication Code）和 SHA-256 加密算法来生成一个签名。
它的主要目的是确保区块内容的完整性和防止篡改。
在 HMAC 中，SecretKey 是一个私密的密钥，应该只有发送方和接收方知道。它用于生成消息的签名（signature），以确保消息的完整性和认证。因为 HMAC 是对称加密算法，签名的生成和验证都使用相同的密钥。
公钥 和 私钥 是非对称加密中的概念，通常用于加密和解密过程，而 HMAC 不涉及加密和解密，只是用于消息验证，所以它需要一个私密的密钥（SecretKey）。
如果你想在区块链中使用公钥和私钥进行数字签名，一般会用到类似 RSA 或 ECDSA 的非对称加密算法。在这种情况下，私钥 用于签名，而 公钥 用于验证签名。
*/
func (bc *BlockChain) Hash(block Block) (string, error) {
	// 将 SecretKey 转换成 JSON 格式的字节切片，并作为 HMAC 算法的密钥。SecretKey 通常是预先定义的一个密钥，用于确保生成的签名是唯一的。
	key, err := json.Marshal(config.SecretKey)
	if err != nil {
		return "", err
	}
	// 创建一个 HMAC 实例，使用 SHA-256 算法和上面生成的密钥
	m := hmac.New(sha256.New, key)
	// 将传入的区块（block）对象转换为 JSON 格式的字节切片。
	b, err := json.Marshal(block)
	if err != nil {
		return "", err
	}
	// m.Write(b)：将区块的字节数据写入 HMAC 对象，开始计算哈希值。
	m.Write(b)
	// 最终生成的哈希值（签名）通过 hex.EncodeToString 转换为十六进制字符串
	signature := hex.EncodeToString(m.Sum(nil))
	return signature, nil
}

func (bc *BlockChain) GetLastBlock() Block {
	return bc.BlockList[len(bc.BlockList)-1]
}

// 这段代码实现了一个 工作量证明（Proof of Work, PoW） 的过程，通常用于区块链中的矿工验证工作，确保新区块的产生需要一定的计算量。
func (bc *BlockChain) ConfirmOfWork() int {
	lastBlock := bc.LastBlock
	confirm := 0
	// lastBlock.PreHash 是前一个区块的哈希值，confirm 是一个递增的整数，用于尝试找到一个符合要求的哈希签名，这个函数每次调用 ValidateConfirm 来验证当前的工作量是否符合要求。
	for !bc.ValidateConfirm(lastBlock.PreHash, confirm) {
		confirm++
	}
	// 一旦找到符合条件的 confirm 值，它就返回这个 confirm 值。
	return confirm
}

func (bc *BlockChain) ValidateConfirm(preHash string, confirm int) bool {
	// 首先，它将 preHash 和 confirm 拼接成一个字符串 s
	s := fmt.Sprintf("%s%d", preHash, confirm)
	key, err := json.Marshal(config.SecretKey)
	if err != nil {
		return false
	}
	// 使用 hmac.New(sha256.New, key) 创建一个 HMAC 对象，使用 SecretKey（即私钥）对字符串进行加密
	m := hmac.New(sha256.New, key)
	m.Write([]byte(s))
	// signature 是生成的哈希签名，最终转为十六进制字符串
	signature := hex.EncodeToString(m.Sum(nil))
	fmt.Println(signature)
	// 比如说前面6位、8位、10位是0的一个数字哈希值是要非常多的时间的
	// 这段代码的关键部分是：if signature[:1] == "0"，它表示哈希签名的 前一个字符 必须是 '0'。这就是工作量证明的条件：找到一个哈希值的前缀符合要求（即以一个零开头）
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
