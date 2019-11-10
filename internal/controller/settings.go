package controller

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
	"net/http"
	"strings"
)

type SettingsMap struct {
	Resource string                 `json:"resource" validate:"required"`
	Values   map[string]interface{} `json:"values" validate:"required"`
}

type SettingsManager struct {
}

var settingsList = [4]string{"combustibles", "carrocerias", "radios_accion", "niveles_servicio"}

func (SettingsManager) ListOptions(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": &settingsList,
	})
}

func (SettingsManager) GetSettingsByResource(c echo.Context) error {
	resource := strings.ToLower(c.Param("resource"))
	db := database.GetConnection()
	defer db.Close()
	if resource == "combustibles" {
		items := new([]models.VehiculoCombustible)
		if err := db.Model(&models.VehiculoCombustible{}).Find(&items, "deleted_at IS NULL"); err.Error != nil {
			panic(err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"resource": resource,
			"items":    &items,
		})
	} else if resource == "radios_accion" {
		items := new([]models.VehiculoRadioAccion)
		if err := db.Model(&models.VehiculoRadioAccion{}).Find(&items, "deleted_at IS NULL"); err.Error != nil {
			panic(err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"resource": resource,
			"items":    &items,
		})
	} else if resource == "carrocerias" {
		items := new([]models.VehiculoCarroceria)
		if err := db.Model(&models.VehiculoCarroceria{}).Find(&items, "deleted_at IS NULL"); err.Error != nil {
			panic(err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"resource": resource,
			"items":    &items,
		})
	} else if resource == "niveles_servicio" {
		items := new([]models.VehiculoNivelServicio)
		if err := db.Model(&models.VehiculoNivelServicio{}).Find(&items, "deleted_at IS NULL"); err.Error != nil {
			panic(err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"resource": resource,
			"items":    &items,
		})
	}
	return echo.ErrNotFound
}

func (SettingsManager) Create(c echo.Context) error {

	var payload = new(SettingsMap)
	if err := c.Bind(&payload); err != nil {
		return echo.ErrInternalServerError
	}
	if err := c.Validate(payload); err != nil {
		return echo.ErrBadRequest
	}
	spew.Dump(payload.Values)
	model, err := generateModel(payload.Resource, payload.Values, false)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"error": fmt.Sprintf("can't save resource '%s', details: %s", payload.Resource, err),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"resource": payload.Resource,
		"item":     &model,
	})
}
func (SettingsManager) Update(c echo.Context) error {

	var payload = new(SettingsMap)
	if err := c.Bind(&payload); err != nil {
		return echo.ErrInternalServerError
	}
	if err := c.Validate(payload); err != nil {
		return echo.ErrBadRequest
	}
	spew.Dump(payload.Values)
	model, err := generateModel(payload.Resource, payload.Values, true)
	if payload.Values["id"] == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": fmt.Sprintf("can't update resource '%s', details: values not have 'id' field for update", payload.Resource),
		})
	}
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]interface{}{
			"error": fmt.Sprintf("can't save resource '%s', details: %s", payload.Resource, err),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"resource": payload.Resource,
		"item":     &model,
	})
}

/*
Function for switch between vehiculo settings models and create or update with `update` flag value.
*/
func generateModel(resource string, input map[string]interface{}, update bool) (interface{}, error) {
	r := strings.ToLower(resource)
	db := database.GetConnection()
	defer db.Close()

	if r == "combustibles" {
		model := new(models.VehiculoCombustible)
		// custom ternary operator
		model.Nombre = map[bool]string{true: input["nombre"].(string), false: ""}[input["nombre"] != nil]
		err := map[bool]error{true: db.Model(&model).Where("id = ?", input["id"]).Update(&model).Error, false: db.Create(&model).Error}[update]
		return model, err
	} else if r == "radios_accion" {
		model := new(models.VehiculoRadioAccion)
		model.Nombre = map[bool]string{true: input["nombre"].(string), false: ""}[input["nombre"] != nil]
		err := map[bool]error{true: db.Model(&model).Where("id = ?", input["id"]).Update(&model).Error, false: db.Create(&model).Error}[update]
		return model, err
	} else if r == "niveles_servicio" {
		model := new(models.VehiculoNivelServicio)
		model.Nombre = map[bool]string{true: input["nombre"].(string), false: ""}[input["nombre"] != nil]
		err := db.Create(&model).Error
		return model, err
	} else if r == "carrocerias" {
		model := new(models.VehiculoCarroceria)
		model.Nombre = map[bool]string{true: input["nombre"].(string), false: ""}[input["nombre"] != nil]
		err := db.Create(&model).Error
		return model, err
	}
	return nil, echo.ErrBadRequest
}
