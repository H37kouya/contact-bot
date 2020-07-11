package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/repository"
	"contact-bot/pkg/helper"
	"os"
	"strconv"
	"time"
)

const (
	defaultDiff int = 1
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
	conf := config.NewConfig()

	contacts, err := cs.contactRepository.GetContactSheet(conf.SpreadSheet.ID)
	if err != nil {
		return nil, err
	}

	now := helper.GetHourTime(helper.GetNowTokyoTime())
	diff := getPollingDiff()
	filterContacts := model.FilterContactByTimeStamp(contacts, getBeforeTime(diff, now), now)

	return filterContacts, nil
}

func getBeforeTime(diff int, t time.Time) time.Time {
	return t.Add(time.Duration(-1*diff) * time.Hour)
}

func getPollingDiff() int {
	str := os.Getenv("POLLING_DIFF")
	if str == "" {
		return defaultDiff
	}

	diff, _ := strconv.Atoi(str)
	return diff
}
