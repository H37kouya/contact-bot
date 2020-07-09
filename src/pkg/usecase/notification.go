package usecase

import (
	"contact-bot/pkg/domain/service"
	"contact-bot/pkg/infra/slack"
	"fmt"
)

// NotificationUsecase NotificationUsecaseのためのinterface
type NotificationUsecase interface {
	SlackNotification()
}

type notificationUsecase struct {
	contactService service.ContactService
}

// NewNotificationUsecase NotificationUsecase DIするために必要
func NewNotificationUsecase(cs service.ContactService) NotificationUsecase {
	return &notificationUsecase{
		contactService: cs,
	}
}

func (nu notificationUsecase) SlackNotification() {
	contacts, err := nu.contactService.GetContactData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(contacts)
	slack.TestNotification()
}
