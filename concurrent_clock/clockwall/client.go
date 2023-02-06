package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
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
