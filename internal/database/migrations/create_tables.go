package migrations

import (
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
)

func CreateTables() {
	dbClient := database.GetConnection()
	defer dbClient.Close()
	dbClient.AutoMigrate(
		&models.User{},
		&models.Municipio{},
		&models.Empresa{},
		&models.EmpresaVehiculo{},
		&models.Vehiculo{},
		&models.VehiculoCarroceria{},
		&models.VehiculoMarca{},
		&models.VehiculoNivelServicio{},
		&models.VehiculoRadioAccion{},
		&models.VehiculoCombustible{},
		&models.Concepto{},
		&models.ConceptoBaseCalculo{},
		&models.Recaudo{},
		&models.SucursalBanco{},
		&models.TarjetaOperacion{},
		&models.ClaseSerivicio{},
		&models.Resolucion{},
		&models.Desvinculacion{},
	)
}
