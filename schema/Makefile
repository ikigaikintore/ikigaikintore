.PHONY: generate

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

generate: ## Generate automated code
	docker build --target backend -t backend-twirp-compiler:latest . && \
	docker run -v `pwd`:/app backend-twirp-compiler:latest && \
	rm -rf ../backend/pkg/proto && \
	mv proto ../backend/pkg/ && \
	docker build --target frontend -t frontend-twirp-compiler:latest . && \
	docker run -v `pwd`:/app frontend-twirp-compiler:latest && \
	rm -rf ../frontend/src/api/* && \
	mv *.pb.ts ../frontend/src/api/ && \
	docker run -v `pwd`:/defs namely/protoc-all -f ./*.proto -l go && \
	cp gen/pb-go/proto/*_grpc.pb.go ../backend/pkg/proto/ && \
	cp gen/pb-go/proto/*.pb.go ../proxybot/pkg/service/ && \
	rm -rf gen