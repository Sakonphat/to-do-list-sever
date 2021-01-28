package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//Env function is for read .env file
func Env(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
