package router

import (
	"github.com/gin-gonic/gin"
	"gorm/gormtest/api"
)

func InitRouter(r *gin.Engine) {
	api.RegisterRouter(r)
}
