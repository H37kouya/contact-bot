package cli

import "contact-bot/pkg/usecase"

// NotificationHandler 通知処理
func NotificationHandler() {
	usecase.NotificationUsecase()
}
