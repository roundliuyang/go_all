package main

import (
	"github.com/gin-gonic/gin"
	"gorm/gormtest/router"
	"log"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
