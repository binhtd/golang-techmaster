package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
check number is 2, 3, 5
*/
func TestIsPrime(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsPrime(2))
	assert.True(IsPrime(3))
	assert.True(IsPrime(5))
}

/*
check number is 4
*/
func TestNotIsPrime(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsPrime(4))
}
