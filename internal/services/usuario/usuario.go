package usuario

import (
	"AlquilerInmuebles/internal/domain"
	"AlquilerInmuebles/internal/services"
	"AlquilerInmuebles/internal/services/common"
	"fmt"
	"strconv"
	"time"
)

func (s *Service) CrearUsuario(usuario domain.Usuario, rol string) (domain.Usuario, error) {

	err := s.Repo.VerificarEmail(usuario.Email)

	if err != nil {
		return usuario, err
	}

	if rol != "encargado" {
		err = common.ComprobarContraseña(usuario.Password)
		if err != nil {
			return usuario, err
		}

		contraseñaHash, err := common.HashPassword(usuario.Password)

		if err != nil {
			return usuario, err
		}

		usuario.Password = contraseñaHash
	}

	usuario.Creacion = time.Now().UTC()
	usuario.Rol = rol

	user, err := s.Repo.Crear(usuario)

	if err != nil {
		return usuario, err
	}
	return user, nil
}

func (s *Service) GetUsuarioByID(id uint) (domain.Usuario, error) {

	usuario, err := s.Repo.GetUsuarioID(id)

	if err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (s *Service) ActualizarUsuario(usuarioActualizado domain.Usuario, idParam uint) (domain.Usuario, error) {

	usuario, err := s.Repo.GetUsuarioID(idParam)

	if err != nil {
		return usuarioActualizado, err
	}

	if usuario.Email != usuarioActualizado.Email {
		err = s.Repo.VerificarEmail(usuarioActualizado.Email)
		if err != nil {
			return usuarioActualizado, err
		}
	}

	usuario.Apellido = usuarioActualizado.Apellido
	usuario.DNI = usuarioActualizado.DNI
	usuario.Nombre = usuarioActualizado.Nombre
	usuario.Email = usuarioActualizado.Email
	usuario.Nacionalidad = usuarioActualizado.Nacionalidad
	usuario.Telefono = usuarioActualizado.Telefono

	err = s.Repo.Actualizar(usuario)

	if err != nil {
		return usuarioActualizado, err
	}

	return usuario, nil

}

func (s *Service) GetEncargados() ([]domain.Usuario, error) {

	encargados, err := s.Repo.DevolverEncargados()

	if err != nil {
		return encargados, err
	}

	return encargados, nil
}

func (s *Service) EstablecerContraseñaEncargado(contraseña domain.CambiarContraseña, id string) error {

	if contraseña.Password != contraseña.PasswordRepeat {
		return domain.ErrorNoCoincideContraseñas()
	}

	err := common.ComprobarContraseña(contraseña.Password)
	if err != nil {
		return err
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	encargado, err := s.Repo.GetUsuarioID(uint(idInt))
	if err != nil {
		return err
	}

	if encargado.Password != "" {
		return domain.ErrorContraseñaEncargado()
	}

	contraseñaHash, err := common.HashPassword(contraseña.Password)

	if err != nil {
		return err
	}

	encargado.Password = contraseñaHash

	err = s.Repo.Actualizar(encargado)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) MandarMailEncargado(encargado domain.Usuario) error {
	subject := "Crear contraseña"
	link := fmt.Sprintf("url/contraseña/encargado/%d", encargado.ID)
	body := fmt.Sprintf(`
Hola %s,

Entrá al siguiente link: %s para establecer tu contraseña y completar tu registro como encargado 

`, encargado.Nombre, link)

	err := services.MailSender.SendMail([]string{encargado.Email}, subject, body)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUsuarioIDParam(id string) (domain.Usuario, error) {

	var usuario domain.Usuario

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return usuario, err
	}

	usuario, err = s.Repo.GetUsuarioID(uint(idInt))
	if err != nil {
		return usuario, err
	}

	return usuario, err
}

func (s *Service) GetInquilinosReserva(id_reserva string) ([]domain.Inquilino, error) {

	idInt, err := strconv.Atoi(id_reserva)
	if err != nil {
		return nil, err
	}

	inquilinos, err := s.Repo.GetInquilinosReserva(uint(idInt))
	if err != nil {
		return nil, err
	}

	for i := range inquilinos {
		inquilinos[i].ImagenUrl = services.BaseUrl + inquilinos[i].ImagenUrl
	}

	return inquilinos, nil
}
