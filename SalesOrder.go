package informer

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// SalesOrder stores SalesOrder from Informer
//
type SalesOrder struct {
	Id                 string
	RelationId         string                    `json:"relation_id"`
	ContactId          string                    `json:"contact_id"`
	ContactName        string                    `json:"contact_name"`
	TemplateId         string                    `json:"template_id"`
	PaymentConditionId string                    `json:"payment_condition_id"`
	Number             string                    `json:"number"`
	Date               string                    `json:"date"`
	ExpiryDays         string                    `json:"expiry_days"`
	ExpiryDate         string                    `json:"expiry_date"`
	Expired            string                    `json:"expired"`
	TotalPriceExclTax  string                    `json:"total_price_excl_tax"`
	TotalPriceInclTax  string                    `json:"total_price_incl_tax"`
	TaxOption          string                    `json:"tax_option"`
	Discount           string                    `json:"discount"`
	Comment            string                    `json:"comment"`
	FooterText         string                    `json:"footer_text"`
	Reference          string                    `json:"reference"`
	Concept            string                    `json:"concept"`
	Accepted           string                    `json:"accepted"`
	AcceptedByRelation string                    `json:"accepted_by_relation"`
	AcceptedBy         string                    `json:"accepted_by"`
	AcceptedDate       string                    `json:"accepted_date"`
	AcceptedIban       string                    `json:"accepted_iban"`
	AcceptedNote       string                    `json:"accepted_note"`
	AcceptedIp         string                    `json:"accepted_ip"`
	Attachments        map[string]string         `json:"attachments"`
	Lines              map[string]SalesOrderLine `json:"line"`
}

type SalesOrderLine struct {
	//Number          string
	Info            string `json:"info"`
	Quantity        string `json:"qty"`
	ProductId       string `json:"product_id"`
	Description     string `json:"description"`
	Amount          string `json:"amount"`
	Discount        string `json:"discount"`
	TaxId           string `json:"tax_id"`
	TaxPercentage   string `json:"tax_percentage"`
	AmountTax       string `json:"amount_tax"`
	LedgerAccountId string `json:"ledger_account_id"`
	CostsId         string `json:"costs_id"`
}

type SalesOrders struct {
	SalesOrders map[string]SalesOrder `json:"sales"`
}

// GetSalesOrders returns all salesOrders
//
func (service *Service) GetSalesOrders() (*[]SalesOrder, *errortools.Error) {
	salesOrders := []SalesOrder{}

	page := 0

	for {
		_salesOrders := SalesOrders{}

		params := url.Values{}
		params.Set("page", fmt.Sprintf("%v", page))

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("invoices/sales?%s", params.Encode())),
			ResponseModel: &_salesOrders,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		for salesOrderId, salesOrder := range _salesOrders.SalesOrders {
			salesOrder.Id = salesOrderId

			salesOrders = append(salesOrders, salesOrder)
		}

		if len(_salesOrders.SalesOrders) == 0 {
			break
		}
		page++
	}

	return &salesOrders, nil
}
