package AuthModule

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Services "github.com/stellayazilim/stella.backend.tenant/services"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type IAuthController interface {
	LoginWithCredentials(ctx *gin.Context)
	GetMe(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	Register(ctx *gin.Context)
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

func (c *authController) Register(ctx *gin.Context) {
	body := Types.RegisterUserRequestBody{}
	ctx.BindJSON(&body)

	if err := passwordvalidator.Validate(string(body.Password), 30); err != nil {
		fmt.Println("pasword weak")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost); err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	} else {
		body.Password = string(hashedPassword)
	}
	if err := c.authService.RegisterUser(body.ConvertToUser()); err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}

func (c *authController) LoginWithCredentials(ctx *gin.Context) {
	body := Types.UserLoginWithCredentialRequest{}

	if err := ctx.ShouldBindJSON(&body); err != nil {

		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	fmt.Println(body)
	if tokens, err := c.authService.LoginWithCredentials(&body); err != nil {
		ctx.Status(http.StatusUnauthorized)
	} else {
		ctx.JSON(http.StatusOK, tokens)
	}
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
