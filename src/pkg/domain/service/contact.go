package service

import (
	"contact-bot/config"
	"contact-bot/pkg/domain/model"
	"contact-bot/pkg/domain/repository"
	"contact-bot/pkg/helper"
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
	spreadSheetConf := config.Conf.SpreadSheet

	contacts, err := cs.contactRepository.GetContactSheet(spreadSheetConf.ID)
	if err != nil {
		return nil, err
	}

	// 情報取得の終了時刻 現在時刻を時間で切り捨てたもの()
	endTime := helper.GetHourTime(helper.GetNowTokyoTime())
	// 情報取得の開始時刻
	startTime := helper.GetBeforeHourTime(pollingDiff, endTime)
	// 時間によってフィルターする
	filterContacts := model.FilterContactByTimeStamp(contacts, startTime, endTime)

	return filterContacts, nil
}
