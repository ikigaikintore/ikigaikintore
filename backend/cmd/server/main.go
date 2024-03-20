package main

import (
	"context"
	"errors"
	"fmt"
	grpc2 "github.com/ikigaikintore/ikigaikintore/backend/internal/input/grpc"
	"github.com/ikigaikintore/ikigaikintore/backend/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikigaikintore/ikigaikintore/backend/internal/config"
	"github.com/ikigaikintore/ikigaikintore/backend/internal/input/twirp"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type grpcAppServer struct {
	l   net.Listener
	srv *grpc.Server
}

func (g grpcAppServer) ListenAndServe() error {
	return g.srv.Serve(g.l)
}

func (g grpcAppServer) Shutdown(_ context.Context) error {
	g.srv.GracefulStop()
	return nil
}

func NewGRPCAppServer(envCfg config.Envs) Server {
	srv := grpc.NewServer()
	proto.RegisterWeatherServer(srv, grpc2.NewWeatherServer())
	port := envCfg.Infra.Port
	if envCfg.App.IsDev() {
		port = +1
	}
	lis, _ := net.Listen("tcp", fmt.Sprintf(":%v", port))
	return &grpcAppServer{srv: srv, l: lis}
}

func NewHTTPAppServer(envCfg config.Envs) Server {
	mux := http.NewServeMux()
	mux.Handle("/v1/weather/", twirp.NewTwirpServer())
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", envCfg.Infra.Port),
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer stop()

	envCfg := config.Load()

	httpSrv := NewHTTPAppServer(envCfg)
	grpcSrv := NewGRPCAppServer(envCfg)

	if envCfg.App.IsHTTP() {
		go func() {
			log.Println("serving http")
			if err := httpSrv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Panic(err)
			}
		}()
	}
	if envCfg.App.IsGRPC() {
		go func() {
			log.Println("serving grpc")
			if err := grpcSrv.ListenAndServe(); err != nil {
				log.Panic(err)
			}
		}()
	}

	if envCfg.App.IsDev() {
		go func() {
			log.Println("serving http")
			if err := httpSrv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Panic(err)
			}
		}()
		go func() {
			log.Println("serving grpc")
			if err := grpcSrv.ListenAndServe(); err != nil {
				log.Panic(err)
			}
		}()
	}

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
	_ = grpcSrv.Shutdown(ctx)
}
