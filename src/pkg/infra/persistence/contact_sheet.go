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
func (cp contactSheetPersistence) GetContactSheet(spreadsheetID string) (*sheets.Spreadsheet, error) {
	jsonKey, err := getJSONKey()
	if err != nil {
		return nil, err
	}

	client, err := getHTTPClient(jsonKey)
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

func getJSONKey() ([]byte, error) {
	b, err := ioutil.ReadFile("credentials.json")
	return b, err
}

// 認証情報からClient情報を取得
func getHTTPClient(jsonKey []byte) (*http.Client, error) {
	conf, err := google.JWTConfigFromJSON(jsonKey, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}

	return conf.Client(oauth2.NoContext), nil
}
