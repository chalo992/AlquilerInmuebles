package common

import (
	"AlquilerInmuebles/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComprobarContraseña(contraseña string) error {
	if len(contraseña) < 8 {
		return domain.ErrorContraseñaNoCumple()
	}
	return nil
}
