package models

type Empresa struct {
	Base
	Nombre                 string            `gorm:"size:200;not null" json:"nombre" validate:"required"`
	Nit                    string            `gorm:"varchar(100);unique_index" json:"nit" validate:"required"`
	RepresentanteNombre    string            `gorm:"size:200" json:"representante_nombre" validate:"required"`
	RepresentanteDocumento uint32            `gorm:"size:11" json:"representante_documento" validate:"required,gte=1"`
	Direccion              string            `json:"direccion" validate:"required"`
	Telefono               uint32            `gorm:"size:11" json:"telefono" validate:"required,gte=5000000"`
	Email                  string            `gorm:"size:100" json:"email" validate:"required,email"`
	EmpresaVehiculos       []EmpresaVehiculo `json:"empresa_vehiculos"`
	CapacidadGeneradora    uint64            `json:"capacidad_generadora" validate:"required,gte=1"`
	EsColectivo            bool              `json:"es_colectivo"`
}

func (Empresa) TableName() string {
	return "empresas"
}
