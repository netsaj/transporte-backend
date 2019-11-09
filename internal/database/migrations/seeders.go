package migrations

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"github.com/netsaj/transporte-backend/internal/utils"
)

// link all models an automigrate to database mapping all changes
// if not exist a user with username`admin` automatically is created.
// with username: admin , password: admin
func CreateAdminIfNotExist() {
	dbClient := database.GetConnection()
	defer dbClient.Close()

	// add admin user if no exist
	var user models.User
	// find User with username `admin`
	if err := dbClient.Where("username = ?", "admin").First(&user).Error; err != nil {
		user.Username = "admin"
		user.Name = "admin"
		user.Role = "Administrator"
		user.Password, _ = utils.HashPassword("admin")
		if err = dbClient.Create(&user).Error; err != nil {
			print(err)
		}
		print("User 'admin' create")
		spew.Dump(&user)

	}
	println("sync success")
}
