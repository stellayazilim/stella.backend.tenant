package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stellayazilim/stella.backend.tenant/misc"
	"github.com/stellayazilim/stella.backend.tenant/modules/AuthModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/CategoryModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ContentModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ValidationModule"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {

	// init .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	/*
		stack module initializers here
	*/
	// init database
	DatabaseModule.InitDatabaseModule()
	// add administrator user
	misc.Initalize()
	// migrate database on startup
	DatabaseModule.MigrateDB()
	// init router
	router := gin.Default()

	router.Use(CORSMiddleware())
	// init user module
	UserModule.InitUserModule(router.Group("users"))
	// init auth module
	AuthModule.InitAuthModule(router.Group("auth"))
	// init category module
	CategoryModule.InitCategoryModule(router.Group("categories"))
	// init product module
	ProductModule.InitProductModule(router.Group("products"))
	// init validation module
	ValidationModule.InitValidationModule(router.Group("validation"))
	// init content module
	ContentModule.InitContentModule(router.Group("public"))
	// listen server
	router.Run(":8080")

}
