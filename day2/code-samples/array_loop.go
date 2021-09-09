package main

import (
	"fmt"
)

func indexLoop() {
	cars := []string{"car1", "car2", "car3"}

	for i := 0; i < len(cars); i++ {
		fmt.Print(i, ", ", cars[i])
	}
}

func whileLoop() {
	cars := []string{"car1", "car2", "car3"}
	i := 0
	for i < len(cars) {
		fmt.Print(i, ", ", cars[i])
		i++
	}
}

func rangeLoop() {
	cars := []string{"car1", "car2", "car3"}

	for _, v := range cars {
		fmt.Print(v, " ")
	}
}

func reserveLoop() {
	cars := []string{"car1", "car2", "car3"}

	for _, v := range cars {
		defer fmt.Println(v)
	}
}

func rawReserveLoop() {
	cars := []string{"car1", "car2", "car3"}

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[len(cars)-i-1])
	}
}
