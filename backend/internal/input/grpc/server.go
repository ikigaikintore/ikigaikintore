package grpc

import (
	"github.com/ervitis/ikigaikintore/backend/pkg/proto"
	"github.com/ervitis/ikigaikintore/backend/pkg/usecase"
	"github.com/twitchtv/twirp"
)

func NewTwirpServer() proto.TwirpServer {
	return proto.NewWeatherServer(usecase.NewWeatherService(), twirp.WithServerPathPrefix("/v1/weather"))
}
