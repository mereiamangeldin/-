package service

import (
	"errors"
	"github.com/mereiamangeldin/OnePlusTask2/model"
	"github.com/mereiamangeldin/OnePlusTask2/repository"
)

type ICurrencyService interface {
	GetCurrencies() ([]model.Currency, error)
	GetCurrencyById(id string) (model.Currency, error)
}

type Manager struct {
	Currency ICurrencyService
}

func NewManager(repo *repository.Repository) (*Manager, error) {
	if repo == nil {
		return nil, errors.New("no storage given")
	}
	currencySrv := NewCurrencyService(repo)
	return &Manager{Currency: currencySrv}, nil
}
