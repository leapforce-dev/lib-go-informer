package informer

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Currency stores Currency from Service
//
type Currency struct {
	ID          string
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Rate        string `json:"rate"`
	AutoUpdate  string `json:"autoupdate"`
	JournalID   string `json:"journal_id "`
	LedgerID    string `json:"ledger_id"`
	BankID      string `json:"bank_id"`
}

type Currencies struct {
	Currencies map[string]Currency `json:"currencies"`
}

// GetCurrencies returns all currencies
//
func (service *Service) GetCurrencies() (*[]Currency, *errortools.Error) {
	currencies := Currencies{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		URL:           service.url("currencies"),
		ResponseModel: &currencies,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	_currencies := []Currency{}
	for id, currency := range currencies.Currencies {
		currency.ID = id
		_currencies = append(_currencies, currency)
	}

	return &_currencies, nil
}
