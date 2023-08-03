package Types

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokensResponse struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenClaims struct {
	Name  string
	Perms []byte
	Role  Role
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}
