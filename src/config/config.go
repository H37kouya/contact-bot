package config

import (
	"os"
	"time"
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
	ID           string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
	Credential   []byte
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

	conf.SpreadSheet.AccessToken = os.Getenv("SHEET_ACCESS_TOKEN")
	conf.SpreadSheet.TokenType = os.Getenv("SHEET_TOKEN_TYPE")
	conf.SpreadSheet.RefreshToken = os.Getenv("SHEET_REFRESH_TOKEN")

	if expiry, err := time.Parse("2006-01-02T15:04:05.999999999Z", os.Getenv("SHEET_EXPIRY")); err != nil {
		conf.SpreadSheet.Expiry = expiry
	}

	if c := os.Getenv("SHEET_CREDENTIAL"); c != "" {
		conf.SpreadSheet.Credential = []byte(c)
	}
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
