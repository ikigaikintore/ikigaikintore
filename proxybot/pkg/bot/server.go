package bot

import (
	"context"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/valyala/fasthttp"
	"golang.ngrok.com/ngrok"
	config2 "golang.ngrok.com/ngrok/config"
	"net"
	"os"
)

type telegram struct {
	bot  *telego.Bot
	wh   Webhooker
	envs config.Envs
}

type Telegram interface {
	ResetWebhook()
	Setup() (*th.BotHandler, error)
	Start() error
	Stop(ctx context.Context) error
}

func New(envs config.Envs, webhooker Webhooker) (Telegram, error) {
	var opts []telego.BotOption
	if envs.App.IsDev() {
		opts = append(opts, telego.WithDefaultDebugLogger())
	}

	bot, err := telego.NewBot(envs.Telegram.Token, opts...)
	if err != nil {
		return nil, err
	}
	return &telegram{bot: bot, wh: webhooker, envs: envs}, nil
}
func (t *telegram) ResetWebhook() {
	info, err := t.bot.GetWebhookInfo()
	if err != nil {
		return
	}
	if info.PendingUpdateCount > 0 {
		_ = t.bot.DeleteWebhook(&telego.DeleteWebhookParams{DropPendingUpdates: true})
	}
}
func (t *telegram) Setup() (*th.BotHandler, error) {
	srv := &fasthttp.Server{}
	updates, err := t.bot.UpdatesViaWebhook(
		t.envs.Telegram.WebhookUriPathBase+t.bot.Token(),
		telego.WithWebhookServer(telego.FuncWebhookServer{
			Server: &telego.FastHTTPWebhookServer{
				Logger:      t.bot.Logger(),
				Server:      srv,
				Router:      router.New(),
				SecretToken: t.envs.Telegram.SecretToken,
			},
			StartFunc: func(_ string) error {
				return srv.Serve(t.wh.Listener())
			},
		}),
		telego.WithWebhookSet(
			&telego.SetWebhookParams{
				URL:         t.wh.URL() + t.envs.Telegram.WebhookUriPathBase + t.bot.Token(),
				SecretToken: t.envs.Telegram.SecretToken,
			},
		),
	)
	if err != nil {
		return nil, err
	}
	bh, _ := th.NewBotHandler(t.bot, updates)
	bh.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		if update.Message.From.IsBot {
			return
		}
		if update.Message.From.ID != t.envs.Telegram.WebhookUserID {
			return
		}
		next(bot, update)
	})
	return bh, nil
}
func (t *telegram) Start() error {
	return t.bot.StartWebhook(t.wh.Port())
}
func (t *telegram) Stop(ctx context.Context) error {
	return t.bot.StopWebhookWithContext(ctx)
}

type Webhooker interface {
	URL() string
	Listener() net.Listener
	HasListener() bool
	Port() string
}

type tunnelWebhook struct {
	tun ngrok.Tunnel
}

func newTunnelWebhook(ctx context.Context) Webhooker {
	if os.Getenv("NGROK_AUTHTOKEN") == "" {
		panic("NGROK_AUTHTOKEN not set")
	}
	tun, err := ngrok.Listen(ctx,
		config2.HTTPEndpoint(config2.WithForwardsTo(":8080")),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		panic(err)
	}
	return &tunnelWebhook{tun: tun}
}
func (w *tunnelWebhook) URL() string {
	return w.tun.URL()
}
func (w *tunnelWebhook) Listener() net.Listener {
	return w.tun
}
func (w *tunnelWebhook) HasListener() bool {
	return w.tun != nil
}
func (w *tunnelWebhook) Port() string {
	return ""
}

type cloudRunWebhook struct {
	uri  string
	port string
	lis  net.Listener
}

func newCloudRunWebhook(uri, port string) Webhooker {
	return &cloudRunWebhook{
		uri:  uri,
		port: port,
	}
}
func (w *cloudRunWebhook) URL() string {
	return w.uri
}
func (w *cloudRunWebhook) Listener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", w.port))
	if err != nil {
		fmt.Println("litener err: ", err)
	}
	return lis
}
func (w *cloudRunWebhook) HasListener() bool {
	return w.lis != nil
}
func (w *cloudRunWebhook) Port() string {
	return w.port
}

func WebhookerFactory(ctx context.Context, envs config.Envs) Webhooker {
	if !envs.App.IsDev() {
		return newCloudRunWebhook(envs.Telegram.WebhookListenUrl, fmt.Sprintf("%v", envs.Infra.Port))
	}

	return newTunnelWebhook(ctx)
}
