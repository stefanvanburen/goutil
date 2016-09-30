package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	x := []int{
		-17, 5, 2, 18, 20, -64, 8,
	}
	y := []int{
		-10, -9, -8, -6,
	}

	assert.Equal(t, Max(x), 20, "Max should be 20")
	assert.Equal(t, Max(y), -6, "Max should be -6")
}
