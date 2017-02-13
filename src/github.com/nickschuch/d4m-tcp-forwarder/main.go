package main

import (
	"io"
	"log"
	"net"

	"gopkg.in/alecthomas/kingpin.v2"
)

const gateway = "192.168.65.1"

var port = kingpin.Flag("port", "Port to forward").Default("9000").OverrideDefaultFromEnvar("PORT").String()

func main() {
	kingpin.Parse()

	l, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Failed to setup listener: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		go forward(conn, gateway+":"+*port)
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
