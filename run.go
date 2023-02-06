package main

import "my_concurrency/concurrent_clock_server"

func main() {
	concurrent_clock_server.Serve("localhost:8000")
}
