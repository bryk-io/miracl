.PHONY: all
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
	go run ./TestECC.go
	go run ./TestBLS.go
	go run ./TestHPKE.go
	go run ./TestNHS.go
	go run ./TestMPIN.go < pins.txt

## benchmark: Run all benchmarks available on the AMCL library
benchmark:
	go run ./BenchtestALL.go

## info: Show package information.
info:
	@go run main.go
