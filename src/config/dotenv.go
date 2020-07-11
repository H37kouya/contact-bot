package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	loadDotEnv()
}

// loadDotEnv .envファイルをロードする
func loadDotEnv() {
	if os.Getenv("CI") == true {
		return
	}

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}
}
