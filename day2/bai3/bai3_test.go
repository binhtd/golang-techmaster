package bai3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
return normal list without duplicate element
*/
func TestRemoveDuplicates1Case1(t *testing.T) {
	assert := assert.New(t)
	nonDuplicateList, _ := RemoveDuplicates1([]int{1, 2, 5, 2, 6, 2, 5})
	assert.ElementsMatch(nonDuplicateList, []int{1, 2, 5, 6})
}

/*
return empty list
*/
func TestRemoveDuplicates1Case2(t *testing.T) {
	assert := assert.New(t)
	nonDuplicateList, _ := RemoveDuplicates1([]int{})
	assert.ElementsMatch(nonDuplicateList, []int{})
}

/*
return normal list without duplicate element
*/
func TestRemoveDuplicates2Case1(t *testing.T) {
	assert := assert.New(t)
	nonDuplicateList, _ := RemoveDuplicates2([]int{1, 2, 5, 2, 6, 2, 5})
	assert.ElementsMatch(nonDuplicateList, []int{1, 2, 5, 6})
}

/*
return empty list
*/
func TestRemoveDuplicates2Case2(t *testing.T) {
	assert := assert.New(t)
	nonDuplicateList, _ := RemoveDuplicates2([]int{})
	assert.ElementsMatch(nonDuplicateList, []int{})
}

/*
benchmark removeDuplicates1
*/
func Benchmark_RemoveDuplicates1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = RemoveDuplicates1([]int{1, 2, 5, 2, 6, 2, 5})
	}
}

/*
benchmark removeDuplicates2
*/
func Benchmark_RemoveDuplicates2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = RemoveDuplicates2([]int{1, 2, 5, 2, 6, 2, 5})
	}
}
