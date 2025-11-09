package reservas

import (
	"AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (r *ReservaHandler) ReservarInmuebleTarjeta(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
	}

	err := r.ReservaService.ReservarInmueble(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "El usuario posee tarjeta",
	})
}

func (r *ReservaHandler) ConfirmarReserva(c *gin.Context) {

	usuarioId, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "Usuario no autorizado"})
		return
	}

	//  Leer formulario multipart
	_, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "No se pudo leer el formulario"})
		return
	}

	//  Obtener campos
	idInmueble := c.PostForm("id_inmueble_reserva")
	fechaInicio := c.PostForm("fecha_inicio")
	fechaFin := c.PostForm("fecha_fin")

	if idInmueble == "" || fechaInicio == "" || fechaFin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "Faltan datos de la reserva"})
		return
	}

	//  Leer inquilinos dinámicamente
	inquilinos := []domain.Inquilino{}
	for i := 0; ; i++ {
		nombre := c.PostForm(fmt.Sprintf("inquilinos[%d][nombre_completo]", i))
		dni := c.PostForm(fmt.Sprintf("inquilinos[%d][dni]", i))
		if nombre == "" || dni == "" {
			break // ya no hay más inquilinos
		}

		inquilino := domain.Inquilino{
			NombreCompleto: nombre,
			DNI:            dni,
		}

		//  Procesar imagen
		file, err := c.FormFile(fmt.Sprintf("imagen_inquilino_%d", i))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"mensaje": fmt.Sprintf("No se encontró la imagen del inquilino %d", i+1),
			})
			return
		}

		savePath, imagenName, err := r.ReservaService.GenerarPathImagenInquilino(file)
		if err != nil {
			status, body := common.ToHTTPError(err)
			c.JSON(status, body)
			return
		}

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"mensaje": "Error guardando archivo"})
			return
		}

		// 🔗 Asociar la URL accesible de la imagen
		inquilino.ImagenUrl = fmt.Sprintf("/static/imagenInquilinos/%s", imagenName)

		inquilinos = append(inquilinos, inquilino)
	}

	if len(inquilinos) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"mensaje": "Debe ingresar al menos un inquilino"})
		return
	}

	//  Construir objeto reserva
	reservaJson := domain.ReservaJson{
		IdUsuarioReserva:  usuarioId,
		IdInmuebleReserva: idInmueble,
		FechaInicio:       fechaInicio,
		FechaFin:          fechaFin,
		Inquilinos:        inquilinos,
	}

	if err := r.ReservaService.ConfirmarReserva(reservaJson); err != nil {
		for _, inq := range inquilinos {
			os.Remove(filepath.Join("imagenInquilinos", filepath.Base(inq.ImagenUrl)))
		}
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "Reserva confirmada correctamente"})
}

func (r *ReservaHandler) ListarReservasDelUsuario(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	reservas, err := r.ReservaService.ReservasDelUsuario(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Reservas del usuario encontradas correctamente",
		"data":    reservas,
	})
}

func (r *ReservaHandler) CancelarReserva(c *gin.Context) {

	id := c.Param("id_reserva")

	inmueble, err := r.ReservaService.CancelarReserva(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	switch inmueble.PoliticaDevolucion {
	case "Devolucion 50%":
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Reserva cancelada con éxito. Se devolverá el 50% del monto pagado.",
		})
	case "Devolucion 100%":
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Reserva cancelada con éxito. Se devolverá el 100% del monto pagado.",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Reserva cancelada con éxito.",
		})
	}

}

func (i *ReservaHandler) ObtenerReserva(c *gin.Context) {

	idReserva := c.Param("id_reserva")

	reserva, err := i.ReservaService.GetReservaId(idReserva)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Reserva encotrada correctamente",
		"data":    reserva,
	})
}

func (i *ReservaHandler) GetReservasTotales(c *gin.Context) {

	reservas, err := i.ReservaService.GetReservas()
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Reservas obtenidas correctamente",
		"data":    reservas,
	})
}

func (i *ReservaHandler) ListarReservasEncargado(c *gin.Context) {

	id, ok := common.ObtenerIdUsuarioClaims(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensaje": "Usuario no autorizado o formato de id inválido",
		})
		return
	}

	inmuebles, err := i.ReservaService.ListarReservasEncargado(id)
	if err != nil {
		status, body := common.ToHTTPError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Reservas encontradas correctamente",
		"data":    inmuebles,
	})
}
