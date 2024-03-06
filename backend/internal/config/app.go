package config

import "github.com/kelseyhightower/envconfig"

type cors struct {
	AllowedDomains string `envconfig:"ALLOWED_DOMAINS" default:"" required:"true"`
}

type app struct {
	Env string `envconfig:"ENV" default:"dev"`
}

type infra struct {
	Port int `envconfig:"PORT" default:"8999"`
}

func (a app) IsDev() bool {
	return a.Env == "dev"
}

type Envs struct {
	Cors  cors
	App   app
	Infra infra
}

func Load() Envs {
	var envCors cors
	envconfig.MustProcess("BACKEND_CORS", &envCors)

	var envApp app
	envconfig.MustProcess("BACKEND_APP", &envApp)

	var infraData infra
	envconfig.MustProcess("", &infraData)

	return Envs{
		Cors:  envCors,
		App:   envApp,
		Infra: infraData,
	}
}
