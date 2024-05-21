# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test

test: 
	$(GOTEST) -v -race -coverprofile=coverage.out -covermode=atomic ./...
