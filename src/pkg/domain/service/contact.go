package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/repository"
	"fmt"
)

// ContactService ContactServiceのためのinterface
type ContactService interface {
	GetContactData()
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

func (cs contactService) GetContactData() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	contacts, err := cs.contactRepository.GetContactSheet(conf.SpreadSheet.ID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contacts)
}
