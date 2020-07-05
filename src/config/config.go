package config

import (
	"os"
)

// SpreedSheetConfig SpreedSheetの設定
type SpreedSheetConfig struct {
	ID string
}

// Config Configの設定
type Config struct {
	SpreedSheet SpreedSheetConfig
}

// NewConfig return configuration struct.
func NewConfig() (Config, error) {
	var config Config

	config.SpreedSheet.ID = os.Getenv("SPREED_SHEET_ID")

	return config, nil
}
