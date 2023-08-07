package main

import (
	"github.com/joho/godotenv"
	"github.com/stellayazilim/stella.backend.tenant/constants/perms"
	"github.com/stellayazilim/stella.backend.tenant/database"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"log"
)

func main() {
	// init .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	database.DB.InitDb()
	database.DB.Migrate()
	db := database.DB.GetDatabase()

	// create root administrator role
	administratorRole := Types.Role{
		Name:        "Administrator",
		Description: "Root admin role",
		Perms: []byte{
			perms.PROMOTE_USER,
			perms.DELETE_USER,
			perms.GET_USERS,
			perms.GET_USER_EMAIL,
			perms.GET_USER_ADDRESS,
			perms.GET_USER_PHONE,
			perms.ADD_ROLE,
			perms.UPDATE_ROLE_NAME,
			perms.UPDATE_ROLE_DESCRIPTION,
			perms.UPDATE_ROLE_PERMS,
			perms.DELETE_ROLE,
			perms.GET_ROLE,
			perms.ADD_PRODUCT,
			perms.GET_PRODUCT,
			perms.UPDATE_PRODUCT_NAME,
			perms.UPDATE_PRODUCT_DESCRIPTION,
			perms.UPDATE_PRODUCT_EXPLANATION,
			perms.UPDATE_PRODUCT_SPECS,
			perms.UPDATE_PRODUCT_CATEGORY,
			perms.DELETE_PRODUCT,
		},
	}
	// create default user role
	userRole := Types.Role{
		Name:        "User",
		Description: "Default user role",
		Perms: []byte{
			perms.GET_PRODUCT,
		},
	}

	// save administrator role to database
	db.Model(&Types.Role{}).Create(&administratorRole)

	// save user role to database
	db.Model(&Types.Role{}).Create(&userRole)

	db.Model(&Types.User{}).Save(&Types.User{
		Name:  "Administrator",
		Email: "administrator@stellasoft.tech",
		Role:  &administratorRole,
	})
}
