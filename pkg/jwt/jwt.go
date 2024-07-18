package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

var jwtSecret = []byte("36468-sasdh-edasns")

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(id int64, email string) (accessToken string, err error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

	// Access Token
	accessTokenClaims := Claims{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间
			Issuer:    "your_issuer_here",             // 签名的发行者 TODO
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(jwtSecret)
	if err != nil {
		return "", errors.Wrap(err, "failed to get accessToken")
	}

	return accessToken, nil
}

// ParseToken 解析JWT
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token")
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		// 检查 JWT 是否过期
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
