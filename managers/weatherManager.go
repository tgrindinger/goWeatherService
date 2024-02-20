package managers

import (
	"strings"

	"github.com/tgrindinger/goWeatherService/models"
	resourceAccess "github.com/tgrindinger/goWeatherService/resource_access"
)

type WeatherManager struct {
	weatherService resourceAccess.WeatherService
}

func NewWeatherManager(weatherService resourceAccess.WeatherService) *WeatherManager {
	return &WeatherManager{
		weatherService: weatherService,
	}
}

func (m *WeatherManager) ReportWeather(location models.Location) (models.WeatherReport, error) {
	current, err := m.weatherService.CurrentWeather(location)
	if err != nil {
		return models.WeatherReport{}, err
	}

	condition := describeConditions(current.Weather)
	temperature := describeTemperature(current.Main.Temp)
	report := models.WeatherReport{
		Condition:   condition,
		Temperature: temperature,
	}
	return report, nil
}

func describeConditions(conditions []resourceAccess.Weather) string {
	var descriptions []string
	for _, condition := range conditions {
		descriptions = append(descriptions, condition.Description)
	}
	return strings.Join(descriptions, ", ")
}

func describeTemperature(degrees float64) string {
	var temperature string
	if degrees < 60 {
		temperature = "cold"
	} else if degrees < 80 {
		temperature = "moderate"
	} else {
		temperature = "hot"
	}
	return temperature
}
