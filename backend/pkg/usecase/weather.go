package usecase

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/backend/internal/output/weather"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/domain"
	"log"
)

type WeatherService interface {
	GetWeather(context.Context, domain.WeatherRequest) (domain.WeatherResponse, error)
}

type weatherService struct {
	client weather.ClientRequest
}

func NewWeatherService() WeatherService {
	return &weatherService{
		client: weather.NewClientRequest(),
	}
}

func (w weatherService) GetWeather(ctx context.Context, req domain.WeatherRequest) (domain.WeatherResponse, error) {
	currWeather, err := w.client.GetCurrentWeather(ctx, &weather.GetCurrentParams{
		Lat:   req.Latitude,
		Lon:   req.Longitude,
		Appid: req.APIKey,
		Mode:  &req.ModeToday,
		Lang:  &req.Lang,
		Units: &req.UnitsToday,
	})
	if err != nil {
		log.Println("error getting current weather", err)
		return domain.WeatherResponse{}, err
	}

	hourlyWeather, err := w.client.GetForecastWeather(ctx, &weather.GetForecast3HourParams{
		Lat:   req.Latitude,
		Lon:   req.Longitude,
		Cnt:   req.Cnt,
		Appid: req.APIKey,
		Mode:  &req.Mode3H,
		Lang:  &req.Lang,
		Units: &req.Units3H,
	})
	if err != nil {
		log.Println("error getting forecast weather", err)
		return domain.WeatherResponse{}, err
	}

	wp := make([]domain.WeatherPoint, len(hourlyWeather))
	for i, v := range hourlyWeather {
		m, n := v.GetTemperatureRange()
		wp[i] = domain.WeatherPoint{
			Timestamp:        uint64(v.GetTime().Unix()),
			Temperature:      v.GetTemperature(),
			Humidity:         int32(v.GetHumidity()),
			Icon:             v.GetIcon(),
			TemperatureRange: struct{ Min, Max float64 }{Min: n, Max: m},
			Weather:          domain.WeatherType(v.GetWeatherType()),
		}
	}

	return domain.WeatherResponse{
		WeatherCurrent: domain.WeatherCurrent{
			Temperature: currWeather.GetTemperature(),
			Icon:        currWeather.GetIcon(),
			WindSpeed:   currWeather.GetWindSpeed(),
			Timestamp:   uint64(currWeather.GetTime().Unix()),
			Humidity:    int32(currWeather.GetHumidity()),
			Weather:     domain.WeatherType(currWeather.GetWeatherType()),
		},
		WeatherPoints: wp,
	}, nil
}
