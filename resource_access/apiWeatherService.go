package resource_access

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/tgrindinger/goWeatherService/models"
)

type ApiWeatherService struct {
	url    string
	apiKey string
}

func NewApiWeatherService(url string, apiKey string) *ApiWeatherService {
	return &ApiWeatherService{
		url:    url,
		apiKey: apiKey,
	}
}

func (s *ApiWeatherService) CurrentWeather(location models.Location) (*CurrentWeather, error) {
	resp, err := s.executeRequest(location)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		allBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to parse error body: %s", err.Error())
		}
		return nil, fmt.Errorf("api call failed: %s", string(allBytes))
	}

	var report CurrentWeather
	err = json.NewDecoder(resp.Body).Decode(&report)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %s", err.Error())
	}
	return &report, nil
}

func (s *ApiWeatherService) executeRequest(location models.Location) (*http.Response, error) {
	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", location.Latitude))
	params.Add("lon", fmt.Sprintf("%f", location.Longitude))
	params.Add("appid", s.apiKey)
	params.Add("units", "imperial")
	url := s.url + "?" + params.Encode()
	resp, err := http.Get(url)
	return resp, err
}
