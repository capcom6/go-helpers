# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

all: deps lint test

test: 
	$(GOTEST) -race -coverprofile=coverage.out -covermode=atomic ./...

clean: 
	$(GOCLEAN)
	rm -f coverage.out

deps:
	$(GOMOD) download

lint:
	golangci-lint run -v

.PHONY: all test clean deps lint
