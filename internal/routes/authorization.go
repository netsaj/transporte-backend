package routes

import (
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/controller"
)

/**

 */
func AuthRoutes(e *echo.Echo) {
	g := e.Group("/api/auth")
	g.POST("/login", controller.AuthMananger{}.Login)
}
