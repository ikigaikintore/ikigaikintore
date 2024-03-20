package twirp

import (
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/usecase"
	"github.com/twitchtv/twirp"
)

func NewTwirpServer() proto.TwirpServer {
	return proto.NewWeatherServer(usecase.NewWeatherService(), twirp.WithServerPathPrefix("/v1/weather"))
}
