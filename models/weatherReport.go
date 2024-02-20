package models

type WeatherReport struct {
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}
