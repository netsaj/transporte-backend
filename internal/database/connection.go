/**
Connection to database with Gorm ORM

*/
package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
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
