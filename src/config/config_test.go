package config_test

import (
	"contact-bot/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("環境変数を取得できる", func(t *testing.T) {
		config := config.Conf

		assert.NotNil(t, config.SpreadSheet.ID)
	})
}

func TestExpiryStrToTime(t *testing.T) {
	t.Run("Expiryの文字列を変換できる", func(t *testing.T) {
		actual := config.ExpiryStrToTime("2020-07-20T03:18:33.0Z")
		expect := time.Date(2020, 7, 20, 3, 18, 33, 0, time.Local)

		assert.Equal(t, true, actual.Equal(expect))
	})
}
