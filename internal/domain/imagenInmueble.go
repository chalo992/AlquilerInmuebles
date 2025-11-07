package domain

type ImagenInmueble struct {
	ID               uint   `json:"id"`
	Url              string `json:"url"`
	PathLocal        string `json:"path_local"`
	IdInmuebleImagen uint   `json:"id_inmueble_imagen"`
}
