package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Resolucion struct {
	Base
	Numero      int32     `gorm:"not null" json:"numero"`
	Fecha       time.Time `gorm:"not null" json:"fecha"`
	Descripcion string    `json:"descripcion"`
	UserID      uuid.UUID `gorm:"not null"  json:"user_id"` // elaborado por
	EmpresaID   uuid.UUID `json:"empresa_id"`
}

func (Resolucion) TableName() string {
	return "resoluciones"
}
