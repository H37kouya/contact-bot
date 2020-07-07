package di

import (
	"contact-bot/pkg/domain/repository"
	"contact-bot/pkg/infra/persistence"
)

// InjectContactSheetPersistence InjectContactSheetPersistenceの依存性注入
func InjectContactSheetPersistence() repository.ContactRepository {
	return persistence.NewContactSheetPersistence()
}
