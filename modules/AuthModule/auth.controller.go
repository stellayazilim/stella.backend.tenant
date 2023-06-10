package AuthModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenant/common/dto"
	"github.com/stellayazilim/stella.backend.tenant/helpers"
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule"
	"log"
	"net/http"
)

type IAuthController interface {
	LoginWithCredentials(ctx *gin.Context)
}
type authController struct {
	authService IAuthService
	userService UserModule.IUserService
}

func AuthController() IAuthController {
	return &authController{
		authService: AuthService(),
		userService: UserModule.UserService(),
	}
}

func (c authController) LoginWithCredentials(ctx *gin.Context) {
	loginData := dto.UserLoginDto{}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": helpers.ListOfErrors(err),
		})
		log.Fatal(err)
		return
	}

	user := models.User{
		Email: loginData.Email,
	}

	// get user if exist
	if err := c.userService.GetUserByEmail(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors": []string{"email does not exist"},
		})
		return
	}
	fmt.Println("istek buraya ulasti : before validate pass", user)
	// check password of user
	if err := c.authService.ValidatePassword(&user, loginData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"errors": []string{"email and password does not match"},
		})
		return
	}
	fmt.Println("istek buraya ulasti : after validate pass")
	// sign tokens
	var (
		signedTokens tokens
		err          error
	)
	fmt.Println("istek buraya ulasti : before sign tokens")
	if signedTokens, err = c.authService.SignTokens(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}
	fmt.Println("istek buraya ulasti : after validate pass")
	// return tokens
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  signedTokens.AccessToken,
		"refreshToken": signedTokens.RefreshToken,
	})

}
