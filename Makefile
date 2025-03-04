# SPDX-License-Identifier: GPL-3.0-or-later

TEMP_DIR := ./.tmp

# Tool Versions #################################
GOSIMPORTS_VERSION := v0.3.8
GOLICENSES_VERSION := v5.0.1

# Command Templates #################################
GOIMPORTS_CMD := $(TEMP_DIR)/gosimports -local github.com/khulnasoft

# Formatting Variables #################################
BOLD := $(shell tput -T linux bold)
PURPLE := $(shell tput -T linux setaf 5)
GREEN := $(shell tput -T linux setaf 2)
CYAN := $(shell tput -T linux setaf 6)
RED := $(shell tput -T linux setaf 1)
RESET := $(shell tput -T linux sgr0)
TITLE := $(BOLD)$(PURPLE)
SUCCESS := $(BOLD)$(GREEN)

DEV_MODULES := all

# Main Targets #################################
.PHONY: all
all: static-analysis test ## Run all checks
	@printf '$(SUCCESS)All checks pass!$(RESET)\n'

.PHONY: static-analysis
static-analysis: check-go-mod-tidy check-licenses ## Run static analysis checks

.PHONY: test
test: unit  ## Run all tests

## Testing targets #################################

.PHONY: unit
unit: $(TEMP_DIR)  ## Run unit tests (with coverage)
	$(call title,Running unit tests)
	go test -coverprofile $(TEMP_DIR)/unit-coverage-details.txt $(shell go list ./... | grep -v anchore/syft/test)
	@.github/scripts/coverage.py $(COVERAGE_THRESHOLD) $(TEMP_DIR)/unit-coverage-details.txt

# Bootstrapping Targets #################################
.PHONY: bootstrap
bootstrap: $(TEMP_DIR) bootstrap-go bootstrap-tools ## Setup dependencies
	@echo "Bootstrapping dependencies..."

.PHONY: bootstrap-tools
bootstrap-tools: $(TEMP_DIR)
	GO111MODULE=on GOBIN=$(realpath $(TEMP_DIR)) go get -u golang.org/x/perf/cmd/benchstat
	curl -sSfL https://raw.githubusercontent.com/khulnasoft/go-licenses/master/golicenses.sh | sh -s -- -b $(TEMP_DIR)/ $(GOLICENSES_VERSION)
	GOBIN="$(realpath $(TEMP_DIR))" go install github.com/rinchsan/gosimports/cmd/gosimports@$(GOSIMPORTS_VERSION)

.PHONY: bootstrap-go
bootstrap-go:
	go mod download

$(TEMP_DIR):
	mkdir -p $(TEMP_DIR)

# Static Analysis #################################
.PHONY: format
format: ## Format source code
	@echo "Running formatting..."
	gofmt -w -s .
	$(GOIMPORTS_CMD) -w .
	go mod tidy

.PHONY: check-licenses
check-licenses: ## Validate license compliance
	@echo "Checking for license compliance..."
	$(TEMP_DIR)/golicenses check ./...

.PHONY: check-go-mod-tidy
check-go-mod-tidy:
	@ .github/scripts/go-mod-tidy-check.sh && echo "go.mod and go.sum are tidy!"

.PHONY: check
check: fmt vet ## Run static code analysis

.PHONY: fmt
fmt:
	@echo "Running formatting..."
	$(FMT_CMD) .
	$(GOIMPORTS_CMD) .
	go mod tidy
	git diff --exit-code

.PHONY: vet
vet:
	go vet ./...

# Build & Clean #################################
.PHONY: build
build: clean ## Build package
	hack/go-build.sh

.PHONY: clean
clean:
	rm -rf bin vendor

.PHONY: release
release: clean download ## Create release artifacts
	hack/go-build.sh all
	hack/go-build.sh configs
	hack/go-build.sh vendor
	cd bin && sha256sum -b * >"sha256sums.txt"

# Development #################################
.PHONY: dev
dev: dev-build dev-up ## Start development build

dev-build:
	docker-compose build

dev-up:
	docker-compose up -d --remove-orphans

.PHONY: dev-exec
dev-exec: ## Enter development environment
	docker-compose exec khulnasoft bash

.PHONY: dev-log
dev-log:
	docker-compose logs -f khulnasoft

.PHONY: dev-run
dev-run: ## Run go.plugin inside development environment
	go run github.com/khulnasoft/go.plugin/cmd/goplugin -d -c conf.d

.PHONY: dev-mock
dev-mock: ## Run go.plugin with mock config
	go run github.com/khulnasoft/go.plugin/cmd/goplugin -d -c ./mocks/conf.d -m $(DEV_MODULES)

# Help #################################
.PHONY: help
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
