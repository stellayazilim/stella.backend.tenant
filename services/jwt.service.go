package Services

import (
	"github.com/golang-jwt/jwt/v5"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"os"
	"time"
)

type ITokenService interface {
	SignAccessToken(user *Types.User) *jwt.Token
	SignRefreshToken(user *Types.User) *jwt.Token
	GetTokenString(tokenType string, token *jwt.Token) string
}

func TokenService() ITokenService {
	return &tokenService{}
}

type tokenService struct {
}

func (s tokenService) SignAccessToken(user *Types.User) *jwt.Token {
	ac := Types.AccessTokenClaims{
		Name: user.Name,
		Role: *user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("TENANT_ID"),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(15))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
}

func (s tokenService) SignRefreshToken(user *Types.User) *jwt.Token {
	rc := Types.RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("TENANT_ID"),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24) * time.Duration(7))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
}

func (s tokenService) GetTokenString(tokenType string, token *jwt.Token) string {
	if ts, err := token.SignedString([]byte(os.Getenv(tokenType))); err != nil {
		return ""
	} else {
		return ts
	}
}
