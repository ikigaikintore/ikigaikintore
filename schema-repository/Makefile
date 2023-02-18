.PHONY: generate-grpc-v1

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'


CROSSFIT_AGENDA_PB_FOLDER = $(PWD)/../crossfitAgenda/adapters/handlers/grpc/


generate-grpc-v1: ## Generate the server and client from proto file
	docker run \
	--rm -v ${PWD}:/defs namely/protoc-all \
	-f ./proto/**/*.proto \
	-l go \
	--go_out=plugins=grpc:/gen/grpc && \
	sleep 1 && \
	mv ./gen/pb-go/* $(CROSSFIT_AGENDA_PB_FOLDER) && \
	sleep 1 && \
	rm -rf ./gen