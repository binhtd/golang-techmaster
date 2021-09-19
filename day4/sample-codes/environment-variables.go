package main

import (
	"fmt"
	"os"
	"strings"
)

func demoEnv() {
	os.Setenv("FOO", "1")
	fmt.Println(os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
