package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/inmueble"
	"AlquilerInmuebles/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func InmuebleRutas(r *gin.Engine, i *inmueble.InmuebleHandler) {
	api := r.Group("/api")
	inmueble := api.Group("/inmueble")

	inmueble.Use(middleware.Autorizacion)
	inmueble.POST("/cargarInmueble", middleware.AdminAutorizacion, i.CargarInmueble)
	inmueble.PUT("/actualizarInmueble", middleware.AdminAutorizacion, i.ActInmueble)
	inmueble.DELETE("/eliminarInmueble/:id_inmueble", middleware.AdminAutorizacion, i.BorrarInmueble)
	inmueble.PATCH("/pausarDespausar/:id_inmueble", middleware.EncargadoAdminAutorizacion, i.PausarDespausarInmueble)
	inmueble.GET("/devolverInmuebleID/:id_inmueble", i.DevolverInmueblePorId)
	inmueble.GET("/devolverTodosInmuebles", middleware.AdminAutorizacion, i.DevolverInmuebles)
	inmueble.GET("/inmublesConFotoNoPausado", i.DevovlerInmueblesConFotoNoPausado)
	inmueble.GET("/buscarInmueble/:localidad", i.BuscarInmueblesPorLocalidadYFechas)
	inmueble.POST("/cargarImagenInmueble/:id_inmueble", middleware.AdminAutorizacion, i.CargarInmuebleImagen)
	inmueble.DELETE("/eliminarImagenInmueble/:id_inmueble", middleware.AdminAutorizacion, i.EliminarImagenInmueble)
	inmueble.GET("/inmueblesEncargado", middleware.EncargadoAutorizacion, i.GetInmueblesEncargados)
	inmueble.PATCH("/registrarCheckOut/:id_reserva", middleware.EncargadoAutorizacion, i.RegistrarCheckOutEncargado)

}
