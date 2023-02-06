package main

import (
	"fmt"
	"time"
)

func main() {
	//fibonacci.InteractiveFibonacci(40)
	now, _ := time.Parse("02-01-2006 15:04:05 -0700", "07-05-2021 12:00:00 +0530")

	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time:       %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("Berlin Time:    %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time:  %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Kolkata")
	fmt.Printf("India Time:     %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Singapore")
	fmt.Printf("Singapore Time: %s\n", now.In(loc))
}
