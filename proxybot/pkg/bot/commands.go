package bot

import (
	th "github.com/mymmrac/telego/telegohandler"
)

type CommandMessage interface {
	Handler() th.MessageHandlerCtx
	Predicates() []th.Predicate
}

type CommandUpdate interface {
	Handler() th.Handler
	Predicates() []th.Predicate
}

type cmdUpdateHandler struct {
	fn   th.Handler
	cmds []th.Predicate
}

func (ch cmdUpdateHandler) Handler() th.Handler {
	return ch.fn
}

func (ch cmdUpdateHandler) Predicates() []th.Predicate {
	return ch.cmds
}

type cmdMessageHandler struct {
	fn   th.MessageHandlerCtx
	cmds []th.Predicate
}

func (ch cmdMessageHandler) Handler() th.MessageHandlerCtx {
	return ch.fn
}

func (ch cmdMessageHandler) Predicates() []th.Predicate {
	return ch.cmds
}
