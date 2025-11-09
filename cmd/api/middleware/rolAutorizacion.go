package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func obtenerRol(c *gin.Context) (string, error) {
	rol, exists := c.Get("rol")
	if !exists {
		return "", fmt.Errorf("Usuario no autorizado: rol no encontrado en el contexto")
	}

	rolUser, ok := rol.(string)
	if !ok {
		return "", fmt.Errorf("Usuario no autorizado: formato de rol inválido")
	}

	return rolUser, nil
}

func AdminAutorizacion(c *gin.Context) {

	rolUser, err := obtenerRol(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": err,
		})
		c.Abort()
		return
	}
	if rolUser != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Acceso denegado: se requiere rol de administrador",
		})
		c.Abort()
		return
	}

	c.Next()
}

func EncargadoAutorizacion(c *gin.Context) {

	rolUser, err := obtenerRol(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": err,
		})
		c.Abort()
		return
	}
	if rolUser != "encargado" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Acceso denegado: se requiere rol de encargado",
		})
		c.Abort()
		return
	}

	c.Next()
}

func EncargadoAdminAutorizacion(c *gin.Context) {
	rolUser, err := obtenerRol(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": err,
		})
		c.Abort()
		return
	}
	if rolUser != "encargado" && rolUser != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Acceso denegado: se requiere rol de encargado o admin",
		})
		c.Abort()
		return
	}

	c.Next()
}
