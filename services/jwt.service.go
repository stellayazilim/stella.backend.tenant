package Services

import "github.com/golang-jwt/jwt/v5"

type accessTokenClaims struct {
	Name  string
	Perms []byte
	jwt.RegisteredClaims
}

type refreshTokenClaims struct {
	jwt.RegisteredClaims
}
