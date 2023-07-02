package handler

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Manager) GetCurrencies(c echo.Context) error {
	currencies, err := h.srv.Currency.GetCurrencies()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, currencies)
}

func (h *Manager) GetCurrencyById(c echo.Context) error {
	id := c.Param("currency_id")
	currency, err := h.srv.Currency.GetCurrencyById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, currency)
}
