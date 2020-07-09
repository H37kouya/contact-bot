package slack

import (
	"contact-bot/pkg/domain/notification"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
)

const (
	CHANNEL  = "botmake"
	USERNAME = "GoBot"
)

// NewSendSlack SendSlack
func NewSendSlack() notification.SlackNotificaion {
	return &sendSlack{}
}

// slackSend slackSend データの構造体
type sendSlack struct{}

// TestNotification Test通知用
func (ss sendSlack) TestNotification() error {
	field1 := slack.Field{Title: "Message", Value: "Hello World!!!!"}
	field2 := slack.Field{Title: "AnythingKey", Value: "AnythingValue"}

	attachment := slack.Attachment{}
	attachment.AddField(field1).AddField(field2)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     CHANNEL,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(os.Getenv("SLACK_WEBHOOK_URL"), "", payload)

	if len(err) == 0 {
		return nil
	}
	return err[0]
}
