package inmueble

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (i *InmuebleHandler) CargarInmueble(c *gin.Context) {

	var inmueble domain.Inmueble
	if err := c.BindJSON(&inmueble); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON inválido",
		})
		return
	}

	inmueble, err := i.InmuebleService.CrearInmueble(inmueble)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmueble cargado correctamente",
		"data":    inmueble,
	})
}

func (i *InmuebleHandler) PausarDespausarInmueble(c *gin.Context) {

	id := c.Param("id_inmueble")

	err := i.InmuebleService.PausarDespausarInmueble(id)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "inmueble pausado/despausado correctamente",
	})
}

func (i *InmuebleHandler) ActInmueble(c *gin.Context) {

	var inmueble domain.Inmueble
	if err := c.BindJSON(&inmueble); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON inválido",
		})
		return
	}

	inmuebleAct, err := i.InmuebleService.ActualizarInmueble(inmueble)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmueble actualziado correctamente",
		"data":    inmuebleAct,
	})
}

func (i *InmuebleHandler) BorrarInmueble(c *gin.Context) {

	id := c.Param("id_inmueble")

	paths, err := i.InmuebleService.EliminarInmueble(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	for _, path := range paths {
		os.Remove(path)
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmueble eliminado correctamente",
	})
}

func (i *InmuebleHandler) DevolverInmueblePorId(c *gin.Context) {

	id := c.Param("id_inmueble")

	inmueble, err := i.InmuebleService.ObtenerInmueblePorID(id)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmueble encontrado correctamente",
		"data":    inmueble,
	})
}

func (i *InmuebleHandler) DevolverInmuebles(c *gin.Context) {

	inmuebles, err := i.InmuebleService.ListarInmuebles()

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmuebles encontrados correctamente",
		"data":    inmuebles,
	})
}

func (i *InmuebleHandler) DevovlerInmueblesConFotoNoPausado(c *gin.Context) {

	inmuebles, err := i.InmuebleService.ListarInmueblesConFotoNoPausado()

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmuebles encontrados correctamente",
		"data":    inmuebles,
	})
}

func (i *InmuebleHandler) BuscarInmueblesPorLocalidadYFechas(c *gin.Context) {

	localidad := c.Param("localidad")
	fechaIni := c.Query("fechaInicio")
	fechaFin := c.Query("fechaFin")

	inmuebles, err := i.InmuebleService.BuscarInmuebleLocalidadYFechas(localidad, fechaIni, fechaFin)

	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmuebles encontrados en la localidad y fechas especificadas",
		"data":    inmuebles,
	})
}

func (i *InmuebleHandler) CargarInmuebleImagen(c *gin.Context) {

	idInmueble := c.PostForm("id_inmueble")
	file, err := c.FormFile("imagen")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "No se ingresó la imagen",
		})
		return
	}

	savePath, imagenName, err := i.InmuebleService.GenerarPathImagen(file)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"mensaje": "Error guardando archivo"})
		return
	}

	err = i.InmuebleService.CargarImagenInmueble(imagenName, idInmueble)
	if err != nil {
		os.Remove(savePath)
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Imagen cargada correctamente",
	})
}

func (i *InmuebleHandler) EliminarImagenInmueble(c *gin.Context) {

	id := c.Param("id_inmueble")

	err := i.InmuebleService.EliminarImagen(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Imagen eliminada correctamente",
	})
}

func (i *InmuebleHandler) GetInmueblesEncargados(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	inmuebles, err := i.InmuebleService.InmueblesPorEncargado(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Inmuebles encontrados correctamente para el encargado",
		"data":    inmuebles,
	})
}

func (i *InmuebleHandler) RegistrarCheckOutEncargado(c *gin.Context) {

	idReserva := c.Param("id_reserva")
	var estado domain.Estado
	if err := c.BindJSON(&estado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "JSON mal formado",
		})
		return
	}

	err := i.InmuebleService.RegistrarCheckOut(idReserva, estado.Estado)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Check-Out registrado correctamente",
	})
}
