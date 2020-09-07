package informer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	types "github.com/Leapforce-nl/go_types"
)

// type
//
type Informer struct {
	ApiURL       string
	ApiKey       string
	SecurityCode string
	IsLive       bool
}

// Response represents highest level of exactonline api response
//
type Response struct {
	Data     *json.RawMessage `json:"data,omitempty"`
	NextPage *NextPage        `json:"next_page,omitempty"`
	Errors   *[]InformerError `json:"errors,omitempty"`
}

// NextPage contains info for batched data retrieval
//
type NextPage struct {
	Offset string `json:"offset"`
	Path   string `json:"path"`
	URI    string `json:"uri"`
}

// InformerError contains error info
//
type InformerError struct {
	Message string `json:"message"`
	Help    string `json:"help"`
}

func New(apiURL string, apiKey string, securityCode string, isLive bool) (*Informer, error) {
	i := new(Informer)

	if apiURL == "" {
		return nil, &types.ErrorString{"Informer ApiUrl not provided"}
	}
	if apiKey == "" {
		return nil, &types.ErrorString{"ApiKey not provided"}
	}
	if securityCode == "" {
		return nil, &types.ErrorString{"SecurityCode not provided"}
	}

	i.ApiURL = apiURL
	i.ApiKey = apiKey
	i.SecurityCode = securityCode
	i.IsLive = isLive

	if !strings.HasSuffix(i.ApiURL, "/") {
		i.ApiURL = i.ApiURL + "/"
	}

	return i, nil
}

// generic Get method
//
func (i *Informer) Get(url string, model interface{}) (*NextPage, *Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("ApiKey", i.ApiKey)
	req.Header.Set("SecurityCode", i.SecurityCode)

	// Send out the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	response := Response{}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, &response, err
	}

	if response.Data == nil {
		return nil, &response, nil
	}

	err = json.Unmarshal(*response.Data, &model)
	if err != nil {
		return nil, &response, err
	}

	return response.NextPage, &response, nil
}
