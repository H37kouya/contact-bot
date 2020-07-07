package service

import "contact-bot/pkg/domain/repository"

// ContactService ContactServiceのためのinterface
type ContactService interface {
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
