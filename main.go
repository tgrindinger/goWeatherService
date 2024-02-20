package main

import (
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
	config "github.com/tgrindinger/goWeatherService/config"
	controllers "github.com/tgrindinger/goWeatherService/controllers"
	_ "github.com/tgrindinger/goWeatherService/docs"
	"github.com/tgrindinger/goWeatherService/managers"
	resourceAccess "github.com/tgrindinger/goWeatherService/resource_access"
)

// @title Current Weather Report
// @version 1.0
// @description This service reports the current weather for a given location
func main() {
	config, err := config.ReadConfig("config/config.json")
	if err != nil {
		fmt.Printf("Unable to read config file: %s", err.Error())
		return
	}
	weatherController := assembleDependencies(config)

	http.HandleFunc("/current", weatherController.CurrentHandler)
	http.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", config.Port)),
	))

	http.ListenAndServe(config.Port, nil)
}

func assembleDependencies(config *config.Config) *controllers.WeatherController {
	weatherService := resourceAccess.NewApiWeatherService(config.ApiUrl, config.ApiKey)
	weatherManager := managers.NewWeatherManager(weatherService)
	weatherController := controllers.NewWeatherController(weatherManager)
	return weatherController
}
