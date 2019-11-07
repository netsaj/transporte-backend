package models

type ConceptoBaseCalculo struct {
	Base
	Sigla  string  `gorm:"size:10;not null" json:"sigla"`
	Nombre string  `gorm:"not null"  json:"nombre"`
	Valor  float64 `json:"valor"`
}

func (ConceptoBaseCalculo) TableName() string {
	return "conceptos_bases_calculo"
}
