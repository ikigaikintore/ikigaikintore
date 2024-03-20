package config

import "github.com/kelseyhightower/envconfig"

type app struct {
	Env    string `envconfig:"ENV" default:"dev"`
	Target string `envconfig:"TARGET"`
}

type infra struct {
	Port int `envconfig:"PORT" default:"8999"`
}

func (a app) IsDev() bool {
	return a.Env == "dev"
}

func (a app) IsHTTP() bool {
	return a.Target == "http"
}

func (a app) IsGRPC() bool {
	return a.Target == "grpc"
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
