/**
Contain all struct for settings and options linked in `Vehiculo` model.
This options required endpoints for user admin can add & remove.
*/

package models

/**

VehiculoMarca Model
--------------------
*/
type VehiculoMarca struct {
	Base
	Nombre string `json:"nombre"`
}

func (VehiculoMarca) TableName() string {
	return "vehiculo_marcas"
}

/**

VehiculoCarroceria Model
--------------------
*/
type VehiculoCarroceria struct {
	Base
	Nombre string `json:"nombre"`
}

func (VehiculoCarroceria) TableName() string {
	return "vehiculo_carrocerias"
}

/**

VehiculoCombustible Model
--------------------
*/
type VehiculoCombustible struct {
	Base
	Nombre string `json:"nombre"`
}

func (VehiculoCombustible) TableName() string {
	return "vehiculo_combustibles"
}

/**

VehiculoRadioAccion Model
--------------------
*/
type VehiculoRadioAccion struct {
	Base
	Nombre string `json:"nombre"`
}

func (VehiculoRadioAccion) TableName() string {
	return "vehiculo_radios_accion"
}

/**

VehiculoNivelServicio Model
--------------------
*/
type VehiculoNivelServicio struct {
	Base
	Nombre string `json:"nombre"`
}

func (VehiculoNivelServicio) TableName() string {
	return "vehiculo_niveles_servicio"
}
