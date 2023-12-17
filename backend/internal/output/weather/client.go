package weather

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type client struct {
	c ClientWithResponsesInterface
}

type Weather struct {
	t                time.Time
	windSpeed        float32
	temperature      float32
	humidity         int
	temperatureRange struct {
		min, max float32
	}
	weatherType int
}
type ListWeather []Weather

func (w Weather) GetTime() time.Time {
	return w.t
}

func (w Weather) GetTemperature() float32 {
	return w.temperature
}

func (w Weather) GetHumidity() int {
	return w.humidity
}

func (w Weather) GetWeatherType() int {
	return w.weatherType
}

func (w Weather) GetTemperatureRange() (float32, float32) {
	return w.temperatureRange.min, w.temperatureRange.max
}

func (w Weather) GetWindSpeed() float32 {
	return w.windSpeed
}

type ClientRequest interface {
	GetCurrentWeather(context.Context) (*Weather, error)
	GetForecastWeather(context.Context) (ListWeather, error)
}

func NewClientRequest() ClientRequest {
	c, _ := NewClientWithResponses("https://api.openweathermap.org")
	return &client{c: c}
}

func (c client) GetCurrentWeather(ctx context.Context) (*Weather, error) {
	resp, err := c.c.GetCurrentWithResponse(
		ctx,
		&GetCurrentParams{
			Appid: os.Getenv("OPENWEATHER_API_KEY"),
			Lat:   35.71,
			Lon:   139.73,
		},
	)
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newErrorResponse(&errorResponse{cr: resp}).GetError()
	}
	return &Weather{
		t:           time.Unix(int64(resp.JSON200.Dt), 0),
		windSpeed:   resp.JSON200.Wind.Speed,
		temperature: resp.JSON200.Main.Temp,
		humidity:    resp.JSON200.Main.Humidity,
		temperatureRange: struct{ min, max float32 }{
			min: resp.JSON200.Main.TempMin,
			max: resp.JSON200.Main.TempMax,
		},
		weatherType: resp.JSON200.Weather[0].Id,
	}, nil
}

func (c client) GetForecastWeather(ctx context.Context) (ListWeather, error) {
	resp, err := c.c.GetForecast3HourWithResponse(
		ctx,
		&GetForecast3HourParams{
			Lat:   35.71,
			Lon:   139.73,
			Cnt:   12,
			Appid: os.Getenv("OPENWEATHER_API_KEY"),
		},
	)
	if err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newErrorResponse(&errorResponse{fcr: resp}).GetError()
	}

	lw := make(ListWeather, resp.JSON200.Cnt)

	for i, v := range resp.JSON200.List {
		lw[i] = Weather{
			t:           time.Unix(int64(v.Dt), 0),
			windSpeed:   v.Wind.Speed,
			temperature: v.Main.Temp,
			humidity:    v.Main.Humidity,
			temperatureRange: struct{ min, max float32 }{
				min: v.Main.TempMin,
				max: v.Main.TempMax,
			},
			weatherType: v.Weather[0].Id,
		}
	}

	return lw, err
}

type errorResponse struct {
	cr  *GetCurrentResponse
	fcr *GetForecast3HourResponse
}

type ErrorWeatherResponse interface {
	GetError() error
}

func newErrorResponse(r *errorResponse) ErrorWeatherResponse {
	if r.cr != nil {
		return r.cr
	}
	return r.fcr
}

func (er GetCurrentResponse) GetError() error {
	switch er.StatusCode() {
	case http.StatusBadRequest:
		return fmt.Errorf(er.JSON400.Message)
	case http.StatusUnauthorized:
		return fmt.Errorf(er.JSON401.Message)
	case http.StatusNotFound:
		return fmt.Errorf(er.JSON404.Message)
	case http.StatusTooManyRequests:
		return fmt.Errorf(er.JSON429.Message)
	default:
		return fmt.Errorf(er.JSON5XX.Message)
	}
}

func (er GetForecast3HourResponse) GetError() error {
	switch er.StatusCode() {
	case http.StatusBadRequest:
		return fmt.Errorf(er.JSON400.Message)
	case http.StatusUnauthorized:
		return fmt.Errorf(er.JSON401.Message)
	case http.StatusNotFound:
		return fmt.Errorf(er.JSON404.Message)
	case http.StatusTooManyRequests:
		return fmt.Errorf(er.JSON429.Message)
	default:
		return fmt.Errorf(er.JSON5XX.Message)
	}
}
