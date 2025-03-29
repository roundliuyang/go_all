package main

import (
	"github.com/gin-gonic/gin"
	"redis-parctice-lesson/01.auth_github/handler"
	"redis-parctice-lesson/global"
)

func init() {
	global.InitGoRedisClient()
}
func main() {
	// https://github.com/login/oauth/authorize?client_id=f6e64e86547e9ff5e98c
	r := gin.Default()
	r.GET("/oauth2/redirect", handler.CodeHandler)
	r.GET("/user", handler.GetGitHbuUser)
	r.Run(":9099")
}
