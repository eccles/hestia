
PWD := $(shell pwd )

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: tools-check
tools-check: ## Check that all tools are available
	@./scripts/tools/check.sh

.PHONY: tools-install
tools-install: ## Ensure that all tools are available
	@./scripts/tools/install.sh

.PHONY: unittests
unittests: ## unittest all source code
	@./scripts/unittests.sh

.PHONY: build
build: ## Compile all source code
	@./scripts/build.sh
