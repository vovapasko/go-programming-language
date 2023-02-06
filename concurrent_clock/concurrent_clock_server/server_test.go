package main

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	host := "localhost:8000"
	buff := &bytes.Buffer{}
	runInTimezone := EuropeLondonTz
	go Serve(buff, host, runInTimezone)
	time.Sleep(100 * time.Millisecond)
	_, err := net.Dial("tcp", host)
	if err != nil {
		t.Fatal("Connection failed")
	}
	want := "Started server and listening on the port " + host + " and timezone " + runInTimezone
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

type TimezoneTest struct {
	Timezone, Time string
}

func TestGetTimeForTimezone(t *testing.T) {
	expectedTimes := []TimezoneTest{
		{Timezone: EuropeLondonTz, Time: "2020-09-13 18:30:00"},
		{Timezone: AsiaTokyoTz, Time: "2020-09-14 02:30:00"},
		{Timezone: UsEasternTz, Time: "2020-09-13 13:30:00"},
	}

	mockTime := time.Date(2020, time.September, 13, 17, 30, 0, 0, time.UTC)

	for _, timeTestCase := range expectedTimes {
		gotLocalTime := getLocalTime(mockTime, timeTestCase.Timezone)
		wantLocalTime := timeTestCase.Time
		if wantLocalTime != gotLocalTime {
			t.Errorf("Wanted %v time, got %v time", wantLocalTime, gotLocalTime)
		}
	}

}
