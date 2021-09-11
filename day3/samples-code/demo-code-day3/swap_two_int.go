package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func DemoSwap() {
	x, y := 5, 10
	fmt.Println("Before swap: ", x, y)
	swap(&x, &y)
	fmt.Println("After swap:", x, y)
}
