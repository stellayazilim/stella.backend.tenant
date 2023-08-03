package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stellayazilim/stella.backend.tenant/dataase"
	"github.com/stellayazilim/stella.backend.tenant/middlewares"
	"github.com/stellayazilim/stella.backend.tenant/modules/AuthModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/CategoryModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ContentModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ProductModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/UserModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ValidationModule"
	"log"
)

func main() {

	// init .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	/*
		stack module initializers here
	*/
	// init database
	dataase.DB.InitDb()
	// migrate database on startup
	dataase.DB.Migrate()
	// init router
	router := gin.Default()

	// apply cors middleware for every request
	router.Use(middlewares.CORSMiddleware("*"))
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
