/**
Municipios model
---------------
For manager al cities in Colombia, for associate to `vehiculo` model
*/

package models

type Municipio struct {
	Base
	Nombre       string `gorm:"size:120;not null"  json:"nombre"`
	Departamento string `gorm:"size:120;not null"  json:"departamento"`
}

func (Municipio) TableName() string {
	return "municipios"
}
