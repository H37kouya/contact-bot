package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/repository"

	"google.golang.org/api/sheets/v4"
)

// ContactService ContactServiceのためのinterface
type ContactService interface {
	GetContactData() (*sheets.Spreadsheet, error)
}

type contactService struct {
	contactRepository repository.ContactRepository
}

// NewContactService ContactService DIするために必要
func NewContactService(cr repository.ContactRepository) ContactService {
	return &contactService{
		contactRepository: cr,
	}
}

func (cs contactService) GetContactData() (*sheets.Spreadsheet, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	spreadSheet, err := cs.contactRepository.GetContactSheet(conf.SpreadSheet.ID)
	if err != nil {
		return nil, err
	}

	return spreadSheet, nil
}
