package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/usuario"
	"AlquilerInmuebles/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func UsuarioRutas(r *gin.Engine, u *usuario.HandlerUsuario) {
	api := r.Group("/api")
	usuario := api.Group("usuario")

	usuario.POST("/registrarCliente", u.CrearUsuarioCliente)
	usuario.POST("/registrarAdmin", u.CrearUsuarioAdmin)
	usuario.POST("/registrarEncargado", middleware.Autorizacion, middleware.AdminAutorizacion, u.CrearUsuarioEncargado)
	usuario.GET("/usuarioID", middleware.Autorizacion, u.GetUsuarioByID)
	usuario.PATCH("/encargadoContra/:id", u.ContraseñaEncargado)
	usuario.GET("/encargados", middleware.Autorizacion, middleware.AdminAutorizacion, u.RetornarEncargados)
	usuario.PATCH("/modificarUsuario", middleware.Autorizacion, u.ActUsuario)
	usuario.GET("/usuarioParam/:id_usuario", middleware.Autorizacion, middleware.AdminAutorizacion, u.GetUsuarioIDParam)
	usuario.GET("/inquilinos/:id_reserva", middleware.Autorizacion, u.DevolverInquilinos)
}
