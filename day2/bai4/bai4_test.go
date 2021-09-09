package bai4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
test sortEmployeeByName
*/
func TestSortEmployeeByName(t *testing.T) {
	assert := assert.New(t)
	people := []Employee{
		{"Gopher", 70, 7},
		{"Alice", 55, 5},
		{"Vera", 24, 2},
		{"Bob", 75, 2},
	}
	SortEmployeeByName(people)
	assert.ElementsMatchf(people, []Employee{
		{"Alice", 55, 5},
		{"Bob", 75, 2},
		{"Gopher", 70, 7},
		{"Vera", 24, 2},
	}, "equal element")
}

/*
test sortEmployeeBySalary
*/
func TestSortEmployeeBySalary(t *testing.T) {
	assert := assert.New(t)
	people := []Employee{
		{"Gopher", 70, 7},
		{"Alice", 55, 5},
		{"Vera", 24, 2},
		{"Bob", 75, 2},
	}
	SortEmployeeByName(people)
	assert.ElementsMatchf(people, []Employee{
		{"Vera", 24, 2},
		{"Alice", 55, 5},
		{"Gopher", 70, 7},
		{"Bob", 75, 2},
	}, "equal element")
}

/*
test listEmployeeIsMax2nd
*/
func TestListEmployeeIsMax2nd(t *testing.T) {
	assert := assert.New(t)
	people := []Employee{
		{"Gopher", 70, 7},
		{"Alice", 70, 7},
		{"Vera", 55, 5},
		{"Bob", 75, 2},
	}
	max2ndSalaryPeople := ListEmployeeIsMax2nd(people)
	assert.ElementsMatchf(max2ndSalaryPeople, []Employee{
		{"Gopher", 70, 7},
		{"Alice", 70, 7},
	}, "equal element")
}
