Docker for Mac: TCP Forwarder
=============================

Sends traffic from the Docker for Mac host, onto our OSX host.

eg.

Xdebug picks up the Docker for Mac VMs network bridge as the source IP and sends it to that instead of the OSX host. This project fixes that.

## Assumptions

* The OSX host IP does not change from `192.168.65.1`

## Usage

To use this approach, merge the following into your existing docker-compose.yml file:

version: '2'
services:
  xdebug:
    image: nickschuch/d4m-tcp-forwarder:latest
    network_mode: host
    environment:
      - PORT=9000

