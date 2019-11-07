package models

type Empresa struct {
	Base
	Nombre                 string            `gorm:"size:200;not null" json:"nombre"`
	Nit                    string            `gorm:"varchar(100);unique_index" json:"nit"`
	RepresentanteNombre    string            `gorm:"size:200" json:"representante_nombre"`
	RepresentanteDocumento uint32            `gorm:"size:11" json:"representante_documento"`
	Direccion              string            `json:"direccion"`
	Telefono               uint32            `gorm:"size:11" json:"telefono"`
	Email                  string            `gorm:"size:100" json:"email"`
	EmpresaVehiculos       []EmpresaVehiculo `json:"empresa_vehiculos"`
}

func (Empresa) TableName() string {
	return "empresas"
}
