package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".devcontainer/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
