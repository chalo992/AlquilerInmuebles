package ports

import "AlquilerInmuebles/internal/domain"

type CredencialesService interface {
	ComprobarCredenciales(credenciales domain.Login) (usuario domain.Usuario, err error)
	GenerarToken(usuario domain.Usuario) (string, error)
	CambiarContraseña(id uint, contraseña domain.CambiarContraseñaUser) error
}

type CredencialesRepository interface {
	ComprobarCredencialesEmail(email string) (usuario domain.Usuario, err error)
}
