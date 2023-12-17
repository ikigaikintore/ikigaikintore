package usecase

import (
	"context"
	"github.com/ervitis/ikigaikintore/backend/internal/output/weather"
	"github.com/ervitis/ikigaikintore/backend/pkg/proto"
)

type WeatherService interface {
	proto.Weather
}

type weatherService struct {
	client weather.ClientRequest
}

func NewWeatherService() WeatherService {
	return &weatherService{
		client: weather.NewClientRequest(),
	}
}

func (w weatherService) GetWeather(ctx context.Context, req *proto.WeatherRequest) (*proto.WeatherReply, error) {
	currWeather, err := w.client.GetCurrentWeather(ctx)
	if err != nil {
		return nil, err
	}

	hourlyWeather, err := w.client.GetForecastWeather(ctx)
	if err != nil {
		return nil, err
	}

	wp := make([]*proto.WeatherDailyPoint, len(hourlyWeather))
	for i, v := range hourlyWeather {
		m, n := v.GetTemperatureRange()
		wp[i] = &proto.WeatherDailyPoint{
			Timestamp:   uint64(v.GetTime().Unix()),
			Temperature: v.GetTemperature(),
			Humidity:    int32(v.GetHumidity()),
			TemperatureRange: &proto.TemperatureRange{
				Max: m,
				Min: n,
			},
			Weather: proto.WeatherType(v.GetWeatherType()),
		}
	}

	return &proto.WeatherReply{
		WeatherCurrent: &proto.WeatherCurrent{
			Temperature: currWeather.GetTemperature(),
			WindSpeed:   currWeather.GetWindSpeed(),
			Timestamp:   uint64(currWeather.GetTime().Unix()),
			Humidity:    int32(currWeather.GetHumidity()),
			Weather:     proto.WeatherType(currWeather.GetWeatherType()),
		},
		WeatherPoint: wp,
	}, nil
}
