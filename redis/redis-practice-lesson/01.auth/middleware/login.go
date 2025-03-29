package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"redis-parctice-lesson/01.auth/conf"
	"redis-parctice-lesson/01.auth/model"
	"time"
)

func DoLogin(user model.User) (*conf.MyJWT, error) {
	if conf.Cfg.OpenJwt {
		token := &conf.MyJWT{}
		customClaims := conf.CustomClaims{
			UserId: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(conf.MaxAge * time.Second).Unix(),
			},
		}

		accessToken, err := customClaims.MakeToken()
		if err != nil {
			return nil, err
		}
		refreshClaims := &conf.CustomClaims{
			UserId: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add((conf.MaxAge + conf.MoreAge) * time.Second).Unix(),
			},
		}
		refreshToken, err := refreshClaims.MakeToken()
		if err != nil {
			return nil, err
		}
		token.Uid = user.ID
		token.AccessToken = accessToken
		token.RefreshToken = refreshToken
		return token, nil
	}
	return nil, nil
}
