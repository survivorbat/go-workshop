TEST?=$$(go list ./... |grep -v 'vendor')

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

r: run
run: ## Run the program, alias: r
	gin --appPort 8080 -i

t: test
test: ## Run unit tests, alias: t
	@go test -i $(TEST) || exit 1
	@echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

.PHONY: run r t test
