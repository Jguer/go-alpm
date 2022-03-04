export GO111MODULE=on

GO ?= go

SOURCES ?= $(shell find . -name "*.go")
GOFLAGS += $(shell pacman -T 'pacman-git' && echo "-tags next")

.PHONY: test
test:
	@test -z "$$(gofmt -l *.go)" || (echo "Files need to be linted. Use make fmt" && false)
	$(GO) test $(GOFLAGS) -v ./...

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: clean
clean:
	go clean --modcache
