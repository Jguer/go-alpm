export GO111MODULE=on

GO ?= go

SOURCES ?= $(shell find . -name "*.go")

.PHONY: test
test:
	@test -z "$$(gofmt -l *.go)" || (echo "Files need to be linted. Use make fmt" && false)
	$(GO) test -race -covermode=atomic -v .

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: clean
clean:
	go clean --modcache
