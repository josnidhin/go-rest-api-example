#
# @author Jose Nidhin
#
VERSION := "0.0.1"
PROJECT_NAME := $(shell basename "$(PWD)")
GO_SRC_FILES := $(shell find . -type f -name '*.go')
GO_SRC_MAIN := $(shell ls *.go)

LDFLAGS=-ldflags "-X=main.AppVersion=$(VERSION) -X=main.AppName=$(PROJECT_NAME)"

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: simplify
simplify:
	gofmt -s -l -w $(GO_SRC_FILES)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	go clean -x

.PHONY: build
build:
	go build -i -v -o $(PROJECT_NAME) $(LDFLAGS) $(GO_SRC_MAIN)
