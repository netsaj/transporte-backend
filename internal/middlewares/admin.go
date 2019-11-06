package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

/*

 */
func ValidateAdminMiddleware() echo.MiddlewareFunc {
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
