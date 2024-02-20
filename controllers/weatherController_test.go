package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/tgrindinger/goWeatherService/managers"
	"github.com/tgrindinger/goWeatherService/models"
	resourceAccess "github.com/tgrindinger/goWeatherService/resource_access"
)

type fakeResponseWriter struct {
	statusCode int
	body       []byte
	header     http.Header
}

func NewFakeResponseWriter() *fakeResponseWriter {
	return &fakeResponseWriter{
		statusCode: 0,
		header:     http.Header{},
		body:       []byte{},
	}
}

func (w *fakeResponseWriter) Header() http.Header {
	return w.header
}

func (w *fakeResponseWriter) Write(bytes []byte) (int, error) {
	w.body = bytes
	return 0, nil
}

func (w *fakeResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func NewRequest(method string, urlStr string) *http.Request {
	parsedUrl, _ := url.Parse(urlStr)
	return &http.Request{
		Method: method,
		URL:    parsedUrl,
	}
}

func TestCurrentHandlerStatusCodes(t *testing.T) {
	cases := []struct {
		desc       string
		method     string
		url        string
		statusCode int
	}{{
		"missing latitude returns bad request",
		"GET",
		"http://localhost.com/current?longitude=123",
		http.StatusBadRequest,
	}, {
		"missing longitude returns bad request",
		"GET",
		"http://localhost.com/current?latitude=45",
		http.StatusBadRequest,
	}, {
		"invalid latitude returns bad request",
		"GET",
		"http://localhost.com/current?latitude=100&longitude=123",
		http.StatusBadRequest,
	}, {
		"invalid longitude returns bad request",
		"GET",
		"http://localhost.com/current?latitude=45&longitude=200",
		http.StatusBadRequest,
	}, {
		"valid latitude and longitude returns OK",
		"GET",
		"http://localhost.com/current?latitude=45&longitude=123",
		http.StatusOK,
	}, {
		"non-GET returns not found",
		"POST",
		"http://localhost.com/current?latitude=45&longitude=123",
		http.StatusNotFound,
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			request := NewRequest(tc.method, tc.url)
			writer := NewFakeResponseWriter()
			controller := NewWeatherController(managers.NewWeatherManager(resourceAccess.NewStubWeatherService()))

			// act
			controller.CurrentHandler(writer, request)

			// assert
			if tc.statusCode != writer.statusCode {
				t.Fatalf("wrong status code: got %d want %d", writer.statusCode, tc.statusCode)
			}
		})
	}
}

func TestCurrentHandlerBody(t *testing.T) {
	cases := []struct {
		desc        string
		condition   string
		temperature string
		url         string
	}{{
		"london is cold and overcast",
		"overcast clouds",
		"cold",
		"http://localhost.com/current?latitude=51.509865&longitude=-0.118092",
	}, {
		"maui is warm and few clouds",
		"few clouds",
		"moderate",
		"http://localhost.com/current?latitude=20.798363&longitude=-156.331924",
	}}
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			// arrange
			request := NewRequest("GET", tc.url)
			writer := NewFakeResponseWriter()
			controller := NewWeatherController(managers.NewWeatherManager(resourceAccess.NewStubWeatherService()))

			// act
			controller.CurrentHandler(writer, request)

			// assert
			var report models.WeatherReport
			reader := bytes.NewReader(writer.body)
			_ = json.NewDecoder(reader).Decode(&report)
			if report.Condition != tc.condition {
				t.Fatalf("wrong condition: got %s want %s", report.Condition, tc.condition)
			}
			if report.Temperature != tc.temperature {
				t.Fatalf("wrong temperature: got %s want %s", report.Temperature, tc.temperature)
			}
		})
	}
}
