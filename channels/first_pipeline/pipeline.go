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
		for x := 0; x < 50; x++ {
			numbers <- x
		}
		close(numbers)
	}()
	go func() {
		for number := range numbers {
			squares <- number * number
		}
		close(squares)
	}()
	for square := range squares {
		fmt.Println(square)
		time.Sleep(10 * time.Millisecond)
	}
}
