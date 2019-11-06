/**
Connection to database with Gorm ORM

*/
package database

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"github.com/netsaj/transporte-backend/internal/utils"
)

// Db uri connection string
var DbUri string

// Setup connections params for postgres database
func init() {
	username := "netsaj"
	password := "fabioe9009"
	dbName := "transporte"
	dbHost := "localhost"
	dbPort := "5432"
	DbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
}

// return a connection instance
// if all is okey return a DB connection instance, else call Panic(error)
func GetConnection() *gorm.DB {
	dbClient, err := gorm.Open(
		"postgres",
		DbUri)

	if err != nil {
		panic("failed connect to database")
	}
	return dbClient
}

// link all models an automigrate to database mapping all changes
// if not exist a user with username`admin` automatically is created.
// with username: admin , password: admin
func SyncModels() {
	dbClient := GetConnection()
	defer dbClient.Close()
	dbClient.AutoMigrate(
		&models.User{},
	)
	// add admin user if no exist
	var user models.User
	// find User with username `admin`
	if err := dbClient.Where("username = ?", "admin").First(&user).Error; err != nil {
		user.Username = "admin"
		user.Name = "admin"
		user.Role = "Administrator"
		user.Password, _ = utils.Crypto{}.HashPassword("admin")
		if err = dbClient.Create(&user).Error; err != nil {
			print(err)
		}
		print("User 'admin' create")
		spew.Dump(&user)

	}
	println("sync success")
}
