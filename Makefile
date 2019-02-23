VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT}"
MODFLAGS=-mod=vendor
TESTFLAGS=-cover

PLATFORMS:=darwin linux windows

all: dev

dev:
	@echo building dist/saga...
	@go build ${MODFLAGS} ${LDFLAGS} -o dist/saga ./cmd/saga

dist: $(PLATFORMS)

$(PLATFORMS):
	@echo building dist/saga-$@-amd64...
	@GOOS=$@ GOARCH=amd64 go build ${MODFLAGS} ${LDFLAGS} -o dist/saga-$@-amd64 ./cmd/saga

clean:
	@rm -fr dist/

test:
	@go test ${MODFLAGS} ${TESTFLAGS} ./...

.PHONY: all dev dist $(PLATFORMS) clean test
