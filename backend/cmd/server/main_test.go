package main

import (
	"fmt"
	"github.com/ikigaikintore/ikigaikintore/backend/internal/config"
	"github.com/ikigaikintore/ikigaikintore/libs/cors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func getBody(res *http.Response) string {
	if res == nil {
		return ""
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("readbody err", err)
	}
	return string(b)
}

func Test_corsHandler(t *testing.T) {
	type args struct {
		method         string
		allowedDomains []string
		host           string
		originHeader   string
	}
	tests := []struct {
		name    string
		args    args
		checker func(rr *httptest.ResponseRecorder) error
	}{
		{
			name: "success",
			args: args{
				method:         http.MethodGet,
				allowedDomains: []string{"kintore.com"},
				host:           "http://kintore.com:80",
				originHeader:   "http://kintore.com:80",
			},
			checker: func(rr *httptest.ResponseRecorder) error {
				s := getBody(rr.Result())
				if rr.Code != http.StatusOK {
					return fmt.Errorf("http code not 200 %v: %v", rr.Code, s)
				}
				return nil
			},
		},
		{
			name: "ko: domain not allowed",
			args: args{
				method:         http.MethodGet,
				allowedDomains: []string{"kintore1.com"},
				host:           "http://kintore.com:80",
				originHeader:   "http://kintore.com:8650",
			},
			checker: func(rr *httptest.ResponseRecorder) error {
				s := getBody(rr.Result())
				if rr.Code != http.StatusMethodNotAllowed {
					return fmt.Errorf("domain allowed when it should not %v: %v", rr.Code, s)
				}
				return nil
			},
		},
		{
			name: "ko: origin header is empty",
			args: args{
				method:         http.MethodGet,
				allowedDomains: []string{"kintore1.com"},
				host:           "http://kintore.com:80",
				originHeader:   "",
			},
			checker: func(rr *httptest.ResponseRecorder) error {
				s := getBody(rr.Result())
				if rr.Code != http.StatusMethodNotAllowed {
					return fmt.Errorf("domain allowed when it should not %v: %v", rr.Code, s)
				}
				return nil
			},
		},
		{
			args: args{
				method:         http.MethodGet,
				host:           "https://ikigaipanel-hello-uc.a.run.app:8080",
				allowedDomains: []string{"ikigaipanel-hello-uc.a.run.app"},
				originHeader:   "https://ikigaipanel-hello-uc.a.run.app:8080",
			},
			name: "ok in cloud run",
			checker: func(rr *httptest.ResponseRecorder) error {
				s := getBody(rr.Result())
				if rr.Code != http.StatusOK {
					return fmt.Errorf("http code not 200 %v: %v", rr.Code, s)
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BACKEND_CORS_ALLOWED_DOMAINS", strings.Join(tt.args.allowedDomains, ","))
			t.Setenv("BACKEND_APP_ENV", "staging")
			cfg := config.Load()
			opts := make([]cors.Option, 0)
			if cfg.App.IsDev() {
				opts = append(opts, cors.LocalEnvironment())
			}
			opts = append(opts, cors.WithAllowedDomains(strings.Split(cfg.Cors.AllowedDomains, ",")...))
			got := cors.NewHandler(opts...)
			req, _ := http.NewRequest(tt.args.method, tt.args.host+"/foo", nil)
			req.Host = tt.args.host
			req.Header.Set("Origin", tt.args.originHeader)
			rr := httptest.NewRecorder()
			handler := cors.DomainAllowed(got, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			handler.ServeHTTP(rr, req)
			if tt.checker == nil {
				t.Skip("skipped because checker func is nil")
			}
			if err := tt.checker(rr); err != nil {
				t.Errorf("checker error: %v", err)
			}
		})
	}
}
