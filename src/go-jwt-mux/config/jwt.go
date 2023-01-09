package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("fjalsdjfk743128974019274sadjf01")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
