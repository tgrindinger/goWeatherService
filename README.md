### Go Weather Service API

This API will provide a brief description of weather conditions for a given latitude and longitude. Responses will generally be of this form:

```
{
  "condition": "clear sky",
  "temperature": "cold"
}
```

A request includes latitude and longitude as query string parameters, such as in this example:

```
http://localhost:8080/current?latitude=39&longitude=-94.58
```

## Setup

* Required: Add your API key for the OpenWeather service to `config/config.json`
* Optional: change the port number if you already have something running on 8080 or you just want it to use a different port
* Optional: Use `go mod tidy` to install dependencies
* Run the app with `go run .`
* Point a web browser to `http://localhost:8080/swagger` to interact with the API in a nice web interface

## Development

* Projects are split across controllers, managers, and resource access layers
  1. Controllers primarily deal with validating user input and converting it to internal representations
  1. Managers are the primary location for business logic and call the resource access layer to obtain necessary external data.
  1. The resource access layer exists to isolate the rest of the projects from external dependencies. Part of this involves providing stub versions of the external dependencies, so unit tests may still be executed against the rest of the projects.
* You should not need to change the URL in the configuration, but it may be useful if one desires to add service virtualization
* If any changes are made to the API contracts, `swag init` will regenerate the swagger docs. You will need to `go install github.com/swaggo/swag/cmd/swag@latest` to do so.

## Testing

* Unit tests are available for most types, cover the vast majority of the code, and are written in a "sociable" manner, meaning they allow collaboration with internal dependencies. External dependencies are stubbed in the resource-access package.
* Integration tests are available, as well, and require the apiKey to be present in config. These cover communication with the OpenWeather service.
* End-to-end tests are also available and require the API service to be running. They will automatically use the same port as the service and so should not need any separate configuration, unless the service is deployed remotely.
