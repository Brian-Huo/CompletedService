package jwtx

import (
	"context"
	"encoding/json"

	"github.com/golang-jwt/jwt"
)

func GetToken(secretKey string, iat int64, seconds int64, uid int64, role int) (string, error) {
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

func GetTokenDetails(ctx context.Context) (int64, int, error) {
	uid, err := ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return -1, -1, err
	}

	role, err := ctx.Value("role").(json.Number).Int64()
	if err != nil {
		return -1, -1, err
	}

	return uid, int(role), err
}
