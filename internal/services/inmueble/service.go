package inmueble

import "AlquilerInmuebles/internal/ports"

type ServiceInmueble struct {
	Repo    ports.InmuebleRepository
	ResRepo ports.ReservaRepository
}
