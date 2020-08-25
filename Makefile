export GO111MODULE=on

SHELL=/bin/bash -o pipefail

GO ?= go
.PHONY: all
all: application
	@echo 'Build successful'

.PHONY: application
application:
	$(GO) build -v  main.go