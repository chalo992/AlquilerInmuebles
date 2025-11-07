package services

import (
	"AlquilerInmuebles/internal/domain"
	"os"

	"github.com/joho/godotenv"
)

var BaseUrl string
var MailSender domain.MailSender

func GetConfigService() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	BaseUrl = os.Getenv("BASE_URL")
	MailSender.Host = os.Getenv("EMAIL_HOST")
	MailSender.Port = os.Getenv("EMAIL_PORT")
	MailSender.Username = os.Getenv("EMAIL_USERNAME")
	MailSender.Password = os.Getenv("EMAIL_PASSWORD")

}
