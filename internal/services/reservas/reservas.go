package reservas

import (
	"AlquilerInmuebles/internal/domain"
	"AlquilerInmuebles/internal/services/common"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func (r *ServiceReserva) ReservarInmueble(id uint) error {

	err := r.Repo.ReservarInmueble(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceReserva) ConfirmarReserva(reservaJson domain.ReservaJson) error {

	fechaIni, fechaFin, err := common.TransformarFecha(reservaJson.FechaInicio, reservaJson.FechaFin)
	if err != nil {
		return err
	}

	idInmueble, err := strconv.Atoi(reservaJson.IdInmuebleReserva)
	if err != nil {
		return err
	}

	reserva := domain.Reserva{
		IdUsuarioReserva:  reservaJson.IdUsuarioReserva,
		FechaInicio:       fechaIni,
		FechaFin:          fechaFin,
		IdInmuebleReserva: uint(idInmueble),
		Inquilinos:        reservaJson.Inquilinos,
	}

	inmueble, err := r.InmRepo.GetInmuebleId(reserva.IdInmuebleReserva)
	if err != nil {
		return err
	}

	ok := common.VerificarReservaFechas(inmueble, reserva.FechaInicio, reserva.FechaFin)
	if !ok {
		return domain.ErrorInmuebleReservado()
	}

	err = r.Repo.ConfirmarReserva(reserva)
	if err != nil {
		return err
	}

	return nil
}

func (r *ServiceReserva) GenerarPathImagenInquilino(file *multipart.FileHeader) (string, string, error) {

	timestamp := time.Now().Format("20060102150405")
	imagenName := fmt.Sprintf("%s_%s", timestamp, file.Filename)

	saveDir := filepath.Join("path/carpeta/inquilinoImagenes", "imagenInquilinos")

	// Construir path completo donde se va a guardar
	savePath := filepath.Join(saveDir, imagenName)

	return savePath, imagenName, nil
}

func (r *ServiceReserva) ReservasDelUsuario(id uint) ([]domain.Reserva, error) {

	reservas, err := r.Repo.ReservasDelUsuario(id)
	if err != nil {
		return nil, err
	}

	if len(reservas) == 0 {
		return nil, domain.ErrorUsuarioNoTieneReservas()
	}

	return reservas, nil
}

func (r *ServiceReserva) CancelarReserva(id string) (*domain.Inmueble, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	reserva, err := r.Repo.GetReserva(uint(idInt))
	if err != nil {
		return nil, err
	}

	inicio := reserva.FechaInicio.Truncate(24 * time.Hour)
	hoy := time.Now().Truncate(24 * time.Hour)

	if inicio.Sub(hoy) < 48*time.Hour {
		return nil, domain.ErrorReservaNoSeCancela()
	}

	reserva.Activa = false

	err = r.Repo.ActualizarReserva(reserva)
	if err != nil {
		return nil, err
	}

	inmueble, err := r.InmRepo.GetInmuebleId(reserva.IdInmuebleReserva)
	if err != nil {
		return nil, err
	}

	return &inmueble, nil
}

func (r *ServiceReserva) GetReservaId(id string) (domain.Reserva, error) {

	var reserva domain.Reserva

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return reserva, err
	}

	reserva, err = r.Repo.GetReservaId(uint(idInt))
	if err != nil {
		return reserva, err
	}

	return reserva, nil
}

func (r *ServiceReserva) GetReservas() (*[]domain.Reserva, error) {

	reservas, err := r.Repo.GetReservas()
	if err != nil {
		return nil, err
	}

	return reservas, nil
}

func (s *ServiceReserva) ListarReservasEncargado(id uint) ([]domain.Inmueble, error) {

	inmuebles, err := s.Repo.ListarReservasEncargado(id)
	if err != nil {
		return inmuebles, err
	}

	if len(inmuebles) == 0 {
		return inmuebles, domain.ErrorEncargadoNoTieneReservas()
	}

	return inmuebles, nil
}
