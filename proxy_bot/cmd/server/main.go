package main

import (
	"context"
	"fmt"
	"github.com/ikigaikintore/ikigaikintore/proxy_bot/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	envs := config.Load()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	bot, err := telego.NewBot(envs.Telegram.Token, telego.WithDefaultDebugLogger())
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
		_, _ = bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				"Got somethin' that might interest ya'!",
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
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	if envs.App.IsDev() {
		bot.StopLongPolling()
	} else {
		_ = bot.StopWebhookWithContext(ctx)
	}
	bh.Stop()
}
