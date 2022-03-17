package informer

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Relation stores Relation from Informer
//
type Relation struct {
	Id                     string
	RelationNumber         string `json:"relation_number"`
	RelationType           string `json:"relation_type"`
	CompanyName            string `json:"company_name"`
	Firstname              string `json:"firstname"`
	SurnamePrefix          string `json:"surname_prefix"`
	Surname                string `json:"surname"`
	Street                 string `json:"street"`
	HouseNumber            string `json:"house_number"`
	HouseNumberSuffix      string `json:"house_number_suffix"`
	Zip                    string `json:"zip"`
	City                   string `json:"city"`
	Country                string `json:"country"`
	PhoneNumber            string `json:"phone_number"`
	FaxNumber              string `json:"fax_number"`
	Web                    string `json:"web"`
	Email                  string `json:"email"`
	Coc                    string `json:"coc"`
	Vat                    string `json:"vat"`
	Iban                   string `json:"iban"`
	Bic                    string `json:"bic"`
	EmailInvoice           string `json:"email_invoice"`
	SalesInvoiceTemplateId string `json:"sales_invoice_template_id"`
	PaymentConditionId     string `json:"payment_condition_id"`
	//Contacts               map[string]Contact `json:"contacts"`
}

type Relations struct {
	Relations map[string]Relation `json:"relation"`
}

// GetRelations returns all relations
//
func (service *Service) GetRelations() (*[]Relation, *errortools.Error) {
	relations := []Relation{}

	page := 0

	for {
		_relations := Relations{}

		params := url.Values{}
		params.Set("page", fmt.Sprintf("%v", page))

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("relations?%s", params.Encode())),
			ResponseModel: &_relations,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		for relationId, relation := range _relations.Relations {
			relation.Id = relationId

			relations = append(relations, relation)
		}

		if len(_relations.Relations) == 0 {
			break
		}
		page++
	}

	return &relations, nil
}
