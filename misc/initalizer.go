package misc

import (
	"github.com/stellayazilim/stella.backend.tenant/models"
	"github.com/stellayazilim/stella.backend.tenant/modules/DatabaseModule"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Initalize() {

	// create root user if it does not exist
	users := []models.User{}
	d := DatabaseModule.DB.Joins("JOIN roles ON users.role_id = roles.id").Where("roles.name = ?", "administrator").Find(&users).RowsAffected

	if d == 0 {
		role := models.Role{
			Name:        "administrator",
			Description: "root role",
			Perms:       []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
		}
		password, err := bcrypt.GenerateFromPassword([]byte("administrator"), 6)

		if err != nil {
			log.Fatal(err)
		}
		user := models.User{
			Name:        "administrator",
			Email:       "administrator@elitasmakina.com",
			Password:    password,
			IsValidated: true,
			Role:        &role,
		}
		DatabaseModule.DB.Create(&user)
	}
}
