package ports

import "AlquilerInmuebles/internal/domain"

type UsuarioService interface {
	CrearUsuario(usuario domain.Usuario, rol string) (domain.Usuario, error)
	ActualizarUsuario(usuario domain.Usuario, idParam uint) (domain.Usuario, error)
	GetUsuarioByID(id uint) (domain.Usuario, error)
	GetUsuarioIDParam(id string) (domain.Usuario, error)
	GetEncargados() ([]domain.Usuario, error)
	MandarMailEncargado(encargado domain.Usuario) error
	EstablecerContraseñaEncargado(contraseña domain.CambiarContraseña, idEncargado string) error
	GetInquilinosReserva(id_reserva string) ([]domain.Inquilino, error)
}

type UsuarioRepository interface {
	Crear(usuario domain.Usuario) (domain.Usuario, error)
	VerificarEmail(email string) error
	GetUsuarioID(idParam uint) (domain.Usuario, error)
	Actualizar(usuario domain.Usuario) error
	DevolverEncargados() ([]domain.Usuario, error)
	GetInquilinosReserva(id_reserva uint) ([]domain.Inquilino, error)
}
