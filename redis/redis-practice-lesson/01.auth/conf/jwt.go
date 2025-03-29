package conf

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey = "abcdefghijklmnopqrstuvwxyz0123456789"
	MaxAge    = 3600 * 24
	MoreAge   = 3600
)

type CustomClaims struct {
	UserId int64
	jwt.StandardClaims
}

func (cc *CustomClaims) MakeToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cc)
	return token.SignedString([]byte(SecretKey))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名错误:%v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
