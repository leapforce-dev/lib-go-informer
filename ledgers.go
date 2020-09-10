package informer

import (
	"fmt"
)

// Ledger stores Ledger from Informer
//
type Ledger struct {
	ID          string
	Number      string  `json:"number"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
	VATCode     *string `json:"vat_code"`
	Costs       *string `json:"costs"`
	RGS         *string `json:"rgs"`
	Blocked     *string `json:"blocked"`
}

type Ledgers struct {
	Ledgers map[string]Ledger `json:"ledgers"`
}

// GetLedgers returns all ledgers
//
func (i *Informer) GetLedgers() ([]Ledger, error) {
	urlStr := "%sledgers"

	ledgers := Ledgers{}
	ledgers_ := []Ledger{}

	url := fmt.Sprintf(urlStr, i.ApiURL)

	err := i.Get(url, &ledgers)
	if err != nil {
		return nil, err
	}
	for id, ledger := range ledgers.Ledgers {
		ledger.ID = id
		ledgers_ = append(ledgers_, ledger)
	}

	return ledgers_, nil
}
