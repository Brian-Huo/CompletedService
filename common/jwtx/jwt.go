package jwtx

import "github.com/golang-jwt/jwt"

func GetToken(secretKey string, iat, seconds, uid int64, role int) (string, error) {
	// role = {0: "company", 1: "employee", 2: "customer"}
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	claims["role"] = role
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
