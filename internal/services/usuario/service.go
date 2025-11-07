package usuario

import (
	"AlquilerInmuebles/internal/ports"
)

type Service struct {
	Repo ports.UsuarioRepository
}
