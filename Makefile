fmt: ## Format code
	go install golang.org/x/tools/cmd/goimports@latest
	find . -type f -name '*.go' | xargs goimports -w && go fmt `go list ./...`
	go mod tidy

lint: ## Execute golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./... --new-from-rev=origin/main --exclude-use-default=false --timeout 10m
