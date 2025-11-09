package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/reservas"
	"AlquilerInmuebles/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func ReservasRutas(r *gin.Engine, i *reservas.ReservaHandler) {
	api := r.Group("/api")
	reserva := api.Group("/reserva")

	reserva.Use(middleware.Autorizacion)
	reserva.POST("/", i.ReservarInmuebleTarjeta)
	reserva.POST("/confirmarReserva", i.ConfirmarReserva)
	reserva.GET("/reservasUsuario", i.ListarReservasDelUsuario)
	reserva.PATCH("/cancelarReserva/:id_reserva", i.CancelarReserva)
	reserva.GET("/obtenerReserva/:id_reserva", i.ObtenerReserva)
	reserva.GET("/obtenerReservasTotales", middleware.EncargadoAdminAutorizacion, i.GetReservasTotales)
	reserva.GET("/reservasEncargado", middleware.EncargadoAutorizacion, i.ListarReservasEncargado)
}
