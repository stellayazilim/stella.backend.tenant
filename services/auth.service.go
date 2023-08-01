package Services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stellayazilim/stella.backend.tenant/Database"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

/*
Only add methods to expose at controllers
*/
type IAuthService interface {
	LoginWithCredentials(data *Types.User)
}

type authService struct {
	Tokens      tokens
	Database    *gorm.DB
	UserService IUserService
}
type tokens struct {
	AccessToken  string
	RefreshToken string
}

func AuthService() IAuthService {
	return &authService{
		UserService: UserService(),
		Database:    Database.DB.GetDatabase(),
	}
}

func (s *authService) SignTokens(user *Types.User) (tokens, error) {
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

func (s *authService) getRefreshTokenString(token *jwt.Token) (string, error) {
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

func (s *authService) getAccessTokenString(token *jwt.Token) (string, error) {
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

func (s *authService) LoginWithCredentials(data *Types.User) {
	// get user with email
	user := Types.User{}
	s.Database.First(&user, "email = ?", user.Email)
	if err := s.ComparePassword(user.Password, data.Password); err != nil {
		// password is not valid
	}
	// login process
}

// compare if password is valid
func (_ *authService) ComparePassword(hash []byte, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return err
	}
	return nil
}
