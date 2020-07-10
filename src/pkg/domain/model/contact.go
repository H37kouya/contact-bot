package model

import (
	"time"
)

// Contact Contact struct
type Contact struct {
	Title     string
	Second    string
	Third     string
	TimeStamp time.Time
}

func (contact Contact) ToNotificationMessage() string {
	meg := contact.Second + contact.Third
	return meg
}

func ContactToNotification(firstTitle, firstMeg string, contacts []Contact) []Notification {
	notifications := make([]Notification, 0, len(contacts))

	notifications = append(notifications, Notification{
		Title: firstTitle,
		Value: firstMeg,
	})

	for _, contact := range contacts {
		notifications = append(notifications, Notification{
			Title: contact.Title,
			Value: contact.ToNotificationMessage(),
		})
	}

	return notifications
}
