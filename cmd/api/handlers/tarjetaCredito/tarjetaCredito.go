package tarjetaCredito

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *HandlerTarjeta) CrearTarjeta(c *gin.Context) {
	var tarjetaJSON domain.TarjetaJSON

	if err := c.BindJSON(&tarjetaJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON invalido",
		})
	}

	idStr, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	tarjeta, err := t.TarjetaService.CrearTarjeta(tarjetaJSON, idStr)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Tarjeta registrada correctamente",
		"data":    tarjeta,
	})
}

func (t *HandlerTarjeta) ActualizarTarjeta(c *gin.Context) {
	var tarjetaJSON domain.TarjetaJSON

	if err := c.BindJSON(&tarjetaJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON invalido",
		})
	}

	idStr, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	tarjeta, err := t.TarjetaService.ActTarjeta(tarjetaJSON, idStr)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Tarjeta actualizada correctamente",
		"data":    tarjeta.ID,
	})
}

func (t *HandlerTarjeta) GetTarjeta(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	tarjeta, err := t.TarjetaService.GetTarjeta(id)

	if err != nil {
		fmt.Println("entra tarjeta")
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Tarjeta encotrada correctamente",
		"data":    tarjeta,
	})
}
