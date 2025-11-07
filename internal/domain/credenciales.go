package domain

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CambiarContraseña struct {
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
}

type CambiarContraseñaUser struct {
	ContraseñaActual string `json:"contraseñaActual"`
	NuevaContraseña  string `json:"nuevaContraseña"`
}
