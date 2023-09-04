package informer

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	t_types "github.com/leapforce-libraries/go_informer/types"
	go_types "github.com/leapforce-libraries/go_types"
)

type Subscriptions struct {
	Subscriptions map[go_types.Int64String]Subscription `json:"subscriptions"`
}

type Subscription struct {
	Id                      go_types.Int64String
	RelationId              go_types.Int64String                      `json:"relation_id"`
	ContactId               go_types.Int64String                      `json:"contact_id"`
	ContactName             string                                    `json:"contact_name"`
	TemplateId              go_types.Int64String                      `json:"template_id"`
	PaymentConditionId      go_types.Int64String                      `json:"payment_condition_id"`
	Number                  go_types.Int64String                      `json:"number"`
	Date                    t_types.DateString                        `json:"date"`
	ExpiryDays              go_types.Int64String                      `json:"expiry_days"`
	ExpiryDate              *t_types.DateString                       `json:"expiry_date"`
	Expired                 go_types.BoolString                       `json:"expired"`
	TotalPriceExclVat       go_types.Float64String                    `json:"total_price_excl_vat"`
	TotalPriceInclVat       go_types.Float64String                    `json:"total_price_incl_vat"`
	VatOption               string                                    `json:"vat_option"`
	Comment                 string                                    `json:"comment"`
	FooterText              string                                    `json:"footer_text"`
	Reference               string                                    `json:"reference"`
	Concept                 go_types.BoolString                       `json:"concept"`
	SubscriptionTypeId      go_types.Int64String                      `json:"subscription_type_id"`
	SubscriptionFrequency   string                                    `json:"subscription_frequency"`
	SubscriptionStartDate   t_types.DateString                        `json:"subscription_start_date"`
	SubscriptionInvoiceDate *t_types.DateString                       `json:"subscription_invoice_date"`
	SubscriptionEndDate     *t_types.DateString                       `json:"subscription_end_date"`
	SubscriptionRestriction string                                    `json:"subscription_restriction"`
	SubscriptionTimes       go_types.Int64String                      `json:"subscription_times"`
	SubscriptionSend        go_types.BoolString                       `json:"subscription_send"`
	ReminderStatus          string                                    `json:"reminder_status"`
	LastReminderDate        *t_types.DateString                       `json:"last_reminder_date"`
	InvoiceUrl              string                                    `json:"invoice_url"`
	Lines                   map[go_types.Int64String]SubscriptionLine `json:"line"`
}

type SubscriptionLine struct {
	Info            go_types.Int64String   `json:"info"`
	Qty             go_types.Float64String `json:"qty"`
	ProductId       go_types.Int64String   `json:"product_id"`
	Description     string                 `json:"description"`
	Amount          go_types.Float64String `json:"amount"`
	Discount        go_types.Float64String `json:"discount"`
	VatId           go_types.Int64String   `json:"vat_id"`
	VatPercentage   go_types.Float64String `json:"vat_percentage"`
	LedgerAccountId go_types.Int64String   `json:"ledger_account_id"`
	CostsId         go_types.Int64String   `json:"costs_id"`
}

// GetSubscriptions returns all subscriptions
func (service *Service) GetSubscriptions() (*[]Subscription, *errortools.Error) {
	subscriptions := []Subscription{}

	page := 0

	for {
		_subscriptions := Subscriptions{}

		params := url.Values{}
		params.Set("page", fmt.Sprintf("%v", page))

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           service.url(fmt.Sprintf("subscriptions?%s", params.Encode())),
			ResponseModel: &_subscriptions,
		}
		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		for subscriptionId, subscription := range _subscriptions.Subscriptions {
			subscription.Id = subscriptionId

			subscriptions = append(subscriptions, subscription)
		}

		if len(_subscriptions.Subscriptions) == 0 {
			break
		}
		page++
	}

	return &subscriptions, nil
}
