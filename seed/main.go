package main

import (
	"github.com/joho/godotenv"
	"github.com/stellayazilim/stella.backend.tenant/database"
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"log"
)

func main() {
	// init .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	role := Types.Role{
		Name:        "User",
		Perms:       []byte{},
		Description: "",
	}
	database.DB.InitDb()
	db := database.DB.GetDatabase()
	db.Model(&Types.Role{}).Create(&role)
}
