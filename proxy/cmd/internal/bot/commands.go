package bot

import (
	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/config"
	"gopkg.in/telebot.v3"
	"log"
	"net/http"
	"time"
)

type botServer struct {
	bot *telebot.Bot
}

type Listener interface {
	Start()
	Stop()
}

func (b *botServer) Start() {
	if b.bot == nil {
		log.Println("bot not loaded, skip")
	}
	b.bot.Handle("/start", b.handlerStart())
	b.bot.Start()
}

func (b *botServer) Stop() {
	if b.bot == nil {
		return
	}
	b.bot.Stop()
}

func NewBot(cfg config.Envs) (Listener, error) {
	var debug bool
	if cfg.App.IsDev() {
		debug = true
	}
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   cfg.Telegram.Token,
		Client:  &http.Client{Timeout: 10 * time.Second, Transport: http.DefaultTransport},
		Poller:  &telebot.LongPoller{Timeout: 15 * time.Second},
		Verbose: debug,
	})
	if err != nil {
		return nil, err
	}
	if err := bot.SetWebhook(&telebot.Webhook{Endpoint: &telebot.WebhookEndpoint{PublicURL: cfg.Telegram.WebhookListenUrl}}); err != nil {
		return nil, err
	}
	return &botServer{bot: bot}, nil
}

func logErrBot(err error) {
	if err == nil {
		return
	}
	log.Println(err)
}

func (b *botServer) handlerStart() telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		_, err := b.bot.Send(ctx.Sender(), "Hello strangar")
		logErrBot(err)
		return err
	}
}
