package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/cors"

	"github.com/ervitis/ikigaikintore/backend/internal/config"
	"github.com/ervitis/ikigaikintore/backend/internal/input/grpc"
)

type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func corsHandler(envCfg config.Envs) *cors.Cors {
	if envCfg.App.IsDev() {
		return cors.AllowAll()
	}
	return cors.New(
		cors.Options{
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodConnect,
				http.MethodOptions,
				http.MethodTrace,
			},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
			AllowOriginVaryRequestFunc: func(r *http.Request, origin string) (bool, []string) {
				if strings.TrimSpace(origin) == "" {
					return false, []string{}
				}
				parsedOrigin, err := url.Parse(origin)
				if err != nil {
					return false, []string{}
				}

				origin = parsedOrigin.Hostname()
				for _, v := range strings.Split(envCfg.Cors.AllowedDomains, ",") {
					if v == origin {
						return true, []string{"Origin"}
					}
				}

				return false, []string{}
			},
		},
	)
}

func domainAllowed(c *cors.Cors, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !c.OriginAllowed(r) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewAppServer(envCfg config.Envs) Server {
	mux := http.NewServeMux()
	mux.Handle("/v1/weather/", domainAllowed(corsHandler(envCfg), grpc.NewTwirpServer()))
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
