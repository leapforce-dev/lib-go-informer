package informer

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// SubscriptionType stores SubscriptionType from Service
type SubscriptionType struct {
	Id          string
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SubscriptionTypes struct {
	SubscriptionTypes map[string]SubscriptionType `json:"subscription_types"`
}

// GetSubscriptionTypes returns all subscription types
func (service *Service) GetSubscriptionTypes() (*[]SubscriptionType, *errortools.Error) {
	subscriptionTypes := SubscriptionTypes{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("subscription-types"),
		ResponseModel: &subscriptionTypes,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	_subscriptionTypes := []SubscriptionType{}
	for id, currency := range subscriptionTypes.SubscriptionTypes {
		currency.Id = id
		_subscriptionTypes = append(_subscriptionTypes, currency)
	}

	return &_subscriptionTypes, nil
}
