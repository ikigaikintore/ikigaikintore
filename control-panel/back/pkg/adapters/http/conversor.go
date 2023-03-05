package http

import "github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/domain"

func (t ReqGoogleToken) To() domain.Credentials {
	return domain.Credentials{
		Token: t.Token,
	}
}

func NewProcessStatus(st domain.Status) GetProcessStatus {
	var s ProcessStatuses
	switch st.ID {
	case domain.Failed:
		s = Failed
	case domain.Working:
		s = Working
	case domain.Finished:
		s = Finished
	}
	return GetProcessStatus{
		Complete: false,
		Date:     uint64(st.Date.Unix()),
		Detail:   st.Description,
		Id:       s,
	}
}

func NewGoogleCredentials(c domain.Credentials) GetGoogleCredentials {
	return GetGoogleCredentials{Link: c.URI}
}
