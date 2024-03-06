.PHONY: emulator

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

emulator: ## Start firebase emulator
	firebase emulators:start

emulator-debug: ## Start firebase emulator with debug
	firebase emulators:start --debug