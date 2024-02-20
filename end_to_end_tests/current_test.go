package end_to_end_tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/tgrindinger/goWeatherService/config"
	"github.com/tgrindinger/goWeatherService/models"
)

func TestCurrent_HappyPath(t *testing.T) {
	cases := []struct {
		desc       string
		latitude   float64
		longitude  float64
		statusCode int64
	}{{
		"kansas city returns valid weather information",
		39.099724,
		-94.578331,
		http.StatusOK,
	}, {
		"london returns valid weather information",
		51.509865,
		-0.118092,
		http.StatusOK,
	}, {
		"maui returns valid weather information",
		20.798363,
		-156.331924,
		http.StatusOK,
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			config, _ := config.ReadConfig("../config/config.json")
			url := fmt.Sprintf("http://localhost%s/current?latitude=%f&longitude=%f", config.Port, tc.latitude, tc.longitude)

			// act
			response, err := http.Get(url)

			// assert
			if err != nil {
				t.Fatalf("error from API: %s", err.Error())
			}
			if response.StatusCode != int(tc.statusCode) {
				t.Fatalf("unexpected status code: got %d want %d", response.StatusCode, tc.statusCode)
			}
			var report models.WeatherReport
			err = json.NewDecoder(response.Body).Decode(&report)
			if err != nil {
				t.Fatalf("error decoding response: %s", err.Error())
			}
			if len(report.Condition) == 0 {
				t.Fatal("condition should be non-empty")
			}
			if len(report.Temperature) == 0 {
				t.Fatal("temperature should be non-empty")
			}
		})
	}
}

func TestCurrent_ValidationErrors(t *testing.T) {
	cases := []struct {
		desc       string
		latitude   string
		longitude  string
		statusCode int64
	}{{
		"missing latitude returns bad request",
		"",
		"-94.578331",
		http.StatusBadRequest,
	}, {
		"missing longitude returns bad request",
		"-0.118092",
		"",
		http.StatusBadRequest,
	}, {
		"invalid latitude returns bad request",
		"120.798363",
		"-156.331924",
		http.StatusBadRequest,
	}, {
		"invalid longitude returns bad request",
		"20.798363",
		"-186.331924",
		http.StatusBadRequest,
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			config, _ := config.ReadConfig("../config/config.json")
			url := fmt.Sprintf("http://localhost%s/current?latitude=%s&longitude=%s", config.Port, tc.latitude, tc.longitude)

			// act
			response, err := http.Get(url)

			// assert
			if err != nil {
				t.Fatalf("error from API: %s", err.Error())
			}
			if response.StatusCode != int(tc.statusCode) {
				t.Fatalf("unexpected status code: got %d want %d", response.StatusCode, tc.statusCode)
			}
		})
	}
}
