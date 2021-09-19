package main

import (
	"flag"
	"fmt"
)

func demoCommandLineFlag() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("numb", 42, "ant int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("work:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
