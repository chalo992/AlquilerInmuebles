package inmueble

import (
	"AlquilerInmuebles/internal/domain"
)

func (db *Repository) Crear(inmueble domain.Inmueble) (domain.Inmueble, error) {

	if err := db.DataBase.Create(&inmueble).Error; err != nil {
		return inmueble, err
	}

	return inmueble, nil
}

func (db *Repository) GetInmuebleId(id uint) (domain.Inmueble, error) {
	var inmueble domain.Inmueble

	if err := db.DataBase.Preload("Imagenes").First(&inmueble, id).Error; err != nil {
		return inmueble, err
	}

	return inmueble, nil
}

func (db *Repository) Actualizar(inmueble domain.Inmueble) (domain.Inmueble, error) {

	if err := db.DataBase.Save(&inmueble).Error; err != nil {
		return inmueble, err
	}

	return inmueble, nil
}

func (db *Repository) ComprobarInmuebleExistente(nombre string) error {
	var inmuebleRegistrado domain.Inmueble
	if err := db.DataBase.Where("nombre = ?", nombre).First(&inmuebleRegistrado).Error; err == nil {
		return domain.ErrorInmuebleYaRegistrado()
	}

	return nil
}

func (db *Repository) Eliminar(id uint) error {

	if err := db.DataBase.Delete(&domain.Inmueble{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) GetInmuebleIdConReservasImagenes(id uint) (domain.Inmueble, error) {

	var inmueble domain.Inmueble

	if err := db.DataBase.Preload("Reservas").Preload("Imagenes").First(&inmueble, id).Error; err != nil {
		return inmueble, err
	}

	return inmueble, nil
}

func (db *Repository) GetInmuebles() ([]domain.Inmueble, error) {
	var inmuebles []domain.Inmueble

	if err := db.DataBase.Preload("Imagenes").Find(&inmuebles).Error; err != nil {
		return inmuebles, err
	}

	return inmuebles, nil
}

func (db *Repository) GetInmueblesConFotoNoPausado() ([]domain.Inmueble, error) {

	var inmuebles []domain.Inmueble

	if err := db.DataBase.Where("pausado = ?", false).Joins("JOIN imagen_inmuebles ON imagen_inmuebles.id_inmueble_imagen = inmuebles.id").
		Preload("Imagenes").Find(&inmuebles).Error; err != nil {
		return inmuebles, err
	}

	return inmuebles, nil
}

func (db *Repository) BuscarInmueblesPorLocalidadConFotoNoPausado(localidad string) ([]domain.Inmueble, error) {
	var inmuebles []domain.Inmueble

	// Hacemos un JOIN con imagen_inmuebles y filtramos por localidad y no pausado
	if err := db.DataBase.
		Table("inmuebles").
		Select("inmuebles.*").
		Joins("JOIN imagen_inmuebles ON imagen_inmuebles.id_inmueble_imagen = inmuebles.id").
		Where("localidad ILIKE ? AND pausado = ?", "%"+localidad+"%", false).
		Group("inmuebles.id").
		Having("COUNT(imagen_inmuebles.id) > 0").Preload("Reservas").
		Preload("Imagenes").
		Find(&inmuebles).Error; err != nil {
		return nil, err
	}

	if len(inmuebles) == 0 {
		return inmuebles, domain.ErrorInmuebleNoEncontradoPorLocalidadFechas()
	}

	return inmuebles, nil
}

func (db *Repository) GuardarImagenInmueble(imagen domain.ImagenInmueble) error {

	if err := db.DataBase.Create(&imagen).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) GetImagen(id uint) (domain.ImagenInmueble, error) {
	var imagen domain.ImagenInmueble

	if err := db.DataBase.First(&imagen, id).Error; err != nil {
		return imagen, err
	}

	return imagen, nil
}

func (db *Repository) EliminarImagen(id uint) error {

	if err := db.DataBase.Delete(domain.ImagenInmueble{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) InmueblesPorEncargado(id uint) ([]domain.Inmueble, error) {

	var inmuebles []domain.Inmueble

	if err := db.DataBase.Where("id_encargado = ?", id).Find(&inmuebles).Error; err != nil {
		return inmuebles, err
	}

	return inmuebles, nil
}

func (db *Repository) RegistrarCheckOut(reserva domain.Reserva) error {

	if err := db.DataBase.Save(&reserva).Error; err != nil {
		return err
	}

	return nil
}
