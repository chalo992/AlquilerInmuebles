package reservas

import "AlquilerInmuebles/internal/ports"

type ServiceReserva struct {
	Repo    ports.ReservaRepository
	InmRepo ports.InmuebleRepository
}
