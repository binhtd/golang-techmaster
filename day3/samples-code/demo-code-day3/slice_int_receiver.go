package main

import (
	"fmt"
)

type SliceInt []int

func (p SliceInt) DoubleIt() {
	for i, v := range p {
		p[i] = v * 2
	}
}

func DemoSliceIntReceiver() {
	var sliceInt SliceInt = []int{1, 2, 3, 4}
	sliceInt.DoubleIt()
	fmt.Println(sliceInt)
}
