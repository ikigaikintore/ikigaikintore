package grpc

import (
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/bounds"
	"time"
)

func (st *Status200Response) Into() bounds.ClientStatus {
	var pst bounds.ClientProcessStatus
	switch st.GetId() {
	case ProcessStatuses_ProcessStatuses_FAILED:
		pst = bounds.Failed
	case ProcessStatuses_ProcessStatuses_FINISHED:
		pst = bounds.Finished
	case ProcessStatuses_ProcessStatuses_WORKING:
		pst = bounds.Working
	}
	return bounds.ClientStatus{
		Complete: st.GetComplete(),
		ID:       pst,
		Date:     time.Unix(int64(st.GetDate()), 0),
		Detail:   st.GetDetail(),
	}
}
