package service

import (
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/notification"
)

// NotificationService NotificationService for Interface
type NotificationService interface {
	ContactNotification(contacts []model.Contact) error
}

type notificationService struct {
	slackNotificaion notification.SlackNotificaion
}

// NewNotificationService NotificationService DIするために必要
func NewNotificationService(sn notification.SlackNotificaion) NotificationService {
	return &notificationService{
		slackNotificaion: sn,
	}
}

// ContactNotification Contactの情報を通知する
func (ns notificationService) ContactNotification(contacts []model.Contact) error {
	err := ns.slackNotificaion.TestNotification()

	return err
}
