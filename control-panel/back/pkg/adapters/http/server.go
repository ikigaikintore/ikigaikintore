package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/infra/config"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/ports"
	"github.com/ikigaikintore/ikigaikintore/control-panel/back/pkg/usecases/crossfit"
)

type (
	AppConfig struct {
		httpSrv *config.HttpServer
	}

	controlPanelServer struct {
		AppConfig
		srv *echo.Echo

		handler Handler
	}

	healthCheckServer struct {
		AppConfig
		srv *echo.Echo
	}

	Server struct {
		*controlPanelServer
		*healthCheckServer
	}
)

func NewServer(cfg *config.HttpServer) *Server {
	return &Server{
		controlPanelServer: newControlPanelServer(cfg, crossfit.NewAgendaCrossfitService()),
		healthCheckServer:  newHealthCheckServer(cfg),
	}
}

func (m *Server) Start() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	RegisterHandlersWithBaseURL(m.controlPanelServer.Router(), m.controlPanelServer.handler, "/v1")

	defer wg.Done()

	go func() {
		log.Panic(m.controlPanelServer.Start())
	}()
	go func() {
		log.Panic(m.healthCheckServer.Start())
	}()
	wg.Wait()
}

func (m *Server) Stop() {
	m.controlPanelServer.Stop()
	m.healthCheckServer.Stop()
}

func newControlPanelServer(cfg *config.HttpServer, svc ports.IAgendaService) *controlPanelServer {
	srv := echo.New()
	srv.HideBanner = true
	srv.Logger.SetLevel(log.INFO)
	srv.Use(middleware.CORS())
	srv.Use(middleware.Recover())
	srv.Use(middleware.RequestID())

	v1Handler := NewHandler(svc)

	return &controlPanelServer{AppConfig{httpSrv: cfg}, srv, v1Handler}
}

func newHealthCheckServer(cfg *config.HttpServer) *healthCheckServer {
	srv := echo.New()
	srv.Logger.SetLevel(log.ERROR)
	srv.HideBanner = true
	srv.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	return &healthCheckServer{AppConfig{httpSrv: cfg}, srv}
}

func (srv *controlPanelServer) Start() error {
	if err := srv.srv.Start(fmt.Sprintf(":%s", srv.AppConfig.httpSrv.AppPort)); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (srv *controlPanelServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.srv.Shutdown(ctx); err != nil {
		log.Error(err)
	}
}

func (srv *controlPanelServer) Router() *echo.Echo {
	return srv.srv
}

func (srv *healthCheckServer) Start() error {
	if err := srv.srv.Start(fmt.Sprintf(":%s", srv.AppConfig.httpSrv.HealthPort)); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (srv *healthCheckServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.srv.Shutdown(ctx); err != nil {
		log.Error(err)
	}
}

func (srv *healthCheckServer) Router() *echo.Echo {
	return srv.srv
}
