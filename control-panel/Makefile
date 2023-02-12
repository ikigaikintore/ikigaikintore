.PHONY: front-api-generate-v1 front-dev-start

OPENAPICLI_VERSION = 6.2.1
OAPICODEGEN_VERSION = 1.11.0
AIR_VERSION = 1.40.4

help: ## Show this help
	@echo "Help"
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[93m %s\n", $$1, $$2}'


front-api-generate-v1: ## Generate API V1 for frontend
	@echo "front-api-generate-v1";
	declare -A APIV1FILES=( ["schema"]="schema.yaml" )\
	; for apitype in "$${!APIV1FILES[@]}" ; do \
		filename=$${APIV1FILES[$${apitype}]}; \
		echo "Cleaning API v1 old files for "$${filename} && \
		rm -rf ./front/src/api/v1/panel && \
		mkdir -p ./api/schemas/v1/apigen && \
		echo "Creating API v1 for "$${filename} && \
		bash runtime_container.sh run -v ./api/schemas/v1:/tmp/ docker.io/openapitools/openapi-generator-cli:v$(OPENAPICLI_VERSION) generate -g typescript-axios -i /tmp/$${filename} -o /tmp/apigen/$${apitype} && \
		mkdir -p ./front/src/api/v1 && \
		echo "Moving file "$${filename}" from ./api/schemas/apigen/v1/"$${apitype}" to ./front/src/api/v1/"$${apitype} && \
		mv -f ./api/schemas/v1/apigen/$${apitype} ./front/src/api/v1/panel && \
		sleep 2; \
	done ; \
	rm -rf ./api/schemas/v1/apigen && \
	rm -rf ./api/schemas/hsperfdata_root

front-dev-start: ## Start server for dev environment
	(cd ./front && yarn dev)

back-api-generate-v1: ## Generate API V1 for backend
	@echo "back-api-generate-v1";
	declare -A APIV1FILES=( ["schema"]="schema.yaml" )\
	; for apitype in "$${!APIV1FILES[@]}" ; do \
		filename=$${APIV1FILES[$${apitype}]}; \
		echo "Cleaning API v1 old files for "$${filename} && \
		rm -rf ./back/api/v1/$${apitype} && \
		echo "Creating API v1 for "$${filename} && \
		mkdir -p ./api/schemas/v1/apigen && \
		bash runtime_container.sh run -v ./api/schemas/v1/:/tmp  docker.io/ervitis/oapi-codegen:v$(OAPICODEGEN_VERSION) oapi-codegen --config /tmp/oapiconfig.yaml /tmp/$${filename} > ./api/schemas/v1/apigen/server$${apitype}.gen.go && \
		mkdir -p ./back/api/v1 && \
		echo "Moving file "$${filename}" from ./api/schemas/v1/apigen/server"$${apitype}".gen.go to ./back/api/v1/server"$${apitype}".gen.go" && \
		mv -f ./api/schemas/v1/apigen/server$${apitype}.gen.go ./back/api/v1/server$${apitype}.gen.go && \
		sleep 2; \
	done ; \
	rm -rf ./api/schemas/v1/apigen && \
	rm -rf ./api/schemas/v1/hsperfdata_root

copy-grpc-v1:


back-dev-start: ## Start server with live reloading
	bash runtime_container.sh build -t control-panel:latest ./back && \
	bash runtime_container.sh run --rm -v ./back:/go/src/github.com/ervitis/control-panel -p 8080:8080 control-panel:latest
