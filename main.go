package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stellayazilim/stella.backend.tenant/misc"
	"github.com/stellayazilim/stella.backend.tenant/modules/CategoryModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/ContentModule"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
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
	DatabaseModule.InitDatabaseModule()

	// add administrator user
	misc.Initalize()
	// migrate database on startup
	DatabaseModule.MigrateDB()
	// init router
	router := gin.Default()
	// init user module
	UserModule.InitUserModule(router.Group("users"))
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
