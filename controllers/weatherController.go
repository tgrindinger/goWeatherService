package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	managers "github.com/tgrindinger/goWeatherService/managers"
	"github.com/tgrindinger/goWeatherService/models"
)

type WeatherController struct {
	manager *managers.WeatherManager
}

func NewWeatherController(weatherManager *managers.WeatherManager) *WeatherController {
	return &WeatherController{
		manager: weatherManager,
	}
}

// @Summary Get current weather report
// @Description Get current weather report
// @Tags Weather
// @Produce json
// @Param latitude query float64 true "Latitude (values between -90 and +90)"
// @Param longitude query float64 true "Longitude (values between -180 and +180)"
// @Success 200 {object} managers.WeatherReport
// @Router /current [get]
func (c *WeatherController) CurrentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	const latitudeLabel = "latitude"
	const longitudeLabel = "longitude"
	latitudeStr := r.URL.Query().Get(latitudeLabel)
	longitudeStr := r.URL.Query().Get(longitudeLabel)
	location, errors := sanitizeParams(latitudeStr, longitudeStr)

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string][]string{"errors": errors})
		return
	}

	report, err := c.manager.ReportWeather(location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Unexpected error: %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(report)
}

func sanitizeParams(latitudeStr string, longitudeStr string) (models.Location, []string) {
	const MAX_LATITUDE = 90
	const MAX_LONGITUDE = 180
	const latitudeLabel = "latitude"
	const longitudeLabel = "longitude"
	errors := []string{}
	latitude, errors := validateValue(latitudeStr, float64(MAX_LATITUDE), latitudeLabel, errors)
	longitude, errors := validateValue(longitudeStr, float64(MAX_LONGITUDE), longitudeLabel, errors)
	return models.Location{Latitude: latitude, Longitude: longitude}, errors
}

func validateValue(value string, maxValue float64, name string, errors []string) (float64, []string) {
	actual, err := strconv.ParseFloat(value, 64)
	if err != nil {
		errors = append(errors, fmt.Sprintf("%s must be numeric, got '%s'", name, value))
	}
	if math.Abs(actual) > maxValue {
		errors = append(errors, fmt.Sprintf("%s must be within -%f and +%f", name, maxValue, maxValue))
	}
	return actual, errors
}
