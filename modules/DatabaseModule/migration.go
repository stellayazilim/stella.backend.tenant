package DatabaseModule

import "github.com/stellayazilim/stella.backend.tenant/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.Settings{})
	DB.AutoMigrate(&models.Session{})
	DB.AutoMigrate(&models.Validation{})
}
