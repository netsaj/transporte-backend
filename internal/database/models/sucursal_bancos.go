package models

type SucursalBanco struct {
	Base
	Nombre string `json:"nombre"`
}

func (SucursalBanco) TableName() string {
	return "sucursal_bancos"
}
