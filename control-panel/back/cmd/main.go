package main

import (
	srvEndpoint "github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/adapters/http"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/infra/config"
	"os"
	"os/signal"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	srv := srvEndpoint.NewServer(&config.HttpServerCfg)

	srv.Start()

	<-quit
	srv.Stop()
}
