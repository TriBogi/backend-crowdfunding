package transaction

import (
	"bogistartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userUD int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}

}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaign.User_ID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil

}

func (s *service) GetTransactionsByUserID(userUD int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userUD)
	if err != nil {
		return transactions, err
	}
	return transactions, nil

}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Campaign_Id = input.CampaignID
	transaction.Amount = input.Amount
	transaction.User_Id = input.User.ID
	transaction.Status = "Pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil

}
