package grpc

import (
	"github.com/ervitis/ikigaikintore/backend/pkg/proto"
	"github.com/ervitis/ikigaikintore/backend/pkg/usecase"
)

func NewTwirpServer() proto.TwirpServer {
	return proto.NewWeatherServer(usecase.NewWeatherService())
}
