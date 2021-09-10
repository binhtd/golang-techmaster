package main

import (
	"day2/bai1"
	"fmt"
)

func main() {
	_, err := bai1.SecondLargestNumber2([]int{2, 1, 4, 4})
	if err != nil {
		fmt.Println("error")
	}
}
