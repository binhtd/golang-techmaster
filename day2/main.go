package main

import (
	"day2/bai4"
	"fmt"
)

func main() {
	people := []bai4.Employee{
		{"Gopher", 70, 7},
		{"Alice", 55, 5},
		{"Vera", 55, 5},
		{"Bob", 75, 2},
	}
	max2ndSalaryPeople := bai4.ListEmployeeIsMax2nd(people)
	fmt.Println(max2ndSalaryPeople)
}
