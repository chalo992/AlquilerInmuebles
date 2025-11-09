package credenciales

import (
	api "AlquilerInmuebles/cmd/api/common"
	"AlquilerInmuebles/internal/domain"
	"AlquilerInmuebles/internal/services/common"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceCredenciales) ComprobarCredenciales(credenciales domain.Login) (domain.Usuario, error) {

	usuario, err := s.Repo.ComprobarCredencialesEmail(credenciales.Email)
	if err != nil {
		return usuario, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(credenciales.Password)); err != nil {
		return usuario, domain.ErrorEmailContraseñaIncorrecto()
	}

	return usuario, nil
}

func (s *ServiceCredenciales) GenerarToken(usuario domain.Usuario) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  usuario.ID,
		"rol": usuario.Rol,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(api.Secret))
	if err != nil {
		fmt.Println(err)
		return "", domain.ErrorCrearToken()
	}

	return tokenString, nil
}

func (s *ServiceCredenciales) CambiarContraseña(id uint, contraseña domain.CambiarContraseñaUser) error {

	usuario, err := s.UserRepo.GetUsuarioID(id)

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(contraseña.ContraseñaActual))
	if err != nil {
		return domain.ErrorContraseñaActual()
	}

	err = common.ComprobarContraseña(contraseña.NuevaContraseña)
	if err != nil {
		return err
	}

	hasPassword, err := common.HashPassword(contraseña.NuevaContraseña)
	if err != nil {
		return err
	}

	usuario.Password = hasPassword

	err = s.UserRepo.Actualizar(usuario)
	if err != nil {
		return err
	}

	return nil
}
