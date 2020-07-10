package di

import (
	"contact-bot/pkg/domain/notification"
	"contact-bot/pkg/infra/slack"
)

func InjectSendSlack() notification.SlackNotificaion {
	return slack.NewSendSlack()
}
