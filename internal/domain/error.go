package domain

import "net/http"

type ErrorNegocio struct {
	Mensaje    string
	HTTPStatus int
}

func (e *ErrorNegocio) Error() string {
	return e.Mensaje
}

func ErrorEmailRegistrado() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El email ya está registrado",
		HTTPStatus: http.StatusConflict,
	}
}

func ErrorContraseñaNoCumple() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "La contraseña debe tener al menos 8 caracteres alfanumericos",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorEmailContraseñaIncorrecto() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El email o contraseña son incorrecto",
		HTTPStatus: http.StatusUnauthorized,
	}
}

func ErrorCrearToken() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "Error al crear el token",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorTarjetaCodigoSeguridad() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El código de seguridad debe tener 3 o 4 dígitos",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorTarjetaNumero() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El número de la tarjeta debe tener 15 o 16 dígitos",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorTarjetaInvalida() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "La tarjeta es inválida",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorTarjetaVencimiento() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "La tarjeta está vencida",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorNoCoincideContraseñas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "Las contraseñas no coinciden",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorContraseñaActual() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "la contraseña actual no coincide",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorContraseñaEncargado() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "La contraseña ya fue registrada",
		HTTPStatus: http.StatusUnauthorized,
	}
}

func ErrorInmuebleYaRegistrado() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El inmueble ya se encuentra registrado",
		HTTPStatus: http.StatusConflict,
	}
}

func ErrorInmuebleConReservas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El inmueble posee reservas activas o futuras, no se puede eliminar",
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

func ErrorInmuebleNoEncontradoPorLocalidadFechas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "No se encontraron resultados",
		HTTPStatus: http.StatusNotFound,
	}
}

func ErrorInmuebleFechasMalIngresadas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "Fechas de reservas mal ingresadas",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorEncargadoNoTieneInmuebles() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "No se tiene inmuebles a cargo hasta el momento",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorEncargadoNoTieneReservas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "No posee reservas",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorNoHayCalificacionesInmueble() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El inmueble no posee calificaciones",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorUsuarioNoTieneTarjeta() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "No hay una tarjeta registrada",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorInmuebleReservado() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "El inmueble ya fue reservado",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorUsuarioNoTieneReservas() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "No se tiene reservas hechas",
		HTTPStatus: http.StatusBadRequest,
	}
}

func ErrorReservaNoSeCancela() *ErrorNegocio {
	return &ErrorNegocio{
		Mensaje:    "La reserva solo se puede cancelar anterior a los dos días de comenzar dicha reserva",
		HTTPStatus: http.StatusBadRequest,
	}
}
