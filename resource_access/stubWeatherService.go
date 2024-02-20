package resource_access

import (
	"encoding/json"
	"strings"

	"github.com/tgrindinger/goWeatherService/models"
)

type StubWeatherService struct {
}

func NewStubWeatherService() *StubWeatherService {
	return &StubWeatherService{}
}

func (s *StubWeatherService) CurrentWeather(location models.Location) (*CurrentWeather, error) {
	cannedResponse := s.cannedResponse(location)
	var report CurrentWeather
	err := json.NewDecoder(strings.NewReader(cannedResponse)).Decode(&report)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

func (s *StubWeatherService) cannedResponse(location models.Location) string {
	const kcText = `{"coord":{"lon":-94.5781,"lat":39.1001},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04n"}],"base":"stations","main":{"temp":45.41,"feels_like":41.83,"temp_min":40.41,"temp_max":47.55,"pressure":1018,"humidity":55,"sea_level":1018,"grnd_level":985},"visibility":10000,"wind":{"speed":6.69,"deg":178,"gust":16.71},"clouds":{"all":100},"dt":1708394029,"sys":{"type":2,"id":2025651,"country":"US","sunrise":1708347935,"sunset":1708387146},"timezone":-21600,"id":4393217,"name":"Kansas City","cod":200}`
	const londonText = `{"coord":{"lon":-0.1181,"lat":51.5099},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04n"}],"base":"stations","main":{"temp":45.1,"feels_like":41.85,"temp_min":40.75,"temp_max":47.32,"pressure":1029,"humidity":89},"visibility":10000,"wind":{"speed":5.99,"deg":231,"gust":7},"clouds":{"all":100},"dt":1708394767,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1708412765,"sunset":1708449767},"timezone":0,"id":2643743,"name":"London","cod":200}`
	const mauiText = `{"coord":{"lon":-156.3319,"lat":20.7984},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"base":"stations","main":{"temp":66.42,"feels_like":65.8,"temp_min":64.22,"temp_max":69.58,"pressure":1017,"humidity":65},"visibility":10000,"wind":{"speed":20.71,"deg":60,"gust":29.93},"clouds":{"all":20},"dt":1708394813,"sys":{"type":2,"id":18862,"country":"US","sunrise":1708361589,"sunset":1708403132},"timezone":-36000,"id":5852697,"name":"Pukalani","cod":200}`
	london := models.Location{
		Latitude:  51.509865,
		Longitude: -0.118092,
	}
	maui := models.Location{
		Latitude:  20.798363,
		Longitude: -156.331924,
	}
	kc := models.Location{
		Latitude:  39.1001,
		Longitude: -94.5781,
	}
	londonDist := distanceSquared(location, london)
	mauiDist := distanceSquared(location, maui)
	kcDist := distanceSquared(location, kc)
	if londonDist < mauiDist && londonDist < kcDist {
		return londonText
	}
	if mauiDist < londonDist && mauiDist < kcDist {
		return mauiText
	}
	return kcText
}

func distanceSquared(location1 models.Location, location2 models.Location) float64 {
	latDist := location1.Latitude - location2.Latitude
	latDist *= latDist
	longDist := location1.Longitude - location2.Longitude
	longDist *= longDist
	return latDist*latDist + longDist*longDist
}
