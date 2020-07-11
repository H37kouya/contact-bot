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

// ToNotificationMessage Notification用のMessageの取得
func (contact Contact) ToNotificationMessage() string {
	meg := contact.Second + contact.Third
	return meg
}

// IsBetweenTimeStamp TimeStampが2つの時間の中間にあるか
func (contact Contact) IsBetweenTimeStamp(before, after time.Time) bool {
	return contact.TimeStamp.After(before) && contact.TimeStamp.Before(after)
}

// ContactToNotification CotactをNotificationへ変換
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

// FilterContactByTimeStamp TimeStampを用いてContactをフィルターする
func FilterContactByTimeStamp(contacts []Contact, before, after time.Time) []Contact {
	newContacts := make([]Contact, 0, len(contacts))

	for _, contact := range contacts {
		if contact.IsBetweenTimeStamp(before, after) {
			newContacts = append(newContacts, contact)
		}
	}

	return newContacts
}
