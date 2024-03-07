package config

import "github.com/kelseyhightower/envconfig"

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
	App   app
	Infra infra
}

func Load() Envs {
	var envApp app
	envconfig.MustProcess("BACKEND_APP", &envApp)

	var infraData infra
	envconfig.MustProcess("", &infraData)

	return Envs{
		App:   envApp,
		Infra: infraData,
	}
}
