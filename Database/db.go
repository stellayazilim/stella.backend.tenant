package Database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDatabaseModule() {
	dsn := os.Getenv("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
		return
	}
	DB = db
}
