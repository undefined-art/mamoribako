# Makefile for Go CLI Password Manager

BINARY_NAME := mamoribako
SRC := ./...

.PHONY: all build test run clean fmt

all: build

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test -v $(SRC)

run:
	go run main.go

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt $(SRC)
