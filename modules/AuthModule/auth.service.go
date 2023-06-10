package AuthModule

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"github.com/stellayazilim/stella.backend.tenant/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type IAuthService interface {
	SignTokens(user *models.User) (tokens, error)
	getRefreshTokenString(token *jwt.Token) (string, error)
	getAccessTokenString(token *jwt.Token) (string, error)
	ValidatePassword(user *models.User, dto dto.UserLoginDto) error
}

type accessTokenClaims struct {
	Name  string
	Perms []byte
	jwt.RegisteredClaims
}

type refreshTokenClaims struct {
	jwt.RegisteredClaims
}
type authService struct {
	Tokens tokens
}
type tokens struct {
	AccessToken  string
	RefreshToken string
}

func AuthService() IAuthService {
	return &authService{}
}

func (s authService) ValidatePassword(user *models.User, dto dto.UserLoginDto) error {
	fmt.Println("istek buraya ulasti : Validate Password")
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(dto.Password)); err != nil {
		return err
	}

	fmt.Println("istek buraya ulasti : Validate password executed")
	return nil
}
func (s authService) SignTokens(user *models.User) (tokens, error) {
	// access token
	fmt.Println("istek buraya ulasti : sign tokens")
	ac := accessTokenClaims{
		Name:  user.Name,
		Perms: user.Role.Perms,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("TENANT_ID"),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(15))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// refresh token
	rc := refreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("TENANT_ID"),
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24) * time.Duration(7))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	/**
	to do convert this a routine
	*/
	var (
		accessToken  string
		refreshToken string
		err          error
	)
	if accessToken, err = s.getAccessTokenString(jwt.NewWithClaims(jwt.SigningMethodHS256, ac)); err != nil {
		log.Fatal(err)
		return tokens{}, err
	}

	if refreshToken, err = s.getRefreshTokenString(jwt.NewWithClaims(jwt.SigningMethodHS256, rc)); err != nil {
		log.Fatal(err)
		return tokens{}, err
	}

	return tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s authService) getRefreshTokenString(token *jwt.Token) (string, error) {
	var (
		ts  string
		err error
	)
	if ts, err = token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET"))); err == nil {

		return ts, nil
	}

	fmt.Println(err)

	return "", fmt.Errorf("signing refresh token failed")
}

func (s authService) getAccessTokenString(token *jwt.Token) (string, error) {
	var (
		ts  string
		err error
	)
	if ts, err = token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET"))); err == nil {
		return ts, nil
	}
	fmt.Println(err)
	return "", fmt.Errorf("signing access token failed")
}

func (s authService) LoginWithCredentials(user models.User) {

}
