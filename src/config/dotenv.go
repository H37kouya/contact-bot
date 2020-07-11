package config

import (
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	loadDotEnv()
}

// loadDotEnv .envファイルをロードする
func loadDotEnv() {
	fmt.Println(os.Getenv("CI"))
	if os.Getenv("CI") == "true" {
		return
	}

	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}
}
