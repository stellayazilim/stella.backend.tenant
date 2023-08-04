package AuthModule

import "github.com/gin-gonic/gin"

func InitAuthModule(r *gin.RouterGroup) {

	c := AuthController()

	r.POST("signin", c.LoginWithCredentials)
	r.GET("/me", c.GetMe)
	r.GET("/refresh", c.Refresh)
	r.POST("/signup", c.Register)
}
