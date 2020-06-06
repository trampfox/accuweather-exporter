package accuweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	clientTimeoutSeconds = 2

	baseUrl              = "http://dataservice.accuweather.com/"
	currentConditionsUrl = baseUrl + "currentconditions/v1/%s?apikey=%s&details=true"
	searchLocationUrl    = baseUrl + "locations/v1/cities/search?apikey=%s&q=%s"
)

type AccuweatherClient interface {
	GetLocation(location string) (*Location, error)
	GetCurrentConditions(cityID string) (*CurrentConditions, error)
}

type accuweatherClient struct {
	apiKey     string
	httpClient http.Client
}

func NewAccuweatherClient(apiKey string) AccuweatherClient {
	return &accuweatherClient{
		apiKey: apiKey,
		httpClient: http.Client{
			Timeout: time.Second * clientTimeoutSeconds,
		},
	}
}

func (ac *accuweatherClient) GetCurrentConditions(locationKey string) (*CurrentConditions, error) {
	url := fmt.Sprintf(currentConditionsUrl, locationKey, ac.apiKey)

	res, err := ac.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	currentConditions := &CurrentConditions{}
	err = json.Unmarshal(body, currentConditions)
	if err != nil {
		return nil, err
	}

	return currentConditions, nil
}

func (ac *accuweatherClient) GetLocation(location string) (*Location, error) {
	q := url.QueryEscape(location)
	url := fmt.Sprintf(searchLocationUrl, ac.apiKey, q)

	res, err := ac.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	locationResults := Locations{}
	err = json.Unmarshal(body, &locationResults)
	if err != nil {
		return nil, err
	}

	if len(locationResults) > 0 {
		return (locationResults)[0], nil
	}

	return nil, nil
}
