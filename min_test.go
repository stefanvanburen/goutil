package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	x := []int{
		-17, 5, 2, 18, 20, -64, 8,
	}
	y := []int{
		10, 9, 8, 7,
	}

	assert.Equal(t, Min(x), -64, "Min should be -64")
	assert.Equal(t, Min(y), -6, "Min should be -6")
}
