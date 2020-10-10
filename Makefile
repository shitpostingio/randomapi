# per gentile concessione: https://gist.github.com/subfuzion/0bd969d08fe0d8b5cc4b23c795854a13

SHELL := /bin/bash

TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

VERSION := $(shell git describe --tags --abbrev=0)
BUILD := $(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean install uninstall fmt simplify check run

all: check build

$(TARGET): $(SRC)
	$(info Building $(TARGET) v${VERSION} Build ${BUILD})
	@env CGO_CFLAGS_ALLOW="-L(.*)|-I(.*)" go build $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	$(info Building $(TARGET) v${VERSION} Build ${BUILD})
	@env CGO_CFLAGS_ALLOW="-L(.*)|-I(.*)" go install $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

run: install
	@$(TARGET)GOBUILD=go