package main

import (
	"fmt"
	"time"
)

func greeting() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good after noon")
	default:
		fmt.Print("Good evening")
	}
}
