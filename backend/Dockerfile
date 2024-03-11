ARG GO_VERSION=1.22

FROM golang:${GO_VERSION} as builder
WORKDIR /ikigaikintore
COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /ikigaikintore/ikigai.app -ldflags="-s -w" cmd/server/main.go

FROM golang:${GO_VERSION} as dev
WORKDIR /tmp/app
COPY ./libs /tmp/libs
COPY ./backend /tmp/app
EXPOSE 8999 18999
RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /tmp/ikigai.app -gcflags="all=-N -l" /tmp/app/cmd/server/main.go
CMD ["/tmp/ikigai.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-builder
COPY --from=builder /ikigaikintore/ikigai.app /app/ikigai.app
EXPOSE 8080
CMD ["/app/ikigai.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy
WORKDIR /app
COPY /tmp/ikigai.app* /app/ikigai.app
EXPOSE 8080
CMD ["/app/ikigai.app"]
