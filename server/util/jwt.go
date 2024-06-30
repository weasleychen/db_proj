package util

import (
	"db_proj/define"
	"db_proj/model"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// rsa 私钥加密, 公钥解密
func GenJWT(extraArgs ...string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Issuer:    define.Issuer,
		Subject:   "RegisterToken",
		Audience:  extraArgs,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(define.ExpireTime)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	result, _ := token.SignedString(define.JWTPrivateToken.PublicKey)
	return result
}

func CheckToken(tokenString string) (bool, model.User) {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return &define.JWTPrivateToken, nil
	})

	user := model.User{}
	if err != nil || !token.Valid {
		return false, user // empty user
	}

	// 此处的err, 追踪到源码发现一定是nil
	claims, _ := token.Claims.GetAudience()
	json.Unmarshal([]byte(claims[0]), &user)

	return true, user
}
