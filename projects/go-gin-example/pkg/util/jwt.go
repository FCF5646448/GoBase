package util

import (
	"time"

	"github.com/fcf/go-gin-example/pkg/setting"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成Token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()

	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // hash算法
	token, err := tokenClaims.SignedString(jwtSecret)                // 生成签名字符串
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	// 对token进行解码和校验
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
