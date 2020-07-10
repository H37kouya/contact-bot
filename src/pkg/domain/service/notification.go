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
	slackNotification notification.SlackNotificaion
}

// NewNotificationService NotificationService DIするために必要
func NewNotificationService(sn notification.SlackNotificaion) NotificationService {
	return &notificationService{
		slackNotification: sn,
	}
}

// ContactNotification Contactの情報を通知する
func (ns notificationService) ContactNotification(contacts []model.Contact) error {
	notifications := model.ContactToNotification("最初のタイトル", "最初のメッセージ", contacts)

	err := ns.slackNotification.TestNotification(notifications)
	return err
}
