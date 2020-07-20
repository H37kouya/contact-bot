package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loadDotEnv .envファイルをロードする
func loadDotEnv() {
	if os.Getenv("CI") == "true" {
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}
}
