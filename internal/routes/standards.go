package routes

import (
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/controller"
	"github.com/netsaj/transporte-backend/internal/middlewares"
)

/*
Init all routes for admins Role and set middlewares for validate
token and authorization token
*/
func StandardUsersRoutes(e *echo.Echo) {
	g := e.Group("/api/standard")
	// Configure middlewares
	g.Use(middlewares.CustomJwtMiddleware(), middlewares.ValidateStandardUserMiddleware(), )
	settingsRoutes(g)
	empresaRoutes(g)
}

// return routes for list options values for each setting table
func settingsRoutes(g *echo.Group) {
	g.GET("/settings/list/:resource", controller.SettingsManager{}.GetSettingsByResource)
}

func empresaRoutes(g *echo.Group) {
	g.POST("/empresa", controller.EmpresaManager{}.Create)
	g.PUT("/empresa", controller.EmpresaManager{}.Update)
	g.GET("/empresa/:id", controller.EmpresaManager{}.Get)
}
