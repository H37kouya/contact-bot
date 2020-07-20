package slack

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/notification"

	"github.com/ashwanthkumar/slack-go-webhook"
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
	conf := config.Conf.ContactSlack

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
		Username:    conf.Username,
		Channel:     conf.ChannelName,
		Attachments: []slack.Attachment{attachment},
	}

	if err := slack.Send(conf.WebhookURL, "", payload); len(err) >= 0 {
		return err[0]
	}

	return nil
}
