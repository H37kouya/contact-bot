package slack

import (
	"contact-bot/pkg/domain/model"
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
func (ss sendSlack) TestNotification(notifications []model.Notification) error {
	attachment := slack.Attachment{}

	for _, notification := range notifications {
		slackField := slack.Field{
			Title: notification.Title,
			Value: notification.Value,
			Short: notification.Short,
		}
		attachment.AddField(slackField)
	}

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
