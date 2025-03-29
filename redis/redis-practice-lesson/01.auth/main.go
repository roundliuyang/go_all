package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"redis-parctice-lesson/01.auth/api"
	"redis-parctice-lesson/01.auth/api/user"
	"redis-parctice-lesson/01.auth/conf"
	"redis-parctice-lesson/01.auth/model"
	"redis-parctice-lesson/01.auth/utils/common"
	"redis-parctice-lesson/01.auth/utils/handle"
	"redis-parctice-lesson/01.auth/utils/request"
	"redis-parctice-lesson/01.auth/utils/response"
	"redis-parctice-lesson/global"
	"time"
)

func init() {
	model.InitDB()
	global.InitGoRedisClient()
}

func Load() {
	c := conf.Config{}
	c.Routes = []string{"/renewal", "/my/info"}
	c.OpenJwt = true
	conf.Set(c)
	handle.InitValidate()
}

func Auth(c *gin.Context) {
	u, err := url.Parse(c.Request.RequestURI)
	if err != nil {
		//TODO 记录错误日志,生产环境，不要轻易写Panic
		panic(err)
	}
	if !common.InProtectedRoutes(u.Path, &conf.Cfg.Routes) {
		c.Next()
		return
	}

	if conf.Cfg.OpenJwt {
		accessToken, has := request.GetParam(c, conf.AccessToken)
		if !has {
			c.Abort()
			response.ShowError(c, "未登录")
			return
		}
		ret, err := conf.ParseToken(accessToken)
		if err != nil {
			c.Abort()
			response.ShowValidatorError(c, err.Error())
			return
		}
		now := time.Now().Unix()
		if ret.ExpiresAt-now < 0 {
			c.Set("uid", ret.UserId)
			c.JSON(999, gin.H{"msg": "accessToken已经过期"})
		} else {
			c.Set("uid", ret.UserId)
			c.Next()
		}
		return
	}
	c.Next()
	return

}

func main() {
	Load()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(Auth)
	r.GET("/my/info", user.Info)
	r.POST("/renewal", user.Renewal)

	r.POST("/login", user.Login)
	r.POST("/login/mobile", user.LoginByMobileCode)
	r.POST("/sendsms", user.SendSms)
	r.POST("/signup/mobile", user.SignupByMobile)
	r.POST("/signup/mobile/exist", user.MobileIsExists)

	r.GET("/", api.Index)

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9100")
}
