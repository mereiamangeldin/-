package web

import (
	"encoding/json"
	"errors"
	"github.com/mereiamangeldin/OnePlusTask2/model"
	"io"
	"log"
	"net/http"
	"strings"
)

type CurrencyRepository struct {
	url string
}

func NewCurrencyRepository(url string) *CurrencyRepository {
	return &CurrencyRepository{url: url}
}

func (r *CurrencyRepository) GetCurrencies() ([]model.Currency, error) {
	resp, err := http.Get(r.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var currencies []model.Currency
	err = json.Unmarshal([]byte(body), &currencies)
	if err != nil {
		log.Println(string(body))
		return nil, err
	}
	return currencies, nil
}

func (r *CurrencyRepository) GetCurrencyById(id string) (model.Currency, error) {
	currencies, err := r.GetCurrencies()
	if err != nil {
		return model.Currency{}, err
	}
	for _, c := range currencies {
		if strings.ToLower(c.Id) == strings.ToLower(id) {
			return c, nil
		}
	}
	return model.Currency{}, errors.New("no such currency")
}
