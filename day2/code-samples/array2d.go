package main

import "fmt"

func demoArray2d() {
	langs := [][]string{
		{"language 1", "language 2", "language 3"},
		{"language 4", "language 5", "language 6"},
		{"language 7", "language 8", "language 9"},
	}

	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang, ", ")
		}
		fmt.Println()
	}
}
