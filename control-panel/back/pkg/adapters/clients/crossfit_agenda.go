package clients

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/bounds"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/clients/grpc"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/infra/config"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/ports"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	client struct {
		c grpc.CrossfitAgendaServiceClient
	}
)

func (c client) SetToken(ctx context.Context, credentials bounds.ClientCredentials) error {
	_, err := c.c.SetTokenGoogle(ctx, &grpc.SetTokenGoogleRequest{Token: credentials.Token})
	return err
}

func (c client) GetURLCredentials(ctx context.Context) (bounds.ClientCredentials, error) {
	resp, err := c.c.CredentialsGoogle(ctx, &emptypb.Empty{})
	if err != nil {
		return bounds.ClientCredentials{}, err
	}
	return bounds.ClientCredentials{Url: resp.GetLink()}, nil
}

func (c client) Book(ctx context.Context) error {
	_, err := c.c.StartCrossfitAgenda(ctx, &emptypb.Empty{})
	return err
}

func (c client) Status(ctx context.Context) (bounds.ClientStatus, error) {
	resp, err := c.c.Status(ctx, &emptypb.Empty{})
	if err != nil {
		return bounds.ClientStatus{}, err
	}
	return resp.Into(), nil
}

func NewCrossfitAgendaClient(cfg config.CrossfitClient) ports.AgendaClient {
	return &client{c: grpc.NewClient(
		cfg,
	)}
}
