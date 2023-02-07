package main

import (
	"fmt"
	"time"
)

func main() {
	// Number generator -> Squarer -> Printer
	numbers := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; ; x++ {
			numbers <- x
		}

	}()
	go func() {
		for {
			x := <-numbers
			squares <- x * x
		}
	}()
	for {
		fmt.Println(<-squares)
		time.Sleep(100 * time.Millisecond)
	}

}
