COMMIT = $(shell git describe --always)
PACKAGES = $(shell go list ./... | grep -v '/vendor/')
VERSION = $(shell grep 'Version string' version.go | sed -E 's/.*"(.+)"$$/\1/')

default: build

build: 
	go build -ldflags "-X main.GitCommit=$(COMMIT)" -o bin/twg

run:
	go run *.go

deps:
	dep ensure -update

deps_update_all:
	rm lock.json & rm -r vendor & dep ensure -update

test-all: vet lint test

test: 
	go test -v -parallel=4 ${PACKAGES}

vet:
	go vet ${PACKAGES}

lint:
	@go get github.com/golang/lint/golint
	go list ./... | grep -v vendor | xargs -n1 golint
