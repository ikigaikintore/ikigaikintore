FROM docker.io/cosmtrek/air:v1.40.4 AS base

WORKDIR /go/src/github.com/ikigaikintore/ikigaikintore/control-panel/back

COPY . .

EXPOSE 8308
EXPOSE 8309

VOLUME /go/src/github.com/ikigaikintore/ikigaikintore/control-panel/back

RUN CGO_ENABLED=0 go build -gcflags 'all=-N -l' -o ./cmd/main ./cmd
