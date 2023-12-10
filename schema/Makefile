.PHONY: generate

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

generate: ## Generate automated code
	docker build -t twirp-compiler:latest . && \
	docker run -v `pwd`:/app twirp-compiler:latest && \
	mv proto ../backend/pkg/input/