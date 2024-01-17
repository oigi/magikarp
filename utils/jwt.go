package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/oigi/Magikarp/global"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(username, secretKey string) (string, error) {
	ep, _ := ParseDuration(global.CONFIG.JWT.ExpiresTime)
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),         // 生效时间
			Issuer:    global.CONFIG.JWT.Issuer,               // 签名的发行者
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

// ParseJwt 解析JWT
func ParseJwt(token, secretKey string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
