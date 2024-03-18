package bot

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"gopkg.in/telebot.v3"

	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/config"
)

type botServer struct {
	bot *telebot.Bot
}

type Listener interface {
	Start()
	Stop()
	Parser(envCfg config.Envs, body io.ReadCloser) error
	Bot() *telebot.Bot
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

func (b *botServer) Bot() *telebot.Bot {
	return b.bot
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
	if err := bot.SetWebhook(&telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: cfg.Telegram.WebhookListenUrl + cfg.Telegram.WebhookUriPathBase,
		},
	}); err != nil {
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

var ErrForbidden = errors.New("forbidden action for the user")

func (b *botServer) secure(envCfg config.Envs, msg *telebot.Message) error {
	flag := false
	flag = flag || msg.Sender.IsBot != true
	flag = flag || msg.Sender.ID == envCfg.Telegram.WebhookUserID

	if flag {
		return nil
	}
	return ErrForbidden
}

func (b *botServer) Parser(envCfg config.Envs, body io.ReadCloser) error {
	var cmd telebot.Update
	if err := json.NewDecoder(body).Decode(&cmd); err != nil {
		log.Println("err decoding body: ", err)
		return err
	}
	if cmd.Message == nil {
		log.Println("no message, skip")
		return nil
	}
	if cmd.Message.IsService() {
		log.Println("service, skip")
		return nil
	}
	if cmd.Message.IsReply() {
		log.Println("reply, skip")
		return nil
	}
	isCommand := func(msg *telebot.Message) bool {
		for _, v := range msg.Entities {
			if v.Type == telebot.EntityCommand {
				return true
			}
		}
		return false
	}
	if !isCommand(cmd.Message) {
		log.Println("no command, skip")
	}
	if err := b.secure(envCfg, cmd.Message); err != nil {
		return err
	}

	return nil
}
