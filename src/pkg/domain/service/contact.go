package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/repository"
	"contact-bot/pkg/helper"
	"time"
)

const (
	defaultDiff int = 1
)

// ContactService ContactServiceのためのinterface
type ContactService interface {
	GetContactData(pollingDiff int) ([]model.Contact, error)
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

// GetContactData Contactのデータを取得する
func (cs contactService) GetContactData(pollingDiff int) ([]model.Contact, error) {
	conf := config.NewConfig()

	contacts, err := cs.contactRepository.GetContactSheet(conf.SpreadSheet.ID)
	if err != nil {
		return nil, err
	}

	// 現在時刻の取得
	now := helper.GetHourTime(helper.GetNowTokyoTime())
	// 時間によってフィルターする
	filterContacts := model.FilterContactByTimeStamp(contacts, getBeforeTime(pollingDiff, now), now)

	return filterContacts, nil
}

func getBeforeTime(diff int, t time.Time) time.Time {
	return t.Add(time.Duration(-1*diff) * time.Hour)
}
