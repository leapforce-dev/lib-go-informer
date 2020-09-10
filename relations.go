package informer

import (
	"fmt"
)

// Relation stores Relation from Informer
//
type Relation struct {
	ID                     string
	RelationNumber         string `json:"relation_number"`
	RelationType           string `json:"relation_type"`
	CompanyName            string `json:"company_name"`
	Firstname              string `json:"firstname"`
	SurnamePrefix          string `json:"surname_prefix"`
	Surname                string `json:"surname"`
	Street                 string `json:"street"`
	HouseNumber            string `json:"house_number"`
	HouseNumberSuffix      string `json:"house_number_suffix"`
	ZIP                    string `json:"zip"`
	City                   string `json:"city"`
	Country                string `json:"country"`
	PhoneNumber            string `json:"phone_number"`
	FaxNumber              string `json:"fax_number"`
	Web                    string `json:"web"`
	Email                  string `json:"email"`
	COC                    string `json:"coc"`
	VAT                    string `json:"vat"`
	IBAN                   string `json:"iban"`
	BIC                    string `json:"bic"`
	EmailInvoice           string `json:"email_invoice"`
	SalesInvoiceTemplateID string `json:"sales_invoice_template_id"`
	PaymentConditionID     string `json:"payment_condition_id"`
	//Contacts               map[string]Contact `json:"contacts"`
}

type Relations struct {
	Relations map[string]Relation `json:"relation"`
}

// GetRelations returns all relations
//
func (i *Informer) GetRelations() ([]Relation, error) {
	urlStr := "%srelations?page=%v"

	relations_ := []Relation{}

	page := 0
	rowCount := 0

	for rowCount > 0 || page == 0 {
		url := fmt.Sprintf(urlStr, i.ApiURL, page)

		relations := Relations{}

		err := i.Get(url, &relations)
		if err != nil {
			fmt.Println("page", page)
			return nil, err
		}
		for relationID, relation := range relations.Relations {
			relation.ID = relationID
			relations_ = append(relations_, relation)
		}

		rowCount = len(relations.Relations)
		page++
	}

	return relations_, nil
}
