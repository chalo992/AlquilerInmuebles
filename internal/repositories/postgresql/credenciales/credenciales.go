package credenciales

import "AlquilerInmuebles/internal/domain"

func (r *Repository) ComprobarCredencialesEmail(email string) (usuario domain.Usuario, err error) {
	var user domain.Usuario

	if err := r.DataBase.Where("email = ?", email).First(&user).Error; err != nil {
		return user, domain.ErrorEmailContraseñaIncorrecto()
	}

	return user, nil
}
