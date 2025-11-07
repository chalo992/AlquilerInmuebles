package common

import (
	"AlquilerInmuebles/internal/domain"
	"time"
)

func VerificarReservaFechas(inmueble domain.Inmueble, fechaInicial, fechaFinal time.Time) bool {

	diferencia := fechaFinal.Sub(fechaInicial)

	if int(diferencia.Hours()/24) < inmueble.DiasMinimosAlquiler {
		return false
	}

	for _, reserva := range inmueble.Reservas {
		if !reserva.Activa {
			continue
		}

		if fechaInicial.Before(reserva.FechaFin) && fechaFinal.After(reserva.FechaInicio) {
			return false
		}
	}
	return true
}
