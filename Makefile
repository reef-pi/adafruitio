SOURCEDIR=.
SOURCES = $(shell find $(SOURCEDIR) -name '*.go')
VERSION=$(shell git describe --always --tags)


.PHONY: test
test:
	go test -cover -v ./...

.PHONY: go-get
go-get:
	go get github.com/google/go-querystring

.PHONY: vet
vet:
	go tool vet -shadow ./...

.PHONY: build
build: go-get test bin
