package managers

import (
	"testing"

	"github.com/tgrindinger/goWeatherService/models"
	resourceAccess "github.com/tgrindinger/goWeatherService/resource_access"
)

func TestReportWeather(t *testing.T) {
	cases := []struct {
		desc        string
		latitude    float64
		longitude   float64
		condition   string
		temperature string
	}{{
		"london is cold and overcast",
		51.509865,
		-0.118092,
		"overcast clouds",
		"cold",
	}, {
		"maui is warm and has a few clouds",
		20.798363,
		-156.331924,
		"few clouds",
		"moderate",
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			manager := NewWeatherManager(resourceAccess.NewStubWeatherService())
			location := models.Location{
				Latitude:  tc.latitude,
				Longitude: tc.longitude,
			}

			// act
			weatherReport, _ := manager.ReportWeather(location)

			// assert
			if tc.condition != weatherReport.Condition {
				t.Fatalf("wrong condition: got %s want %s", weatherReport.Condition, tc.condition)
			}
			if tc.temperature != weatherReport.Temperature {
				t.Fatalf("wrong temperature: got %s want %s", weatherReport.Temperature, tc.temperature)
			}
		})
	}
}
