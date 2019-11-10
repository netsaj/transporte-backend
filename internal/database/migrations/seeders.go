package migrations

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"github.com/netsaj/transporte-backend/internal/utils"
)

// link all models an automigrate to database mapping all changes
// if not exist a user with username`admins` automatically is created.
// with username: admins , password: admins
func CreateAdminIfNotExist() {
	dbClient := database.GetConnection()
	defer dbClient.Close()

	// add admins user if no exist
	var user models.User
	// find User with username `admins`
	if err := dbClient.Where("username = ?", "admins").First(&user).Error; err != nil {
		user.Username = "admins"
		user.Name = "admins"
		user.Role = "Administrator"
		user.Password, _ = utils.HashPassword("admins")
		if err = dbClient.Create(&user).Error; err != nil {
			print(err)
		}
		print("User 'admins' create")
		spew.Dump(&user)

	}
	println("sync success")
}
