#!/usr/bin/make -f

GO=go
GB=gb

darwin:
	env GOOS=darwin GOARCH=amd64 $(GB) build

linux:
	env GOOS=linux GOARCH=amd64 $(GB) build

docker: linux
	docker build -t nickschuch/d4m-tcp-forwarder .

all: test clean darwin linux

clean:
	rm -fR pkg bin

test:
	$(GB) test -test.v
