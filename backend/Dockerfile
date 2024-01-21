FROM golang:1.21 as builder

WORKDIR /ikigaikintore

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /ikigaikintore/iki.app -ldflags="-s -w" cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-builder

COPY --from=builder /ikigaikintore/iki.app /
EXPOSE 8999
CMD ["/iki.app"]

FROM gcr.io/distroless/static-debian12:nonroot-amd64 as app-deploy

WORKDIR /
COPY iki.app /iki.app
EXPOSE 8999
ENTRYPOINT ["/iki.app"]