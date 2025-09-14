package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func APIUrl(envName string) string {
	_ = fmt.Sprintf("../%s", envName)
	err := godotenv.Load(fmt.Sprintf("./%s", envName))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("API_URL")
}
