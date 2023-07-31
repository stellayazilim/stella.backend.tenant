package Database

import "github.com/stellayazilim/stella.backend.tenant/types"

func MigrateDB() {
	DB.AutoMigrate(&types.User{})
	DB.AutoMigrate(&types.Category{})
	DB.AutoMigrate(&types.Product{})
	DB.AutoMigrate(&types.Role{})
	DB.AutoMigrate(&types.Settings{})
	DB.AutoMigrate(&types.Session{})
	DB.AutoMigrate(&types.Validation{})
	DB.AutoMigrate(&types.Image{})
}
