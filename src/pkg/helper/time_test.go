package helper_test

import (
	"contact-bot/pkg/helper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSheetTimeStampStrToTime(t *testing.T) {
	t.Run("文字列をTime型に変換できる", func(t *testing.T) {
		actual := helper.SheetTimeStampStrToTime("2006/01/03 15:04:05")
		expect := time.Date(2006, 1, 3, 15, 4, 5, 0, time.Local)

		assert.Equal(t, true, actual.Equal(expect))
	})
}

func TestGetHourTime(t *testing.T) {
	t.Run("15:00:00を取得できる", func(t *testing.T) {
		actual := helper.GetHourTime(helper.SheetTimeStampStrToTime("2006/01/03 15:04:05"))

		assert.Equal(t, true, actual.Equal(helper.SheetTimeStampStrToTime("2006/01/03 15:00:00")))
	})
}
