//jwt的使用
package utils

import (
	"time"

	"errors"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/mazhaoyong/zero/settings"
)

const (
	tokenDuration = 72
	expireOffset  = 3600
)

//NewJWTToken 生成新的JWT token
func NewJWTToken(data jwt.MapClaims) (string, error) {
	data["exp"] = time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix()
	data["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(settings.Get().SecretKey))
}

//TokenRemainingValidity token还有多长时间的有效期
func TokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

//JWTMiddleWare 生成jwt的中间件
func JWTMiddleWare() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//校验token有效时间
			if data, ok := token.Claims.(jwt.MapClaims); ok {
				end := data["exp"] //jwt token中保存的过期时间
				//start := data["iat"]
				now := time.Now().Unix()
				if iend, y := end.(int64); y {
					if iend < now {
						return nil, errors.New("Token is out of date.Please refresh the token!")
					}
				}
			}
			return []byte(settings.Get().SecretKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
