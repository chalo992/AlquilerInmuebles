package domain

type Calificacion struct {
	ID                     uint   `json:"id" gorm:"primaryKey"`
	IdInmuebleCalificacion uint   `json:"id_inmueble_calificacion"`
	IdReservaCalificacion  uint   `json:"id_reserva_calificacion"`
	IdUsuarioCalificacion  uint   `json:"id_usuario_calificacion"`
	Limpieza               int    `json:"limpieza"`
	Comodidad              int    `json:"comodidad"`
	Ubicacion              int    `json:"ubicacion"`
	Distribucion           int    `json:"distribucion"`
	Comentario             string `json:"comentario"`
}
