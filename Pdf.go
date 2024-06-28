package informer

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Ledger stores Ledger from Informer
type Pdf struct {
	Pdf struct {
		Base64 string `json:"base64"`
	} `json:"pdf"`
}

type PdfType string

const (
	PdfTypeSales    PdfType = "sales"
	PdfTypePurchase PdfType = "purchase"
	PdfTypeReceipt  PdfType = "receipt"
)

// GetPdf returns a base64 encoded pdf file for a certain invoice or receipt
func (service *Service) GetPdf(pdfType PdfType, id string) (*Pdf, *errortools.Error) {
	pdf := Pdf{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("pdf/%v/%s", pdfType, id)),
		ResponseModel: &pdf,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &pdf, nil
}
