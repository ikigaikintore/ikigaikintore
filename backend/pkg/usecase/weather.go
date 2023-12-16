package usecase

import (
	"context"
	"github.com/ervitis/ikigaikintore/backend/pkg/proto"
)

type WeatherService interface {
	proto.Weather
}

type weatherService struct {
	client interface{}
}

func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (w weatherService) GetWeather(context.Context, *proto.WeatherRequest) (*proto.WeatherReply, error) {
	return nil, nil
}
