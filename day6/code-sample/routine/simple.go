package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func demoGoRoutineSimple() {
	f("direct")

	go f("gorountine")

	go func(msg string) {
		time.Sleep(8 * time.Millisecond)
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
