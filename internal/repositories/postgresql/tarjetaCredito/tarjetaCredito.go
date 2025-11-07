package tarjetaCredito

import (
	"AlquilerInmuebles/internal/domain"
	"errors"

	"gorm.io/gorm"
)

func (db *Repository) RegistrarTarjeta(tarjeta domain.Tarjeta) (domain.Tarjeta, error) {

	if err := db.DataBase.Create(&tarjeta).Error; err != nil {
		return tarjeta, err
	}

	return tarjeta, nil
}

func (db *Repository) GetTarjetaUsuario(id uint) (domain.Tarjeta, error) {
	var tarjeta domain.Tarjeta

	if err := db.DataBase.Where("id_usuario_tarjeta = ?", id).First(&tarjeta).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tarjeta, nil
		}
		return tarjeta, err
	}

	return tarjeta, nil
}

func (db *Repository) ActualizarTarjeta(tarjeta domain.Tarjeta) (domain.Tarjeta, error) {

	if err := db.DataBase.Save(&tarjeta).Error; err != nil {
		return tarjeta, err
	}

	return tarjeta, nil
}
