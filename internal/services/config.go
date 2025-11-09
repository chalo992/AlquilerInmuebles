package services

import (
	"AlquilerInmuebles/internal/domain"
	"os"

	"github.com/joho/godotenv"
)

var MailSender domain.MailSender
var BaseUrl string

func GetConfigService() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	MailSender.Host = os.Getenv("EMAIL_HOST")
	MailSender.Port = os.Getenv("EMAIL_PORT")
	MailSender.Username = os.Getenv("EMAIL_USERNAME")
	MailSender.Password = os.Getenv("EMAIL_PASSWORD")
	BaseUrl = os.Getenv("BASE_URL")

}
