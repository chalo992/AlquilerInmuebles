package ports

import (
	"AlquilerInmuebles/internal/domain"
	"mime/multipart"
)

type ReservaService interface {
	ReservarInmueble(id uint) error
	ConfirmarReserva(reservaJson domain.ReservaJson) error
	GenerarPathImagenInquilino(file *multipart.FileHeader) (string, string, error)
	ReservasDelUsuario(id uint) ([]domain.Reserva, error)
	CancelarReserva(id string) (*domain.Inmueble, error)
	GetReservaId(id string) (domain.Reserva, error)
	GetReservas() (*[]domain.Reserva, error)
	ListarReservasEncargado(id uint) ([]domain.Inmueble, error)
}

type ReservaRepository interface {
	ReservarInmueble(id uint) error
	ConfirmarReserva(reserva domain.Reserva) error
	ReservasDelUsuario(id uint) ([]domain.Reserva, error)
	GetReserva(id uint) (*domain.Reserva, error)
	ActualizarReserva(reserva *domain.Reserva) error
	GetReservaId(id uint) (domain.Reserva, error)
	GetReservas() (*[]domain.Reserva, error)
	ListarReservasEncargado(id uint) ([]domain.Inmueble, error)
}
