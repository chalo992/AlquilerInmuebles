package calificacion

import "AlquilerInmuebles/internal/domain"

func (db *Repository) CalificarInmueble(calificacion domain.Calificacion) error {

	if err := db.DataBase.Create(&calificacion).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) EliminarCalificacion(id uint) error {

	if err := db.DataBase.Delete(domain.Calificacion{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) EditarCalificacion(calificacion domain.Calificacion) error {

	if err := db.DataBase.Save(&calificacion).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) GetCalificacion(id uint) (domain.Calificacion, error) {
	var califiacion domain.Calificacion

	if err := db.DataBase.First(&califiacion, id).Error; err != nil {
		return califiacion, err
	}

	return califiacion, nil
}

func (db *Repository) ObtenerCalificacionesInmueble(id uint) ([]domain.Calificacion, error) {

	var calificaciones []domain.Calificacion

	if err := db.DataBase.Where("id_inmueble_calificacion = ?", id).Find(&calificaciones).Error; err != nil {
		return calificaciones, err
	}

	return calificaciones, nil
}

func (db *Repository) ObtenerCalificacionReservaId(id uint) (*domain.Calificacion, error) {

	var califiacion = new(domain.Calificacion)
	if err := db.DataBase.Where("id_reserva_calificacion = ?", id).First(&califiacion).Error; err != nil {
		return nil, err
	}

	return califiacion, nil
}
