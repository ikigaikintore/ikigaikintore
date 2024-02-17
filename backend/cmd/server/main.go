package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ervitis/ikigaikintore/backend/internal/input/grpc"
	"github.com/rs/cors"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func NewAppServer() Server {
	mux := http.NewServeMux()
	mux.Handle("/v1/weather/", grpc.NewTwirpServer())
	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	return &http.Server{
		Addr:              ":8999",
		Handler:           cors.Default().Handler(mux),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}
}

func NewFileServer() Server {
	return &http.Server{
		Addr:              ":3000",
		Handler:           cors.Default().Handler(http.StripPrefix("/", http.FileServer(http.Dir("/static")))),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       time.Minute,
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := NewAppServer()
	fileSrv := NewFileServer()

	go func() {
		log.Println("serving app")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}()

	go func() {
		log.Println("serving static files")
		if err := fileSrv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
	if err := fileSrv.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
	}
}
