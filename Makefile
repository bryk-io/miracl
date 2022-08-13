.PHONY: *
.DEFAULT_GOAL:=help

help:
	@echo "Commands available"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /' | sort

## update: Update the AMCL library from the latest commit on the fork repo
update:
	@./build.sh

## test: Run all tests available on the AMCL library
test:
	go test -v ./...

## benchmark: Run all benchmarks available on the AMCL library
benchmark:
	go test -v -bench=. ./...

## version: Show package version
version:
	@go run main.go
