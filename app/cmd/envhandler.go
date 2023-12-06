package cmd

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadDotEnvVariables() error {
	err := godotenv.Load("./.env")
	if err != nil {
		return fmt.Errorf(".env was not provided.\nUsing ambient values\n%w", err) //Add color
	}

	return nil
}
