package common

import "time"

func TransformarFecha(fechaini, fechaFin string) (time.Time, time.Time, error) {

	var fechaFinal, fechaInicial time.Time
	var err error

	fechaInicial, err = time.Parse("02/01/2006", fechaini)
	if err != nil {
		return fechaInicial, fechaFinal, err
	}

	fechaFinal, err = time.Parse("02/01/2006", fechaFin)
	if err != nil {
		return fechaInicial, fechaFinal, err
	}

	return fechaInicial, fechaFinal, nil

}
