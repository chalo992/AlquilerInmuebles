package common

import (
	"AlquilerInmuebles/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToHTTPError(err error) (int, gin.H) {
	var e *domain.ErrorNegocio
	if errors.As(err, &e) {
		return e.HTTPStatus, gin.H{
			"estado":  "Error",
			"mensaje": e.Mensaje,
		}
	}

	return http.StatusInternalServerError, gin.H{
		"estado":  "Error",
		"mensaje": "Ocurrió un error inesperado",
	}
}
