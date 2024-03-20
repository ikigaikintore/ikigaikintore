package grpc

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/usecase"
)

type WeatherServer proto.WeatherServer

type weatherServer struct {
	weatherService usecase.WeatherService
}

func (w weatherServer) GetWeather(ctx context.Context, request *proto.WeatherRequest) (*proto.WeatherReply, error) {
	return w.weatherService.GetWeather(ctx, request)
}

func NewWeatherServer() WeatherServer {
	return &weatherServer{weatherService: usecase.NewWeatherService()}
}
