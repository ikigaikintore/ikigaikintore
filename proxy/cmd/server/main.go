package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"
	"syscall"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"golang.org/x/oauth2/google"
	"golang.org/x/time/rate"
	"google.golang.org/api/idtoken"
	gOpt "google.golang.org/api/option"

	"github.com/ikigaikintore/ikigaikintore/libs/cors"
	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/bot"
	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/config"
)

func newReverseProxy(target *url.URL, token string) *httputil.ReverseProxy {
	if credentials, err := google.FindDefaultCredentials(context.Background()); err != nil {
		fmt.Println("FindDefaultCredentials", err)
	} else {
		if ts, err := idtoken.NewTokenSource(context.Background(), target.String(), gOpt.WithCredentials(credentials)); err != nil {
			fmt.Println("NewTokenSource", err)
		} else {
			if t, err := ts.Token(); err != nil {
				fmt.Println("Token", err)
			} else {
				token = t.AccessToken
			}
		}
	}
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Header.Set("X-Forwarded-For", req.RemoteAddr)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
		req.Header.Set("X-Forwarded-Uri", req.URL.Path)
		req.Host = target.Host
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
	}
	return &httputil.ReverseProxy{Director: director, Transport: http.DefaultTransport}
}

func NewProxy(envConfig config.Envs, authClient *auth.Client, botClient bot.Listener) http.Handler {
	pr := http.NewServeMux()
	pr.HandleFunc("/v1/*", func(w http.ResponseWriter, r *http.Request) {
		validateToken := func(r *http.Request, rawToken string) bool {
			token, err := authClient.VerifyIDToken(r.Context(), rawToken)
			if err != nil {
				log.Println("token error ", err)
				return false
			}

			emailRaw, ok := token.Claims["email"]
			if !ok {
				log.Println("email not in token")
				return false
			}

			email, ok := emailRaw.(string)
			if !ok {
				log.Println("email not found in token")
				return false
			}
			if !slices.Contains(strings.Split(os.Getenv("NEXT_PUBLIC_USER_AUTH"), ","), email) {
				log.Println("email not valid in env")
				return false
			}
			return true
		}
		rawToken := strings.TrimSpace(strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1))
		if rawToken != "" && !validateToken(r, rawToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		scheme := "http"
		if !envConfig.App.IsDev() {
			scheme = "https"
		}
		urlTarget := &url.URL{
			Scheme: scheme,
			Host:   envConfig.App.TargetBackend,
			Path:   r.URL.Path,
		}
		fmt.Println("target", urlTarget.String(), urlTarget.Scheme, urlTarget.Host, urlTarget.Path, urlTarget.RequestURI())

		proxy := newReverseProxy(urlTarget, rawToken)

		proxy.ServeHTTP(w, r)
	})
	pr.HandleFunc(envConfig.Telegram.WebhookUriPathBase, func(w http.ResponseWriter, r *http.Request) {
		err := botClient.Parser(envConfig, r.Body)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			return
		}
		log.Println("parsing body: ", err)
		if errors.Is(err, bot.ErrForbidden) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	})
	return pr
}

type ipRateLimiter struct {
	ips        map[string]*rate.Limiter
	tokenPerIp int
	mtx        sync.RWMutex
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
		mtx:        sync.RWMutex{},
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
	lim := rate.NewLimiter(ir.rl, ir.tokenPerIp)
	ir.mtx.Lock()
	ir.ips[ip] = lim
	ir.mtx.Unlock()
	return lim
}

func (ir *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
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
		fmt.Println(r.RemoteAddr, r.Proto, r.Host, r.RequestURI, r.ContentLength, r.Method, r.URL)
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

	appBot, err := bot.NewBot(envCfg)
	if err != nil {
		log.Println("bot not loaded: ", err)
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", envCfg.Infra.Port),
		Handler: cors.NewHandler(opts...).Handler(
			ipRateLimiterMid(limCli)(logger(NewProxy(envCfg, authClient, appBot))),
		),
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	go func() {
		log.Println("serving telegram bot")
		appBot.Start()
	}()

	go func() {
		log.Println("serving proxy")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ch
	appBot.Stop()
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
