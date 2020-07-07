package repository

import "google.golang.org/api/sheets/v4"

// ContactRepository Contactリポジトリ
type ContactRepository interface {
	GetContactSheet(spreadsheetID, credentialFilePath string) (*sheets.Spreadsheet, error)
}
