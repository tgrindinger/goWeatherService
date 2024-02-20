package resource_access

import (
	"testing"

	"github.com/tgrindinger/goWeatherService/models"
)

func TestStubWeatherService_ReportWeather(t *testing.T) {
	cases := []struct {
		desc      string
		latitude  float64
		longitude float64
		main      string
		temp      float64
	}{{
		"kansas city is overcast and cold",
		39.099724,
		-94.578331,
		"overcast clouds",
		45.41,
	}, {
		"london is overcast and cold",
		51.509865,
		-0.118092,
		"overcast clouds",
		45.1,
	}, {
		"maui has few clouds and is warm",
		20.798363,
		-156.331924,
		"few clouds",
		66.42,
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			service := NewStubWeatherService()
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
			if tc.main != currentWeather.Weather[0].Description {
				t.Fatalf("wrong weather: got %s want %s", currentWeather.Weather[0].Main, tc.main)
			}
			if tc.temp != currentWeather.Main.Temp {
				t.Fatalf("wrong temperature: got %f want %f", currentWeather.Main.Temp, tc.temp)
			}
		})
	}
}
