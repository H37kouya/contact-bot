package usecase

import (
	"contact-bot/pkg/domain/service"
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
	fmt.Println("Hello World")

	nu.contactService.GetContactData()
}
