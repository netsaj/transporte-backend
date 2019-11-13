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
	username := "knifpeviokvltf"
	password := "7d75ad000e2f5aca266780785f2839e142cc3d277ba8b45aca5bd69aa685063c"
	dbName := "d5bqhoj50lgam1"
	dbHost := "ec2-107-22-160-102.compute-1.amazonaws.com"
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
