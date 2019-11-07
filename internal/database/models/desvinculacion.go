package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Desvinculacion struct {
	Base
	FechaNotificacion time.Time `json:"fecha_notificacion"`
	//ya esta vinculado con la tarjeta de operacion | VehiculoID uuid.UUID
	TipoDesvinculacion  string `json:"tipo_desvinculacion"`
	FormaDesvinculacion string `json:"forma_desvinculacion"`
	//fecha y numero no es necesario si hago una relacion
	ResolucionID       uuid.UUID `json:"resolucion_id"`
	EmpresaID          uuid.UUID `json:"empresa_id"`
	TarjetaOperacionID uuid.UUID `json:"tarjeta_operacion_id"`
	UserID           uuid.UUID `gorm:"not null" json:"user_id"`
}

func (Desvinculacion) TableName() string {
	return "desvinculaciones"
}
