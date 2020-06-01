package accuweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	clientTimeoutSeconds = 2

	baseUrl              = "http://dataservice.accuweather.com/"
	currentConditionsUrl = baseUrl + "currentconditions/v1/%s?apikey=%s&language=it-it&details=true"
)

type AccuweatherClient interface {
	GetCurrentConditions(locationKey string) (*CurrentConditions, error)
}

type accuweatherClient struct {
	apiKey string
}

func NewAccuweatherClient(apiKey string) AccuweatherClient {
	return &accuweatherClient{
		apiKey: apiKey,
	}
}

func (ac *accuweatherClient) GetCurrentConditions(locationKey string) (*CurrentConditions, error) {
	url := fmt.Sprintf(currentConditionsUrl, locationKey, ac.apiKey)

	awClient := http.Client{
		Timeout: time.Second * clientTimeoutSeconds,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := awClient.Do(req)
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
