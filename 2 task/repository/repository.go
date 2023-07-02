package repository

import (
	"github.com/mereiamangeldin/OnePlusTask2/model"
	"github.com/mereiamangeldin/OnePlusTask2/repository/web"
)

type ICurrencyRepository interface {
	GetCurrencies() ([]model.Currency, error)
	GetCurrencyById(id string) (model.Currency, error)
}

type Repository struct {
	Currency ICurrencyRepository
}

func New(url string) (*Repository, error) {
	currency := web.NewCurrencyRepository(url)
	return &Repository{Currency: currency}, nil
}
