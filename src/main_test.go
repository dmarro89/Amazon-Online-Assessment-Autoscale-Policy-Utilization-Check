package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, getInstancesNumber(2, []int{25, 23, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 76, 80}), 2)
	assert.Equal(t, getInstancesNumber(5, []int{30, 5, 4, 8, 19, 89}), 3)
}
