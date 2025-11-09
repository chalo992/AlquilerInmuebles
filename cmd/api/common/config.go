package common

import (
	"os"

	"github.com/joho/godotenv"
)

var Secret string

func GetConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	Secret = os.Getenv("SECRET")
}
