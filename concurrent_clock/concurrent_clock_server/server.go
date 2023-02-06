package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var clientNum = 0

const defaultTimeFormat = "2006-01-02 15:04:05"
const defaultHost = "localhost:8000"

func main() {
	host := ""
	if len(os.Args) < 2 {
		host = defaultHost
	} else {
		host = os.Args[1]
	}
	Serve(os.Stdout, host)
}

func Serve(w io.Writer, host string) {
	_, _ = fmt.Fprint(w, "Started server and listening on the port "+host)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		clientNum += 1
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(connection, clientNum)
	}
}

func handleConnection(conn net.Conn, clientNum int) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	for {
		currentTime := getCurrentTime(defaultTimeFormat, time.Now)
		_, err := fmt.Fprintf(conn, "[%s]: hello from client %v\n", currentTime, clientNum)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func getCurrentTime(format string, timeNow func() time.Time) string {
	t := timeNow()
	return t.Format(format)
}
