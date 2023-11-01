package cmd

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadDotEnvVariables() error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}

	return nil
}
