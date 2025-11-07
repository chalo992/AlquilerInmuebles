package domain

type Inquilino struct {
	ID                 uint   `json:"id" gorm:"primarykey"`
	IdReservaInquilino uint   `json:"id_reserva_inquilino"`
	DNI                string `json:"dni"`
	NombreCompleto     string `json:"nombre_completo"`
	ImagenUrl          string `json:"imagen_url"`
}
