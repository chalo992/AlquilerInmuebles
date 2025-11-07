package domain

import (
	"time"
)

type Usuario struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Nombre       string    `json:"nombre"`
	Apellido     string    `json:"apellido"`
	Nacionalidad string    `json:"nacionalidad"`
	Telefono     int       `json:"telefono"`
	DNI          int       `json:"dni"`
	Rol          string    `json:"rol"`
	Tarjeta      *Tarjeta  `gorm:"foreignKey:IdUsuarioTarjeta"`
	Reservas     []Reserva `gorm:"foreignKey:IdUsuarioReserva"`
	Creacion     time.Time `json:"-"`
}
