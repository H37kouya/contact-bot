package cli

import (
	"contact-bot/pkg/di"
)

// NotificationHandler 通知処理
func NotificationHandler(pollingDiff int) {
	nu := di.InjectNotificationUsecase()

	nu.SlackNotification(pollingDiff)
}
