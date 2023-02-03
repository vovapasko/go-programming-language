package fibonacci

import (
	"fmt"
	"time"
)

func InteractiveFibonacci(n int) int {
	go spinner(100 * time.Millisecond)
	wantedNumber := fib(n)
	fmt.Printf("\r %d", wantedNumber)
	return wantedNumber
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
