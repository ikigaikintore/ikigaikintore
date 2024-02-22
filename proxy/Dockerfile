FROM golang:1.22 as builder
WORKDIR /ikigaikintore

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /proxy/proxy.app -ldflags="-s -w" cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-builder
COPY --from=builder /proxy/proxy.app /
EXPOSE 8080
CMD ["/proxy.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy
WORKDIR /
COPY /tmp/proxy.app* /proxy.app
EXPOSE 8080
CMD ["/proxy.app"]
