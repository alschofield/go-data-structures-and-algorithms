GO ?= go
NAME ?=
PACKAGE = ./src/$(NAME)
OLD ?=
NEW ?=

.DEFAULT_GOAL := help
.PHONY: help test test-all contract benchmark race benchmark-compare

help:
	@echo "Run safe default tests: make test-all"
	@echo "Run one leaf's default tests: make test NAME=data-structures/linear/stacks/stack"
	@echo "Run an implemented leaf's contracts: make contract NAME=data-structures/linear/stacks/stack"
	@echo "Benchmark an implemented leaf: make benchmark NAME=data-structures/linear/stacks/stack"
	@echo "Run race-enabled default tests: make race"
	@echo "Compare saved benchmark output: make benchmark-compare OLD=before.txt NEW=after.txt"

test:
	@if "$(NAME)"=="" (echo Set NAME to a taxonomy topic. & exit /b 1)
	$(GO) test $(PACKAGE) || $(GO) test ./...

test-all:
	$(GO) test ./...

contract:
	@if "$(NAME)"=="" (echo Set NAME to a taxonomy topic. & exit /b 1)
	$(GO) test -tags=contract $(PACKAGE)

benchmark:
	@if "$(NAME)"=="" (echo Set NAME to a taxonomy topic. & exit /b 1)
	$(GO) test -tags=contract -run '^$$' -bench . -benchmem $(PACKAGE)

race:
	$(GO) test -race ./...

benchmark-compare:
	@if "$(OLD)"=="" (echo Set OLD and NEW benchmark output files. & exit /b 1)
	@if "$(NEW)"=="" (echo Set OLD and NEW benchmark output files. & exit /b 1)
	benchstat $(OLD) $(NEW)
