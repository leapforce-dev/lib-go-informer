package informer

import (
	"fmt"
)

// LedgerEntry stores LedgerEntry from Informer
//
type LedgerEntry struct {
	InvoiceID        string      `json:"invoice_id"`
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

// GetLedgerEntries returns all ledgerEntries
//
func (i *Informer) GetLedgerEntries(ledgerID string, yearFrom int, yearTo int, periodFrom int, periodTo int) ([]LedgerEntry, error) {
	urlStr := "%sreports/ledger?ledger_id=%s&year_from=%v&year_to=%v&period_from=%v&period_to=%v"
	url := fmt.Sprintf(urlStr, i.ApiURL, ledgerID, yearFrom, yearTo, periodFrom, periodTo)
	//fmt.Println(url)

	ledgerEntries := LedgerEntries{}

	err := i.Get(url, &ledgerEntries)
	if err != nil {
		//fmt.Println("page", page)
		return nil, err
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

	return ledgerEntries.LedgerEntries, nil
}
