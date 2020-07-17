package config

import (
	"os"
)

const (
	// CHANNEL Slackチャンネル名
	CHANNEL = "botmake"
	// USERNAME Slackのユーザー名
	USERNAME = "GoBot"
)

// Conf 設定系の情報
var Conf *Config

// SpreadSheetConfig SpreedSheetの設定
type SpreadSheetConfig struct {
	ID string
}

// SlackConfig Slackの設定1
type SlackConfig struct {
	ChannelName string
	WebhookURL  string
	Username    string
}

// Config Configの設定
type Config struct {
	SpreadSheet  SpreadSheetConfig
	ContactSlack SlackConfig
}

func init() {
	Conf = NewConfig()
}

// NewConfig return configuration struct.
func NewConfig() *Config {
	config := &Config{}
	config.setSpreadSheetConf()
	config.setSlack()
	return config
}

func (conf *Config) setSpreadSheetConf() {
	conf.SpreadSheet.ID = os.Getenv("SPREAD_SHEET_ID")
}

func (conf *Config) setSlack() {

	if str := os.Getenv("SLACK_CHANNEL_NAME"); str != "" {
		conf.ContactSlack.ChannelName = str
	} else {
		conf.ContactSlack.ChannelName = CHANNEL
	}

	if str := os.Getenv("SLACK_USERNAME"); str != "" {
		conf.ContactSlack.Username = str
	} else {
		conf.ContactSlack.Username = USERNAME
	}

	conf.ContactSlack.WebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
}
