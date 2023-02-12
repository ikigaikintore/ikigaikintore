FROM docker.io/cosmtrek/air:v1.40.4 AS base

WORKDIR /go/src/github.com/ervitis/control-panel

COPY . .

EXPOSE 8080

VOLUME /go/src/github.com/ervitis/control-panel

RUN CGO_ENABLED=0 go build -gcflags 'all=-N -l' -o ./cmd/main ./cmd
