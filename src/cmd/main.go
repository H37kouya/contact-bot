package main

import (
	"contact-bot/pkg/handler/cli"
	"flag"
)

var (
	// pollingDiff 取得する時間間隔
	pollingDiff int
)

func init() {
	flag.IntVar(&pollingDiff, "diff", 1, "取得する時間間隔(h)")
	flag.Parse()
}

func main() {
	cli.NotificationHandler(pollingDiff)
}
