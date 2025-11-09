package calificacion

import (
	"AlquilerInmuebles/internal/domain"
	"strconv"
)

func (c *ServiceCalificacion) CalificarInmueble(calificacion domain.Calificacion, id uint) error {

	calificacion.IdUsuarioCalificacion = id

	err := c.Repo.CalificarInmueble(calificacion)
	if err != nil {
		return err
	}

	return nil
}

func (c *ServiceCalificacion) GetCalificacionInmueble(id string) (*domain.Calificacion, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	calificacion, err := c.Repo.GetCalificacion(uint(idInt))
	if err != nil {
		return nil, err
	}

	return &calificacion, nil

}

func (c *ServiceCalificacion) EliminarCalificacion(id string) error {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = c.Repo.EliminarCalificacion(uint(idInt))
	if err != nil {
		return err
	}

	return nil
}

func (c *ServiceCalificacion) EditarCalificacion(newCalificacion domain.Calificacion) error {

	calificacionActual, err := c.Repo.GetCalificacion(newCalificacion.ID)
	if err != nil {
		return err
	}

	calificacionActual.Comentario = newCalificacion.Comentario
	calificacionActual.Comodidad = newCalificacion.Comodidad
	calificacionActual.Distribucion = newCalificacion.Distribucion
	calificacionActual.Limpieza = newCalificacion.Limpieza
	calificacionActual.Ubicacion = newCalificacion.Ubicacion

	err = c.Repo.EditarCalificacion(calificacionActual)
	if err != nil {
		return err
	}

	return nil
}

func (c *ServiceCalificacion) ObtenerCalificacionesInmueble(id string) ([]domain.Calificacion, error) {

	var calificaciones []domain.Calificacion

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return calificaciones, err
	}

	calificaciones, err = c.Repo.ObtenerCalificacionesInmueble(uint(idInt))
	if err != nil {
		return calificaciones, err
	}

	if len(calificaciones) == 0 {
		return calificaciones, domain.ErrorNoHayCalificacionesInmueble()
	}

	return calificaciones, nil
}

func (c *ServiceCalificacion) ObtenerCalificacionReserva(id string) (*domain.Calificacion, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	calificacion, err := c.Repo.ObtenerCalificacionReservaId(uint(idInt))
	if err != nil {
		return nil, err
	}

	return calificacion, nil
}
