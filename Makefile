VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"
MODFLAGS=-mod=vendor

SAGA_PACKAGE=./cmd/saga
SAGA_BINARY=saga

all: dev

clean:
	rm -fr dist/

dev:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/${SAGA_BINARY} ${SAGA_PACKAGE}

test:
	go test ${MODFLAGS} ./...

coverage:
	go test ${MODFLAGS} -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o dist/coverage.html

.PHONY: all clean dev test coverage
