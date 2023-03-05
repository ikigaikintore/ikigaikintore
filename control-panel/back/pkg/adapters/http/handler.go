package http

import "github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/ports"

type Handler struct {
	agendaService ports.IAgendaService
}

func NewHandler(agendaService ports.IAgendaService) Handler {
	return Handler{agendaService: agendaService}
}
