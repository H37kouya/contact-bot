package usecase

import (
	"contact-bot/pkg/domain/service"
	"fmt"
)

// NotificationUsecase NotificationUsecaseのためのinterface
type NotificationUsecase interface {
	SlackNotification(pollingDiff int)
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

// SlackNotification Slackへ通知する
func (nu notificationUsecase) SlackNotification(pollingDiff int) {
	contacts, err := nu.contactService.GetContactData(pollingDiff)
	if err != nil {
		fmt.Println(err)
		nu.notificationService.ErrorNotification(err)
		return
	}

	if len(contacts) > 0 {
		if err = nu.notificationService.ContactNotification(contacts); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Slackに通知しました！")
		return
	}

	fmt.Println("最新の通知は0件でした")
}
