package concurrent_clock_server

import (
	"fmt"
	"log"
	"net"
	"time"
)

var clientNum = 0

func Serve(host string) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	for {
		_, err := fmt.Fprintf(conn, "hello from client %v\n", clientNum)
		if err != nil {
			clientNum += 1
			return
		}
		time.Sleep(1 * time.Second)
	}
}
