package config

import "github.com/kelseyhightower/envconfig"

type cors struct {
	AllowedDomains string `envconfig:"ALLOWED_DOMAINS" default:"" required:"true"`
}

type app struct {
	Env string `envconfig:"ENV" default:"dev"`
}

func (a app) IsDev() bool {
	return a.Env == "dev"
}

type Envs struct {
	Cors cors
	App  app
}

func Load() Envs {
	var envCors cors
	envconfig.MustProcess("BACKEND_CORS", &envCors)

	var envApp app
	envconfig.MustProcess("BACKEND_APP", &envApp)

	return Envs{
		Cors: envCors,
		App:  envApp,
	}
}
