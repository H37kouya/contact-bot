package repository

import (
	"contact-bot/pkg/domain/model"
)

// ContactRepository Contactリポジトリ
type ContactRepository interface {
	GetContactSheet(spreadsheetID string) ([]model.Contact, error)
}
