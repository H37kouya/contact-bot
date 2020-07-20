package service

import (
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/notification"
	"fmt"
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
	title := "お問い合わせがあります"
	meg := fmt.Sprintf("%d件あります。", len(contacts))
	notifications := model.ContactToNotification(title, meg, contacts)

	err := ns.slackNotification.TestNotification(notifications)
	return err
}
