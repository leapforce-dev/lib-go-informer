package informer

import (
	"fmt"
)

// Currency stores Currency from Informer
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
func (i *Informer) GetCurrencies() ([]Currency, error) {
	urlStr := "%scurrencies"

	currencies := Currencies{}
	currencies_ := []Currency{}

	url := fmt.Sprintf(urlStr, i.ApiURL)

	err := i.Get(url, &currencies)
	if err != nil {
		return nil, err
	}
	for id, currency := range currencies.Currencies {
		currency.ID = id
		currencies_ = append(currencies_, currency)
	}

	return currencies_, nil
}
