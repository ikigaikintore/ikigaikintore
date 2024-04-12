package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	bot2 "github.com/ikigaikintore/ikigaikintore/proxybot/pkg/bot"
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service/proto"
)

func main() {
	envs := config.Load()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	conn := service.BackendClient(config.Load())
	defer func() {
		_ = conn.Close()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	webhook := bot2.WebhookerFactory(ctx, envs)
	telegramBot, err := bot2.New(envs, webhook)
	if err != nil {
		panic(err)
	}
	weatherClient := proto.NewWeatherClient(conn)
	commands := []bot2.CommandUpdate{
		bot2.NewHandlerTodayWeather(weatherClient),
		bot2.NewHandlerFuture(weatherClient),
		bot2.NewHandlerLocation(),
	}

	messageHandlers := []bot2.CommandMessage{
		bot2.NewHandlerResponseLocation(),
	}

	bh, err := telegramBot.Setup()
	if err != nil {
		panic(err)
	}

	for _, handle := range commands {
		bh.Handle(handle.Handler(), handle.Predicates()...)
	}
	for _, handle := range messageHandlers {
		bh.HandleMessageCtx(handle.Handler(), handle.Predicates()...)
	}

	go func() {
		if err := telegramBot.Start(); !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()
	go func() {
		bh.Start()
	}()

	<-ch
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = telegramBot.Stop(ctx)
	bh.Stop()
}
