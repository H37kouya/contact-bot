package di

import "contact-bot/pkg/domain/service"

// InjectContactService InjectContactServiceの依存性注入
func InjectContactService() service.ContactService {
	return service.NewContactService(
		InjectContactSheetPersistence(),
	)
}

// InjectNotificationService InjectNotificationServiceの依存性注入
func InjectNotificationService() service.NotificationService {
	return service.NewNotificationService(
		InjectSendSlack(),
	)
}
