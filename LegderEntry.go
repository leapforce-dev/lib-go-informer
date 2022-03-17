package informer

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// LedgerEntry stores LedgerEntry from Service
//
type LedgerEntry struct {
	InvoiceId        string      `json:"invoice_id"`
	Number           string      `json:"number"`
	Type             string      `json:"type"`
	Date             string      `json:"date"`
	Period           string      `json:"period"`
	Year             string      `json:"year"`
	Costs            *string     `json:"costs"`
	EntryDescription string      `json:"entry_description"`
	LineDescription  string      `json:"line_description"`
	Debit_           interface{} `json:"debit"`
	Debit            *string     `json:"-"`
	Credit_          interface{} `json:"credit"`
	Credit           *string     `json:"-"`
}

type LedgerEntries struct {
	LedgerEntries []LedgerEntry `json:"ledger_entries"`
}

type GetLedgerEntriesConfig struct {
	LedgerId   string
	YearFrom   int
	YearTo     int
	PeriodFrom int
	PeriodTo   int
}

// GetLedgerEntries returns all ledgerEntries
//
func (service *Service) GetLedgerEntries(config *GetLedgerEntriesConfig) (*[]LedgerEntry, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("GetLedgerEntriesConfig must not be nill")
	}

	params := url.Values{}
	params.Set("ledger_id", config.LedgerId)
	params.Set("year_from", fmt.Sprintf("%v", config.YearFrom))
	params.Set("year_to", fmt.Sprintf("%v", config.YearTo))
	params.Set("period_from", fmt.Sprintf("%v", config.PeriodFrom))
	params.Set("period_to", fmt.Sprintf("%v", config.PeriodTo))

	ledgerEntries := LedgerEntries{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("reports/ledger?%s", params.Encode())),
		ResponseModel: &ledgerEntries,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	for i := range ledgerEntries.LedgerEntries {
		debit_ := fmt.Sprintf("%v", ledgerEntries.LedgerEntries[i].Debit_)
		if debit_ != "0" {
			ledgerEntries.LedgerEntries[i].Debit = &debit_
		}

		credit_ := fmt.Sprintf("%v", ledgerEntries.LedgerEntries[i].Credit_)
		if credit_ != "0" {
			ledgerEntries.LedgerEntries[i].Credit = &credit_
		}
	}

	return &ledgerEntries.LedgerEntries, nil
}
