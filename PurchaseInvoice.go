package informer

import (
	"fmt"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// PurchaseInvoice stores PurchaseInvoice from Service
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
func (service *Service) GetPurchaseInvoices() (*[]PurchaseInvoice, *errortools.Error) {
	purchaseInvoices := []PurchaseInvoice{}

	page := 0

	for {
		_purchaseInvoices := PurchaseInvoices{}

		params := url.Values{}
		params.Set("page", fmt.Sprintf("%v", page))

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("invoices/purchase?%s", params.Encode())),
			ResponseModel: &_purchaseInvoices,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		for purchaseInvoiceID, purchaseInvoice := range _purchaseInvoices.PurchaseInvoices {
			purchaseInvoice.ID = purchaseInvoiceID

			paid_ := fmt.Sprintf("%v", purchaseInvoice.Paid_)
			if paid_ != "0" {
				//fmt.Println(paid_)
				purchaseInvoice.Paid = &paid_
			}

			purchaseInvoices = append(purchaseInvoices, purchaseInvoice)
		}

		if len(_purchaseInvoices.PurchaseInvoices) == 0 {
			break
		}
		page++
	}

	return &purchaseInvoices, nil
}
