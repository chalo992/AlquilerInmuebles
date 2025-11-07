package common

import (
	"AlquilerInmuebles/internal/domain"
	"strconv"
	"time"
)

func ValidarTarjeta(tarjetaJSON domain.TarjetaJSON) (error, time.Time) {

	fecha, err := time.Parse("01/06", tarjetaJSON.Vencimiento)

	if err != nil {
		return err, fecha
	}

	if fecha.Before(time.Now()) {
		return domain.ErrorTarjetaVencimiento(), fecha
	}

	numeroStr := strconv.Itoa(tarjetaJSON.Numero)
	codigoStr := strconv.Itoa(tarjetaJSON.CodigoSeguridad)

	if len(numeroStr) != 15 && len(numeroStr) != 16 {
		return domain.ErrorTarjetaNumero(), fecha
	}

	if len(codigoStr) != 3 && len(codigoStr) != 4 {
		return domain.ErrorTarjetaCodigoSeguridad(), fecha
	}

	if len(codigoStr)+len(numeroStr) != 19 {
		return domain.ErrorTarjetaInvalida(), fecha
	}

	return nil, fecha

}
