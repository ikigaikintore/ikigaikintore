package config

import "github.com/kelseyhightower/envconfig"

type app struct {
	Env           string `envconfig:"ENV" default:"dev"`
	TargetBackend string `envconfig:"TARGET_BACKEND" default:"0.0.0.0:9000"`
}

type Telegram struct {
	Token              string `envconfig:"BOT_TOKEN"`
	WebhookListenUrl   string `envconfig:"WEBHOOK_URL"`
	WebhookUriPathBase string `envconfig:"WEBHOOK_PATH_BASE"`
	WebhookUserID      int64  `envconfig:"WEBHOOK_USER_ID"`
	SecretToken        string `envconfig:"SECRET_TOKEN"`
}

type infra struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func (a app) IsDev() bool {
	return a.Env == "dev"
}

type Envs struct {
	App      app
	Infra    infra
	Telegram Telegram
}

func Load() Envs {
	var envApp app
	envconfig.MustProcess("PROXY_BOT_APP", &envApp)

	var telegramBot Telegram
	envconfig.MustProcess("PROXY_BOT_TELEGRAM", &telegramBot)

	var infraData infra
	envconfig.MustProcess("", &infraData)

	return Envs{
		App:      envApp,
		Infra:    infraData,
		Telegram: telegramBot,
	}
}
