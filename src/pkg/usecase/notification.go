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
	contactService      service.ContactService
	notificationService service.NotificationService
}

// NewNotificationUsecase NotificationUsecase DIするために必要
func NewNotificationUsecase(cs service.ContactService, ns service.NotificationService) NotificationUsecase {
	return &notificationUsecase{
		contactService:      cs,
		notificationService: ns,
	}
}

func (nu notificationUsecase) SlackNotification() {
	contacts, err := nu.contactService.GetContactData()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(contacts)
	err = nu.notificationService.ContactNotification(contacts)
	if err != nil {
		fmt.Println(err)
	}
}
