package bai2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
return normal array
*/
func TestFindListElementHasMaxLength1Case1(t *testing.T) {
	assert := assert.New(t)
	listElementHasMaxLength, _ := FindListElementHasMaxLength1([]string{"aba", "aa", "ad", "c", "vcd"})
	assert.ElementsMatch(listElementHasMaxLength, []string{"aba", "vcd"})
}

/*
return empty array
*/
func TestFindListElementHasMaxLength1Case2(t *testing.T) {
	assert := assert.New(t)
	listElementHasMaxLength, _ := FindListElementHasMaxLength1([]string{})
	assert.ElementsMatch(listElementHasMaxLength, []string{})
}

/*
return normal array
*/
func TestFindListElementHasMaxLength2Case1(t *testing.T) {
	assert := assert.New(t)
	listElementHasMaxLength, _ := FindListElementHasMaxLength2([]string{"aba", "aa", "ad", "c", "vcd"})
	assert.ElementsMatch(listElementHasMaxLength, []string{"aba", "vcd"})
}

/*
return empty array
*/
func TestFindListElementHasMaxLength2Case2(t *testing.T) {
	assert := assert.New(t)
	listElementHasMaxLength, _ := FindListElementHasMaxLength2([]string{})
	assert.ElementsMatch(listElementHasMaxLength, []string{})
}

/*
benchmark FindListElementHasMaxLength1
*/
func Benchmark_FindListElementHasMaxLength1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = FindListElementHasMaxLength1([]string{"aba", "aa", "ad", "c", "vcd"})
	}
}

/*
benchmark FindListElementHasMaxLength2
*/
func Benchmark_FindListElementHasMaxLength2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, _ = FindListElementHasMaxLength2([]string{"aba", "aa", "ad", "c", "vcd"})
	}
}
