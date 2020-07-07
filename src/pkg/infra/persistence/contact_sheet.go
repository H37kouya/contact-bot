package persistence

import (
	"contact-bot/pkg/domain/repository"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// NewContactSheetPersistence ContactSheetPersistence
func NewContactSheetPersistence() repository.ContactRepository {
	return &contactSheetPersistence{}
}

// contactSheetPersistence Image データの構造体
type contactSheetPersistence struct{}

// GetContactSheetPersistence ContactSheetを取得する
func (cp contactSheetPersistence) GetContactSheet(spreadsheetID, credentialFilePath string) (*sheets.Spreadsheet, error) {
	client, err := httpClient(credentialFilePath)
	if err != nil {
		return nil, err
	}

	sheetService, err := sheets.New(client)
	if err != nil {
		return nil, err
	}

	spreadsheet, err := sheetService.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		return nil, err
	}

	return spreadsheet, nil
}

func httpClient(credentialFilePath string) (*http.Client, error) {
	data, err := ioutil.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext), nil
}
