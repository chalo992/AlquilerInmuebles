package inmueble

import (
	"AlquilerInmuebles/internal/domain"
	"AlquilerInmuebles/internal/services"
	"AlquilerInmuebles/internal/services/common"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func (i *ServiceInmueble) CrearInmueble(inmueble domain.Inmueble) (domain.Inmueble, error) {

	err := i.Repo.ComprobarInmuebleExistente(inmueble.Nombre)
	if err != nil {
		return inmueble, err
	}

	inmueble, err = i.Repo.Crear(inmueble)

	if err != nil {
		return inmueble, err
	}

	return inmueble, nil
}

func (i *ServiceInmueble) PausarDespausarInmueble(id string) error {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	inmueble, err := i.Repo.GetInmuebleId(uint(idInt))
	if err != nil {
		return err
	}

	inmueble.Pausado = !inmueble.Pausado

	inmueble, err = i.Repo.Actualizar(inmueble)
	if err != nil {
		return err
	}

	return nil
}

func (i *ServiceInmueble) ActualizarInmueble(inmuebleActualizado domain.Inmueble) (domain.Inmueble, error) {

	inmueble, err := i.Repo.GetInmuebleId(inmuebleActualizado.ID)
	if err != nil {
		return inmuebleActualizado, err
	}

	if inmueble.Nombre != inmuebleActualizado.Nombre {
		err = i.Repo.ComprobarInmuebleExistente(inmuebleActualizado.Nombre)
		if err != nil {
			return inmuebleActualizado, err
		}
	}

	inmueble.Camas = inmuebleActualizado.Camas
	inmueble.DiasMinimosAlquiler = inmuebleActualizado.DiasMinimosAlquiler
	inmueble.Direccion = inmuebleActualizado.Direccion
	inmueble.Habitaciones = inmuebleActualizado.Habitaciones
	inmueble.Localidad = inmuebleActualizado.Localidad
	inmueble.MaxInquilinos = inmuebleActualizado.MaxInquilinos
	inmueble.Nombre = inmuebleActualizado.Nombre
	inmueble.PoliticaDevolucion = inmuebleActualizado.PoliticaDevolucion
	inmueble.PoliticaPago = inmuebleActualizado.PoliticaPago
	inmueble.PrecioDia = inmuebleActualizado.PrecioDia
	inmueble.TieneAire = inmuebleActualizado.TieneAire
	inmueble.TieneCable = inmuebleActualizado.TieneCable
	inmueble.TieneGasNatural = inmuebleActualizado.TieneGasNatural
	inmueble.TieneInternet = inmuebleActualizado.TieneInternet
	inmueble.TienePileta = inmuebleActualizado.TienePileta
	inmueble.IdEncargado = inmuebleActualizado.IdEncargado

	inmueble, err = i.Repo.Actualizar(inmueble)
	if err != nil {
		return inmuebleActualizado, err
	}

	return inmueble, err
}

func (i *ServiceInmueble) EliminarInmueble(id string) ([]string, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	inmueble, err := i.Repo.GetInmuebleIdConReservasImagenes(uint(idInt))

	for _, reserva := range inmueble.Reservas {
		if reserva.Activa {
			if reserva.FechaInicio.After(time.Now()) || reserva.FechaFin.After(time.Now()) {
				return nil, domain.ErrorInmuebleConReservas()
			}
		}
	}

	var paths []string
	for _, imagen := range inmueble.Imagenes {
		paths = append(paths, fmt.Sprintf("path/carpeta/inmuebleImagenes\\%s", imagen.PathLocal))
	} //el string debe ser el path a la carpeta donde se guardan localmente las imagenes de los inmuebles

	err = i.Repo.Eliminar(uint(idInt))
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (i *ServiceInmueble) ObtenerInmueblePorID(id string) (domain.Inmueble, error) {

	var inmueble domain.Inmueble

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return inmueble, err
	}

	inmueble, err = i.Repo.GetInmuebleId(uint(idInt))

	if err != nil {
		return inmueble, err
	}

	for idx := range inmueble.Imagenes {
		inmueble.Imagenes[idx].Url = services.BaseUrl + inmueble.Imagenes[idx].Url
	}

	return inmueble, nil
}

func (i *ServiceInmueble) ListarInmuebles() ([]domain.Inmueble, error) {

	inmuebles, err := i.Repo.GetInmuebles()
	if err != nil {
		return inmuebles, err
	}

	return inmuebles, nil
}

func (i *ServiceInmueble) ListarInmueblesConFotoNoPausado() ([]domain.Inmueble, error) {

	inmuebles, err := i.Repo.GetInmueblesConFotoNoPausado()

	if err != nil {
		return inmuebles, err
	}

	for _, inmueble := range inmuebles {
		for idx := range inmueble.Imagenes {
			inmueble.Imagenes[idx].Url = services.BaseUrl + inmueble.Imagenes[idx].Url
		}
	}

	return inmuebles, err
}

func (i *ServiceInmueble) BuscarInmuebleLocalidadYFechas(localidad, fechaini, fechaFin string) ([]domain.Inmueble, error) {

	var inmuebles []domain.Inmueble
	var inmueblesLista []domain.Inmueble

	fechaInicial, fechaFinal, err := common.TransformarFecha(fechaini, fechaFin)
	if err != nil {
		return inmueblesLista, err
	}

	if fechaFinal.Before(fechaInicial) {
		return inmueblesLista, domain.ErrorInmuebleFechasMalIngresadas()
	}

	inmuebles, err = i.Repo.BuscarInmueblesPorLocalidadConFotoNoPausado(localidad)
	if err != nil {
		return inmueblesLista, err
	}

	for _, inmueble := range inmuebles {
		if common.VerificarReservaFechas(inmueble, fechaInicial, fechaFinal) {
			inmueblesLista = append(inmueblesLista, inmueble)
		}
	}

	if len(inmueblesLista) == 0 {
		return inmueblesLista, domain.ErrorInmuebleNoEncontradoPorLocalidadFechas()
	}

	for _, inmueble := range inmueblesLista {
		for idx := range inmueble.Imagenes {
			// Si en BD tenés "/static/imagenesInmuebles/xxxx.jpg"
			// entonces concatenás el host
			inmueble.Imagenes[idx].Url = services.BaseUrl + inmueble.Imagenes[idx].Url
		}
	}

	return inmueblesLista, nil
}

func (i *ServiceInmueble) GenerarPathImagen(file *multipart.FileHeader) (string, string, error) {
	timestamp := time.Now().Format("20060102150405")
	imagenName := fmt.Sprintf("%s_%s", timestamp, file.Filename)

	// Ruta absoluta hacia la carpeta go/imagenesInmuebles
	saveDir := filepath.Join("path/carpeta/inmuebleImagenes", "imagenesInmuebles")

	savePath := filepath.Join(saveDir, imagenName)

	return savePath, imagenName, nil
}

func (i *ServiceInmueble) CargarImagenInmueble(imagenName, idInmueble string) error {

	InmuebleID, err := strconv.Atoi(idInmueble)
	if err != nil {
		return err
	}

	publicURL := "/static/imagenesInmuebles/" + imagenName
	img := domain.ImagenInmueble{
		IdInmuebleImagen: uint(InmuebleID),
		Url:              publicURL,
		PathLocal:        filepath.Join("./imagenesInmuebles", imagenName),
	}

	if err := i.Repo.GuardarImagenInmueble(img); err != nil {
		return err
	}

	return nil
}

func (i *ServiceInmueble) EliminarImagen(id string) (string, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	imagen, err := i.Repo.GetImagen(uint(idInt))

	err = i.Repo.EliminarImagen(uint(idInt))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("path/carpeta/inmuebleImagenes\\%s", imagen.PathLocal), err

}

func (i *ServiceInmueble) InmueblesPorEncargado(id uint) ([]domain.Inmueble, error) {

	inmuebles, err := i.Repo.InmueblesPorEncargado(id)
	if err != nil {
		return inmuebles, err
	}

	if len(inmuebles) == 0 {
		return inmuebles, domain.ErrorEncargadoNoTieneInmuebles()
	}

	return inmuebles, nil
}

func (i *ServiceInmueble) RegistrarCheckOut(id, estado string) error {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	reserva, err := i.ResRepo.GetReservaId(uint(idInt))
	if err != nil {
		return err
	}

	reserva.Estado = estado
	reserva.CheckOut = true

	err = i.Repo.RegistrarCheckOut(reserva)
	if err != nil {
		return err
	}

	return nil
}
