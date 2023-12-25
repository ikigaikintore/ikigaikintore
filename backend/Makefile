.PHONY: generate

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

generate: ## Generate go files
	go generate ./...
	make install

install: ## Install dependencies and clean
	go mod tidy

run: ## Run server
	go install github.com/cosmtrek/air@latest
	air server -d --port 8999