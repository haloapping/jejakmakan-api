package config

import (
	"os"
)

func APIUrl(envName string) string {
	return os.Getenv("API_URL")
}
