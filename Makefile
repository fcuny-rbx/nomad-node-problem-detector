BINARY ?= npd
BINDIR ?= $(DESTDIR)/usr/local/bin
ifndef $(GOLANG)
    GOLANG=$(shell which go)
    export GOLANG
endif

export GO111MODULE=on
export NOMAD_ADDR=http://localhost:4646
export NOMAD_E2E=1

GIT_COMMIT ?= $(shell git rev-parse HEAD)
BUILD_DATE ?= $(shell date +%s)

LDFLAGS += -X 'main.Timestamp=${BUILD_DATE}'
LDFLAGS += -X 'main.GitCommit=${GIT_COMMIT}'

default: build

.PHONY: clean
clean:
	rm -f $(BINARY)

.PHONY: build
build:
	$(GOLANG) build -ldflags "$(LDFLAGS)" -o $(BINARY) .

.PHONY: install
install:
	$(GOLANG) build -o $(BINARY) .
	install -m 755 $(BINARY) $(BINDIR)/$(BINARY)

.PHONY: test
test:
	$(GOLANG) test -count=1 -v ./...
