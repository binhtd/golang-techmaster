package main

import (
	"fmt"
)

func ArgumentToPrintlnEscapeToHeap() {
	x := 100
	y := x
	fmt.Println(y)
	tom := Person{Name: "Tom", Age: 18}
	tom.Age += 1
	fmt.Println(tom)

}
