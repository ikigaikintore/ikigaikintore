package config

import "github.com/kelseyhightower/envconfig"

type (
	HttpServer struct {
		AppPort    string `envconfig:"APP_PORT" default:"8308"`
		HealthPort string `envconfig:"HEALTH_PORT" default:"8309"`
	}
)

var (
	HttpServerCfg HttpServer
)

func LoadConfigParameters() {
	envconfig.MustProcess("HTTP", &HttpServerCfg)
}
