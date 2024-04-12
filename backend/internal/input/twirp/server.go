package twirp

import (
	"context"

	"github.com/ikigaikintore/ikigaikintore/backend/internal/input/common"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"github.com/twitchtv/twirp"
)

func NewTwirpServer() proto.TwirpServer {
	return proto.NewWeatherServer(NewTwirpWeatherService(), twirp.WithServerPathPrefix("/v1/weather"))
}

type twirpServer struct {
	proto.Weather

	common common.Handler
}

func NewTwirpWeatherService() *twirpServer {
	return &twirpServer{
		common: common.NewHandler(),
	}
}

func (w twirpServer) GetWeather(ctx context.Context, request *proto.WeatherRequest) (*proto.WeatherReply, error) {
	return w.common.GetWeather(ctx, request)
}
