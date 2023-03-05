package grpc

import (
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/infra/config"
	lib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(cfg config.CrossfitClient) CrossfitAgendaServiceClient {
	con, err := lib.Dial(
		cfg.FullAddress(),
		lib.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	return NewCrossfitAgendaServiceClient(con)
}
