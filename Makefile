SHELL := /bin/zsh

.PHONY: build swag-init run all

GOPATH := $(shell go env GOPATH)
BINARY := cmd/main

build:
	go build -o ./$(BINARY) .

swag-init:
	$(GOPATH)/bin/swag init

fmt:
	$(GOPATH)/bin/swag fmt

air:
	$(GOPATH)/bin/air -- -mode=dev -port=4000 -file=false

all: swag-init build
