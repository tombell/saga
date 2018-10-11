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

.PHONY: all clean dev test
