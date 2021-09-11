package main

import (
	"fmt"
)

func DemoBasicPointer() {
	age := new(int)
	*age = 26

	fmt.Printf("value of age %d\n", *age)
	fmt.Printf("address of age pointer %p\n", &age)
	fmt.Printf("address of pointer age reference %p\n", age)

	tuoi := 27
	age = &tuoi
	fmt.Printf("value of tuoi %d\n", *age)
	(*age) += 1
	fmt.Println(tuoi)

}
