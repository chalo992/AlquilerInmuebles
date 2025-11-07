package tarjetaCredito

import (
	"AlquilerInmuebles/internal/domain"
	"AlquilerInmuebles/internal/services/common"
)

func (t *ServiceTarjeta) CrearTarjeta(tarjetaJSON domain.TarjetaJSON, id uint) (domain.Tarjeta, error) {

	var tarjeta domain.Tarjeta

	err, fecha := common.ValidarTarjeta(tarjetaJSON)

	if err != nil {
		return tarjeta, err
	}

	tarjeta.CodigoSeguridad = tarjetaJSON.CodigoSeguridad
	tarjeta.Numero = tarjetaJSON.Numero
	tarjeta.Vencimiento = fecha
	tarjeta.IdUsuarioTarjeta = id

	tarjeta, err = t.Repo.RegistrarTarjeta(tarjeta)

	if err != nil {
		return tarjeta, err
	}

	return tarjeta, nil
}

func (t *ServiceTarjeta) ActTarjeta(tarjetaJSON domain.TarjetaJSON, idUsuario uint) (domain.Tarjeta, error) {

	var tarjeta domain.Tarjeta

	err, fecha := common.ValidarTarjeta(tarjetaJSON)
	if err != nil {
		return tarjeta, err
	}

	tarjeta, err = t.Repo.GetTarjetaUsuario(idUsuario)
	if err != nil {
		return tarjeta, err
	}

	tarjeta.CodigoSeguridad = tarjetaJSON.CodigoSeguridad
	tarjeta.Numero = tarjetaJSON.Numero
	tarjeta.Vencimiento = fecha

	tarjeta, err = t.Repo.ActualizarTarjeta(tarjeta)

	if err != nil {
		return tarjeta, err
	}

	return tarjeta, nil

}

func (t *ServiceTarjeta) GetTarjeta(idUsuario uint) (domain.Tarjeta, error) {

	tarjeta, err := t.Repo.GetTarjetaUsuario(idUsuario)
	if err != nil {
		return tarjeta, err
	}

	return tarjeta, nil
}
