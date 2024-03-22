package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	proto "github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/valyala/fasthttp"
	"golang.ngrok.com/ngrok"
	config2 "golang.ngrok.com/ngrok/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func backendClient(envs config.Envs) *grpc.ClientConn {
	cred := insecure.NewCredentials()
	if !envs.App.IsDev() {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}
		cred = credentials.NewTLS(&tls.Config{RootCAs: certPool, MinVersion: tls.VersionTLS13})
	}
	addr := envs.App.TargetBackend
	if len(strings.Split(addr, ":")) < 2 && !envs.App.IsDev() {
		addr = addr + ":443"
	}
	conn, err := grpc.Dial(envs.App.TargetBackend, grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}
	return conn
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

func newTunnelWebhook() Webhooker {
	if os.Getenv("NGROK_AUTHTOKEN") == "" {
		panic("NGROK_AUTHTOKEN not set")
	}
	tun, err := ngrok.Listen(context.Background(),
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

func webhookerFactory(envs config.Envs) Webhooker {
	if !envs.App.IsDev() {
		return newCloudRunWebhook(envs.Telegram.WebhookListenUrl, fmt.Sprintf("%v", envs.Infra.Port))
	}
	return newTunnelWebhook()
}

type handler struct {
	fn   func() th.Handler
	cmds []th.Predicate
}

func main() {
	envs := config.Load()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	conn := backendClient(envs)
	defer func() {
		_ = conn.Close()
	}()
	grpcClient := proto.NewWeatherClient(conn)

	var opts []telego.BotOption
	if envs.App.IsDev() {
		opts = append(opts, telego.WithDefaultDebugLogger())
	}

	bot, err := telego.NewBot(envs.Telegram.Token, opts...)
	if err != nil {
		panic(err)
	}

	info, _ := bot.GetWebhookInfo()
	if info != nil {
		_ = bot.DeleteWebhook(&telego.DeleteWebhookParams{DropPendingUpdates: true})
	}

	webhook := webhookerFactory(envs)

	url := webhook.URL() + envs.Telegram.WebhookUriPathBase + bot.Token()
	webhookUrl := envs.Telegram.WebhookUriPathBase + bot.Token()

	srv := &fasthttp.Server{}
	updates, err := bot.UpdatesViaWebhook(
		webhookUrl,
		telego.WithWebhookServer(telego.FuncWebhookServer{
			Server: &telego.FastHTTPWebhookServer{
				Logger:      bot.Logger(),
				Server:      srv,
				Router:      router.New(),
				SecretToken: envs.Telegram.SecretToken,
			},
			StartFunc: func(_ string) error {
				return srv.Serve(webhook.Listener())
			},
		}),
		telego.WithWebhookSet(
			&telego.SetWebhookParams{
				URL:         url,
				SecretToken: envs.Telegram.SecretToken,
			},
		),
	)
	if err != nil {
		fmt.Println(err)
	}

	bh, _ := th.NewBotHandler(bot, updates)
	bh.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		if update.Message.From.IsBot {
			return
		}
		if update.Message.From.ID != envs.Telegram.WebhookUserID {
			return
		}
		next(bot, update)
	})

	handlers := []handler{
		{
			fn: func() th.Handler {
				return func(bot *telego.Bot, update telego.Update) {
					resp, err := grpcClient.GetWeather(update.Context(), &proto.WeatherRequest{WeatherFilter: &proto.WeatherFilter{Location: "Tokyo"}})
					if err != nil {
						log.Println(err)
						return
					}
					_, _ = bot.SendMessage(
						tu.Messagef(
							tu.ID(update.Message.Chat.ID),
							"Got somethin' that might interest ya'! %v",
							resp.GetWeatherCurrent().GetTemperature(),
						),
					)
				}
			},
			cmds: []th.Predicate{th.CommandEqual("start")},
		},
	}

	for _, handle := range handlers {
		bh.Handle(handle.fn(), handle.cmds...)
	}

	go func() {
		log.Panic(bot.StartWebhook(webhook.Port()))
	}()
	go func() {
		bh.Start()
	}()

	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if envs.App.IsDev() {
		bot.StopLongPolling()
	} else {
		_ = bot.StopWebhookWithContext(ctx)
	}
	bh.Stop()
}
