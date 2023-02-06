package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection lost")
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("Something happened with the connection")
		}
	}(conn)
	giveOutput(conn)
}

func giveOutput(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
