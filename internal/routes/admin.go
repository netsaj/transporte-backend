package routes

import (
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/controller"
	"github.com/netsaj/transporte-backend/internal/middlewares"
)

/*
Init all routes for admin Role and set middlewares for validate
token and authorization token
*/
func AdminRoutes(e *echo.Echo) {
	g := e.Group("/admin")
	// Configure middlewares
	g.Use(middlewares.CustomJwtMiddleware(), middlewares.ValidateAdminMiddleware(), )
	users_routes(g)
}

func users_routes(g *echo.Group) {
	g.POST("/user", controller.UsersManager{}.Create)
	g.GET("/user/:id", controller.UsersManager{}.Read)
	g.PUT("/user/:id", controller.UsersManager{}.Update)
	g.PATCH("/user/:id", controller.UsersManager{}.UpdatePassword)
	g.GET("/users/:active", controller.UsersManager{}.List)
}
