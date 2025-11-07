package usuario

import (
	"AlquilerInmuebles/internal/domain"
)

func (db *Repository) Crear(usuario domain.Usuario) (domain.Usuario, error) {

	if err := db.DataBase.Create(&usuario).Error; err != nil {
		return usuario, err
	}
	return usuario, nil
}

func (db *Repository) Actualizar(usuario domain.Usuario) error {

	if err := db.DataBase.Save(&usuario).Error; err != nil {
		return err
	}

	return nil
}

func (db *Repository) VerificarEmail(email string) error {
	var usuarioExiste domain.Usuario
	if err := db.DataBase.Where("email = ?", email).First(&usuarioExiste).Error; err == nil {
		return domain.ErrorEmailRegistrado()
	}
	return nil
}

func (db *Repository) GetUsuarioID(idParam uint) (domain.Usuario, error) {
	var usuario domain.Usuario

	if err := db.DataBase.First(&usuario, idParam).Error; err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (db *Repository) DevolverEncargados() ([]domain.Usuario, error) {
	var encargados []domain.Usuario

	if err := db.DataBase.Where("rol = ?", "encargado").Find(&encargados).Error; err != nil {
		return encargados, err
	}

	return encargados, nil
}

func (db *Repository) GetInquilinosReserva(id_reserva uint) ([]domain.Inquilino, error) {

	var inquilinos []domain.Inquilino

	if err := db.DataBase.Where("id_reserva_inquilino = ?", id_reserva).Find(&inquilinos).Error; err != nil {
		return nil, err
	}

	return inquilinos, nil
}
