FROM golang:1.21

ARG PROTOCGEN_VERSION="1.31.0"
ARG PROTOCGENTWIRP_VERSION="8.1.3"

RUN apt update && \
    apt install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOCGEN_VERSION} && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@v${PROTOCGENTWIRP_VERSION} && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

CMD find . -name "*.proto" -exec protoc --go_out=:./proto --go_opt=paths=source_relative --twirp_out=:. {} \;