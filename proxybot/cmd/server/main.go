package main

import (
	"context"
	"fmt"
	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	proto "github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func backendClient(envs config.Envs) *grpc.ClientConn {
	conn, err := grpc.Dial(envs.App.TargetBackend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
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

	var updates <-chan telego.Update
	if envs.App.IsDev() {
		updates, _ = bot.UpdatesViaLongPolling(nil)
	} else {
		_ = bot.SetWebhook(&telego.SetWebhookParams{
			SecretToken: envs.Telegram.SecretToken,
			URL:         envs.Telegram.WebhookListenUrl,
		})
		updates, _ = bot.UpdatesViaWebhook(envs.Telegram.WebhookUriPathBase)
	}

	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()

	bh.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		if update.Message.From.IsBot {
			return
		}
		if update.Message.From.ID != envs.Telegram.WebhookUserID {
			return
		}
		next(bot, update)
	})

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
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
	}, th.CommandEqual("start"))

	if envs.App.IsDev() {
		go func() {
			bh.Start()
		}()
	}

	if !envs.App.IsDev() {
		go func() {
			bh.Start()
		}()

		go func() {
			log.Panic(bot.StartWebhook(fmt.Sprintf(":%v", envs.Infra.Port)))
		}()
	}

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
