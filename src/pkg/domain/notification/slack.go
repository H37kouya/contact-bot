package notification

import "contact-bot/pkg/domain/model"

type SlackNotificaion interface {
	TestNotification(notifications []model.Notification) error
	ErrorNotification(err error) error
}
