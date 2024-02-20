package resource_access

import (
	"testing"

	"github.com/tgrindinger/goWeatherService/config"
	"github.com/tgrindinger/goWeatherService/models"
)

func TestApiWeatherService_ReportWeather(t *testing.T) {
	cases := []struct {
		desc      string
		latitude  float64
		longitude float64
		name      string
		timezone  int64
	}{{
		"kansas city returns correct metadata",
		39.099724,
		-94.578331,
		"Kansas City",
		-21600,
	}, {
		"london returns correct metadata",
		51.509865,
		-0.118092,
		"London",
		0,
	}, {
		"maui returns correct metadata",
		20.798363,
		-156.331924,
		"Pukalani",
		-36000,
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			config, _ := config.ReadConfig("../config/config.json")
			service := NewApiWeatherService(config.ApiUrl, config.ApiKey)
			location := models.Location{
				Latitude:  tc.latitude,
				Longitude: tc.longitude,
			}

			// act
			currentWeather, err := service.CurrentWeather(location)

			// assert
			if err != nil {
				t.Fatalf("error: %s", err.Error())
			}
			if tc.name != currentWeather.Name {
				t.Fatalf("wrong location: got %s want %s", currentWeather.Name, tc.name)
			}
			if tc.timezone != int64(currentWeather.Timezone) {
				t.Fatalf("wrong timezone: got %d want %d", currentWeather.Timezone, tc.timezone)
			}
		})
	}
}
