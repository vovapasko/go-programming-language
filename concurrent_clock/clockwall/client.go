package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

const defaultHost = "localhost:8000"

func main() {
	hostsList := make([]string, 0)
	if len(os.Args) == 1 {
		hostsList = append(hostsList, defaultHost)
	} else {
		hostsList = os.Args[1:]
	}
	RunOnHosts(hostsList)
}

func RunOnHosts(hosts []string) {
	for _, host := range hosts {
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
		go handleOutput(conn)
	}
	for {
		time.Sleep(time.Second)
	}
}

func handleOutput(conn net.Conn) {
	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
