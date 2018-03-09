package main

import (
	"io"
	"net"

	"github.com/prometheus/common/log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliPort    = kingpin.Flag("port", "Port to forward").Default("9000").Envar("PORT").String()
	cliGateway = kingpin.Flag("gateway", "Host IP address which is running xdebug").Default("docker.for.mac.localhost").Envar("GATEWAY").String()
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
			log.Errorf("Failed to accept listener: %v", err)
			continue
		}

		go forward(conn, *cliGateway+":"+*cliPort)
	}
}

func forward(conn net.Conn, b string) {
	client, err := net.Dial("tcp", b)
	if err != nil {
		log.Errorf("Fail to dial: %v", err)
	}

	go func() {
		defer client.Close()
		defer conn.Close()

		_, err := io.Copy(client, conn)
		if err != nil {
			log.Errorf("Fail to copy connection: %v", err)
		}
	}()

	go func() {
		defer client.Close()
		defer conn.Close()

		_, err := io.Copy(conn, client)
		if err != nil {
			log.Errorf("Fail to copy connection: %v", err)
		}
	}()
}
