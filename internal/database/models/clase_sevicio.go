package models

type ClaseSerivicio struct {
	Base
	Nombre string `gorm:"not null" json:"nombre"`
}

func (ClaseSerivicio) TableName() string {
	return "clases_servicio"
}
