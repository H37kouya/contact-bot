package config

import (
	"os"
)

// SpreadSheetConfig SpreedSheetの設定
type SpreadSheetConfig struct {
	ID string
}

// Config Configの設定
type Config struct {
	SpreadSheet SpreadSheetConfig
}

// NewConfig return configuration struct.
func NewConfig() (Config, error) {
	var config Config

	config.SpreadSheet.ID = os.Getenv("SPREAD_SHEET_ID")

	return config, nil
}
