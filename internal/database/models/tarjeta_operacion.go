package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type TarjetaOperacion struct {
	Base
	VigenciaInicio  time.Time `json:"vigencia_inicio"`
	VigenciaFinal   time.Time `json:"vigencia_final"`
	FechaExpedicion time.Time `json:"fecha_expedicion"`
	NumeroTarjeta   string    `json:"numero_tarjeta"`
	FirmadoPor      string    `json:"firmado_por"`
	Radicado        string    `json:"radicado"`
	//relationships
	VehiculoID       uuid.UUID `json:"vehiculo_id"`
	EmpresaID        uuid.UUID `json:"empresa_id"`
	ClaseSerivicioID uuid.UUID `json:"clase_serivicio_id"`
	SedeMunicipio    uuid.UUID `json:"sede_municipio"`
	UserID           uuid.UUID `gorm:"not null" json:"user_id"`
}

func (TarjetaOperacion) TableName() string {
	return "tarjetas_operacion"
}
