package Services

import (
	"github.com/stellayazilim/stella.backend.tenant/database"
	"github.com/stellayazilim/stella.backend.tenant/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
Only add methods to expose at controllers
*/
type IAuthService interface {
	LoginWithCredentials(data *Types.UserLoginWithCredentialRequest) (Types.TokensResponse, error)
}

type authService struct {
	Database     *gorm.DB
	UserService  IUserService
	TokenService ITokenService
}

func AuthService() IAuthService {
	return &authService{
		UserService:  UserService(),
		TokenService: TokenService(),
		Database:     database.DB.GetDatabase(),
	}
}

func (s *authService) LoginWithCredentials(data *Types.UserLoginWithCredentialRequest) (Types.TokensResponse, error) {
	// get user with email
	user := Types.User{}
	s.Database.Joins("Role").First(&user, "email = ?", data.Email)
	if err := s.ComparePassword(user.Password, []byte(data.Password)); err != nil {
		return Types.TokensResponse{}, err
	}

	// if user has correct credentials sign tokens
	accessToken := s.TokenService.SignAccessToken(&user)
	refreshToken := s.TokenService.SignRefreshToken(&user)

	return Types.TokensResponse{
		AccessToken:  s.TokenService.GetTokenString("ACCESS_TOKEN_SECRET", accessToken),
		RefreshToken: s.TokenService.GetTokenString("REFRESH_TOKEN_SECRET", refreshToken),
	}, nil

}

// compare if password is valid
func (_ *authService) ComparePassword(hash []byte, password []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return err
	}
	return nil
}
