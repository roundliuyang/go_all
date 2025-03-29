package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"redis-parctice-lesson/01.auth/conf"
	"redis-parctice-lesson/01.auth/middleware"
	"redis-parctice-lesson/01.auth/model"
	"redis-parctice-lesson/01.auth/utils/common"
	"redis-parctice-lesson/01.auth/utils/handle"
	"redis-parctice-lesson/01.auth/utils/request"
	"redis-parctice-lesson/01.auth/utils/response"
	"redis-parctice-lesson/01.auth/utils/sms"
	"strconv"
	"time"
)

type MobilePasswordCode struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Passwd string `json:"passwd" form:"passwd" bindging:"required,max=20,min=6"`
	Code   string `json:"code" form:"code" binding:"required,len=6"`
}

type MobileCode struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Code   string `json:"code" form:"code" binding:"required,len=6"`
}

type MobilePassword struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
	Passwd string `json:"passwd" form:"passwd" bindging:"required,max=20,min=6"`
}

type Mobile struct {
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}

var MobileTrans = map[string]string{"mobile": "手机号"}
var UserMobileTrans = map[string]string{"Mobile": "手机号", "Passwd": "密码", "Code": "验证码"}

func Login(c *gin.Context) {
	var mobilePasswd MobilePassword
	if err := c.BindJSON(&mobilePasswd); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}
	b := model.IsExistsMobile(mobilePasswd.Mobile)
	if b.ID < 1 {
		response.ShowError(c, "手机号不存在")
		return
	}

	passwd := common.HashPassword(mobilePasswd.Passwd, []byte(b.Salt))
	if common.DoPasswordsMatch(passwd, b.Passwd, []byte(b.Salt)) {
		response.ShowError(c, "登录错误")
		return
	}
	token, err := middleware.DoLogin(b)
	if err != nil {
		response.ShowError(c, "失败")
		return
	}
	response.ShowSuccess(c, token)
	return
}

func LoginByMobileCode(c *gin.Context) {
	var userMobile MobileCode
	if err := c.BindJSON(&userMobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}

	//TODO 验证code

	b := model.IsExistsMobile(userMobile.Mobile)
	if b.ID < 1 {
		response.ShowError(c, "手机号不存在")
		return
	}

	token, err := middleware.DoLogin(b)
	if err != nil {
		response.ShowError(c, "失败")
		return
	}
	response.ShowSuccess(c, token)
	return
}

func MobileIsExists(c *gin.Context) {
	var mobile Mobile
	if err := c.BindJSON(&mobile); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}
	//TODO 验证手机号合法性

	var data = map[string]bool{"is_exist": true}
	u := model.IsExistsMobile(mobile.Mobile)
	if u.ID < 1 {
		data["is_exist"] = false
	}
	response.ShowData(c, data)
	return
}

func SendSms(c *gin.Context) {
	var p Mobile
	err := c.BindJSON(&p)
	if err != nil {
		response.ShowError(c, "失败")
		return
	}
	code := common.GetRandomNum(6)
	err = sms.SendSms(p.Mobile, code)
	if err != nil {
		response.ShowError(c, "失败")
		return
	}
	response.ShowError(c, "成功")
	return
}

func SignupByMobile(c *gin.Context) {
	var mobilePasswordCode MobilePasswordCode
	if err := c.BindJSON(&mobilePasswordCode); err != nil {
		msg := handle.TransTagName(&UserMobileTrans, err)
		response.ShowValidatorError(c, msg)
		return
	}
	u := model.IsExistsMobile(mobilePasswordCode.Mobile)
	if u.ID > 1 {
		response.ShowError(c, "手机号码已存在")
		return
	}
	num := rand.Intn(10000)
	salt := strconv.Itoa(num)
	pwd := common.HashPassword(mobilePasswordCode.Passwd, []byte(salt))
	u.Add("", mobilePasswordCode.Mobile, mobilePasswordCode.Mobile, pwd, salt)
	c.JSON(http.StatusOK, "添加成功")
	return
}

func Renewal(c *gin.Context) {
	accessToken, has := request.GetParam(c, conf.AccessToken)
	if !has {
		response.ShowValidatorError(c, "access token不存在")
		return
	}
	refreshToken, has := request.GetParam(c, conf.RefreshToken)
	if !has {
		response.ShowValidatorError(c, "refresh token不存在")
		return
	}

	ret, err := conf.ParseToken(refreshToken)
	if err != nil {
		response.ShowError(c, "refresh_token")
		return
	}
	now := time.Now().Unix()
	if ret.ExpiresAt-now < 0 {
		//证明我们的refresh_token过期了，需要去登录页面。
		return
	}

	maxAge := time.Duration(conf.MaxAge) * time.Second
	nowTime := time.Now()
	customClaims := &conf.CustomClaims{
		UserId: ret.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowTime.Add(maxAge).Unix(),
		},
	}

	accessToken, err = customClaims.MakeToken()
	if err != nil {
		response.ShowError(c, "失败")
		return
	}

	refreshToken, err = customClaims.MakeToken()
	if err != nil {
		response.ShowError(c, "失败")
		return
	}
	token := &conf.MyJWT{
		Uid:          ret.UserId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	response.ShowSuccess(c, token)
	return
}

func Info(c *gin.Context) {
	uid := c.MustGet("uid").(int64)
	user := model.User{}
	row := user.GetRowById(uid)
	if row.ID < 1 {
		err := errors.New("验证失败")
		response.ShowValidatorError(c, err)
		return
	}
	s := row.Mobile
	row.Mobile = string([]byte(s)[0:3]) + "****" + string([]byte(s)[6:])
	response.ShowData(c, row)
	return
}
