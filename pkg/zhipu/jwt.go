package zhipu

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type ZhipuClaims struct {
	APIKey    string `json:"api_key"`
	Timestamp int64  `json:"timestamp"`
	jwt.RegisteredClaims
}

// GenerateToken 通过zhipu apikey生成jwt token\n
// expSeconds 的单位是秒
func GenerateToken(key string, expSeconds int) (string, error) {
	parts := strings.Split(key, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid apikey")
	}
	id, secret := parts[0], parts[1]
	claims := ZhipuClaims{
		APIKey:    id,
		Timestamp: time.Now().UnixMilli(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expSeconds) * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
