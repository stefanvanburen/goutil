package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestEditDistance(t *testing.T) {
	is := is.New(t)

	tests := []struct {
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
		is.Equal(editdistance1, x.out)
		is.Equal(editdistance2, x.out)
	}
}
