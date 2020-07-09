package di

import (
	"contact-bot/pkg/usecase"
)

// InjectNotificationUsecase InjectContactServiceの依存性注入
func InjectNotificationUsecase() usecase.NotificationUsecase {
	return usecase.NewNotificationUsecase(
		InjectContactService(),
		InjectNotificationService(),
	)
}
