package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
)

func UserStandardMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			id := claims["id"].(bool)
			DB := database.GetConnection()
			defer DB.Close()
			DB.Where("id = ?", id)

			return echo.ErrUnauthorized
		}
	}
}


