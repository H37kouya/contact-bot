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
	fileSuffix := ""

	if os.Getenv("GO_ENV") != "" {
		fileSuffix = "." + os.Getenv("GO_ENV")
	}

	if godotenv.Load(fmt.Sprintf(".env%s", fileSuffix)) != nil {
		log.Fatal(fmt.Sprintf(".env%s", fileSuffix))
		log.Fatal("Error loading .env file")
	}
}
