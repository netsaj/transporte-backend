package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CreateAdminGroup(path string, e *echo.Echo) *echo.Group {
	r := e.Group("/admin")

	// Configure middlewares
	r.Use(middleware.JWT([]byte("secret")), ValidateAdminRole(),)

	return r
}

func ValidateAdminRole() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			isAdmin := claims["admin"].(bool)
			if isAdmin {
				println("is user admin")
				return next(c)
			}
			return echo.ErrUnauthorized
		}
	}
}
