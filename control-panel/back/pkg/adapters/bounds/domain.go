package bounds

import (
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/domain"
	"time"
)

type ClientProcessStatus int

const (
	Working ClientProcessStatus = iota + 1
	Finished
	Failed
)

type (
	ClientCredentials struct {
		Token string
		Url   string
	}

	ClientStatus struct {
		Complete bool
		ID       ClientProcessStatus
		Date     time.Time
		Detail   string
	}
)

func ToCredentials(crd domain.Credentials) ClientCredentials {
	return ClientCredentials{
		Token: crd.Token,
		Url:   crd.URI,
	}
}
