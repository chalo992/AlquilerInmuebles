package ports

import (
	"AlquilerInmuebles/internal/domain"
	"mime/multipart"
)

type InmuebleService interface {
	CrearInmueble(inmueble domain.Inmueble) (domain.Inmueble, error)
	PausarDespausarInmueble(id string) error //solo los encargados pueden pausar y despausar los inmuebles
	ActualizarInmueble(inmueble domain.Inmueble) (domain.Inmueble, error)
	EliminarInmueble(id string) ([]string, error)
	ListarInmuebles() ([]domain.Inmueble, error)                 //todos los inmuebles, solo los admin pueden acceder a estos
	ListarInmueblesConFotoNoPausado() ([]domain.Inmueble, error) //los inmuebles que estan publicados (los que pueden ver los clientes)
	ObtenerInmueblePorID(id string) (domain.Inmueble, error)
	BuscarInmuebleLocalidadYFechas(localidad, fechaIni, fechaFin string) ([]domain.Inmueble, error)
	CargarImagenInmueble(imagenName, idInmueble string) error
	GenerarPathImagen(file *multipart.FileHeader) (string, string, error)
	EliminarImagen(id string) (string, error)
	InmueblesPorEncargado(id uint) ([]domain.Inmueble, error)
	RegistrarCheckOut(id, estado string) error
}

type InmuebleRepository interface {
	Crear(inmueble domain.Inmueble) (domain.Inmueble, error)
	GetInmuebleId(id uint) (domain.Inmueble, error)
	GetInmuebles() ([]domain.Inmueble, error)
	GetInmuebleIdConReservasImagenes(id uint) (domain.Inmueble, error)
	GetInmueblesConFotoNoPausado() ([]domain.Inmueble, error)
	Actualizar(inmueble domain.Inmueble) (domain.Inmueble, error)
	ComprobarInmuebleExistente(nombre string) error
	Eliminar(id uint) error
	BuscarInmueblesPorLocalidadConFotoNoPausado(localidad string) ([]domain.Inmueble, error)
	GuardarImagenInmueble(imagen domain.ImagenInmueble) error
	GetImagen(id uint) (domain.ImagenInmueble, error)
	EliminarImagen(id uint) error
	InmueblesPorEncargado(id uint) ([]domain.Inmueble, error)
	RegistrarCheckOut(reserva domain.Reserva) error
}
