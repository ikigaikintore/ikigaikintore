package grpc

import "github.com/ervitis/crossfitAgenda/adapters/domain"

func Into(st domain.Status) ProcessStatuses {
	switch st {
	case domain.Finished:
		return ProcessStatuses_ProcessStatuses_FINISHED
	case domain.Working:
		return ProcessStatuses_ProcessStatuses_WORKING
	default:
		return ProcessStatuses_ProcessStatuses_FAILED
	}
}
