package informer

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Ledger stores Ledger from Informer
//
type Ledger struct {
	Id          string
	Number      string  `json:"number"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	VatCode     *string `json:"vat_code"`
	Costs       *string `json:"costs"`
	Rgs         *string `json:"rgs"`
	Blocked     *string `json:"blocked"`
}

type Ledgers struct {
	Ledgers map[string]Ledger `json:"ledgers"`
}

// GetLedgers returns all ledgers
//
func (service *Service) GetLedgers() (*[]Ledger, *errortools.Error) {
	ledgers := Ledgers{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("ledgers"),
		ResponseModel: &ledgers,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	_ledgers := []Ledger{}
	for id, ledger := range ledgers.Ledgers {
		ledger.Id = id
		_ledgers = append(_ledgers, ledger)
	}

	return &_ledgers, nil
}
