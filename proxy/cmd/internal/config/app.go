package config

import "github.com/kelseyhightower/envconfig"

type cors struct {
	AllowedDomains string `envconfig:"ALLOWED_DOMAINS" default:"" required:"true"`
}

type app struct {
	Env           string `envconfig:"ENV" default:"dev"`
	TargetBackend string `envconfig:"TARGET_BACKEND" default:"localhost:8999"`
}

type infra struct {
	Port int `envconfig:"PORT" default:"8997"`
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
	envconfig.MustProcess("PROXY_CORS", &envCors)

	var envApp app
	envconfig.MustProcess("PROXY_APP", &envApp)

	var infraData infra
	envconfig.MustProcess("", &infraData)

	return Envs{
		Cors:  envCors,
		App:   envApp,
		Infra: infraData,
	}
}
