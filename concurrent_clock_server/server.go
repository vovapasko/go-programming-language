package concurrent_clock_server

import (
	"fmt"
	"log"
	"net"
	"time"
)

var clientNum = 0

const defaultTimeFormat = "2006-01-02 15:04:05"

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
		currentTime := getCurrentTime(defaultTimeFormat)
		_, err := fmt.Fprintf(conn, "[%s]: hello from client %v\n", currentTime, clientNum)
		if err != nil {
			clientNum += 1
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func getCurrentTime(format string) string {
	t := time.Now()
	return t.Format(format)
}
