package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

var jwtSecret = []byte("38324-search-engine") // TODO 从配置文件读取

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(id int64, username string) (accessToken, refreshToken string, err error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)
	rtExpireTime := now.Add(10 * 24 * time.Hour)

	// Access Token
	accessTokenClaims := Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 过期时间
			Issuer:    "your_issuer_here",             // 签名的发行者 TODO
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to get accessToken")
	}

	// Refresh Token
	refreshTokenClaims := Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(rtExpireTime), // 过期时间
			Issuer:    "your_issuer_here",               // 签名的发行者 TODO
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to get refreshToken")
	}

	return accessToken, refreshToken, nil
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
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ParseRefreshToken 验证用户token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to parse accessToken")
	}

	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to parse refreshToken")
	}

	if accessClaim.ExpiresAt.Unix() > time.Now().Unix() {
		// 如果 access_token 没过期,每一次请求都刷新 refresh_token 和 access_token
		return GenerateJWT(accessClaim.ID, accessClaim.Username)
	}

	if refreshClaim.ExpiresAt.Unix() > time.Now().Unix() {
		// 如果 access_token 过期了,但是 refresh_token 没过期, 刷新 refresh_token 和 access_token
		return GenerateJWT(accessClaim.ID, accessClaim.Username)
	}

	// 如果两者都过期了,重新登陆
	return "", "", errors.New("身份过期，重新登陆")
}
