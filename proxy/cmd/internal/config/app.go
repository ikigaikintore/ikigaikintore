package config

import "github.com/kelseyhightower/envconfig"

type cors struct {
	AllowedDomains string `envconfig:"ALLOWED_DOMAINS" default:"" required:"true"`
}

type app struct {
	Env           string `envconfig:"ENV" default:"dev"`
	TargetBackend string `envconfig:"TARGET_BACKEND" default:"localhost:8999"`
}

type Telegram struct {
	Token              string `envconfig:"BOT_TOKEN"`
	WebhookListenUrl   string `envconfig:"WEBHOOK_URL"`
	WebhookUriPathBase string `envconfig:"WEBHOOK_PATH_BASE"`
	WebhookUserID      int64  `envconfig:"WEBHOOK_USER_ID"`
}

type infra struct {
	Port int `envconfig:"PORT" default:"8997"`
}

func (a app) IsDev() bool {
	return a.Env == "dev"
}

type Envs struct {
	Cors     cors
	App      app
	Infra    infra
	Telegram Telegram
}

func Load() Envs {
	var envCors cors
	envconfig.MustProcess("PROXY_CORS", &envCors)

	var envApp app
	envconfig.MustProcess("PROXY_APP", &envApp)

	var telegramBot Telegram
	envconfig.MustProcess("PROXY_TELEGRAM", &telegramBot)

	var infraData infra
	envconfig.MustProcess("", &infraData)

	return Envs{
		Cors:     envCors,
		App:      envApp,
		Infra:    infraData,
		Telegram: telegramBot,
	}
}
