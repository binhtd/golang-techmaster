package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
has x1, x2 and both number is different
*/
func TestQuadraticEquation1(t *testing.T) {
	assert := assert.New(t)
	x1, x2 := QuadraticEquation(1, -3, 2)
	assert.Equal(x1, 2.000000, "x1 = 2.000000")
	assert.Equal(x2, 1.000000, "x2 = 1.000000")
}

/*
has x1, x2 and both number is the same
*/
func TestQuadraticEquation2(t *testing.T) {
	assert := assert.New(t)
	x1, x2 := QuadraticEquation(1, -2, 1)
	assert.Equal(x1, 1.000000, "x1 = 1.000000")
	assert.Equal(x2, 1.000000, "x2 = 1.000000")
}

/*
does not have x1, x2
*/
/*
func TestQuadraticEquation3(t *testing.T) {
	assert.Panics(t, QuadraticEquation(1, 1, 1), Delta_Less_Than_Zero_Mesage)
}
*/
