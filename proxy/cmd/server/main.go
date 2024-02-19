package server

import (
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewProxy(target string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, r.Body)
	}))
	return mux
}

func main() {
	target := os.Getenv("PROXY_TARGET_BACKEND")
	if target == "" {
		panic("no target")
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app, err := firebase.NewApp(ctx)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           NewProxy(target),
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ch
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
