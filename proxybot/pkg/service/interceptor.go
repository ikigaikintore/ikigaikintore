package service

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	gOpt "google.golang.org/api/option"
	"google.golang.org/grpc"
	grpc_metadata "google.golang.org/grpc/metadata"
)

func tokenSetInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		creds, err := google.FindDefaultCredentials(ctx)
		if err != nil {
			fmt.Println("cannot find credentials: ", err)
			return err
		}
		target := strings.Split(cc.Target(), ":")[0]
		if !strings.HasPrefix(target, "https://") {
			target = "https://" + target
		}
		tSource, err := idtoken.NewTokenSource(ctx, target, gOpt.WithCredentials(creds))
		if err != nil {
			fmt.Println("error creating token source:", err)
			return err
		}
		token, err := tSource.Token()
		if err != nil {
			fmt.Println("cannot create token: ", err)
			return err
		}
		ctx = grpc_metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func loggerInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println(method, cc.Target(), ctx)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
