package controller

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"github.com/netsaj/transporte-backend/internal/utils"
	"net/http"
)

type UsersManager struct {
}

func (UsersManager) Create(c echo.Context) (err error) {
	user := new(models.User)
	DB := database.GetConnection()
	defer DB.Close()
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}

func (UsersManager) Read(c echo.Context) (err error) {
	id := c.Param("id")
	DB := database.GetConnection()
	defer DB.Close()
	var user models.User
	err = DB.Where("id = ? ", id).First(&user).Error;
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (UsersManager) List(c echo.Context) (err error) {
	active := c.Param("active") == "true"
	DB := database.GetConnection()
	defer DB.Close()
	var users [] models.User
	DB.Where("active = ?", active).Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (UsersManager) Update(c echo.Context) error {
	var user models.User
	type updateUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Role     string `json:"role"`
		Active   bool   `json:"active"`
	}
	payload := new(updateUser)
	id := c.Param("id")
	fmt.Printf("user id to update:%s", id)
	if err := c.Bind(payload); err != nil {
		println(err.Error())
		return echo.ErrBadRequest
	}
	println(payload)
	DB := database.GetConnection()
	defer DB.Close()
	result := DB.Model(&user).Where("id = ?", id).Update(&payload).First(&user, "id = ?", id)
	if result.Error != nil {
		println(result.Error)
		return result.Error
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (UsersManager) UpdatePassword(c echo.Context) error {
	id := c.Param("id")
	type updatePassword struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	payload := new(updatePassword)
	if err := c.Bind(payload); err != nil {
		return echo.ErrBadRequest
	}
	spew.Dump(payload)
	DB := database.GetConnection()
	defer DB.Close()
	user := new(models.User)
	result := DB.Where("id = ? ", id).First(&user)
	if result.Error != nil {
		println(result.Error)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": result,
		})
	}

	if user.CheckPassword(payload.OldPassword) {
		println("password match")
		newPassword, _ := utils.Crypto{}.HashPassword(payload.NewPassword)
		result := DB.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword)
		if result.Error != nil {
			panic(result.Error.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "update done",
		})
	}
	return echo.ErrUnauthorized

}
