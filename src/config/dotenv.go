package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	loadDotEnv()
}

// loadDotEnv .envファイルをロードする
func loadDotEnv() {
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}
}
