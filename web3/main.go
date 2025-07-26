package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"web3/handlers"
)

var InPutPort = flag.Int("port", 9999, "input port")

func main() {
	flag.Parse()
	engine := gin.Default()
	engine.POST("/transaction/new", handlers.TransactionNew)
	engine.GET("/mint", handlers.Mint)
	engine.GET("/show/chain", handlers.ShowBlockChain)
	engine.POST("/nodes/add", handlers.Add)
	engine.GET("/consensus", handlers.Consensus)
	port := fmt.Sprintf(":%d", *InPutPort)
	engine.Run(port)
}
