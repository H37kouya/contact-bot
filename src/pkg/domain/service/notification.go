package service

import (
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/notification"
	"strconv"
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
	notifications := contactToNotification("最初のタイトル", "最初のメッセージ", contacts)

	err := ns.slackNotification.TestNotification(notifications)
	return err
}

func contactToNotification(firstTitle, firstMeg string, contacts []model.Contact) []model.Notification {
	notifications := make([]model.Notification, 0, len(contacts))

	notifications = append(notifications, model.Notification{
		Title: firstTitle,
		Value: firstMeg,
	})

	for _, contact := range contacts {
		notifications = append(notifications, model.Notification{
			Title: contact.Title,
			Value: strconv.Itoa(contact.ID),
		})
	}

	return notifications
}
