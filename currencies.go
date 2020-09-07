package informer

import (
	"fmt"

	sentry "github.com/getsentry/sentry-go"
)

// Currency stores Currency from Informer
//
type Currency struct {
	ID          string
	Currency    string `json:"currency"`
	Description string `json:"description"`
	Rate        string `json:"rate"`
	AutoUpdate  string `json:"autoupdate"`
	JournalID   string `json:"journal_id"`
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

	currencies := []Currency{}

	url := fmt.Sprintf(urlStr, i.ApiURL)
	//fmt.Println(url)

	_, response, err := i.Get(url, &currencies)
	if err != nil {
		return nil, err
	}

	if response != nil {
		if response.Errors != nil {
			for _, e := range *response.Errors {
				message := fmt.Sprintf("Error in %v: %v", url, e.Message)
				if i.IsLive {
					sentry.CaptureMessage(message)
				}
				fmt.Println(message)
			}
		}
	}

	return currencies, nil
}
