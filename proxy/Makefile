.PHONY: install run fmt tests lint

BUILD_FLAGS := -ldflags="-s -w"
BUILD_ENV := CGO_ENABLED=0 GOARCH="amd64" GOOS="linux"


help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

install: ## Install dependencies and clean
	go mod tidy

run: ## Run server
	go install github.com/cosmtrek/air@latest
	air server -d --port 8999

fmt: ## Format code
	go install golang.org/x/tools/cmd/goimports@latest
	find . -type f -name '*.go' | xargs goimports -w && go fmt `go list ./...`
	go mod tidy

compile: ## Compile project
	@mkdir -p tmp
	$(BUILD_ENV) go build -o ./tmp/ikigai.app $(BUILD_FLAGS) ./cmd/server/main.go

tests: ## Execute tests
	go test -race -v ./...

lint: ## Execute golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./... --new-from-rev=origin/main --exclude-use-default=false --timeout 10m
