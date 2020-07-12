package persistence

import (
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/repository"
	"contact-bot/pkg/helper"
	"contact-bot/pkg/infra/sheet"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2/google"
)

// NewContactSheetPersistence ContactSheetPersistence
func NewContactSheetPersistence() repository.ContactRepository {
	return &contactSheetPersistence{}
}

// contactSheetPersistence contactSheetPersistence データの構造体
type contactSheetPersistence struct{}

// GetContactSheetPersistence ContactSheetを取得する
func (cp contactSheetPersistence) GetContactSheet(spreadsheetID string) ([]model.Contact, error) {
	b, err := getCredentials()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		return nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	srv, err := sheet.NewSheet(config)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
		return nil, err
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	readRange := "フォームの回答 1!A2:D"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return nil, nil
	}

	contacts := convertToContacts(resp.Values)
	return contacts, nil
}

func convertToContacts(rows [][]interface{}) []model.Contact {
	contacts := make([]model.Contact, 0, len(rows))

	for _, row := range rows {
		timestamp := helper.SheetTimeStampStrToTime(row[0].(string))

		if len(row) == 3 {
			contact := model.Contact{
				TimeStamp: timestamp,
				Title:     interfaceToStr(row[1]),
				Second:    interfaceToStr(row[2]),
				Third:     "",
			}
			contacts = append(contacts, contact)
		} else if len(row) == 4 {
			contact := model.Contact{
				TimeStamp: timestamp,
				Title:     interfaceToStr(row[1]),
				Second:    interfaceToStr(row[2]),
				Third:     interfaceToStr(row[3]),
			}
			contacts = append(contacts, contact)
		}
	}

	return contacts
}

func getCredentials() ([]byte, error) {
	c := os.Getenv("SHEET_CREDENTIAL")

	if c != "" {
		b := []byte(c)
		return b, nil
	}

	return ioutil.ReadFile("credentials.json")
}

func interfaceToStr(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	default:
		return ""
	}
}
