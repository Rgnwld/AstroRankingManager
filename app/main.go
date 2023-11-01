package main

import (
	"Astro/cmd"
	"log"
)

func main() {

	cmd.LoadDotEnvVariables()

	log.Fatal(cmd.Execute())
}
