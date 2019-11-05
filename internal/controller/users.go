package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"net/http"
)

type UsersManager struct {
}

func (UsersManager) Init(g *echo.Group) {

	g.POST("/user", UsersManager{}.Create)
	g.GET("/user/:id", UsersManager{}.Read)
	g.PUT("/user/:id", UsersManager{}.Update)
	g.GET("/users/:active", UsersManager{}.List)
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
	type UpdateUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Role     string `json:"role"`
		Active   bool   `json:"active"`
	}
	payload := new(UpdateUser)
	id := c.Param("id")
	fmt.Printf("user id to update:%s", id)
	if err := c.Bind(payload); err != nil {
		println(err.Error())
		return echo.ErrBadRequest
	}
	println(payload)
	DB := database.GetConnection()
	defer DB.Close()
	result := DB.Model(&user).Where("id = ?", id).Update(&payload).First(&user,"id = ?", id)
	if result.Error != nil {
		println(result.Error)
		return result.Error
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
