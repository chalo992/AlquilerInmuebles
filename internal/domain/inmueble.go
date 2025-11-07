package domain

type Inmueble struct {
	ID                  uint             `json:"id" gorm:"primaryKey"`
	Nombre              string           `json:"nombre"`
	Localidad           string           `json:"localidad"`
	PrecioDia           int              `json:"precio_dia"`
	Direccion           string           `json:"direccion"`
	Habitaciones        int              `json:"habitaciones"`
	Camas               int              `json:"camas"`
	DiasMinimosAlquiler int              `json:"dias_minimos_alquiler"`
	TienePileta         bool             `json:"tiene_pileta"`
	TieneInternet       bool             `json:"tiene_internet"`
	TieneGasNatural     bool             `json:"tiene_gas_natural"`
	TieneCable          bool             `json:"tiene_cable"`
	TieneAire           bool             `json:"tiene_aire"`
	MaxInquilinos       int              `json:"max_inquilinos"`
	IdEncargado         uint             `json:"id_encargado"`
	PoliticaDevolucion  string           `json:"politica_devolucion"`
	PoliticaPago        string           `json:"politica_pago"`
	Pausado             bool             `json:"pausado" gorm:"default:false"`
	Imagenes            []ImagenInmueble `json:"imagenes" gorm:"constraint:OnDelete:CASCADE;foreignKey:IdInmuebleImagen"`
	Reservas            []Reserva        `gorm:"constraint:OnDelete:SET NULL;foreignKey:IdInmuebleReserva"`
	Calificaciones      []Calificacion   `gorm:"constraint:OnDelete:CASCADE;foreignKey:IdInmuebleCalificacion"`
}

type Estado struct {
	Estado string `json:"estado"`
}
