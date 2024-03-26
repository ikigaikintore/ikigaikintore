package bot

import (
	th "github.com/mymmrac/telego/telegohandler"
)

type Command interface {
	Handler() th.Handler
	Predicates() []th.Predicate
}

type CommandHandler func() Command

type cmdHandler struct {
	fn   th.Handler
	cmds []th.Predicate
}

func (ch cmdHandler) Handler() th.Handler {
	return ch.fn
}

func (ch cmdHandler) Predicates() []th.Predicate {
	return ch.cmds
}
