SHELL = /usr/bin/env bash

# go environment settings
GO               ?= $(shell which go 2>/dev/null)
GOLANG_CI        ?= $(shell which golangci-lint 2>/dev/null)
GO_TEST_PACKAGES := $(sort $(dir $(shell find . -type f -name '*_test.go')))

# terminal colors
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all
all: help

## Coding Style

.PHONY: lint
lint: $(GOLANG_CI) ## lint all code
	@$(GOLANG_CI) run

## Testing

.PHONY: test
test: ## run unit tests
	@$(GO) test -v -race $(GO_TEST_PACKAGES)

## Help

.PHONY: help
help: ## show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		   if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
			 else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
			 }' $(MAKEFILE_LIST)
