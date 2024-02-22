package main

import (
	"golang.org/x/time/rate"
	"sync"
	"testing"
)

func Test_ipRateLimiter_ipAddress(t *testing.T) {
	type fields struct {
		ips        map[string]*rate.Limiter
		tokenPerIp int
		mtx        sync.Locker
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
				mtx:        &sync.Mutex{},
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
				mtx:        &sync.Mutex{},
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
				mtx:        tt.fields.mtx,
				rl:         tt.fields.rl,
			}
			if got := ir.ipAddress(tt.args.ip); got != tt.want {
				t.Errorf("ipAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
