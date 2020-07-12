package helper_test

import (
	"contact-bot/pkg/helper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetBeforeHourTime(t *testing.T) {
	testCases := []struct {
		name string
		want time.Time
		arg1 int
		arg2 time.Time
	}{
		{
			name: "1時間前を取得できるか",
			want: time.Date(2020, 7, 12, 15, 0, 0, 0, time.Local),
			arg1: 1,
			arg2: time.Date(2020, 7, 12, 16, 0, 0, 0, time.Local),
		},
		{
			name: "1時間前を取得できるか(日付をまたぐ)",
			want: time.Date(2020, 7, 11, 23, 0, 0, 0, time.Local),
			arg1: 1,
			arg2: time.Date(2020, 7, 12, 0, 0, 0, 0, time.Local),
		},
		{
			name: "6時間前を取得できるか",
			want: time.Date(2020, 7, 12, 10, 0, 0, 0, time.Local),
			arg1: 6,
			arg2: time.Date(2020, 7, 12, 16, 0, 0, 0, time.Local),
		},
		{
			name: "6時間前を取得できるか(日付をまたぐ)",
			want: time.Date(2020, 7, 11, 21, 0, 0, 0, time.Local),
			arg1: 6,
			arg2: time.Date(2020, 7, 12, 3, 0, 0, 0, time.Local),
		},
		{
			name: "24時間前を取得できるか",
			want: time.Date(2020, 7, 11, 16, 0, 0, 0, time.Local),
			arg1: 24,
			arg2: time.Date(2020, 7, 12, 16, 0, 0, 0, time.Local),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := helper.GetBeforeHourTime(testCase.arg1, testCase.arg2)

			assert.Equal(t, true, actual.Equal(testCase.want))
		})
	}
}

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
