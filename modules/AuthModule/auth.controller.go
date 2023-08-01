package AuthModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	"net/http"
)

type IAuthController interface {
	LoginWithCredentials(ctx *gin.Context)
	GetMe(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}
type authController struct {
	authService Services.IAuthService
	userService Services.IUserService
}

func AuthController() IAuthController {
	us := Services.UserService()
	return &authController{
		authService: Services.AuthService(),
		userService: us,
	}
}

func (c *authController) LoginWithCredentials(ctx *gin.Context) {

}

func (c *authController) GetMe(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")

	fmt.Println(token)
	ctx.JSON(http.StatusOK, gin.H{
		"username": "stella",
	})
}

func (c *authController) Refresh(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")

	fmt.Println(token)

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  token,
		"refreshToken": token,
	})
}
