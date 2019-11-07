package migrations

import (
	"github.com/netsaj/transporte-backend/internal/database"
	"github.com/netsaj/transporte-backend/internal/database/models"
)

func CreateIndexes() {
	db := database.GetConnection()
	defer db.Close()
	// set Users primary key
	_ = db.Exec("alter table users add constraint users_pk primary key (id);")
	// EmpresaVehiculos relationships
	db.Model(models.EmpresaVehiculo{}).AddForeignKey("empresa_id", "empresas(id)", "RESTRICT", "RESTRICT")
	db.Model(models.EmpresaVehiculo{}).AddForeignKey("vehiculo_id", "vehiculos(id)", "RESTRICT", "RESTRICT")

	// Vehiculo unique keys
	// - not repeat vehiculo by vehicle license plate
	db.Model(models.Vehiculo{}).AddUniqueIndex("vehiculo_placa_uk", "placa")

	// Vehiculo relationships
	db.Model(models.Vehiculo{}).AddForeignKey("marca_id", "vehiculo_marcas(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Vehiculo{}).AddForeignKey("tipo_carroceria_id", "vehiculo_carrocerias(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Vehiculo{}).AddForeignKey("tipo_combustible_id", "vehiculo_combustibles(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Vehiculo{}).AddForeignKey("radio_accion_id", "vehiculo_radios_accion(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Vehiculo{}).AddForeignKey("nivel_servicio", "vehiculo_niveles_servicio(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Vehiculo{}).AddForeignKey("matricula_municipio", "municipios(id)", "RESTRICT", "RESTRICT")

	//Concepto relationships
	db.Model(models.Concepto{}).AddForeignKey("concepto_base_calculo_id", "conceptos_bases_calculo(id)", "RESTRICT", "RESTRICT")

	//Recaudo relationships
	db.Model(models.Recaudo{}).AddForeignKey("empresa_id", "empresas(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Recaudo{}).AddForeignKey("vehiculo_id", "vehiculos(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Recaudo{}).AddForeignKey("concepto_id", "conceptos(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Recaudo{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	//TarjetaOperacion relationships
	db.Model(models.TarjetaOperacion{}).AddForeignKey("vehiculo_id", "vehiculos(id)", "RESTRICT", "RESTRICT")
	db.Model(models.TarjetaOperacion{}).AddForeignKey("vehiculo_id", "vehiculos(id)", "RESTRICT", "RESTRICT")
	db.Model(models.TarjetaOperacion{}).AddForeignKey("clase_serivicio_id", "clases_servicio(id)", "RESTRICT", "RESTRICT")
	db.Model(models.TarjetaOperacion{}).AddForeignKey("sede_municipio", "municipios(id)", "RESTRICT", "RESTRICT")
	db.Model(models.TarjetaOperacion{}).AddForeignKey("sede_municipio", "municipios(id)", "RESTRICT", "RESTRICT")
	db.Model(models.TarjetaOperacion{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	//Resolucion relationships
	db.Model(models.Resolucion{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Resolucion{}).AddForeignKey("empresa_id", "empresas(id)", "RESTRICT", "RESTRICT")

	//Desvinculacion relationships
	db.Model(models.Desvinculacion{}).AddForeignKey("resolucion_id", "resoluciones(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Desvinculacion{}).AddForeignKey("empresa_id", "empresas(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Desvinculacion{}).AddForeignKey("tarjeta_operacion_id", "tarjetas_operacion(id)", "RESTRICT", "RESTRICT")
	db.Model(models.Desvinculacion{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
