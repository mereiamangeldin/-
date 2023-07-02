package service

import (
	"github.com/mereiamangeldin/OnePlusTask2/model"
	"github.com/mereiamangeldin/OnePlusTask2/repository"
)

type CurrencyService struct {
	Repo *repository.Repository
}

func NewCurrencyService(repo *repository.Repository) *CurrencyService {
	return &CurrencyService{Repo: repo}
}

func (s *CurrencyService) GetCurrencies() ([]model.Currency, error) {
	return s.Repo.Currency.GetCurrencies()
}

func (s *CurrencyService) GetCurrencyById(id string) (model.Currency, error) {
	return s.Repo.Currency.GetCurrencyById(id)
}
