package repository

// ContactRepository Contactリポジトリ
type ContactRepository interface {
	GetContactSheet(spreadsheetID string)
}
