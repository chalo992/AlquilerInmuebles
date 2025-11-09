package usuario

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerUsuario) CrearUsuarioCliente(c *gin.Context) {
	h.crearUsuarioPorTipo(c, "cliente")
}

func (h *HandlerUsuario) CrearUsuarioAdmin(c *gin.Context) {
	h.crearUsuarioPorTipo(c, "admin")
}

func (h *HandlerUsuario) CrearUsuarioEncargado(c *gin.Context) {
	h.crearUsuarioPorTipo(c, "encargado")
}

func (h *HandlerUsuario) crearUsuarioPorTipo(c *gin.Context, tipo string) {
	var usuarioParams domain.Usuario

	if err := c.BindJSON(&usuarioParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"estado":  "Error",
			"mensaje": "JSON inválido",
		})
		return
	}

	usuario, err := h.UsuarioService.CrearUsuario(usuarioParams, tipo)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	if tipo == "encargado" {
		if err := h.UsuarioService.MandarMailEncargado(usuario); err != nil {
			status, body := common.ToHTTPError(err)
			c.JSON(status, body)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"estado":  "completado",
		"data":    usuario,
		"mensaje": "Usuario creado correctamente",
	})
}

func (h *HandlerUsuario) GetUsuarioByID(c *gin.Context) {
	id, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	usuario, err := h.UsuarioService.GetUsuarioByID(id)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Usuario encontrado correctamente",
		"data":    usuario,
	})
}

func (h *HandlerUsuario) ActUsuario(c *gin.Context) {
	var usuarioActualizado domain.Usuario

	if err := c.BindJSON(&usuarioActualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON invalido",
		})
		return
	}

	idStr, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	usuario, err := h.UsuarioService.ActualizarUsuario(usuarioActualizado, idStr)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    usuario,
		"mensaje": "Usuario actualizado correctamente",
	})
}

func (h *HandlerUsuario) RetornarEncargados(c *gin.Context) {

	encargados, err := h.UsuarioService.GetEncargados()

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "encargados encontrados correctamente",
		"data":    encargados,
	})
}

func (h *HandlerUsuario) ContraseñaEncargado(c *gin.Context) {

	var contraseña domain.CambiarContraseña
	id := c.Param("id")

	if err := c.BindJSON(&contraseña); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON invalido",
		})
		return
	}

	err := h.UsuarioService.EstablecerContraseñaEncargado(contraseña, id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Contraseña registrada correctamente",
	})
}

func (h *HandlerUsuario) GetUsuarioIDParam(c *gin.Context) {

	id := c.Param("id_usuario")
	usuario, err := h.UsuarioService.GetUsuarioIDParam(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Usuario encontrado correctamente",
		"data":    usuario,
	})
}

func (h *HandlerUsuario) DevolverInquilinos(c *gin.Context) {

	id_reserva := c.Param("id_reserva")
	inquilinos, err := h.UsuarioService.GetInquilinosReserva(id_reserva)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inquilinos encontrados correstamente",
		"data":    inquilinos,
	})
}
