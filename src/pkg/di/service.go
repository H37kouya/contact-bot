package di

import "contact-bot/pkg/domain/service"

// InjectContactService InjectContactServiceの依存性注入
func InjectContactService() service.ContactService {
	return service.NewContactService(
		InjectContactInfra(),
	)
}
