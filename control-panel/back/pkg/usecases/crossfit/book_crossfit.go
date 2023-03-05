package crossfit

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/bounds"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/clients"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/domain"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/infra/config"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/ports"
)

type agendaService struct {
	client ports.AgendaClient
}

func (a agendaService) CrossfitStatus(ctx context.Context) (domain.Status, error) {
	st, err := a.client.Status(ctx)
	if err != nil {
		return domain.Status{}, err
	}
	return ports.ToStatus(st), nil
}

func (a agendaService) StartBooking(ctx context.Context) error {
	return a.client.Book(ctx)
}

func (a agendaService) GetToken(ctx context.Context) (domain.Credentials, error) {
	cred, err := a.client.GetURLCredentials(ctx)
	if err != nil {
		return domain.Credentials{}, err
	}
	return ports.ToCredentials(cred), nil
}

func (a agendaService) SetToken(ctx context.Context, credentials domain.Credentials) error {
	return a.client.SetToken(ctx, bounds.ToCredentials(credentials))
}

func NewAgendaCrossfitService() ports.IAgendaService {
	return &agendaService{
		client: clients.NewCrossfitAgendaClient(config.AgendaCfg),
	}
}
