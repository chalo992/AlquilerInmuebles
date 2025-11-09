package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/calificacion"
	"AlquilerInmuebles/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func CalificacionRutas(r *gin.Engine, c *calificacion.CalificacionHandler) {

	api := r.Group("/api")
	calificacion := api.Group("/calificacion")

	calificacion.Use(middleware.Autorizacion)
	calificacion.POST("/calificarInmueble", c.CalificarInmueble)
	calificacion.GET("/obtenerCalificacionInmueble/:id_calificacion", c.GetCalificacionInmueble)
	calificacion.DELETE("/eliminarCalificacion/:id_calificacion", c.EliminarCalificacion)
	calificacion.PUT("/editarCalificacion", c.EditarCalificacion)
	calificacion.GET("/calificacionesInmueble/:id_inmueble", c.ObtenerCalificacionesInmueble)
	calificacion.GET("/calificacionReserva/:id_reserva", c.GetCalificacionReserva)
}
