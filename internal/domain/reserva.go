package domain

import "time"

type Reserva struct {
	ID                uint        `json:"id" gorm:"primaryKey"`
	IdUsuarioReserva  uint        `json:"id_usuario_reserva"`
	IdInmuebleReserva uint        `json:"id_inmueble_reserva"`
	FechaInicio       time.Time   `json:"fecha_inicio"`
	FechaFin          time.Time   `json:"fecha_fin"`
	Inquilinos        []Inquilino `json:"inquilinos" gorm:"foreignKey:IdReservaInquilino"`
	Activa            bool        `json:"activa" gorm:"default:true"`
	CheckOut          bool        `json:"check_out" gorm:"default:false"`
	Estado            string      `json:"estado"`
}

type ReservaJson struct {
	IdUsuarioReserva  uint        `json:"id_usuario_reserva"`
	IdInmuebleReserva string      `json:"id_inmueble_reserva"`
	FechaInicio       string      `json:"fecha_inicio"`
	FechaFin          string      `json:"fecha_fin"`
	Inquilinos        []Inquilino `json:"inquilinos" gorm:"foreignKey:IdReservaInquilino"`
}
