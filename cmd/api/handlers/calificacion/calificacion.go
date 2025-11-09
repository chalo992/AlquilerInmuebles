package calificacion

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CalificacionHandler) CalificarInmueble(c *gin.Context) {

	var calificacion domain.Calificacion
	if err := c.BindJSON(&calificacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON inválido",
		})
	}

	id, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	err := h.CalificacionService.CalificarInmueble(calificacion, id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Califiacion registrada correctamente",
	})
}

func (h *CalificacionHandler) GetCalificacionInmueble(c *gin.Context) {

	id := c.Param("id_calificacion")

	calificacion, err := h.CalificacionService.GetCalificacionInmueble(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Calificación obtenida exitosamente",
		"data":    calificacion,
	})
}

func (h *CalificacionHandler) EliminarCalificacion(c *gin.Context) {

	idCalificacion := c.Param("id_calificacion")

	err := h.CalificacionService.EliminarCalificacion(idCalificacion)

	if err != nil {

		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Califiacion eliminada correctamente",
	})
}

func (h *CalificacionHandler) EditarCalificacion(c *gin.Context) {

	var califiacion domain.Calificacion

	if err := c.BindJSON(&califiacion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON inválido",
		})
	}

	err := h.CalificacionService.EditarCalificacion(califiacion)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Calificacion editada correctamente",
	})
}

func (h *CalificacionHandler) ObtenerCalificacionesInmueble(c *gin.Context) {

	id := c.Param("id_inmueble")

	calificaciones, err := h.CalificacionService.ObtenerCalificacionesInmueble(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "calificaciones del inmueble obtenidas correctamente",
		"data":    calificaciones,
	})
}

func (h *CalificacionHandler) GetCalificacionReserva(c *gin.Context) {

	id := c.Param("id_reserva")

	calificacion, err := h.CalificacionService.ObtenerCalificacionReserva(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Calificacion de la reserva encontrada correctamente",
		"data":    calificacion,
	})
}
