package domain

import "time"

type Tarjeta struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	IdUsuarioTarjeta uint      `json:"id_usuario_tarjeta"`
	Numero           int       `json:"numero"`
	CodigoSeguridad  int       `json:"codigoseguridad"`
	Vencimiento      time.Time `json:"vencimiento"`
}

type TarjetaJSON struct {
	Numero          int    `json:"numero"`
	CodigoSeguridad int    `json:"codigoseguridad"`
	Vencimiento     string `json:"vencimiento"`
}
