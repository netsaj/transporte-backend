/**
`Vehiculo` model.
---
applied for collective and individual service.
*/
package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Vehiculo struct {
	Base
	Placa               string    `gorm:"size:5;not null" json:"placa"`
	Modelo              string    `gorm:"size:50;not null" json:"modelo"`
	Serie               string    `json:"serie"`
	MarcaID             uuid.UUID `json:"marca_id"`
	Capacidad           string    `json:"capacidad"`
	Linea               string    `json:"linea"`
	NumeroMotor         string    `json:"numero_motor"`
	TipoCarroceriaID    uuid.UUID `json:"tipo_carroceria_id"`
	TipoCombustibleID   uuid.UUID `json:"tipo_combustible_id"`
	RadioAccionID       uuid.UUID `json:"radio_accion_id"`
	NivelServicio       uuid.UUID `json:"nivel_servicio"`
	PropietarioNombre   string    `json:"propietario_nombre"`
	PropietarioCedula   uint64    `json:"propietario_cedula"`
	TarjetaPropiedad    string    `json:"tarjeta_propiedad"`
	MatriculaMunicipio  uuid.UUID `json:"matricula_municipio"`
	MatriculaFecha      time.Time `json:"matricula_fecha"`
	NumeroMetropolitano string    `gorm:"size:4" json:"numero_metropolitano"`
	NumeroChip          string    `gorm:"size:6" json:"numero_chip"`
	// NumeroInterno : created table empresaVehiculo for this relation cuz is only in collective type
	EsColectivo bool `json:"es_colectivo"`
}

func (Vehiculo) TableName() string {
	return "vehiculos"
}
