package service

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ikigaikintore/ikigaikintore/proxybot/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

func BackendClient(envs config.Envs) *grpc.ClientConn {
	cred := insecure.NewCredentials()
	interceptors := []grpc.UnaryClientInterceptor{loggerInterceptor()}
	if !envs.App.IsDev() {
		certPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}
		cred = credentials.NewTLS(&tls.Config{RootCAs: certPool, MinVersion: tls.VersionTLS13})
		interceptors = append(interceptors, tokenSetInterceptor())
	}
	addr := envs.App.TargetBackend
	if len(strings.Split(addr, ":")) < 2 && !envs.App.IsDev() {
		addr = addr + ":443"
	}
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(cred),
		grpc.WithChainUnaryInterceptor(interceptors...),
	)
	if err != nil {
		panic(err)
	}
	return conn
}
