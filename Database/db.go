package Database

import (
	Types "github.com/stellayazilim/stella.backend.tenant/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type database struct {
	db *gorm.DB
}
type IDatabase interface {
	InitDb()
	GetDatabase() *gorm.DB
	Migrate()
}

var DB IDatabase = &database{}

// connect to instance
func (d *database) InitDb() {
	dsn := os.Getenv("DATABASE_URI")
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Can not connect to database %v", err)
	}
	d.db = db
}

// get instance
func (d *database) GetDatabase() *gorm.DB {
	return d.db
}

// migrate database
func (d *database) Migrate() {
	d.db.AutoMigrate(&Types.Category{})
	d.db.AutoMigrate(&Types.Product{})
	d.db.AutoMigrate(&Types.Role{})
	d.db.AutoMigrate(&Types.Settings{})
	d.db.AutoMigrate(&Types.Session{})
	d.db.AutoMigrate(&Types.Validation{})
	d.db.AutoMigrate(&Types.Image{})
}
