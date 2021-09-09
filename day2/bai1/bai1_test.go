package bai1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
max 2nd number is distinct
*/
func TestSecondLargestNumber1Case1(t *testing.T) {
	assert := assert.New(t)
	max2nd, _ := SecondLargestNumber1([]int{2, 1, 3, 4})
	assert.Equal(max2nd, 3, "max2d == 3")
}

/*
max 2nd number is equal max number
*/
func TestSecondLargestNumber1Case2(t *testing.T) {
	assert := assert.New(t)
	max2nd, _ := SecondLargestNumber1([]int{2, 1, 4, 4})
	assert.Equal(max2nd, 4, "max2d == 4")
}

/*
total element is less than 2
*/
func TestSecondLargestNumber1Case3(t *testing.T) {
	assert := assert.New(t)
	_, err := SecondLargestNumber1([]int{2})
	assert.True(err != nil)
}

/*
max 2nd number is distinct
*/
func TestSecondLargestNumber2Case1(t *testing.T) {
	assert := assert.New(t)
	max2nd, _ := SecondLargestNumber2([]int{2, 1, 3, 4})
	assert.Equal(max2nd, 3, "max2d == 3")
}

/*
max 2nd number is equal max number
*/
func TestSecondLargestNumber2Case2(t *testing.T) {
	assert := assert.New(t)
	max2nd, _ := SecondLargestNumber2([]int{2, 1, 4, 4})
	assert.Equal(max2nd, 4, "max2d == 4")
}

/*
total element is less than 2
*/
func TestSecondLargestNumber2Case3(t *testing.T) {
	assert := assert.New(t)
	_, err := SecondLargestNumber2([]int{2})
	assert.True(err != nil)
}

/*
benchmark 1st solution
*/
func Benchmark_SecondLargestNumber1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = SecondLargestNumber1([]int{2, 1, 3, 4, 5, 9, 8, 4, 2, 3})
	}
}

/*
benchmark 2st solution
*/
func Benchmark_SecondLargestNumber2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = SecondLargestNumber2([]int{2, 1, 3, 4, 5, 9, 8, 4, 2, 3})
	}
}
