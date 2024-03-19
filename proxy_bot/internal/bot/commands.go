package bot

import (
	"errors"
	"github.com/ikigaikintore/ikigaikintore/proxy_bot/config"
	"io"
	"log"
)

type botServer struct {
}

type Listener interface {
	Start()
	Stop()
	Parser(envCfg config.Envs, body io.ReadCloser) error
}

func (b *botServer) Start() {
}

func (b *botServer) Stop() {
}

func NewBot(cfg config.Envs) (Listener, error) {
	if cfg.App.IsDev() {
	}
	return &botServer{}, nil
}

func logErrBot(err error) {
	if err == nil {
		return
	}
	log.Println(err)
}

func (b *botServer) handlerStart() {

}

var ErrForbidden = errors.New("forbidden action for the user")

func (b *botServer) secure(envCfg config.Envs) error {
	return ErrForbidden
}

func (b *botServer) Parser(envCfg config.Envs, body io.ReadCloser) error {

	return nil
}
