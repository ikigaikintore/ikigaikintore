FROM golang:1.22 as builder
WORKDIR /ikigaikintore

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /ikigaikintore/ikigai.app -ldflags="-s -w" cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-builder
COPY --from=builder /ikigaikintore/ikigai.app /
EXPOSE 8999
CMD ["/ikigai.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy
WORKDIR /
COPY /tmp/ikigai.app* /ikigai.app
EXPOSE 8999
CMD ["/ikigai.app"]
