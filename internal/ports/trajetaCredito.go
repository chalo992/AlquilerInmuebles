package ports

import "AlquilerInmuebles/internal/domain"

type TarjetaService interface {
	CrearTarjeta(tarjeta domain.TarjetaJSON, id uint) (domain.Tarjeta, error)
	ActTarjeta(tarjeta domain.TarjetaJSON, idStr uint) (domain.Tarjeta, error)
	GetTarjeta(id uint) (domain.Tarjeta, error)
}

type TarjetaRepository interface {
	RegistrarTarjeta(tarjeta domain.Tarjeta) (domain.Tarjeta, error)
	ActualizarTarjeta(tarjeta domain.Tarjeta) (domain.Tarjeta, error)
	GetTarjetaUsuario(id uint) (domain.Tarjeta, error)
}
