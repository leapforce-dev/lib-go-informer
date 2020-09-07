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

// InformerError contains error info
//
type InformerError struct {
	Error []string `json:"error"`
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
func (i *Informer) Get(url string, model interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("ApiKey", i.ApiKey)
	req.Header.Set("SecurityCode", i.SecurityCode)

	// Send out the HTTP request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &model)
	if err != nil {
		return err
	}

	return nil
}
