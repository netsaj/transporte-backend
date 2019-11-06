/**

 */
package main;

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/routes"
	_ "go/doc"
	"net/http"
)

// Instance echo server and start
func main() {
	// sync postgres database
	database.SyncModels()

	e := echo.New()
	e.Debug = true
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// authentication routes
	routes.AuthRoutes(e)
	// administrative routes
	routes.AdminRoutes(e)

	// Start server
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} \n",
	}))
	e.Logger.Fatal(e.Start(":3000"))
}
