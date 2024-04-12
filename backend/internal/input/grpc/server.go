package grpc

import (
	"context"

	"github.com/ikigaikintore/ikigaikintore/backend/internal/input/common"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
)

type WeatherServer proto.WeatherServer

type weatherServer struct {
	common common.Handler
}

func (w weatherServer) GetWeather(ctx context.Context, request *proto.WeatherRequest) (*proto.WeatherReply, error) {
	return w.common.GetWeather(ctx, request)
}

func NewWeatherServer() WeatherServer {
	return &weatherServer{common: common.NewHandler()}
}
