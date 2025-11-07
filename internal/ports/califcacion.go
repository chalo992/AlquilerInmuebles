package ports

import "AlquilerInmuebles/internal/domain"

type CalificacionService interface {
	CalificarInmueble(calificacion domain.Calificacion, id uint) error
	GetCalificacionInmueble(id string) (*domain.Calificacion, error)
	EliminarCalificacion(id string) error
	EditarCalificacion(calificacion domain.Calificacion) error
	ObtenerCalificacionesInmueble(id string) ([]domain.Calificacion, error)
	ObtenerCalificacionReserva(id string) (*domain.Calificacion, error)
}

type CalificacionRepository interface {
	CalificarInmueble(calificacion domain.Calificacion) error
	EditarCalificacion(calificacion domain.Calificacion) error
	EliminarCalificacion(id uint) error
	GetCalificacion(id uint) (domain.Calificacion, error)
	ObtenerCalificacionesInmueble(id uint) ([]domain.Calificacion, error)
	ObtenerCalificacionReservaId(id uint) (*domain.Calificacion, error)
}
