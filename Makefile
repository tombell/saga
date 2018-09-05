BINARY=saga
GOARCH=amd64

VERSION?=dev
COMMIT=$(shell git rev-parse HEAD | cut -c -8)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT}"
MODFLAGS=-mod=vendor

PACKAGE=./cmd/saga

all: dev

clean:
	rm -fr dist/

dev:
	go build ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

cibuild:
	go build ${MODFLAGS} ${LDFLAGS} -o dist/${BINARY} ${PACKAGE}

dist: darwin windows

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-darwin-${GOARCH} ${PACKAGE}

windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o dist/${BINARY}-windows-${GOARCH} ${PACKAGE}

test:
	@go test ${MODFLAGS} github.com/tombell/saga

.PHONY: all clean dev cibuild dist darwin windows test
