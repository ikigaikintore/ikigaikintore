package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type CrossfitClient struct {
	Address string `envconfig:"ADDRESS" default:"localhost"`
	Port    string `envconfig:"PORT" default:"9090"`
}

var (
	AgendaCfg CrossfitClient
)

func (cc CrossfitClient) FullAddress() string {
	return fmt.Sprintf("%s:%s", cc.Address, cc.Port)
}

func LoadConfigAgendaClient() {
	envconfig.MustProcess("CROSSFIT_AGENDA_CLIENT", &AgendaCfg)
}
