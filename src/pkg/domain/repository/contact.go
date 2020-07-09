package repository

import "google.golang.org/api/sheets/v4"

// ContactRepository Contactリポジトリ
type ContactRepository interface {
	GetContactSheet(spreadsheetID string) (*sheets.Spreadsheet, error)
}
