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

var DbUri string

func init() {
	username := "netsaj"
	password := "fabioe9009"
	dbName := "transporte"
	dbHost := "localhost"
	dbPort := "5432"
	DbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
}

func GetConnection() *gorm.DB {
	dbClient, err := gorm.Open(
		"postgres",
		DbUri)

	if err != nil {
		panic("failed connect to database")
	}
	return dbClient
}

func SyncModels() {
	dbClient := GetConnection()
	defer dbClient.Close()
	dbClient.AutoMigrate(
		&models.User{},
	)
	// add admin user if no exist
	var user models.User
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
