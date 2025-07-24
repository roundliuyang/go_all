package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"strings"
	"time"
	"web3/block_chain"
	"web3/model"
	"web3/transaction"
)

var BlockChain *block_chain.BlockChain

func init() {
	BlockChain = &block_chain.BlockChain{}
	BlockChain.Nodes = make(map[string]struct{})

	BlockChain.NewBlock(time.Now().UnixMilli(), 666, "")
}

func TransactionNew(c *gin.Context) {
	t := transaction.Transaction{}
	err := c.ShouldBind(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
	}
	if t.From == "" || t.To == "" || t.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	index := BlockChain.NewTransaction(t.From, t.To, t.Amount)
	msg := fmt.Sprintf("即将添加到{%d}的区块链上", index)
	c.JSON(http.StatusCreated, gin.H{
		"msg": msg,
	})
}

func Mint(c *gin.Context) {
	confirm := BlockChain.ConfirmOfWork()
	v4, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成ID失败",
		})
		return
	}
	id := strings.ReplaceAll(v4.String(), "-", "")
	BlockChain.NewTransaction("", id, 12)
	newBlock := BlockChain.NewBlock(time.Now().UnixMilli(), confirm, "")
	c.JSON(http.StatusOK, gin.H{
		"msg":         "成功",
		"index":       newBlock.Index,
		"transaction": newBlock.Transactions,
		"confirm":     newBlock.Confirm,
		"preHash":     newBlock.PreHash,
	})
}

func ShowBlockChain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"blockList": BlockChain.BlockList,
		"length":    len(BlockChain.BlockList),
	})
}

func Add(c *gin.Context) {
	var ask model.Ask
	err := c.ShouldBind(&ask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	if len(ask.Nodes) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	b := true
	s := ""
	for _, item := range ask.Nodes {
		err := BlockChain.AddNode(item)
		if err != nil {
			fmt.Println(err)
			b = false
			s += item + "注册失败"
			/// TODO 自己处理业务逻辑
		}
	}
	if b {
		s = "全部节点注册成功"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   s,
		"items": ask.Nodes,
	})
}

func Consensus(c *gin.Context) {
	b := BlockChain.Consensus()
	if b {
		c.JSON(http.StatusOK, gin.H{
			"msg":       "区块链已更新为最新版本",
			"blockList": BlockChain.BlockList,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":       "区块链未改变，已是最新版本",
			"blockList": BlockChain.BlockList,
		})
	}
}
