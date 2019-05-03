#
# @author Jose Nidhin
#
VERSION := "0.0.1"
PROJECT_NAME := $(shell basename "$(PWD)")
GO_SRC_FILES := $(shell find . -type f -name '*.go')
GO_SRC_MAIN := $(shell ls *.go)

LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

.PHONY: fmt
fmt:
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: simplify
simplify:
	gofmt -s -l -w $(GO_SRC_FILES)

.PHONY: build
build:
	go build $(LDFLAGS) -o $(PROJECT_NAME) $(GO_SRC_MAIN)

.PHONY: clean
clean:
	rm -f $(PROJECT_NAME)
