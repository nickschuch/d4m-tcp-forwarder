FROM       gliderlabs/alpine:3.1
MAINTAINER Nick Schuch

RUN apk --update add ca-certificates

ADD bin/d4m-tcp-forwarder-linux-amd64 /d4m-tcp-forwarder
RUN chmod a+x /d4m-tcp-forwarder

CMD ["/d4m-tcp-forwarder"]
