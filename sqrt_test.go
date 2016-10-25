package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtInt(t *testing.T) {
	var tests = []struct {
		in  int
		out int
	}{
		{4, 2},
		{9, 3},
		{100, 10},
		{144, 12},
	}

	for _, x := range tests {
		assert.Equal(t, SqrtInt(x.in), x.out, "Should be equal")
	}
}

func TestSqrtInt64(t *testing.T) {
	var tests = []struct {
		in  int64
		out int64
	}{
		{4, 2},
		{9, 3},
		{100, 10},
		{144, 12},
	}

	for _, x := range tests {
		assert.Equal(t, SqrtInt64(x.in), x.out, "Should be equal")
	}
}
