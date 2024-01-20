package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.GET("/save", SaveUser)
	r.GET("/get", GetUser)
	r.GET("/getAll", GetAll)
}
