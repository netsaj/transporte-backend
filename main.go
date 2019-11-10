/**

 */
package main;

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/netsaj/transporte-backend/internal/database/migrations"
	"github.com/netsaj/transporte-backend/internal/routes"
	"github.com/netsaj/transporte-backend/internal/utils"
	_ "go/doc"
	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Instance echo server and start
func main() {
	// sync postgres database
	migrations.CreateTables()
	migrations.CreateIndexes()
	migrations.CreateAdminIfNotExist()

	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Debug = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// validator request
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// authentication routes
	routes.AuthRoutes(e)
	// administrative routes
	routes.AdminRoutes(e)
	// routes for all logged users (admins and standards)
	routes.StandardUsersRoutes(e)

	// Start server
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} \n",
	}))
	e.Logger.Fatal(e.Start(":3000"))
}
