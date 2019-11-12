package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"net/http"
)

type EmpresaManager struct {
}

func (EmpresaManager) Create(c echo.Context) error {
	var payload models.Empresa

	if err := c.Bind(&payload); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	if err := c.Validate(&payload); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	db := database.GetConnection()
	defer db.Close()
	if result := db.Create(&payload); result.Error != nil {
		fmt.Println(result.Error.Error())
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"error": "can't create 'empresa', details: " + result.Error.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"empresa": &payload,
	})
}

func (EmpresaManager) Update(c echo.Context) error {
	var payload models.Empresa
	if err := c.Bind(&payload); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "can't load payload. Details: " + err.Error(),
		})
	}
	if err := c.Validate(&payload); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	db := database.GetConnection()
	defer db.Close()
	if result := db.Model(&models.Empresa{}).Update(&payload).First(&payload, "id = ?", &payload.ID); result.Error != nil {
		fmt.Println(result.Error.Error())
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"error": "can't update 'empresa', details: " + result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"empresa": &payload,
	})
}

func (EmpresaManager) Get(c echo.Context) error {
	id := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	var payload models.Empresa
	if result := db.Where("id = ?", id).First(&payload); result.Error != nil {
		fmt.Println(result.Error)
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"payload": &payload,
	})
}
