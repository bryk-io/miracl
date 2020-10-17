.PHONY: *
.DEFAULT_GOAL:=help

## help: Prints this help message
help:
	@echo "Commands available"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | sort

## build: Build the AMCL library from the latest commit on the fork repo
build:
	@./build.sh

## test: Run all tests available on the AMCL library
test:
	go test -v ./...

## benchmark: Run all benchmarks available on the AMCL library
benchmark:
	go test -v -bench=. ./...

## info: Show package information.
info:
	@go run main.go
