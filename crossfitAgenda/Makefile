.PHONY: tests

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

install: ## Install dependencies
	go mod tidy

generate-api-v1: generate-http-api-v1 install ## Generate http API v1

generate-http-api-v1: ## Generates the endpoint for full v1 API
	rm -f ./service/handlers/http/server.gen.go && \
	mkdir -p ./out/api/ && \
	docker run \
	--rm -v ${PWD}:/local \
	docker.io/ervitis/oapi-codegen:v1.12.4 oapi-codegen --config /local/service/handlers/oapiconfig.yaml /local/service/handlers/schema.yml > ./out/api/server.gen.go && \
	sleep 3 && \
	mv ./out/api/server.gen.go ./service/handlers/http/ && \
	sleep 1 && rm -rf ./out
