package cli

import (
	"contact-bot/pkg/di"
)

// NotificationHandler 通知処理
func NotificationHandler() {
	nu := di.InjectNotificationUsecase()

	nu.SlackNotification()
}
