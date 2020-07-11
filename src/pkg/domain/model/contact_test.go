package model_test

import (
	"contact-bot/pkg/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsBetweenTimeStamp(t *testing.T) {
	testCases := []struct {
		name string
		want bool
		arg  []time.Time
	}{
		{
			name: "二つの時間の間にあるとき、trueであるか", want: true, arg: []time.Time{strToTime("14:59"), strToTime("15:01")},
		},
		{
			name: "二つの時間が等しいとき、falseであるか", want: false, arg: []time.Time{strToTime("15:00"), strToTime("15:00")},
		},
		{
			name: "二つの時間が早いとき、falseであるか", want: false, arg: []time.Time{strToTime("14:50"), strToTime("14:55")},
		},
		{
			name: "二つの時間が遅いとき、falseであるか", want: false, arg: []time.Time{strToTime("15:05"), strToTime("15:10")},
		},
	}

	contact := model.Contact{
		Title:     "タイトル",
		Second:    "セカンド",
		Third:     "サード",
		TimeStamp: strToTime("15:00"),
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.want, contact.IsBetweenTimeStamp(testCase.arg[0], testCase.arg[1]))
		})
	}
}

func TestFilterContactByTimeStamp(t *testing.T) {

	contacts := []model.Contact{
		{
			Title:     "タイトル",
			Second:    "セカンド",
			Third:     "サード",
			TimeStamp: strToTime("15:05"),
		},
		{
			Title:     "タイトル",
			Second:    "セカンド",
			Third:     "サード",
			TimeStamp: strToTime("15:10"),
		},
		{
			Title:     "タイトル",
			Second:    "セカンド",
			Third:     "サード",
			TimeStamp: strToTime("16:10"),
		},
		{
			Title:     "タイトル",
			Second:    "セカンド",
			Third:     "サード",
			TimeStamp: strToTime("16:20"),
		},
	}

	t.Run("フィルターできるか", func(t *testing.T) {
		filter := model.FilterContactByTimeStamp(contacts, strToTime("15:00"), strToTime("16:00"))
		assert.Equal(t, 2, len(filter))
	})
}

func strToTime(s string) time.Time {
	t, err := time.Parse("15:04", s)
	if err != nil {
		panic(err)
	}

	return t
}
