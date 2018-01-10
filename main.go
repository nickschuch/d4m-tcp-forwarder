package main

import (
	"io"
	"log"
	"net"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliPort    = kingpin.Flag("port", "Port to forward").Default("9000").OverrideDefaultFromEnvar("PORT").String()
	cliGateway = kingpin.Flag("gateway", "Host IP address which is running xdebug").Default("docker.for.mac.localhost").OverrideDefaultFromEnvar("GATEWAY").String()
)

func main() {
	kingpin.Parse()

	l, err := net.Listen("tcp", ":"+*cliPort)
	if err != nil {
		log.Fatalf("Failed to setup listener: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		go forward(conn, *cliGateway+":"+*cliPort)
	}
}

func forward(conn net.Conn, b string) {
	client, err := net.Dial("tcp", b)
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(client, conn)
	}()
	go func() {
		defer client.Close()
		defer conn.Close()
		io.Copy(conn, client)
	}()
}
