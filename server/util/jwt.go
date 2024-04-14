package util

import (
	"db_proj/define"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// rsa 私钥加密, 公钥解密
func GenJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Issuer:    define.Issuer,
		Subject:   "RegisterToken",
		Audience:  nil,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(define.ExpireTime)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	result, _ := token.SignedString(define.JWTPrivateToken)
	return result
}

func CheckToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return &define.JWTPrivateToken.PublicKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}
