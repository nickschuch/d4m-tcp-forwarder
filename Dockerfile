FROM golang:1.8
ADD . /go/src/github.com/nickschuch/d4m-tcp-forwarder
WORKDIR /go/src/github.com/nickschuch/d4m-tcp-forwarder
RUN go get github.com/mitchellh/gox
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/nickschuch/d4m-tcp-forwarder/bin/d4m-tcp-forwarder_linux_amd64 /usr/local/bin/d4m-tcp-forwarder

ENTRYPOINT ["/usr/local/bin/d4m-tcp-forwarder"]
