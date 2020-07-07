package config_test

import (
	"contact-bot/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("環境変数を取得できる", func(t *testing.T) {
		config, err := config.NewConfig()

		if err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, config.SpreadSheet.ID)
	})
}
