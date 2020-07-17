package config_test

import (
	"contact-bot/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("環境変数を取得できる", func(t *testing.T) {
		config := config.Conf

		assert.NotNil(t, config.SpreadSheet.ID)
	})
}
