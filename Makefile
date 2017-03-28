SOURCEDIR=.
SOURCES = $(shell find $(SOURCEDIR) -name '*.go')
VERSION=$(shell git describe --always --tags)

.PHONY: build
build: go-get test vet

.PHONY: test
test:
	go test -cover -v ./...

.PHONY: go-get
go-get:
	go get github.com/google/go-querystring/query
	go get -u github.com/golang/lint/golint

.PHONY: vet
vet:
	go vet ./...
