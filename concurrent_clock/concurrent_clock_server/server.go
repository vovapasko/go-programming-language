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
const (
	UsEasternTz    = "US/Eastern"
	AsiaTokyoTz    = "Asia/Tokyo"
	EuropeLondonTz = "Europe/London"
)

func main() {
	host, timezone := "", ""

	if len(os.Args) < 2 {
		host = defaultHost
	} else {
		host = os.Args[1]
	}
	if len(os.Args) < 3 {
		timezone = EuropeLondonTz
	} else {
		timezone = os.Args[2]
	}
	Serve(os.Stdout, host, timezone)
}

func Serve(w io.Writer, host, timezone string) {
	_, _ = fmt.Fprintf(w, "Started server and listening on the port %v and timezone %v", host, timezone)
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
		go handleConnection(connection, clientNum, timezone)
	}
}

func handleConnection(conn net.Conn, clientNum int, timezone string) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
	for {
		currentTime := getLocalTime(time.Now(), timezone)
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

func getLocalTime(timeUTC time.Time, timeZoneCode string) string {
	loc, _ := time.LoadLocation(timeZoneCode)
	return timeUTC.In(loc).Format(defaultTimeFormat)
}
