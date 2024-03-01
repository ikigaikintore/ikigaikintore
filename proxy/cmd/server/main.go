package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"
	"syscall"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"

	"github.com/ikigaikintore/ikigaikintore/libs/cors"
	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/config"
)

func NewProxy(target string, authClient *auth.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := authClient.VerifyIDToken(r.Context(), strings.TrimSpace(strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1)))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		emailRaw, ok := token.Claims["email"]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		email, ok := emailRaw.(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !slices.Contains(strings.Split(os.Getenv("NEXT_PUBLIC_USER_AUTH"), ","), email) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req, err := http.NewRequest(r.Method, target+r.URL.String(), r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req.Header = r.Header
		req.Header.Set("X-Forwarded-For", r.RemoteAddr)
		req.Header.Set("X-Forwarded-Host", r.Host)

		client := http.DefaultTransport

		resp, err := client.RoundTrip(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		w.WriteHeader(resp.StatusCode)
		_, _ = io.Copy(w, resp.Body)
	})
}

type ipRateLimiter struct {
	ips        map[string]*rate.Limiter
	tokenPerIp int
	mtx        sync.Locker
	rl         rate.Limit
}

type option func(*ipRateLimiter)

func withTokensPerIp(tokens int) option {
	return func(limiter *ipRateLimiter) {
		limiter.tokenPerIp = tokens
	}
}

func withLimit(limit rate.Limit) option {
	return func(limiter *ipRateLimiter) {
		limiter.rl = limit
	}
}

func newIpRateLimiter(opts ...option) *ipRateLimiter {
	def := &ipRateLimiter{
		tokenPerIp: 10,
		mtx:        &sync.Mutex{},
		rl:         1,
	}

	for _, opt := range opts {
		opt(def)
	}
	def.ips = make(map[string]*rate.Limiter)
	return def
}

func (ir *ipRateLimiter) ipAddress(ip string) string {
	// ip address has :
	if idx := strings.Index(ip, ":"); idx != -1 {
		// 192.168.2.3:9098 -> idx = 11
		ip = ip[:idx]
	}
	return ip
}

func (ir *ipRateLimiter) addIp(ip string) *rate.Limiter {
	ir.mtx.Lock()
	lim := rate.NewLimiter(ir.rl, ir.tokenPerIp)
	ir.ips[ip] = lim
	ir.mtx.Unlock()
	return lim
}

func (ir *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
	ir.mtx.Lock()
	defer ir.mtx.Unlock()
	lim, ok := ir.ips[ip]
	if !ok {
		return ir.addIp(ip)
	}
	return lim
}

func ipRateLimiterMid(cl *ipRateLimiter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lim := cl.getLimiter(cl.ipAddress(r.RemoteAddr))
			if !lim.Allow() {
				fmt.Printf("ip blocked: %s %s %s", r.RemoteAddr, r.Method, r.URL)
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	envCfg := config.Load()

	opts := make([]cors.Option, 0)
	if envCfg.App.IsDev() {
		opts = append(opts, cors.LocalEnvironment())
	}
	opts = append(opts, cors.WithAllowedDomains(strings.Split(envCfg.Cors.AllowedDomains, ",")...))

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: os.Getenv("PROJECT_ID"),
	})
	if err != nil {
		panic(err)
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	limCli := newIpRateLimiter()

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           cors.DomainAllowed(cors.NewHandler(opts...), ipRateLimiterMid(limCli)(logger(NewProxy(envCfg.App.TargetBackend, authClient)))),
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
