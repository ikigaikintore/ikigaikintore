package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikigaikintore/ikigaikintore/backend/internal/config"
	"github.com/ikigaikintore/ikigaikintore/backend/internal/input/grpc"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func NewAppServer(envCfg config.Envs) Server {
	mux := http.NewServeMux()
	mux.Handle("/v1/weather/", grpc.NewTwirpServer())
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
