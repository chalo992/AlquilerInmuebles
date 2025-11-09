package credenciales

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *CredencialesHanlder) Login(c *gin.Context) {
	var credenciales domain.Login

	if err := c.BindJSON(&credenciales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON mal formado o inválido",
		})
		return
	}

	usuario, err := h.CredencialesService.ComprobarCredenciales(credenciales)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	token, err := h.CredencialesService.GenerarToken(usuario)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Usuario logueado correctamente",
		"token":   token,
		"id":      usuario.ID,
		"rol":     usuario.Rol,
	})
}

func (h *CredencialesHanlder) CambiarContraseña(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	var contraseña domain.CambiarContraseñaUser
	if err := c.BindJSON(&contraseña); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "JSON mal formado",
		})
		return
	}

	err := h.CredencialesService.CambiarContraseña(id, contraseña)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Contraseña cambiada exitosamente",
	})
}
