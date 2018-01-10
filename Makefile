#!/usr/bin/make -f

export CGO_ENABLED=0

PROJECT=github.com/nickschuch/d4m-tcp-forwarder

# Builds the project
build:
	gox -os='linux darwin' -arch='amd64' -output='bin/d4m-tcp-forwarder_{{.OS}}_{{.Arch}}' -ldflags='-extldflags "-static"' $(PROJECT)

# Run all lint checking with exit codes for CI
lint:
	golint -set_exit_status `go list ./... | grep -v /vendor/`

# Run tests with coverage reporting
test:
	go test -cover ./server/...
	go test -cover ./cmd/...

IMAGE=nickschuch/d4m-tcp-forwarder
VERSION=$(shell git describe --tags --always)

# Releases the project Docker Hub
release:
	# Building M8s versioned image...
	docker build -t ${IMAGE}:${VERSION} .
	docker push ${IMAGE}:${VERSION}
	# Building M8s latest image...
	docker build -t ${IMAGE}:latest .
	docker push ${IMAGE}:latest

.PHONY: build lint test release
