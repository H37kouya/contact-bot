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
	notifications := make([]model.Notification, 0, 3)
	notifications = append(notifications, model.Notification{
		Title: "Message",
		Value: "Hello World!!!!",
	})
	notifications = append(notifications, model.Notification{
		Title: "AnythingKey",
		Value: "AnythingValue",
	})
	notifications = append(notifications, model.Notification{
		Title: "Yahho",
		Value: "Is it Perfect??",
	})

	err := ns.slackNotificaion.TestNotification(notifications)
	return err
}
