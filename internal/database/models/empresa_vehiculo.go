package models

import uuid "github.com/satori/go.uuid"

type EmpresaVehiculo struct {
	Base
	Empresa       Empresa
	EmpresaID     uuid.UUID `json:"empresa_id"`
	Vehiculo      Vehiculo `json:"vehiculo"`
	VehiculoID	uuid.UUID `json:"vehiculo_id"`
	NumeroInterno string    `gorm:"not null"  json:"numero_interno"`
}
