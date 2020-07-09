package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/repository"
)

// ContactService ContactServiceのためのinterface
type ContactService interface {
	GetContactData() ([]model.Contact, error)
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

func (cs contactService) GetContactData() ([]model.Contact, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	contacts, err := cs.contactRepository.GetContactSheet(conf.SpreadSheet.ID)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}
