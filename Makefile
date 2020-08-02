VERSION := 1.0

ifdef GOBIN
PATH := $(GOBIN):$(PATH)
else
PATH := $(subst :,/bin:,$(shell go env GOPATH))/bin:$(PATH)
endif

LDFLAGS := $(LDFLAGS) -X main.version=$(VERSION)

.PHONY: deps
deps:
	go mod download

.PHONY: acme
acme: deps
	@echo "Building platform binary..."
	go build -ldflags "$(LDFLAGS)" ./cmd/acme
	mv ./acme ./acme_osx

.PHONY: test
test: deps
	go test -short ./...

.PHONY: linux
linux: deps
	@echo "Building static linux binary..."
	@CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags "$(LDFLAGS)" ./cmd/acme
	mv ./acme ./acme_linux
