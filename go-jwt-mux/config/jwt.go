package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("adsgasdfgasdg12341")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}