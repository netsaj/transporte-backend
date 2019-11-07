package models

import uuid "github.com/satori/go.uuid"

type Concepto struct {
	Base
	Year   uint16 `json:"año"`
	Nombre string `json:"nombre"`
	//Esto esta en veremos porque varia con el año
	ConceptoBaseCalculoID uuid.UUID `gorm:"not null" json:"concepto_base_calculo_id"`
	Porcentaje            float32   `gorm:"not null" json:"porcentaje"`
}

func (Concepto) TableName() string {
	return "conceptos"
}
