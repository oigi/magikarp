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
func GenerateJWT(username string, privateKey []byte) (string, error) {
	exp, _ := ParseDuration(global.CONFIG.JWT.ExpiresTime)
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)), // 过期时间7天
			IssuedAt:  jwt.NewNumericDate(time.Now()),          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),          // 生效时间
			Issuer:    global.CONFIG.JWT.Issuer,                // 签名的发行者
		},
	}
	// 使用RS256签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := token.SignedString(privateKey)

	return s, err
}

// ParseJwt 解析JWT
func ParseJwt(tokenString string, publicKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
