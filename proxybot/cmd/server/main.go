package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	bot2 "github.com/ikigaikintore/ikigaikintore/proxybot/pkg/bot"
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/service"
)

func main() {
	envs := config.Load()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	conn := service.BackendClient(config.Load())
	defer func() {
		_ = conn.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	webhook := bot2.WebhookerFactory(ctx, envs)
	telegramBot, err := bot2.New(envs, webhook)
	if err != nil {
		panic(err)
	}

	handlers := []bot2.Command{
		bot2.NewHandlerTodayWeather(service.NewWeatherClient(conn)),
	}

	bh, err := telegramBot.Setup()
	if err != nil {
		panic(err)
	}

	for _, handle := range handlers {
		bh.Handle(handle.Handler(), handle.Predicates()...)
	}

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Panic(err)
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
