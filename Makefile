.DEFAULT_GOAL := help

# Target creators
.PHONY: lint
.PHONY: go_get go_test_integration
.PHONY: help

#
# Linter targets.
#
lint: ## Run set of preconfigured linters
	@docker build -f build/linter/Dockerfile -t golangci-linter . && docker run golangci-linter

#
# Go targets.
#
go_get: ## Get all project dependencies
	@echo '>>> Getting go modules.'
	go mod download

go_test_unit: ## Run unit tests only over Fake API client
	@echo ">>> Running unit tests."
	go test -v -tags unit ./...

go_test_integration: ## Run integration tests over Fake API client and Fake API
	@echo ">>> Running integration tests."
	go test -v -p 1 -tags="integration" ./tests/integration/...

help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo
