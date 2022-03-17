package informer

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// SalesInvoice stores SalesInvoice from Informer
//
type SalesInvoice struct {
	Id                 string
	RelationId         string                      `json:"relation_id"`
	ContactId          string                      `json:"contact_id"`
	ContactName        string                      `json:"contact_name"`
	TemplateId         string                      `json:"template_id"`
	PaymentConditionId string                      `json:"payment_condition_id"`
	Number             string                      `json:"number"`
	Date               string                      `json:"date"`
	ExpiryDays         string                      `json:"expiry_days"`
	ExpiryDate         string                      `json:"expiry_date"`
	Expired            string                      `json:"expired"`
	TotalPriceExclTax  string                      `json:"total_price_excl_tax"`
	TotalPriceInclTax  string                      `json:"total_price_incl_tax"`
	Paid_              interface{}                 `json:"paid"`
	Paid               *string                     `json:"-"`
	VatOption          string                      `json:"vat_option"`
	Comment            string                      `json:"comment"`
	FooterText         string                      `json:"footer_text"`
	Reference          string                      `json:"reference"`
	Concept            string                      `json:"concept"`
	ReminderStatus     string                      `json:"reminder_status"`
	LastReminderDate   string                      `json:"last_reminder_date"`
	Attachments        map[string]string           `json:"attachments"`
	Lines              map[string]SalesInvoiceLine `json:"line"`
}

type SalesInvoiceLine struct {
	//Number          string
	Info            string `json:"info"`
	Quantity        string `json:"qty"`
	ProductId       string `json:"product_id"`
	Description     string `json:"description"`
	Amount          string `json:"amount"`
	Discount        string `json:"discount"`
	VatId           string `json:"vat_id"`
	VatPercentage   string `json:"vat_percentage"`
	LedgerAccountId string `json:"ledger_account_id"`
	CostsId         string `json:"costs_id"`
}

type SalesInvoices struct {
	SalesInvoices map[string]SalesInvoice `json:"sales"`
}

// GetSalesInvoices returns all salesInvoices
//
func (service *Service) GetSalesInvoices() (*[]SalesInvoice, *errortools.Error) {
	salesInvoices := []SalesInvoice{}

	page := 0

	for {
		_salesInvoices := SalesInvoices{}

		params := url.Values{}
		params.Set("page", fmt.Sprintf("%v", page))

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("invoices/sales?%s", params.Encode())),
			ResponseModel: &_salesInvoices,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		for salesInvoiceId, salesInvoice := range _salesInvoices.SalesInvoices {
			salesInvoice.Id = salesInvoiceId

			paid_ := fmt.Sprintf("%v", salesInvoice.Paid_)
			if paid_ != "0" {
				//fmt.Println(paid_)
				salesInvoice.Paid = &paid_
			}

			salesInvoices = append(salesInvoices, salesInvoice)
		}

		if len(_salesInvoices.SalesInvoices) == 0 {
			break
		}
		page++
	}

	return &salesInvoices, nil
}
