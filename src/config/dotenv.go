package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	loadDotEnv()
}

// loadDotEnv .envファイルをロードする
func loadDotEnv() {
	if godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV"))) != nil {
		log.Fatal("Error loading .env file")
	}
}
