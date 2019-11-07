package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Recaudo struct {
	Valor              float64   `gorm:"not null" json:"valor"`
	NumeroConsignacion string    `gorm:"not null" json:"numero_consignacion"`
	SucursalBancoID    string    `gorm:"not null" json:"sucursal_banco_id"`
	FechaRecaudo       time.Time `json:"fecha_recaudo"`
	FechaConsignacion  time.Time `json:"fecha_consignacion"`
	// relationships
	EmpresaID  uuid.UUID `json:"empresa_id"`
	VehiculoID uuid.UUID `gorm:"not null" json:"vehiculo_id"`
	ConceptoID uuid.UUID `gorm:"not null" json:"concepto_id"`
	UserID     uuid.UUID `gorm:"not null" json:"user_id"`
}

func (Recaudo) TableName() string {
	return "recaudos"
}
