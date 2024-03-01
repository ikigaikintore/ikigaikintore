package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ervitis/ikigaikintore/libs/cors"

	"github.com/ervitis/ikigaikintore/backend/internal/config"
	"github.com/ervitis/ikigaikintore/backend/internal/input/grpc"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func NewAppServer(envCfg config.Envs) Server {
	mux := http.NewServeMux()
	opts := make([]cors.Option, 0)
	if envCfg.App.IsDev() {
		opts = append(opts, cors.LocalEnvironment())
	}
	opts = append(opts, cors.WithAllowedDomains(strings.Split(envCfg.Cors.AllowedDomains, ",")...))
	mux.Handle("/v1/weather/", cors.DomainAllowed(cors.NewHandler(opts...), grpc.NewTwirpServer()))
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	return &http.Server{
		Addr:              ":8080",
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

	srv := NewAppServer(envCfg)

	go func() {
		log.Println("serving app")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
}
