package routes

import (
	"AlquilerInmuebles/cmd/api/handlers/credenciales"

	"github.com/gin-gonic/gin"
)

func CredencialesRutas(r *gin.Engine, c *credenciales.CredencialesHanlder) {
	api := r.Group("/api")

	login := api.Group("login")
	login.POST("/", c.Login)

	cambiar := api.Group("/cambiar")
	cambiar.PATCH("/contrasena", c.CambiarContraseña)

}
