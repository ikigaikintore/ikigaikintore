package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func Test_ipRateLimiter_ipAddress(t *testing.T) {
	type fields struct {
		ips        map[string]*rate.Limiter
		tokenPerIp int
		rl         rate.Limit
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "IP address with port",
			fields: fields{
				ips:        make(map[string]*rate.Limiter),
				tokenPerIp: 10,
				rl:         1,
			},
			args: args{
				ip: "192.168.2.3:9098",
			},
			want: "192.168.2.3",
		},
		{
			name: "IP address without port",
			fields: fields{
				ips:        make(map[string]*rate.Limiter),
				tokenPerIp: 10,
				rl:         1,
			},
			args: args{
				ip: "192.168.2.3",
			},
			want: "192.168.2.3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir := &ipRateLimiter{
				ips:        tt.fields.ips,
				tokenPerIp: tt.fields.tokenPerIp,
				mtx:        sync.RWMutex{},
				rl:         tt.fields.rl,
			}
			if got := ir.ipAddress(tt.args.ip); got != tt.want {
				t.Errorf("ipAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipRateLimiterMid(t *testing.T) {
	type args struct {
		cl *ipRateLimiter
	}
	tests := []struct {
		name    string
		n       int
		args    args
		checker func(*httptest.ResponseRecorder, *http.Request) error
	}{
		{
			name: "not block",
			args: args{
				cl: newIpRateLimiter(),
			},
			n: 1,
			checker: func(rr *httptest.ResponseRecorder, req *http.Request) error {
				if rr.Code != http.StatusOK {
					return fmt.Errorf("not 200 code")
				}
				return nil
			},
		},
		{
			name: "block",
			args: args{
				cl: newIpRateLimiter(withLimit(1), withTokensPerIp(1)),
			},
			n: 10,
			checker: func(rr *httptest.ResponseRecorder, req *http.Request) error {
				if rr.Code == http.StatusTooManyRequests {
					fmt.Println("ip blocked")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.checker == nil {
				t.Errorf("checkere is nil")
			}
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			mid := ipRateLimiterMid(tt.args.cl)(next)

			req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
			req.RemoteAddr = "128.2.2.2:7777"

			for i := 0; i < tt.n; i++ {
				rr := httptest.NewRecorder()
				if i == 0 {
					if rr.Code != http.StatusOK {
						t.Errorf("code not 200")
					}
				}

				mid.ServeHTTP(rr, req)
				if err := tt.checker(rr, req); err != nil {
					t.Error(err)
				}
			}
		})
	}
}
