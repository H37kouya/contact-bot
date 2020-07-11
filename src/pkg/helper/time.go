package helper

import (
	"fmt"
	"time"
)

const (
	// timeLayout Timeのレイアウト
	timeLayout string = "2006/01/02 15:04:05"
)

// GetNowTokyoTime Tokyoの時間を取得
func GetNowTokyoTime() time.Time {
	now := time.Now()
	return ToTokyoTime(now)
}

// GetHourTime 15:01:00だったら15:00:00を取得する
func GetHourTime(t time.Time) time.Time {
	hour := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, time.Local)
	return hour
}

// ToTokyoTime Tokyoの時間を取得
func ToTokyoTime(t time.Time) time.Time {
	nowUTC := t.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	return nowJST
}

// SheetTimeStampStrToTime SpreadSheetのTimeStampをTime型に変換
func SheetTimeStampStrToTime(sheetTimeStamp string) time.Time {
	timestamp, err := time.Parse(timeLayout, sheetTimeStamp)

	if err != nil {
		fmt.Println(err)
	}

	return timestamp
}
