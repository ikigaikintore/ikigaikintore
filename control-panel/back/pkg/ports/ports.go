package ports

import (
	"context"
	"github.com/ervitis/control-panel/pkg/domain"
)

type (
	IUserService interface {
		Filter(context.Context, domain.Filter) []domain.User
	}
)
