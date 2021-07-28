/**
 * @Author:  jager
 * @Email:   lhj168os@gmail.com
 * @File:    jwt
 * @Date:    2021/6/21 11:45 上午
 * @package: jwt
 * @Version: v1.0.0
 *
 * @Description:
 *
 */

package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jageros/hawos/internal/conf"
	"time"
)

var (
	secret  = conf.JWT_SECRET
	timeout = time.Hour * 10
)

func SetOption(jwtSecret string, tokenTimeout time.Duration) {
	secret = jwtSecret
	timeout = tokenTimeout
}

type Claims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid string) (string, error) {
	expireTime := time.Now().Add(timeout).Unix()
	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "HawOs",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func RefreshToken(token string) (newToken string, err error) {
	claims, err := ParseToken(token)
	if err != nil {
		return
	}
	token, err = GenerateToken(claims.Uid)
	return
}
