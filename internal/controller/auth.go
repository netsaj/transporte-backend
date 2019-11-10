package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type AuthMananger struct{}

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Admin    bool      `json:"admins"`
	jwt.StandardClaims
}

func (AuthMananger) Login(c echo.Context) (err error) {

	type Params struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}

	u := new(Params)

	if err = c.Bind(u); err != nil {
		print(err)
		return echo.ErrBadRequest
	}
	DB := database.GetConnection()
	defer DB.Close()

	var user models.User
	if err := DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		print(err)
	}
	if user.CheckPassword(u.Password) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set custom claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["username"] = user.Username
		claims["admins"] = user.IsAdmin()
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Create token with claims
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		// Generate encoded token and send it as response.
		c.Response().Header().Set("Authorization", t)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
