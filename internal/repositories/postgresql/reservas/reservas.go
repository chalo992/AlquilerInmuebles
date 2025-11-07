package reservas

import (
	"AlquilerInmuebles/internal/domain"

	"gorm.io/gorm"
)

func (db *Repository) ReservarInmueble(id uint) error {

	var user domain.Usuario

	if err := db.DataBase.Preload("Tarjeta").First(&user, id).Error; err != nil {
		return err
	}

	if user.Tarjeta == nil {
		return domain.ErrorUsuarioNoTieneTarjeta()
	}

	return nil
}

func (db *Repository) ConfirmarReserva(reserva domain.Reserva) error {

	err := db.DataBase.Transaction(func(tx *gorm.DB) error {
		// Crear la orden
		if err := tx.Create(&reserva).Error; err != nil {
			return err
		}

		// Crear los inquilinos asociados automáticamente
		if err := tx.Model(&reserva).Association("Inquilinos").Append(reserva.Inquilinos); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (db *Repository) ReservasDelUsuario(id uint) ([]domain.Reserva, error) {

	var reservas = new([]domain.Reserva)

	if err := db.DataBase.Where("id_usuario_reserva = ?", id).Find(&reservas).Error; err != nil {
		return nil, err
	}

	return *reservas, nil
}

func (db *Repository) GetReserva(id uint) (*domain.Reserva, error) {

	var reserva domain.Reserva

	if err := db.DataBase.First(&reserva, id).Error; err != nil {
		return nil, err
	}

	return &reserva, nil
}

func (db *Repository) ActualizarReserva(reserva *domain.Reserva) error {

	if err := db.DataBase.Save(&reserva).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) GetReservaId(id uint) (domain.Reserva, error) {

	var reserva domain.Reserva

	if err := db.DataBase.First(&reserva, id).Error; err != nil {
		return reserva, err
	}

	return reserva, nil

}

func (db *Repository) GetReservas() (*[]domain.Reserva, error) {

	var reservas []domain.Reserva

	if err := db.DataBase.Find(&reservas).Error; err != nil {
		return nil, err
	}

	return &reservas, nil
}

func (db *Repository) ListarReservasEncargado(id uint) ([]domain.Inmueble, error) {

	var inmuebles []domain.Inmueble

	if err := db.DataBase.Joins("JOIN reservas ON reservas.id_inmueble_reserva = inmuebles.id").
		Where("inmuebles.id_encargado = ?", id).Preload("Reservas").
		Find(&inmuebles).Error; err != nil {
		return inmuebles, err
	}

	return inmuebles, nil
}
