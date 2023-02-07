package main

import (
	"io"
	"log"
	"net"
	"os"
)

const defaultHost = "localhost:8000"

func main() {
	host := defaultHost
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	StartConnection(host)
}

func StartConnection(host string) {
	conn, err := net.Dial("tcp", host)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Could not close a connection")
		}
	}(conn)
	if err != nil {
		log.Fatal("Connection lost")
	}
	handleOutput(conn)
}

func handleOutput(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
