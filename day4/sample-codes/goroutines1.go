package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func demoGoroutines1() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)

	}("going")
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
