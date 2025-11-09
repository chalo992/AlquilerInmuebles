package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/tarjetaCredito"
	"AlquilerInmuebles/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func TarjetaRutas(r *gin.Engine, t *tarjetaCredito.HandlerTarjeta) {
	api := r.Group("/api")
	tarjeta := api.Group("/tarjeta")

	tarjeta.Use(middleware.Autorizacion)
	tarjeta.POST("/registrarTarjeta", t.CrearTarjeta)
	tarjeta.PUT("/actualizarTarjeta", t.ActualizarTarjeta)
	tarjeta.GET("/getTarjetaUser", t.GetTarjeta)

}
