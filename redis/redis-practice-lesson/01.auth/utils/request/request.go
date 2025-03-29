package request

import (
	"github.com/gin-gonic/gin"
)

func GetParam(c *gin.Context, key string) (string, bool) {
	val := c.GetHeader(key)
	if val != "" {
		return val, true
	}
	val, err := c.Cookie(key)
	if err != nil {
		return "", false
	}
	return val, true
}
