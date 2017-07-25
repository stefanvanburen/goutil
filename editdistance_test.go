package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditDistance(t *testing.T) {
	var tests = []struct {
		a   string
		b   string
		out int
	}{
		{"hello", "hello", 0},
		{"hello", "hell0", 1},
		{"hello", "hllo", 1},
		{"hello", "h3ll0", 2},
	}

	for _, x := range tests {
		editdistance1, _ := EditDistance1(x.a, x.b)
		editdistance2, _ := EditDistance2(x.a, x.b)
		assert.Equal(t, editdistance1, x.out, "Should be equal")
		assert.Equal(t, editdistance2, x.out, "Should be equal")
	}
}
