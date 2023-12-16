package main

import (
	"context"
	"errors"
	"github.com/ervitis/ikigaikintore/backend/internal/input/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func NewServer() Server {
	return &http.Server{
		Addr:              ":8999",
		Handler:           grpc.NewTwirpServer(),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := NewServer()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
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
