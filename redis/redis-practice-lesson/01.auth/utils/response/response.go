package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redis-parctice-lesson/01.auth/conf"
)

func ShowError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
	})
}
func ShowValidatorError(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

func ShowSuccess(c *gin.Context, msg *conf.MyJWT) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}
func ShowData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
