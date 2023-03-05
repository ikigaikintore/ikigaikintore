package ports

import (
	"context"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/bounds"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/domain"
)

type (
	IUserService interface {
		Filter(context.Context, domain.Filter) []domain.User
	}

	IAgendaService interface {
		StartBooking(context.Context) error
		CrossfitStatus(context.Context) (domain.Status, error)
		GetToken(context.Context) (domain.Credentials, error)
		SetToken(context.Context, domain.Credentials) error
	}

	AgendaClient interface {
		SetToken(context.Context, bounds.ClientCredentials) error
		GetURLCredentials(context.Context) (bounds.ClientCredentials, error)
		Book(context.Context) error
		Status(context.Context) (bounds.ClientStatus, error)
	}
)

func ToStatus(st bounds.ClientStatus) domain.Status {
	var s domain.ProcessStatus
	switch st.ID {
	case bounds.Failed:
		s = domain.Failed
	case bounds.Working:
		s = domain.Working
	case bounds.Finished:
		s = domain.Finished
	}
	return domain.Status{
		ID:          s,
		Date:        st.Date,
		Description: st.Detail,
		Complete:    st.Complete,
	}
}

func ToCredentials(crd bounds.ClientCredentials) domain.Credentials {
	return domain.Credentials{
		URI:   crd.Url,
		Token: crd.Token,
	}
}
