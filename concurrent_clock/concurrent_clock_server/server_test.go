package main

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	// test starting server and check the connection
	host := "localhost:8000"
	buff := &bytes.Buffer{}
	go Serve(buff, host)
	time.Sleep(100 * time.Millisecond)
	_, err := net.Dial("tcp", host)
	if err != nil {
		t.Fatal("Connection failed")
	}

	want := "Started server and listening on the port " + host
	got := buff.String()
	if err != nil {
		t.Fatal("Failed to receive a message")
	}
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestGetCurrentTime(t *testing.T) {
	mockTime := time.Date(2020, time.September, 13, 17, 30, 0, 0, time.UTC)
	timeNow := func() time.Time {
		return mockTime
	}
	got := getCurrentTime(defaultTimeFormat, timeNow)
	want := "2020-09-13 17:30:00"
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
