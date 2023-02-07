package main

import (
	"fmt"
	"time"
)

func main() {
	// Number generator -> Squarer -> Printer
	numbers := make(chan int)
	squares := make(chan int)
	go counter(numbers)
	go squarer(numbers, squares)

	// method 1
	//go printer(squares)
	//time.Sleep(time.Second)

	// or just start in main goroutine
	printer(squares)
}

func printer(in <-chan int) {
	for square := range in {
		fmt.Println(square)
		time.Sleep(10 * time.Millisecond)
	}
}

func counter(out chan<- int) {
	for x := 0; x < 50; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for number := range in {
		out <- number * number
	}
	close(out)
}
