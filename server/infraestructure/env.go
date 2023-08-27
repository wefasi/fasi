package infraestructure

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	return err
}
