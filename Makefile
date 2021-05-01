export GO111MODULE=on

GOFLAGS := -v -modcacherw
EXTRA_GOFLAGS ?=
LDFLAGS := $(LDFLAGS)
GO ?= go

SOURCES ?= $(shell find . -name "*.go")
GOFLAGS += $(shell pacman -T 'pacman>6' && echo "-tags six")

.PHONY: default
default: build

.PHONY: build
build:
	$(GO) build $(GOFLAGS) -ldflags '-s -w $(LDFLAGS)' $(EXTRA_GOFLAGS)

.PHONY: test
test:
	@test -z "$$(gofmt -l *.go)" || (echo "Files need to be linted. Use make fmt" && false)
	$(GO) test $(GOFLAGS)  .

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: clean
clean:
	go clean --modcache
