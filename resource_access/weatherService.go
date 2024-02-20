package resource_access

import (
	"github.com/tgrindinger/goWeatherService/models"
)

type WeatherService interface {
	CurrentWeather(location models.Location) (*CurrentWeather, error)
}
