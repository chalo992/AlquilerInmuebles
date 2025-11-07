package credenciales

import (
	"AlquilerInmuebles/internal/ports"
)

type ServiceCredenciales struct {
	Repo     ports.CredencialesRepository
	UserRepo ports.UsuarioRepository
}
