package informer

import (
	"fmt"
)

// PurchaseInvoice stores PurchaseInvoice from Informer
//
type PurchaseInvoice struct {
	ID                string
	RelationID        string                         `json:"relation_id"`
	Number            string                         `json:"number"`
	Date              string                         `json:"date"`
	TotalPriceExclTax string                         `json:"total_price_excl_tax"`
	TotalPriceInclTax string                         `json:"total_price_incl_tax"`
	VATAmount         string                         `json:"vat_amount"`
	VATOption         string                         `json:"vat_option"`
	Exported          string                         `json:"exported"`
	ExportDate        string                         `json:"export_date"`
	ExpiryDate        string                         `json:"expiry_date"`
	Paid_             interface{}                    `json:"paid"`
	Paid              *string                        `json:"-"`
	Lines             map[string]PurchaseInvoiceLine `json:"line"`
}

type PurchaseInvoiceLine struct {
	Description     string `json:"description"`
	Amount          string `json:"amount"`
	VATID           string `json:"vat_id"`
	VATPercentage   string `json:"vat_percentage"`
	LedgerAccountID string `json:"ledger_account_id"`
	CostsID         string `json:"costs_id"`
}

type PurchaseInvoices struct {
	PurchaseInvoices map[string]PurchaseInvoice `json:"purchase"`
}

// GetPurchaseInvoices returns all purchaseInvoices
//
func (i *Informer) GetPurchaseInvoices() ([]PurchaseInvoice, error) {
	urlStr := "%sinvoices/purchase?page=%v"

	purchaseInvoices_ := []PurchaseInvoice{}

	page := 0
	rowCount := 0

	for rowCount > 0 || page == 0 {
		url := fmt.Sprintf(urlStr, i.ApiURL, page)

		purchaseInvoices := PurchaseInvoices{}

		err := i.Get(url, &purchaseInvoices)
		if err != nil {
			//fmt.Println("page", page)
			return nil, err
		}
		for purchaseInvoiceID, purchaseInvoice := range purchaseInvoices.PurchaseInvoices {
			purchaseInvoice.ID = purchaseInvoiceID

			paid_ := fmt.Sprintf("%v", purchaseInvoice.Paid_)
			if paid_ != "0" {
				//fmt.Println(paid_)
				purchaseInvoice.Paid = &paid_
			}

			purchaseInvoices_ = append(purchaseInvoices_, purchaseInvoice)
		}

		rowCount = len(purchaseInvoices.PurchaseInvoices)
		page++
	}

	return purchaseInvoices_, nil
}
