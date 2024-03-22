ARG GO_VERSION=1.22

FROM golang:${GO_VERSION} as builder
ARG GO_VERSION
WORKDIR /ikigaikintore
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /proxybot/proxybot.app -ldflags="-s -w" cmd/server/main.go

FROM golang:${GO_VERSION} as dev
ARG GO_VERSION
WORKDIR /tmp/app
COPY ./libs /tmp/libs
COPY ./proxybot /tmp/app
EXPOSE 7080 17080 8080
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /tmp/proxybot.app -gcflags="all=-N -l" /tmp/app/cmd/server/main.go
CMD ["/tmp/proxybot.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-builder
COPY --from=builder /proxybot/proxybot.app /
EXPOSE 8080
CMD ["/proxybot.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy
WORKDIR /
COPY /tmp/proxybot.app* /proxybot.app
EXPOSE 8080
CMD ["/proxybot.app"]